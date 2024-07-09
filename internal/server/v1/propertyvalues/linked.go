// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package propertyvalues

import (
	"context"
	"strings"

	pb "github.com/avsolatorio/datacommons-mixer/internal/proto"
	pbv1 "github.com/avsolatorio/datacommons-mixer/internal/proto/v1"
	"github.com/avsolatorio/datacommons-mixer/internal/server/placein"
	"github.com/avsolatorio/datacommons-mixer/internal/store"
	"github.com/avsolatorio/datacommons-mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LinkedPropertyValues implements mixer.LinkedPropertyValues handler.
func LinkedPropertyValues(
	ctx context.Context,
	in *pbv1.LinkedPropertyValuesRequest,
	store *store.Store,
) (*pbv1.PropertyValuesResponse, error) {
	nodeProperty := in.GetNodeProperty()
	parts := strings.Split(nodeProperty, "/")
	if len(parts) < 2 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request URI")
	}
	property := parts[len(parts)-1]
	node := strings.Join(parts[0:len(parts)-1], "/")
	valueNodeType := in.GetValueNodeType()
	// Check arguments
	if property != "containedInPlace" {
		return nil, status.Errorf(
			codes.InvalidArgument, "only support property 'containedInPlace'")
	}
	if valueNodeType == "" {
		return nil, status.Errorf(
			codes.InvalidArgument, "missing argument: value_node_type")
	}
	if err := util.CheckValidDCIDs([]string{node}); err != nil {
		return nil, err
	}
	resp, err := placein.GetPlacesIn(
		ctx,
		store,
		[]string{node},
		valueNodeType,
	)
	if err != nil {
		return nil, err
	}
	valueDcids := resp[node]
	// Fetch names
	data, _, err := Fetch(
		ctx,
		store,
		valueDcids,
		[]string{"name"},
		0,
		"",
		"out",
	)
	if err != nil {
		return nil, err
	}
	result := &pbv1.PropertyValuesResponse{}
	for _, dcid := range valueDcids {
		var name string
		if tmp, ok := data[dcid]["name"]; ok {
			name = tmp[""][0].Value
		}
		result.Values = append(result.Values,
			&pb.EntityInfo{
				Dcid: dcid,
				Name: name,
			},
		)
	}
	return result, nil
}

// BulkLinkedPropertyValues implements mixer.BulkLinkedPropertyValues handler.
func BulkLinkedPropertyValues(
	ctx context.Context,
	in *pbv1.BulkLinkedPropertyValuesRequest,
	store *store.Store,
) (*pbv1.BulkPropertyValuesResponse, error) {
	property := in.GetProperty()
	nodes := in.GetNodes()
	valueNodeType := in.GetValueNodeType()
	// Check arguments
	if property != "containedInPlace" {
		return nil, status.Errorf(
			codes.InvalidArgument, "only support property 'containedInPlace'")
	}
	if valueNodeType == "" {
		return nil, status.Errorf(
			codes.InvalidArgument, "missing argument: value_node_type")
	}
	if err := util.CheckValidDCIDs(nodes); err != nil {
		return nil, err
	}
	resp, err := placein.GetPlacesIn(
		ctx,
		store,
		nodes,
		valueNodeType,
	)
	if err != nil {
		return nil, err
	}
	valueDcids := []string{}
	for _, e := range resp {
		valueDcids = append(valueDcids, e...)
	}
	// Fetch names
	data, _, err := Fetch(
		ctx,
		store,
		valueDcids,
		[]string{"name"},
		0,
		"",
		"out",
	)
	if err != nil {
		return nil, err
	}
	result := &pbv1.BulkPropertyValuesResponse{}
	for _, n := range nodes {
		children := resp[n]
		oneNodeResult := &pb.NodePropertyValues{
			Node: n,
		}
		for _, dcid := range children {
			var name string
			if nameValues, ok := data[dcid]["name"]; ok {
				if len(nameValues) > 0 {
					name = nameValues[""][0].Value
				} else {
					name = ""
				}

			}
			oneNodeResult.Values = append(oneNodeResult.Values,
				&pb.EntityInfo{
					Dcid: dcid,
					Name: name,
				},
			)
		}
		result.Data = append(result.Data, oneNodeResult)
	}
	return result, nil
}
