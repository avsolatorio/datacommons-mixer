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

// Package observation is for V2 observation API
package observation

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/avsolatorio/datacommons-mixer/internal/merger"
	pb "github.com/avsolatorio/datacommons-mixer/internal/proto"
	pbv2 "github.com/avsolatorio/datacommons-mixer/internal/proto/v2"
	"github.com/avsolatorio/datacommons-mixer/internal/server/ranking"
	"github.com/avsolatorio/datacommons-mixer/internal/server/resource"
	"github.com/avsolatorio/datacommons-mixer/internal/server/stat"
	"github.com/avsolatorio/datacommons-mixer/internal/server/v2/shared"
	"github.com/avsolatorio/datacommons-mixer/internal/store/bigtable"
	"github.com/avsolatorio/datacommons-mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/avsolatorio/datacommons-mixer/internal/store"
)

// Direct response are from child entities list. No need to have an entity in
// the response if it has no observation.
func trimDirectResp(resp *pbv2.ObservationResponse) *pbv2.ObservationResponse {
	result := &pbv2.ObservationResponse{
		ByVariable: map[string]*pbv2.VariableObservation{},
		Facets:     map[string]*pb.Facet{},
	}
	for variable, variableData := range resp.ByVariable {
		for entity, entityData := range variableData.ByEntity {
			if len(entityData.OrderedFacets) == 0 {
				delete(variableData.ByEntity, entity)
			}
		}
		result.ByVariable[variable] = variableData
	}
	for facet, res := range resp.Facets {
		result.Facets[facet] = res
	}
	return result
}

// FetchContainedIn fetches data for child places contained in an ancestor place.
func FetchContainedIn(
	ctx context.Context,
	store *store.Store,
	metadata *resource.Metadata,
	sqlProvenances map[string]*pb.Facet,
	httpClient *http.Client,
	remoteMixer string,
	variables []string,
	ancestor string,
	childType string,
	queryDate string,
	filter *pbv2.FacetFilter,
) (*pbv2.ObservationResponse, error) {
	// Need to use child places to for direct fetch.
	var result *pbv2.ObservationResponse
	var childPlaces []string
	var err error
	if store.BtGroup != nil {
		readCollectionCache := false
		if queryDate != "" {
			result = &pbv2.ObservationResponse{
				ByVariable: map[string]*pbv2.VariableObservation{},
				Facets:     map[string]*pb.Facet{},
			}
			readCollectionCache = util.HasCollectionCache(ancestor, childType)
			if readCollectionCache {
				var btData map[string]*pb.ObsCollection
				var err error
				btData, err = stat.ReadStatCollection(
					ctx, store.BtGroup, bigtable.BtObsCollection,
					ancestor, childType, variables, queryDate,
				)
				if err != nil {
					return nil, err
				}
				for _, variable := range variables {
					result.ByVariable[variable] = &pbv2.VariableObservation{
						ByEntity: map[string]*pbv2.EntityObservation{},
					}
					// Create a short alias
					obsByEntity := result.ByVariable[variable].ByEntity
					data, ok := btData[variable]
					if !ok || data == nil {
						continue
					}
					cohorts := data.SourceCohorts
					// Sort cohort first, so the preferred source is populated first.
					sort.Sort(ranking.CohortByRank(cohorts))
					for _, cohort := range cohorts {
						facet := util.GetFacet(cohort)
						// If there is a facet filter, check that the cohort matches the
						// filter. Otherwise, skip.
						if filter != nil && !shouldKeepSourceSeries(filter, facet) {
							continue
						}
						facetID := util.GetFacetID(facet)
						for entity, val := range cohort.Val {
							if _, ok := obsByEntity[entity]; !ok {
								obsByEntity[entity] = &pbv2.EntityObservation{}
							}
							// When date is in the request, response date is the given date.
							// Otherwise, response date is the latest date from the cache.
							respDate := queryDate
							if respDate == shared.LATEST {
								respDate = cohort.PlaceToLatestDate[entity]
							}
							// If there is higher quality facet, then do not pick from the inferior
							// facet even it could have more recent data.
							if len(obsByEntity[entity].OrderedFacets) > 0 && stat.IsInferiorFacetPb(cohort) {
								continue
							}
							obsByEntity[entity].OrderedFacets = append(
								obsByEntity[entity].OrderedFacets,
								&pbv2.FacetObservation{
									FacetId: facetID,
									Observations: []*pb.PointStat{
										{
											Date:  respDate,
											Value: proto.Float64(val),
										},
									},
								},
							)
							result.Facets[facetID] = facet
						}
					}
				}
			}
		}
		if !readCollectionCache {
			childPlaces, err = shared.FetchChildPlaces(
				ctx, store, metadata, httpClient, remoteMixer, ancestor, childType)
			if err != nil {
				return nil, err
			}
			totalSeries := len(variables) * len(childPlaces)
			if totalSeries > shared.MaxSeries {
				return nil, status.Errorf(
					codes.Internal,
					"Stop processing large number of concurrent observation series: %d",
					totalSeries,
				)
			}
			log.Println("Fetch series cache in contained-in observation query")
			directResp, err := FetchDirectBT(
				ctx,
				store.BtGroup,
				variables,
				childPlaces,
				queryDate,
				filter,
			)
			if err != nil {
				return nil, err
			}
			result = trimDirectResp(directResp)
		}
	}

	// Fetch Data from SQLite database.
	var sqlResult *pbv2.ObservationResponse
	if store.SQLClient != nil {
		if ancestor == childType {
			sqlResult = initObservationResult(variables)
			variablesStr := "'" + strings.Join(variables, "', '") + "'"
			query := fmt.Sprintf(
				`
					SELECT entity, variable, date, value, provenance FROM observations as o
					JOIN triples as t ON o.entity = t.subject_id
					AND t.predicate = 'typeOf'
					AND t.object_id = '%s'
					AND o.value != ''
					AND o.variable IN (%s)
				`,
				childType,
				variablesStr,
			)
			if queryDate != "" && queryDate != shared.LATEST {
				query += fmt.Sprintf("AND date = (%s) ", queryDate)
			}
			query += "ORDER BY date ASC;"
			rows, err := store.SQLClient.Query(query)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			tmp, err := handleSQLRows(rows, variables)
			if err != nil {
				return nil, err
			}
			sqlResult = processSqlData(sqlResult, tmp, queryDate, sqlProvenances)
		} else {
			if len(childPlaces) == 0 {
				childPlaces, err = shared.FetchChildPlaces(
					ctx, store, metadata, httpClient, remoteMixer, ancestor, childType)
				if err != nil {
					return nil, err
				}
			}
			directResp, err := FetchDirectSQL(
				ctx,
				store.SQLClient,
				sqlProvenances,
				variables,
				childPlaces,
				queryDate,
				filter,
			)
			if err != nil {
				return nil, err
			}
			sqlResult = trimDirectResp(directResp)
		}
		// Prefer SQL data over BT data, so put sqlResult first.
		result = merger.MergeObservation(sqlResult, result)
	}
	return result, nil
}
