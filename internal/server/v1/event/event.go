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

// Package event contain code for event.
package event

import (
	"context"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/store"
	"github.com/datacommonsorg/mixer/internal/store/bigtable"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Collection implements API for Mixer.EventCollection.
func Collection(
	ctx context.Context,
	in *pb.EventCollectionRequest,
	store *store.Store,
) (*pb.EventCollectionResponse, error) {
	if in.GetEventType() == "" || in.GetDate() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing required arguments")
	}
	if !util.CheckValidDCIDs([]string{in.GetAffectedPlaceDcid()}) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid DCID")
	}

	btDataList, err := bigtable.Read(
		ctx,
		store.BtGroup,
		bigtable.BtEventCollection,
		[][]string{{in.GetEventType()}, {in.GetAffectedPlaceDcid()}, {in.GetDate()}},
		func(jsonRaw []byte) (interface{}, error) {
			var eventCollection pb.EventCollection
			err := proto.Unmarshal(jsonRaw, &eventCollection)
			return &eventCollection, err
		},
	)
	if err != nil {
		return nil, err
	}

	resp := &pb.EventCollectionResponse{}

	// Go through (ordered) import groups one by one, stop when data is found.
	for _, btData := range btDataList {
		for _, row := range btData {
			resp.EventCollection = row.Data.(*pb.EventCollection)
			break
		}
	}

	return resp, nil
}

// CollectionDate implements API for Mixer.EventCollectionDate.
func CollectionDate(
	ctx context.Context,
	in *pb.EventCollectionDateRequest,
	store *store.Store,
) (*pb.EventCollectionDateResponse, error) {
	if in.GetEventType() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing required arguments")
	}
	if !util.CheckValidDCIDs([]string{in.GetAffectedPlaceDcid()}) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid DCID")
	}

	btDataList, err := bigtable.Read(
		ctx,
		store.BtGroup,
		bigtable.BtEventCollectionDate,
		[][]string{{in.GetEventType()}, {in.GetAffectedPlaceDcid()}},
		func(jsonRaw []byte) (interface{}, error) {
			var d pb.EventCollectionDate
			err := proto.Unmarshal(jsonRaw, &d)
			return &d, err
		},
	)
	if err != nil {
		return nil, err
	}

	resp := &pb.EventCollectionDateResponse{}

	// Go through (ordered) import groups one by one, stop when data is found.
	for _, btData := range btDataList {
		for _, row := range btData {
			resp.EventCollectionDate = row.Data.(*pb.EventCollectionDate)
			break
		}
	}

	return resp, nil
}
