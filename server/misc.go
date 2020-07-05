// Copyright 2019 Google LLC
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

package server

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/datacommonsorg/mixer/proto"
	"github.com/datacommonsorg/mixer/sparql"
	"github.com/datacommonsorg/mixer/translator"
	"github.com/datacommonsorg/mixer/util"
)

// Translate implements API for Mixer.Translate.
func (s *Server) Translate(ctx context.Context,
	in *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	if in.GetSchemaMapping() == "" || in.GetSparql() == "" {
		return nil, fmt.Errorf("missing required arguments")
	}

	out := pb.TranslateResponse{}
	mappings, err := translator.ParseMapping(in.GetSchemaMapping())
	if err != nil {
		return nil, err
	}
	nodes, queries, opts, err := sparql.ParseQuery(in.GetSparql())
	if err != nil {
		return nil, err
	}
	trans, err := translator.Translate(
		mappings, nodes, queries, s.metadata.SubTypeMap, opts)
	if err != nil {
		return nil, err
	}
	out.Sql = trans.SQL
	translation, err := json.MarshalIndent(trans, "", "  ")
	if err != nil {
		return nil, err
	}
	out.Translation = string(translation)
	return &out, nil
}

// GetPopCategory implements API for Mixer.GetPopCategory.
func (s *Server) GetPopCategory(ctx context.Context,
	in *pb.GetPopCategoryRequest) (*pb.GetPopCategoryResponse, error) {
	if in.GetPlaceType() == "" {
		return nil, fmt.Errorf("missing required arguments")
	}
	btRow, err := s.btTable.ReadRow(
		ctx, util.BtPopCategoryPrefix+in.GetPlaceType())
	if err != nil {
		return nil, err
	}
	out := pb.GetPopCategoryResponse{}
	if len(btRow[util.BtFamily]) > 0 {
		out.Payload = string(btRow[util.BtFamily][0].Value)
	}
	return &out, nil
}
