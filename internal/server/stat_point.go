// Copyright 2020 Google LLC
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
	"sort"

	"cloud.google.com/go/bigtable"
	pb "github.com/datacommonsorg/mixer/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetStatValue implements API for Mixer.GetStatValue.
// Endpoint: /stat (/stat/value)
func (s *Server) GetStatValue(ctx context.Context, in *pb.GetStatValueRequest) (
	*pb.GetStatValueResponse, error) {
	place := in.GetPlace()
	statVar := in.GetStatVar()
	svobsMode := s.metadata.SvObsMode

	if place == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: place")
	}
	if statVar == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: stat_var")
	}
	date := in.GetDate()
	filterProp := &ObsProp{
		Mmethod: in.GetMeasurementMethod(),
		Operiod: in.GetObservationPeriod(),
		Unit:    in.GetUnit(),
		Sfactor: in.GetScalingFactor(),
	}

	var rowList bigtable.RowList
	var keyTokens map[string]*placeStatVar
	if svobsMode {
		rowList, keyTokens = buildStatsKeyNew([]string{place}, []string{statVar})
	} else {
		// Read triples for the statistical variable.
		triplesRowList := buildTriplesKey([]string{statVar})
		triples, err := readTriples(ctx, s.store, triplesRowList)
		if err != nil {
			return nil, err
		}
		// Get the StatisticalVariable
		if triples[statVar] == nil {
			return nil, status.Errorf(codes.NotFound,
				"No statistical variable found for %s", statVar)
		}
		statVarObject, err := triplesToStatsVar(statVar, triples[statVar])
		if err != nil {
			return nil, err
		}
		// Construct BigTable row keys.
		rowList, keyTokens = buildStatsKey(
			[]string{place},
			map[string]*StatisticalVariable{statVar: statVarObject})
	}

	var obsTimeSeries *ObsTimeSeries
	btData, err := readStats(ctx, s.store, rowList, keyTokens)
	if err != nil {
		return nil, err
	}
	obsTimeSeries = btData[place][statVar]
	if obsTimeSeries == nil {
		return nil, status.Errorf(
			codes.NotFound, "No data for %s, %s", place, statVar)
	}
	obsTimeSeries.SourceSeries = filterSeries(obsTimeSeries.SourceSeries, filterProp)
	result, err := getValueFromBestSource(obsTimeSeries, date)
	if err != nil {
		return nil, err
	}
	return &pb.GetStatValueResponse{Value: result}, nil
}

// GetStatSet implements API for Mixer.GetStatSet.
// Endpoint: /stat/set
func (s *Server) GetStatSet(ctx context.Context, in *pb.GetStatSetRequest) (
	*pb.GetStatSetResponse, error) {
	places := in.GetPlaces()
	statVars := in.GetStatVars()
	date := in.GetDate()
	svobsMode := s.metadata.SvObsMode

	if len(places) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: places")
	}
	if len(statVars) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: stat_vars")
	}

	// Initialize result with stat vars and place dcids.
	result := &pb.GetStatSetResponse{
		Data: make(map[string]*pb.PlacePointStat),
	}
	for _, statVar := range statVars {
		result.Data[statVar] = &pb.PlacePointStat{
			Stat: make(map[string]*pb.PointStat),
		}
		for _, place := range places {
			result.Data[statVar].Stat[place] = nil
		}
	}

	var rowList bigtable.RowList
	var keyTokens map[string]*placeStatVar
	if svobsMode {
		rowList, keyTokens = buildStatsKeyNew(places, statVars)
	} else {
		// TODO(shifucun): Merge this with the logic in GetStatAll()
		// Read triples for statistical variable.
		triplesRowList := buildTriplesKey(statVars)
		triples, err := readTriples(ctx, s.store, triplesRowList)
		if err != nil {
			return nil, err
		}
		statVarObject := map[string]*StatisticalVariable{}
		for statVar, triplesCache := range triples {
			if triplesCache != nil {
				statVarObject[statVar], err = triplesToStatsVar(statVar, triplesCache)
				if err != nil {
					return nil, err
				}
			}
		}
		// Construct BigTable row keys.
		rowList, keyTokens = buildStatsKey(places, statVarObject)
	}
	cacheData, err := readStatsPb(ctx, s.store, rowList, keyTokens)
	if err != nil {
		return nil, err
	}
	for place, placeData := range cacheData {
		for statVar, data := range placeData {
			if data != nil {
				result.Data[statVar].Stat[place], err = getValueFromBestSourcePb(data, date)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return result, nil
}

// GetStatCollection implements API for Mixer.GetStatCollection.
// Endpoint: /stat/set/within-place
// Endpoint: /stat/collection
func (s *Server) GetStatCollection(
	ctx context.Context, in *pb.GetStatCollectionRequest) (
	*pb.GetStatCollectionResponse, error) {
	parentPlace := in.GetParentPlace()
	statVars := in.GetStatVars()
	childType := in.GetChildType()
	date := in.GetDate()
	svobsMode := s.metadata.SvObsMode

	if parentPlace == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: parent_place")
	}
	if len(statVars) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: stat_vars")
	}
	if date == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: date")
	}
	if childType == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: child_type")
	}

	// Initialize result.
	result := &pb.GetStatCollectionResponse{
		Data: make(map[string]*pb.SourceSeries),
	}
	// Initialize with nil to help check if data is in mem-cache. The nil field
	// will be populated with empty pb.ObsCollection struct in the end.
	for _, sv := range statVars {
		result.Data[sv] = nil
	}

	var rowList bigtable.RowList
	var keyTokens map[string]string
	if svobsMode {
		rowList, keyTokens = buildStatCollectionKeyNew(parentPlace, childType, date, statVars)
	} else {
		// Read triples for statistical variable.
		triplesRowList := buildTriplesKey(statVars)
		triples, err := readTriples(ctx, s.store, triplesRowList)
		if err != nil {
			return nil, err
		}
		statVarObject := map[string]*StatisticalVariable{}
		for statVar, triplesCache := range triples {
			if triplesCache != nil {
				statVarObject[statVar], err = triplesToStatsVar(statVar, triplesCache)
				if err != nil {
					return nil, err
				}
			}
		}
		// Construct BigTable row keys.
		rowList, keyTokens = buildStatCollectionKey(
			parentPlace, childType, date, statVarObject)
	}
	cacheData, err := readStatCollection(ctx, s.store, rowList, keyTokens)
	if err != nil {
		return nil, err
	}
	for sv, data := range cacheData {
		if data != nil {
			cohorts := data.SourceCohorts
			sort.Sort(SeriesByRank(cohorts))
			result.Data[sv] = cohorts[0]
		}
	}
	for sv := range result.Data {
		if result.Data[sv] == nil {
			result.Data[sv] = &pb.SourceSeries{}
		}
	}
	return result, nil
}
