# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#!/bin/bash


export PROJECT_ID=$1
export DOMAIN=$2

IMAGE=$(cat "docker_image.txt")
export IMAGE
echo "Docker image: ${IMAGE}"

SERVICE_ACCOUNT=mixer-robot@$PROJECT_ID.iam.gserviceaccount.com
if [ "$PROJECT_ID" == "datcom-mixer" ]; then
  echo "Set datacommons domain"
  DOMAIN="api.datacommons.org"
fi


cp template_deployment.yaml deployment.yaml
cp template_api_config.yaml api_config.yaml

# Set project id
perl -i -pe's/PROJECT_ID/$ENV{PROJECT_ID}/g' deployment.yaml api_config.yaml

# Set api title
if [ "$PROJECT_ID" == "datcom-mixer" ]; then
  perl -i -pe's/TITLE/Data Commons API/g' api_config.yaml
else
  perl -i -pe's/TITLE/Data Commons API ($ENV{PROJECT_ID})/g' api_config.yaml
fi

# Set docker image
perl -i -pe's/IMAGE/$ENV{IMAGE}/g' deployment.yaml

# Set endpints domain
if [[ $DOMAIN ]]; then
  perl -i -pe's/#_c\|//g' deployment.yaml
  perl -i -pe's/DOMAIN/$ENV{DOMAIN}/g' deployment.yaml
else
  perl -i -pe's/#_d\|//g' deployment.yaml
fi

# Set BQ dataset
if [ "$PROJECT_ID" == "datcom-mixer" ]; then
  bq_dataset_input_file="prod_bq_dataset.txt"
else
  bq_dataset_input_file="staging_bq_dataset.txt"
fi
BQ_DATASET=$(cat $bq_dataset_input_file)
export BQ_DATASET
perl -i -pe's/BQ_DATASET/$ENV{BQ_DATASET}/g' deployment.yaml
# Set BQ dataset in versioned_mapping MCFs
mkdir versioned_mapping
rm versioned_mapping/*
cp mapping/* versioned_mapping/
perl -i -pe's/BQ_DATASET/$ENV{BQ_DATASET}/g' versioned_mapping/base.mcf
perl -i -pe's/BQ_DATASET/$ENV{BQ_DATASET}/g' versioned_mapping/weather.mcf

# Set BT_INSTANCE, same for prod and staging.
perl -i -pe's/BT_INSTANCE/prophet-cache/g' deployment.yaml

# Set BT table
if [ "$PROJECT_ID" == "datcom-mixer" ]; then
  bt_table_input_file="prod_bt_table.txt"
else
  bt_table_input_file="staging_bt_table.txt"
fi
BT_TABLE=$(cat $bt_table_input_file)
export BT_TABLE
perl -i -pe's/BT_TABLE/$ENV{BT_TABLE}/g' deployment.yaml

# Get a static ip address
if ! [[ $(gcloud compute addresses list --global --filter='name:mixer-ip' --format=yaml) ]]; then
 gcloud compute addresses create mixer-ip --global
fi
ip=$(gcloud compute addresses list --global --filter='name:mixer-ip' --format='value(ADDRESS)')


# Deploy endpoints
perl -i -pe's/IP_ADDRESS/'"$ip"'/g' api_config.yaml

if [[ $DOMAIN ]]; then
 perl -i -pe's/#_c\|//g' api_config.yaml
 perl -i -pe's/DOMAIN/$ENV{DOMAIN}/g' api_config.yaml
else
 perl -i -pe's/#_d\|//g' api_config.yaml
fi

gcloud endpoints services deploy out.pb api_config.yaml


# GKE setup
gcloud components install kubectl
gcloud services enable container.googleapis.com

# Create GKE instance
# Use custom machine type with 1cpu and 5G memory per instance. There are 3 instance by default.
CLUSTER_NAME="mixer-cluster-high-mem"
if [[ $(gcloud container clusters list --filter="$CLUSTER_NAME" --format=yaml) ]]; then
  echo "$CLUSTER_NAME already exists, continue..."
else
  gcloud container clusters create $CLUSTER_NAME \
    --zone=us-central1-c \
    --machine-type=custom-2-5120
fi

gcloud container clusters get-credentials $CLUSTER_NAME

# Create namespace
kubectl create namespace mixer

if [[ $(kubectl get secret bigquery-key --namespace mixer -o yaml | grep 'key.json') ]]; then
  echo "The secret bigquery-key already exists..."
else
  echo "Creating new bigquery-key..."
  # Create service account key and mount secret
    key_ids=$(gcloud iam service-accounts keys list --iam-account "$SERVICE_ACCOUNT" --managed-by=user --format="value(KEY_ID)")
    while read -r key_id; do
      if [[ $key_id ]]; then
        gcloud iam service-accounts keys delete $key_id --iam-account $SERVICE_ACCOUNT
      fi
    done <<< "$key_ids"
  gcloud iam service-accounts keys create key.json --iam-account $SERVICE_ACCOUNT

  # Mount secrete
  kubectl create secret generic bigquery-key --from-file=key.json=key.json --namespace=mixer
fi

# Mount nginx config
kubectl create configmap nginx-config --from-file=nginx.conf --namespace=mixer

# Mount schema mapping volumes
kubectl delete configmap schema-mapping --namespace mixer
kubectl create configmap schema-mapping --from-file=versioned_mapping/ --namespace=mixer

# Create certificate
if [ $DOMAIN ]; then
cat <<EOT > custom-certificate.yaml
apiVersion: networking.gke.io/v1beta1
kind: ManagedCertificate
metadata:
  name: custom-certificate
  namespace: mixer
spec:
  domains:
    - $DOMAIN
EOT
kubectl apply -f custom-certificate.yaml
else
cat <<EOT > certificate.yaml
apiVersion: networking.gke.io/v1beta1
kind: ManagedCertificate
metadata:
  name: mixer-certificate
  namespace: mixer
spec:
  domains:
    - datacommons.endpoints.$PROJECT_ID.cloud.goog
EOT
kubectl apply -f certificate.yaml
fi


# Bring up service and pods
kubectl apply -f service.yaml
kubectl apply -f deployment.yaml


# Bring ingress with certificate
if [ "$PROJECT_ID" == "datcom-mixer" ]; then
  perl -i -pe's/#__//g' ingress-ssl.yaml
fi
kubectl apply -f ingress-ssl.yaml
