// Copyright 2021 Google LLC
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

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/server/api/translator"
)

func (s *Server) Translate(
	ctx context.Context,
	in *pb.TranslateRequest,
) (*pb.TranslateResponse, error) {
	return translator.Translate(ctx, in, s.metadata)
}

func (s *Server) Query(
	ctx context.Context,
	in *pb.QueryRequest,
) (*pb.QueryResponse, error) {
	return translator.Query(ctx, in, s.metadata, s.store)
}
