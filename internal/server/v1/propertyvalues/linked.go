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

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/server/placein"
	"github.com/datacommonsorg/mixer/internal/store"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LinkedPropertyValues implements mixer.LinkedPropertyValues handler.
func LinkedPropertyValues(
	ctx context.Context,
	in *pb.LinkedPropertyValuesRequest,
	store *store.Store,
) (*pb.LinkedPropertyValuesResponse, error) {
	property := in.GetProperty()
	entity := in.GetEntity()
	valueEntityType := in.GetValueEntityType()
	// Check arguments
	if property != "containedInPlace" {
		return nil, status.Errorf(
			codes.InvalidArgument, "only support property 'containedInPlace'")
	}
	if !util.CheckValidDCIDs([]string{entity}) {
		return nil, status.Errorf(
			codes.InvalidArgument, "invalid entity %s", entity)
	}
	resp, err := placein.GetPlacesIn(
		ctx,
		store,
		[]string{entity},
		valueEntityType,
	)
	if err != nil {
		return nil, err
	}
	valueDcids := resp[entity]
	// Fetch names
	data, _, err := Fetch(
		ctx,
		store,
		[]string{"name"},
		valueDcids,
		0,
		"",
		"out",
	)
	if err != nil {
		return nil, err
	}
	result := &pb.LinkedPropertyValuesResponse{}
	for _, dcid := range valueDcids {
		var name string
		if tmp, ok := data["name"][dcid]; ok {
			name = tmp[0].Value
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
