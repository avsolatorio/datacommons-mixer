// Copyright 2023 Google LLC
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

// Package server
package server

import (
	"context"

	v2 "github.com/datacommonsorg/mixer/internal/server/v2"
	v2observation "github.com/datacommonsorg/mixer/internal/server/v2/observation"
	v2p "github.com/datacommonsorg/mixer/internal/server/v2/properties"
	v2pv "github.com/datacommonsorg/mixer/internal/server/v2/propertyvalues"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
)

// V2Node implements API for mixer.V2Node.
func (s *Server) V2Node(
	ctx context.Context, in *pbv2.NodeRequest,
) (*pbv2.NodeResponse, error) {
	arcs, err := v2.ParseProperty(in.GetProperty())
	if err != nil {
		return nil, err
	}
	// TODO: abstract this out to a router module.
	// Simple Property Values
	// Examples:
	//   ->name
	//   <-containedInPlace
	//   ->[name, address]
	if len(arcs) == 1 {
		arc := arcs[0]
		direction := util.DirectionOut
		if !arc.Out {
			direction = util.DirectionIn
		}
		if arc.SingleProp != "" {
			if arc.Wildcard == "" {
				// Examples:
				//   ->name
				//   <-containedInPlace
				return v2pv.API(
					ctx,
					s.store,
					in.GetNodes(),
					[]string{arc.SingleProp},
					direction,
					int(in.GetLimit()),
					in.NextToken,
				)
			}

			if arc.Wildcard == "+" && !arc.Out {
				// Examples:
				//   <-containedInPlace+{typeOf:City}
				return v2pv.LinkedPropertyValues(
					ctx, s.store, in.GetNodes(), arc.SingleProp, arc.Filter)
			}
		} else { // arc.SingleProp == ""
			if len(arc.BracketProps) == 0 {
				// Examples:
				//   ->
				//   <-
				return v2p.API(ctx, s.store, in.GetNodes(), direction)
			}

			// Examples:
			//   ->[name, address]
			return v2pv.API(
				ctx,
				s.store,
				in.GetNodes(),
				arc.BracketProps,
				direction,
				int(in.GetLimit()),
				in.GetNextToken(),
			)
		}
	}
	return nil, nil
}

// V2Observation implements API for mixer.V2Observation.
func (s *Server) V2Observation(
	ctx context.Context, in *pbv2.ObservationRequest,
) (*pbv2.ObservationResponse, error) {
	// (TODO): The routing logic here is very rough. This needs more work.
	var queryDate, queryValue, queryVariable, queryEntity bool
	for _, item := range in.GetSelect() {
		if item == "date" {
			queryDate = true
		} else if item == "value" {
			queryValue = true
		} else if item == "variable" {
			queryVariable = true
		} else if item == "entitie" {
			queryEntity = true
		}
	}
	if !queryVariable || !queryEntity {
		return nil, status.Error(
			codes.InvalidArgument, "Must select 'variable' and 'entitie'")
	}
	// Observation date and value query.
	if queryDate && queryValue {
		if len(in.GetVariable().GetDcids()) > 0 && len(in.GetEntity().GetDcids()) > 0 {
			return v2observation.FetchFromSeries(
				ctx,
				s.store,
				in.GetVariable().GetDcids(),
				in.GetEntity().GetDcids(),
				in.GetDate(),
			)
		}
		if len(in.GetVariable().GetDcids()) > 0 && in.GetEntity().GetExpression() != "" {
			// Example of expression
			// "geoId/06<-containedInPlace+{typeOf: City}"
			expr := in.GetEntity().GetExpression()
			g, err := v2.ParseLinkedNodes(expr)
			if err != nil {
				return nil, err
			}
			if len(g.Arcs) != 1 {
				return nil, status.Errorf(
					codes.InvalidArgument, "invalid expression string: %s", expr)
			}
			arc := g.Arcs[0]
			if arc.SingleProp != "containedInPlace" ||
				arc.Wildcard != "+" ||
				arc.Filter == nil ||
				arc.Filter["typeOf"] == "" {
				return nil, status.Errorf(
					codes.InvalidArgument, "invalid expression string: %s", expr)
			}
			return v2observation.FetchFromCollection(
				ctx,
				s.store,
				in.GetVariable().GetDcids(),
				g.Subject,
				arc.Filter["typeOf"],
				in.GetDate(),
			)
		}
	}

	// Get existence of <variable, entity> pair.
	if !queryDate && !queryValue {
		if len(in.GetEntity().GetDcids()) > 0 {
			if len(in.GetVariable().GetDcids()) > 0 {
				// Have both entity.dcids and variable.dcids. Check existence cache.
				return v2observation.Existence(
					ctx, s.store, in.GetVariable().GetDcids(), in.GetEntity().GetDcids())
			} else {
				// TODO: Support appending entities from entity.expression
				// Only have entity.dcids, fetch variables for each entity.
				return v2observation.Variable(ctx, s.store, in.GetEntity().GetDcids())
			}
		}
	}
	return &pbv2.ObservationResponse{}, nil
}
