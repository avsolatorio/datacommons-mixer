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

package observation

import (
	"context"

	pbv2 "github.com/avsolatorio/datacommons-mixer/internal/proto/v2"
	"github.com/avsolatorio/datacommons-mixer/internal/server/cache"
	"github.com/avsolatorio/datacommons-mixer/internal/server/count"
	"github.com/avsolatorio/datacommons-mixer/internal/store"
)

// Existence implements logic to check existence for entity, variable pair.
// This function fetches data from both Bigtable and SQLite database.
func Existence(
	ctx context.Context,
	store *store.Store,
	cachedata *cache.Cache,
	variables []string,
	entities []string,
) (*pbv2.ObservationResponse, error) {
	result := &pbv2.ObservationResponse{
		ByVariable: map[string]*pbv2.VariableObservation{},
	}
	countMap, err := count.Count(ctx, store, cachedata, variables, entities)
	if err != nil {
		return nil, err
	}
	obsByVar := result.ByVariable // Short alias
	for v := range countMap {
		if _, ok := obsByVar[v]; !ok {
			obsByVar[v] = &pbv2.VariableObservation{
				ByEntity: map[string]*pbv2.EntityObservation{},
			}
		}
		for e := range countMap[v] {
			obsByVar[v].ByEntity[e] = &pbv2.EntityObservation{}

		}
	}
	return result, nil
}
