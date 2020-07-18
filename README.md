# DataCommons Mixer Development

## Development Process

In Mixer GitHub [repo](https://github.com/datacommonsorg/mixer), click on "Fork"
button to fork the repo.

Clone your forked repo to your desktop.

Add datacommonsorg/mixer repo as a remote.

```shell
git remote add dc https://github.com/datacommonsorg/mixer.git
```

Every time when you want to send a Pull Request, do the following steps:

```shell
git checkout master
git pull dc master
git checkout -b new_branch_name
# Make some code change
git add .
git commit -m "commit message"
git push -u origin new_branch_name
```

Then in your forked repo, can send a Pull Request.

Wait for approval of the Pull Request and merge the change.

## Generate Protobuf go code

Install protoc by following
[this](http://google.github.io/proto-lens/installing-protoc.html).

Run the following code to generate golang proto files.

```bash
./prepare-proto.sh
protoc \
    --proto_path=proto \
    --go_out=. \
    --go-grpc_out=. \
    --go-grpc_opt=requireUnimplementedServers=false \
    proto/mixer.proto
```

## E2E test locally

Install `cloud-build-local` following
[the guide](https://cloud.google.com/cloud-build/docs/build-debug-locally), then
run:

```bash
cloud-build-local --config=cloudbuild.test.yaml --dryrun=false .
```

## Run grpc server and examples locally

```bash
gcloud auth application-default login

go run main.go \
    --bq_dataset=$(head -1 deployment/bigquery.txt) \
    --bt_table=$(head -1 deployment/bigtable.txt) \
    --bt_project=google.com:datcom-store-dev \
    --bt_instance=prophet-cache \
    --project_id=datcom-mixer-staging

# Open a new shell
cd examples/get_place_obs
go run main.go
```

## Update golden files on staging

Run the following commands to update golden files in ./golden_response/staging.

```bash
./update-golden.sh
```
