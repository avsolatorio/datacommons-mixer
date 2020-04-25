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

package store

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/bigtable"

	"encoding/json"

	mapset "github.com/deckarep/golang-set"

	pb "github.com/datacommonsorg/mixer/proto"
	"github.com/datacommonsorg/mixer/translator"
	"github.com/datacommonsorg/mixer/util"

	"golang.org/x/sync/errgroup"

	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	obsAncestorTypeObservedNode = "0"
	obsAncestorTypeComparedNode = "1"
)

// Triple represents a triples entry in the BT triples cache.
type Triple struct {
	SubjectID    string   `json:"subjectId,omitempty"`
	SubjectName  string   `json:"subjectName,omitempty"`
	SubjectTypes []string `json:"subjectTypes,omitempty"`
	Predicate    string   `json:"predicate,omitempty"`
	ObjectID     string   `json:"objectId,omitempty"`
	ObjectName   string   `json:"objectName,omitempty"`
	ObjectValue  string   `json:"objectValue,omitempty"`
	ObjectTypes  []string `json:"objectTypes,omitempty"`
	ProvenanceID string   `json:"provenanceId,omitempty"`
}

// Node represents a information about a node.
type Node struct {
	Dcid   string   `json:"dcid,omitempty"`
	Name   string   `json:"name,omitempty"`
	ProvID string   `json:"provenanceId,omitempty"`
	Value  string   `json:"value,omitempty"`
	Types  []string `json:"types,omitempty"`
}

// TriplesCache represents the json structure returned by the BT triples cache
type TriplesCache struct {
	Triples []*Triple `json:"triples"`
}

// PropValueCache represents the json structure returned by the BT PropVal cache
type PropValueCache struct {
	Nodes []*Node `json:"entities,omitempty"`
}

// PropLabelCache represents the json structure returned by the BT Prop cache
type PropLabelCache struct {
	InLabels  []string `json:"inLabels"`
	OutLabels []string `json:"outLabels"`
}

func (s *store) GetPropertyValues(ctx context.Context,
	in *pb.GetPropertyValuesRequest, out *pb.GetPropertyValuesResponse) error {
	if in.GetLimit() == 0 {
		in.Limit = util.BtCacheLimit
	}

	direction := in.GetDirection()
	var inArc, outArc bool
	switch direction {
	case "in":
		inArc = true
		outArc = false
	case "out":
		inArc = false
		outArc = true
	default:
		inArc = true
		outArc = true
	}

	var err error
	var inRes, outRes map[string][]*Node
	if in.GetLimit() > util.BtCacheLimit {
		if inArc {
			inRes, err = s.bqGetPropertyValues(ctx, in, false)
			if err != nil {
				return err
			}
		}
		if outArc {
			outRes, err = s.bqGetPropertyValues(ctx, in, true)
			if err != nil {
				return err
			}
		}
	} else {
		if inArc {
			inRes, err = s.btGetPropertyValues(ctx, in, false)
			if err != nil {
				return err
			}
		}
		if outArc {
			outRes, err = s.btGetPropertyValues(ctx, in, true)
			if err != nil {
				return err
			}
		}
	}

	nodeRes := make(map[string]map[string][]*Node)
	for _, dcid := range in.GetDcids() {
		nodeRes[dcid] = map[string][]*Node{}
	}
	for dcid, data := range inRes {
		nodeRes[dcid]["in"] = data
	}
	for dcid, data := range outRes {
		nodeRes[dcid]["out"] = data
	}

	jsonRaw, err := json.Marshal(nodeRes)
	if err != nil {
		return err
	}
	out.Payload = string(jsonRaw)
	return err
}

func (s *store) bqGetPropertyValues(ctx context.Context,
	in *pb.GetPropertyValuesRequest, arcOut bool) (map[string][]*Node, error) {
	// TODO(antaresc): Fix the ValueType not being used in the triple query
	dcids := in.GetDcids()

	// Get request parameters
	valueType := in.GetValueType()
	prop := in.GetProperty()
	limit := in.GetLimit()
	triples := []*Triple{}

	// Get triples from the triples table
	var srcIDCol, valIDCol string
	if arcOut {
		srcIDCol = "t.subject_id"
		valIDCol = "t.object_id"
	} else {
		srcIDCol = "t.object_id"
		valIDCol = "t.subject_id"
	}

	// Perform the SQL query
	// POTENTIAL BUG: If the value is not a DCID then this joins may return an
	// empty result. Hopefully, users will only specify a type when the predicate
	// does not point to a leaf node.
	var qStr string
	if valueType != "" {
		qStr = fmt.Sprintf("SELECT t.prov_id, t.subject_id, t.predicate, t.object_id, t.object_value "+
			"FROM `%s.Triple` AS t JOIN `%s.Instance` AS i ON %s = i.id "+
			"WHERE %s IN (%s) "+
			"AND t.predicate = \"%s\" "+
			"AND i.type = \"%s\" "+
			"LIMIT %d",
			s.bqDb, s.bqDb, valIDCol, srcIDCol, util.StringList(dcids), prop, valueType, limit)
	} else {
		qStr = fmt.Sprintf("SELECT t.prov_id, t.subject_id, t.predicate, t.object_id, t.object_value "+
			"FROM `%s.Triple` AS t "+
			"WHERE %s IN (%s) "+
			"AND t.predicate = \"%s\" "+
			"LIMIT %d",
			s.bqDb, srcIDCol, util.StringList(dcids), prop, limit*util.LimitFactor)
	}
	queryStrs := []string{qStr}
	tripleRes, err := queryTripleTable(ctx, s.bqClient, queryStrs)
	if err != nil {
		return nil, err
	}
	triples = append(triples, tripleRes...)

	// Get the nodeType to use for special table queries
	nodeType := in.GetValueType()
	if nodeType == "" {
		nodeType, err = getNodeType(ctx, s.bqClient, s.bqDb, dcids[0])
		if err != nil {
			return nil, err
		}
	}

	// Get out in/out arc triples depending on the direction given.
	if arcOut {
		outArcInfo, err := s.bqGetOutArcInfo(nodeType)
		if err != nil {
			return nil, err
		}
		tripleRes, err = getOutArcFromSpecialTable(ctx, s.bqClient, s.bqDb, dcids, outArcInfo, prop)
		if err != nil {
			return nil, err
		}
		triples = append(triples, tripleRes...)
	} else {
		inArcInfo, err := s.bqGetInArcInfo(nodeType)
		if err != nil {
			return nil, err
		}
		tripleRes, err = getInArcFromSpecialTable(ctx, s.bqClient, s.bqDb, dcids, inArcInfo, prop)
		if err != nil {
			return nil, err
		}
		triples = append(triples, tripleRes...)
	}

	// Get node info for all reference dcids. First, build a list of dcids to
	// getNodeInfo for
	nodeIds := make([]string, 0)
	for _, t := range triples {
		if arcOut && t.ObjectID != "" {
			nodeIds = append(nodeIds, t.ObjectID)
		} else if !arcOut {
			nodeIds = append(nodeIds, t.SubjectID)
		}
	}
	nodeInfo, err := getNodeInfo(ctx, s.bqClient, s.bqDb, nodeIds)
	if err != nil {
		return nil, err
	}

	// Populate nodes from the final list of triples. First initialize the map
	// with the given parameters
	nodeRes := make(map[string][]*Node, 0)
	for _, dcid := range dcids {
		nodeRes[dcid] = []*Node{}
	}
	// Then copy over all triples
	for _, t := range triples {
		// Get the node's contents
		var srcID string
		var node Node
		if arcOut && t.ObjectValue != "" {
			srcID = t.SubjectID
			node = Node{
				ProvID: t.ProvenanceID,
				Value:  t.ObjectValue,
			}
		} else {
			var valID string
			if arcOut {
				srcID = t.SubjectID
				valID = t.ObjectID
			} else {
				srcID = t.ObjectID
				valID = t.SubjectID
			}
			if valNode, valOk := nodeInfo[valID]; valOk {
				node = *valNode
			} else {
				node = Node{
					Dcid:   valID,
					ProvID: t.ProvenanceID,
				}
			}
		}
		// Map the node accordingly
		nodeRes[srcID] = append(nodeRes[srcID], &node)
	}

	return nodeRes, nil
}

func getPropertyValue(
	valType string,
	cacheString string,
	limit int,
) ([]*Node, error) {
	btJSONRaw, err := util.UnzipAndDecode(cacheString)
	if err != nil {
		return nil, err
	}
	result := []*Node{}
	// Parse the JSON and filter the nodes by type and apply limit.
	var btPropVals PropValueCache
	json.Unmarshal(btJSONRaw, &btPropVals)
	// Filter nodes if value type is specified.
	if valType != "" {
		for _, node := range btPropVals.Nodes {
			for _, nType := range node.Types {
				if nType == valType {
					result = append(result, node)
					break
				}
			}
		}
	} else {
		result = btPropVals.Nodes
	}
	// Limit the number of nodes.
	if len(result) > limit {
		result = result[:limit]
	}
	return result, nil
}

func (s *store) btGetPropertyValues(ctx context.Context,
	in *pb.GetPropertyValuesRequest, arcOut bool) (map[string][]*Node, error) {
	dcids := in.GetDcids()
	prop := in.GetProperty()
	valType := in.GetValueType()
	limit := int(in.GetLimit())

	keyPrefix := map[bool]string{
		true:  util.BtOutPropValPrefix,
		false: util.BtInPropValPrefix,
	}

	rowList := bigtable.RowList{}
	for _, dcid := range dcids {
		rowKey := fmt.Sprintf("%s%s^%s", keyPrefix[arcOut], dcid, prop)
		rowList = append(rowList, rowKey)
	}

	nodeRes := map[string][]*Node{}
	if err := bigTableReadRowsParallel(ctx, s.btTable, rowList,
		func(btRow bigtable.Row) error {
			// Extract DCID from row key.
			rowKey := btRow.Key()
			parts := strings.Split(rowKey, "^")
			dcid := strings.TrimPrefix(parts[0], keyPrefix[arcOut])
			btResult := btRow[util.BtFamily][0]
			nodes, err := getPropertyValue(
				valType,
				string(btResult.Value),
				limit,
			)
			nodeRes[dcid] = nodes
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
		return nil, err
	}

	// Add branch cache data
	if in.GetOption().GetCacheChoice() != pb.Option_BASE_CACHE_ONLY {
		errs, _ := errgroup.WithContext(ctx)
		for _, rowKey := range rowList {
			rowKey := rowKey
			errs.Go(func() error {
				if branchString, ok := s.cache.Read(rowKey); ok {
					parts := strings.Split(rowKey, "^")
					dcid := strings.TrimPrefix(parts[0], keyPrefix[arcOut])
					nodes, err := getPropertyValue(
						valType,
						branchString,
						limit,
					)
					baseNodes, exist := nodeRes[dcid]
					if !exist {
						nodeRes[dcid] = nodes
					} else if len(nodes) > 0 {
						// Merge branch cache into base cache.
						itemKeys := mapset.NewSet()
						for _, n := range baseNodes {
							itemKeys.Add(n.Dcid + n.Value)
						}
						for _, n := range nodes {
							if itemKeys.Contains(n.Dcid + n.Value) {
								continue
							}
							nodeRes[dcid] = append(nodeRes[dcid], n)
						}
					}
					if err != nil {
						return err
					}
				}
				return nil
			})
		}
		// Wait for completion and return the first error (if any)
		err := errs.Wait()
		if err != nil {
			return nil, err
		}
	}
	return nodeRes, nil
}

func (s *store) GetTriples(ctx context.Context,
	in *pb.GetTriplesRequest, out *pb.GetTriplesResponse) error {
	return s.btGetTriples(ctx, in, out)
}

// TODO(*): Deprecate this method once the BT route is stable.
func (s *store) bqGetTriples(
	ctx context.Context, in *pb.GetTriplesRequest, out *pb.GetTriplesResponse) error {
	// Get parameters from the request and create the triples channel
	dcids := in.GetDcids()
	limit := in.GetLimit()
	resultsChan := make(chan map[string][]*Triple, len(dcids))

	// Perform the BigQuery queries for each dcid asynchronously.
	// TODO(antaresc): replace all instances of ctx.
	errs, errCtx := errgroup.WithContext(ctx)
	for _, dcid := range dcids {
		// Asynchronously handle each dcid
		dcid := dcid
		errs.Go(func() error {
			// Maintain a list of triples to return
			resTrips := make([]*Triple, 0)

			// First get the node type associated with this dcid.
			nodeType, err := getNodeType(errCtx, s.bqClient, s.bqDb, dcid)
			if err != nil {
				return err
			}
			resTrips = append(resTrips, &Triple{
				SubjectID:   dcid,
				Predicate:   "typeOf",
				ObjectValue: nodeType,
			})

			// Get all the in/out arc nodes/values from the Triples table.
			qStrs := []string{
				// Get all the out arc nodes/values from Triples table.
				fmt.Sprintf(
					"SELECT t.prov_id, t.subject_id, t.predicate, t.object_id, t.object_value  "+
						"FROM `%s.Triple` AS t "+
						"WHERE t.subject_id = \"%s\"", s.bqDb, dcid),
				// Get all the in arc nodes from Triples table.
				fmt.Sprintf(
					"SELECT t.prov_id, t.subject_id, t.predicate, t.object_id, t.object_value  "+
						"FROM `%s.Triple` AS t "+
						"WHERE t.predicate != \"typeOf\" AND t.object_id = \"%s\" LIMIT %d",
					s.bqDb, dcid, limit),
			}
			// Send the query and append new triples
			tt, err := queryTripleTable(errCtx, s.bqClient, qStrs)
			if err != nil {
				return err
			}
			resTrips = append(resTrips, tt...)

			// Get all the out arc nodes/values from special tables (Place, etc.).
			outArcInfo, err := s.bqGetOutArcInfo(nodeType)
			if err != nil {
				return err
			}
			outTriples, err := getOutArcFromSpecialTable(errCtx, s.bqClient, s.bqDb, []string{dcid}, outArcInfo)
			if err != nil {
				return err
			}
			resTrips = append(resTrips, outTriples...)

			// Get all the in arc nodes from Special tables (Place etc...).
			inArcInfo, err := s.bqGetInArcInfo(nodeType)
			if err != nil {
				return err
			}
			inTriples, err := getInArcFromSpecialTable(errCtx, s.bqClient, s.bqDb, []string{dcid}, inArcInfo)
			if err != nil {
				return err
			}
			resTrips = append(resTrips, inTriples...)

			// Get typeOf in arc nodes from Instance table.
			if nodeType == "Class" {
				instanceTriples, err := getInstances(errCtx, s.bqClient, s.bqDb, dcid, limit)
				if err != nil {
					return err
				}
				resTrips = append(resTrips, instanceTriples...)
			}

			// Populate all other fields in the triple (i.e. name, type, etc.)
			baseNodeInfo, err := getNodeInfo(errCtx, s.bqClient, s.bqDb, []string{dcid})
			if err != nil {
				return err
			}
			name := baseNodeInfo[dcid].Name
			types := baseNodeInfo[dcid].Types

			allDcids := []string{}
			for _, t := range resTrips {
				if t.SubjectID == dcid {
					if t.ObjectID != "" {
						allDcids = append(allDcids, t.ObjectID)
					}
				} else {
					allDcids = append(allDcids, t.SubjectID)
				}
			}
			nodeInfo, err := getNodeInfo(errCtx, s.bqClient, s.bqDb, allDcids)
			if err != nil {
				return err
			}
			for _, t := range resTrips {
				if t.SubjectID == dcid {
					t.SubjectName = name
					t.SubjectTypes = types
					if t.ObjectID != "" {
						t.ObjectName = nodeInfo[t.ObjectID].Name
						t.ObjectTypes = nodeInfo[t.ObjectID].Types
					}
				} else {
					t.SubjectName = nodeInfo[t.SubjectID].Name
					t.SubjectTypes = nodeInfo[t.SubjectID].Types
					t.ObjectName = name
					t.ObjectTypes = types
				}
			}

			// Send the list of triples to the channel
			triplesMap := map[string][]*Triple{dcid: resTrips}
			resultsChan <- triplesMap

			return nil
		})
	}

	// Block on threads performing the cache read.
	err := errs.Wait()
	if err != nil {
		return err
	}

	// Copy over the contents of the results channel
	close(resultsChan)
	resultsMap := map[string][]*Triple{}
	for triplesMap := range resultsChan {
		for dcid := range triplesMap {
			// Need only copy over the dcids because an empty list of triples is
			// created for each dcid given.
			resultsMap[dcid] = triplesMap[dcid]
		}
	}

	// Format the json response and encode it in base64 as necessary.
	jsonRaw, err := json.Marshal(resultsMap)
	if err != nil {
		return err
	}
	jsonStr := string(jsonRaw)
	out.Payload = jsonStr

	return nil
}

// Returns information for inwards facing arcs towards the given node type.
func (s *store) bqGetInArcInfo(nodeType string) ([]translator.InArcInfo, error) {
	// Get parent type
	parentType, exists := s.subTypeMap[nodeType]
	if !exists {
		parentType = nodeType
	}

	// Get in arc info
	var err error
	inArcInfo, exists := s.inArcInfo[parentType]
	if !exists {
		inArcInfo, err = translator.GetInArcInfo(s.bqMapping, parentType)
		if err != nil {
			return nil, err
		}
		s.inArcInfo[parentType] = inArcInfo
	}

	return inArcInfo, nil
}

// Returns information for outwards facing arcs towards the given node type.
func (s *store) bqGetOutArcInfo(nodeType string) (map[string][]translator.OutArcInfo, error) {
	// Get parent type
	parentType, exists := s.subTypeMap[nodeType]
	if !exists {
		parentType = nodeType
	}

	// Get out arc info
	var err error
	outArcInfo, exists := s.outArcInfo[parentType]
	if !exists {
		outArcInfo, err = translator.GetOutArcInfo(s.bqMapping, parentType)
		if err != nil {
			return nil, err
		}
		s.outArcInfo[parentType] = outArcInfo
	}

	return outArcInfo, nil
}

func (s *store) GetPropertyLabels(
	ctx context.Context,
	in *pb.GetPropertyLabelsRequest,
	out *pb.GetPropertyLabelsResponse) error {
	dcids := in.GetDcids()

	rowList := bigtable.RowList{}
	for _, dcid := range dcids {
		rowList = append(rowList, fmt.Sprintf("%s%s", util.BtArcsPrefix, dcid))
	}

	resultsMap := map[string]*PropLabelCache{}
	if err := bigTableReadRowsParallel(ctx, s.btTable, rowList,
		func(btRow bigtable.Row) error {
			// Extract DCID from row key.
			dcid := strings.TrimPrefix(btRow.Key(), util.BtArcsPrefix)

			if len(btRow[util.BtFamily]) > 0 {
				btRawValue := btRow[util.BtFamily][0].Value
				btJSONRaw, err := util.UnzipAndDecode(string(btRawValue))
				if err != nil {
					return err
				}
				var btPropLabels PropLabelCache
				json.Unmarshal(btJSONRaw, &btPropLabels)
				resultsMap[dcid] = &btPropLabels

				// Fill in InLabels / OutLabels with an empty list if not present.
				if resultsMap[dcid].InLabels == nil {
					resultsMap[dcid].InLabels = []string{}
				}
				if resultsMap[dcid].OutLabels == nil {
					resultsMap[dcid].OutLabels = []string{}
				}
			}

			return nil
		}); err != nil {
		return err
	}

	// Iterate through all dcids to make sure they are present in resultsMap.
	for _, dcid := range dcids {
		if _, exists := resultsMap[dcid]; !exists {
			resultsMap[dcid] = &PropLabelCache{
				InLabels:  []string{},
				OutLabels: []string{},
			}
		}
	}

	jsonRaw, err := json.Marshal(resultsMap)
	if err != nil {
		return err
	}
	jsonStr := string(jsonRaw)
	out.Payload = jsonStr

	return nil
}

func addObsTriple(
	key string,
	btRawValue string,
	resultsMap map[string][]*Triple,
	objPlaceIDNameMap map[string]string,
) error {
	parts := strings.Split(key, "^")
	dcid := strings.TrimPrefix(parts[0], util.BtObsAncestorPrefix)
	var pred string
	if parts[1] == obsAncestorTypeObservedNode {
		pred = "observedNode"
	} else if parts[1] == obsAncestorTypeComparedNode {
		pred = "comparedNode"
	} else {
		return fmt.Errorf("unsupported predicate")
	}
	val, err := util.UnzipAndDecode(btRawValue)
	if err != nil {
		return err
	}
	objID := string(val)

	if !strings.HasPrefix(objID, "dc/p/") { // Not StatisticalPopulation.
		objPlaceIDNameMap[objID] = ""
	}

	if _, ok := resultsMap[dcid]; !ok {
		resultsMap[dcid] = []*Triple{}
	}
	for _, t := range resultsMap[dcid] {
		if t.ObjectID == objID {
			return nil
		}
	}
	resultsMap[dcid] = append(resultsMap[dcid], &Triple{
		SubjectID: dcid,
		Predicate: pred,
		ObjectID:  objID,
	})
	return nil
}

func (s *store) btGetTriples(
	ctx context.Context, in *pb.GetTriplesRequest, out *pb.GetTriplesResponse) error {
	dcids := in.GetDcids()
	var regDcids, obsDcids, popDcids []string
	for _, dcid := range dcids {
		if strings.HasPrefix(dcid, "dc/o/") {
			obsDcids = append(obsDcids, dcid)
		} else {
			regDcids = append(regDcids, dcid)
			if strings.HasPrefix(dcid, "dc/p/") {
				popDcids = append(popDcids, dcid)
			}
		}
	}

	resultsMap := map[string][]*Triple{}

	// Regular DCIDs.
	if len(regDcids) > 0 {
		rowList := bigtable.RowList{}
		for _, dcid := range regDcids {
			rowList = append(rowList, fmt.Sprintf("%s%s", util.BtTriplesPrefix, dcid))
		}
		if err := bigTableReadRowsParallel(ctx, s.btTable, rowList,
			func(btRow bigtable.Row) error {
				// Extract DCID from row key.
				dcid := strings.TrimPrefix(btRow.Key(), util.BtTriplesPrefix)

				if len(btRow[util.BtFamily]) > 0 {
					btRawValue := btRow[util.BtFamily][0].Value
					btJSONRaw, err := util.UnzipAndDecode(string(btRawValue))
					if err != nil {
						return err
					}
					var btTriples TriplesCache
					json.Unmarshal(btJSONRaw, &btTriples)
					resultsMap[dcid] = filterTriplesLimit(dcid, btTriples.Triples, int(in.GetLimit()))
				}
				return nil
			}); err != nil {
			return err
		}
	}

	// Observation DCIDs.
	objPlaceIDNameMap := map[string]string{}
	if len(obsDcids) > 0 {
		obsRowList := bigtable.RowList{}
		for _, dcid := range obsDcids {
			for _, pred := range []string{
				obsAncestorTypeObservedNode,
				obsAncestorTypeComparedNode,
			} {
				obsRowList = append(obsRowList,
					fmt.Sprintf("%s%s^%s", util.BtObsAncestorPrefix, dcid, pred))
			}
		}
		if err := bigTableReadRowsParallel(ctx, s.btTable, obsRowList,
			func(btRow bigtable.Row) error {
				if len(btRow[util.BtFamily]) > 0 {
					btRawValue := btRow[util.BtFamily][0].Value
					err := addObsTriple(
						btRow.Key(),
						string(btRawValue),
						resultsMap,
						objPlaceIDNameMap)
					if err != nil {
						return err
					}
				}
				return nil
			}); err != nil {
			return err
		}

		// If using branch cache, then check the branch cache as well.
		if in.GetOption().GetCacheChoice() != pb.Option_BASE_CACHE_ONLY {
			for _, key := range obsRowList {
				if branchString, ok := s.cache.Read(key); ok {
					err := addObsTriple(
						key,
						branchString,
						resultsMap,
						objPlaceIDNameMap)
					if err != nil {
						return err
					}
				}
			}
		}

		// Get name for places that are ancestors of obseravtions.
		objPlaceDCIDs := []string{}
		for dcid := range objPlaceIDNameMap {
			objPlaceDCIDs = append(objPlaceDCIDs, dcid)
		}
		if len(objPlaceDCIDs) > 0 {
			res, err := s.btGetPropertyValues(ctx, &pb.GetPropertyValuesRequest{
				Dcids:    objPlaceDCIDs,
				Property: "name",
				Limit:    util.BtCacheLimit,
			}, true)
			if err != nil {
				return err
			}
			for dcid, v1 := range res {
				if len(v1) == 0 {
					continue
				}
				objPlaceIDNameMap[dcid] = v1[0].Value
			}
		}
	}

	// Iterate through all dcids to make sure they are present in resultsMap.
	for _, dcid := range dcids {
		if _, exists := resultsMap[dcid]; !exists {
			resultsMap[dcid] = []*Triple{}
		}
	}

	// Add place names to observation triples as ObjectName.
	if len(obsDcids) > 0 {
		for _, dcid := range obsDcids {
			for _, t := range resultsMap[dcid] {
				if v, ok := objPlaceIDNameMap[t.ObjectID]; ok {
					t.ObjectName = v
				}
			}
		}
	}

	// Add PVs for populations.
	if len(popDcids) > 0 {
		popPVRowList := bigtable.RowList{}
		for _, dcid := range popDcids {
			popPVRowList = append(popPVRowList, fmt.Sprintf("%s%s", util.BtPopPVPrefix, dcid))
		}
		if err := bigTableReadRowsParallel(ctx, s.btTable, popPVRowList,
			func(btRow bigtable.Row) error {
				// Extract DCID from row key.
				dcid := strings.TrimPrefix(btRow.Key(), util.BtPopPVPrefix)

				if len(btRow[util.BtFamily]) > 0 {
					btRawValue := btRow[util.BtFamily][0].Value
					val, err := util.UnzipAndDecode(string(btRawValue))
					if err != nil {
						return err
					}
					parts := strings.Split(string(val), "^")
					if len(parts) == 0 || len(parts)%2 != 0 {
						return fmt.Errorf("wrong number of PVs: %v", string(val))
					}
					resultsMap[dcid] = append(resultsMap[dcid], &Triple{
						SubjectID:   dcid,
						Predicate:   "numConstraints",
						ObjectValue: strconv.Itoa(len(parts) / 2),
					})
					for i := 0; i < len(parts); i = i + 2 {
						resultsMap[dcid] = append(resultsMap[dcid], &Triple{
							SubjectID: dcid,
							Predicate: parts[i],
							ObjectID:  parts[i+1],
						})
					}
				}
				return nil
			}); err != nil {
			return err
		}
	}

	// Format the json response and encode it in base64 as necessary.
	jsonRaw, err := json.Marshal(resultsMap)
	if err != nil {
		return err
	}
	jsonStr := string(jsonRaw)
	out.Payload = jsonStr

	return nil
}

// ObsTimeSeries contains information about observation time series.
type ObsTimeSeries struct {
	Val           map[string]float64 `json:"val,omitempty"`
	Unit          string             `json:"unit,omitempty"`
	PlaceName     string             `json:"placeName,omitempty"`
	IsDcAggregate bool               `json:"isDcAggregate,omitempty"`
}

// ChartStore contains ObsTimeSeries.
// TODO(*): Add ObsCollection when needed.
type ChartStore struct {
	ObsTimeSeries *ObsTimeSeries `json:"obsTimeSeries,omitempty"`
}

func (s *store) GetChartData(ctx context.Context,
	in *pb.GetChartDataRequest, out *pb.GetChartDataResponse) error {
	rowList := bigtable.RowList{}
	for _, key := range in.GetKeys() {
		rowList = append(rowList, fmt.Sprintf("%s%s", util.BtChartDataPrefix, key))
	}

	results := map[string]*ChartStore{}
	if err := bigTableReadRowsParallel(ctx, s.btTable, rowList,
		func(btRow bigtable.Row) error {
			key := strings.TrimPrefix(btRow.Key(), util.BtChartDataPrefix)

			if len(btRow[util.BtFamily]) > 0 {
				btRawValue := btRow[util.BtFamily][0].Value
				btJSONRaw, err := util.UnzipAndDecode(string(btRawValue))
				if err != nil {
					return err
				}
				var chartStore ChartStore
				json.Unmarshal(btJSONRaw, &chartStore)
				results[key] = &chartStore
			}

			return nil
		}); err != nil {
		return err
	}

	jsonRaw, err := json.Marshal(results)
	if err != nil {
		return err
	}
	out.Payload = string(jsonRaw)

	return nil
}

// ----------------------------- HELPER FUNCTIONS -----------------------------

func filterTriplesLimit(dcid string, triples []*Triple, limit int) []*Triple {
	if triples == nil {
		return []*Triple{}
	}
	if limit == 0 { // Default limit value means no further limit.
		return triples
	}

	store := map[string][]*Triple{} // Key is {isOut + predicate + neighborType}.
	for _, t := range triples {
		isOut := "0"
		neighborTypes := t.SubjectTypes
		if t.SubjectID == dcid {
			isOut = "1"
			neighborTypes = t.ObjectTypes
		}

		for _, nt := range neighborTypes {
			key := isOut + t.Predicate + nt
			if _, ok := store[key]; !ok {
				store[key] = []*Triple{}
			}
			store[key] = append(store[key], t)
		}
	}

	triplesMap := map[*Triple]struct{}{}
	for _, triples := range store {
		filteredTriples := triples
		if len(triples) > limit {
			filteredTriples = triples[:limit]
		}
		for _, t := range filteredTriples {
			triplesMap[t] = struct{}{}
		}
	}

	result := []*Triple{}
	for t := range triplesMap {
		result = append(result, t)
	}

	return result
}

func readTripleFromBq(it *bigquery.RowIterator) ([]*Triple, error) {
	result := []*Triple{}
	for {
		t := Triple{}
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		for idx, cell := range row {
			// Columns are: prov_id, subject_id, predicate, object_id, object_value
			v, _ := cell.(string)
			switch idx {
			case 0:
				t.ProvenanceID = v
			case 1:
				t.SubjectID = v
			case 2:
				t.Predicate = v
			case 3:
				if cell != nil {
					t.ObjectID = v
				}
			case 4:
				if cell != nil {
					t.ObjectValue = v
				}
			default:
				return nil, status.Errorf(codes.InvalidArgument, "Unexpected column index %d", idx)
			}
		}
		result = append(result, &t)
	}
	return result, nil
}

func queryTripleTable(ctx context.Context, client *bigquery.Client, qStrs []string) ([]*Triple, error) {
	tripleChan := make(chan []*Triple, 2)
	errs, errCtx := errgroup.WithContext(ctx)
	for _, qStr := range qStrs {
		qStr := qStr
		errs.Go(func() error {
			log.Printf("Query: %v\n", qStr)
			q := client.Query(qStr)
			it, err := q.Read(errCtx)
			if err != nil {
				return err
			}
			triples, err := readTripleFromBq(it)
			if err != nil {
				return err
			}
			tripleChan <- triples
			return nil
		})
	}
	// Wait for completion and return the first error (if any)
	err := errs.Wait()
	if err != nil {
		return nil, err
	}
	close(tripleChan)

	result := []*Triple{}
	for triples := range tripleChan {
		result = append(result, triples...)
	}
	return result, nil
}

func getOutArcFromSpecialTable(
	ctx context.Context,
	client *bigquery.Client,
	db string,
	dcid []string,
	outArcInfo map[string][]translator.OutArcInfo,
	predicate ...string) ([]*Triple, error) {
	result := []*Triple{}
	for table, pcs := range outArcInfo {
		hasQuery := false
		qStr := "SELECT id"
		for _, pc := range pcs {
			if len(predicate) > 0 && predicate[0] != pc.Pred {
				continue
			}
			hasQuery = true
			qStr += fmt.Sprintf(", t.%s AS %s ", pc.Column, pc.Pred)
		}
		if !hasQuery {
			continue
		}
		qStr += fmt.Sprintf("FROM %s AS t WHERE id IN (%s)", table, util.StringList(dcid))
		log.Printf("Query: %v\n", qStr)
		q := client.Query(qStr)
		it, err := q.Read(ctx)
		if err != nil {
			return nil, err
		}
		for {
			var row []bigquery.Value
			err := it.Next(&row)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			var rowDcid string
			for i, cell := range row {
				v, _ := cell.(string)
				if i == 0 {
					rowDcid = v
					continue
				}
				if cell == nil {
					continue
				}
				pcIndex := i - 1
				t := Triple{SubjectID: rowDcid, Predicate: pcs[pcIndex].Pred}
				if pcs[pcIndex].IsNode {
					t.ObjectID = v
				} else {
					t.ObjectValue = v
				}
				result = append(result, &t)
			}
		}
	}
	return result, nil
}

func getInArcFromSpecialTable(
	ctx context.Context,
	client *bigquery.Client,
	db string,
	dcid []string,
	inArcInfo []translator.InArcInfo,
	predicate ...string) ([]*Triple, error) {
	result := []*Triple{}
	for _, info := range inArcInfo {
		if info.ObjCol == "place_key" || info.ObjCol == "observed_node_key" {
			continue
		}
		if len(predicate) > 0 && predicate[0] != info.Pred {
			continue
		}
		qStr := fmt.Sprintf("SELECT id, %s FROM %s WHERE %s IN (%s)", info.SubCol, info.Table, info.ObjCol, util.StringList(dcid))
		log.Printf("Query: %v\n", qStr)
		q := client.Query(qStr)
		it, err := q.Read(ctx)
		if err != nil {
			return nil, err
		}
		for {
			var row []bigquery.Value
			err := it.Next(&row)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			var rowDcid string
			for i, cell := range row {
				v, _ := cell.(string)
				if i == 0 {
					rowDcid = v
					continue
				}
				if cell == nil {
					continue
				}
				t := Triple{SubjectID: rowDcid, Predicate: info.Pred}
				t.ObjectID = v
				result = append(result, &t)
			}
		}
	}
	return result, nil
}

func getInstances(
	ctx context.Context,
	client *bigquery.Client,
	db string,
	dcid string,
	limit int32) ([]*Triple, error) {
	result := []*Triple{}
	qStr := fmt.Sprintf("SELECT id From `%s.Instance` where type = \"%s\" LIMIT %d", db, dcid, limit)
	log.Printf("Query: %v\n", qStr)
	q := client.Query(qStr)
	it, err := q.Read(ctx)
	if err != nil {
		return nil, err
	}
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		inst := row[0].(string)
		result = append(
			result,
			&Triple{
				SubjectID:   inst,
				Predicate:   "typeOf",
				ObjectValue: dcid,
			})
	}
	return result, nil
}

func getNodeType(ctx context.Context, client *bigquery.Client, db string, dcid string) (string, error) {
	qStr := fmt.Sprintf("SELECT type FROM `%s.Instance` where id = \"%s\"", db, dcid)
	log.Printf("Query: %v\n", qStr)
	q := client.Query(qStr)
	it, err := q.Read(ctx)
	if err != nil {
		return "", err
	}
	var nodeType string
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
		nodeType = row[0].(string)
	}
	return nodeType, nil
}

func getNodeInfo(ctx context.Context, client *bigquery.Client, db string, dcid []string) (map[string]*Node, error) {
	result := map[string]*Node{}

	// Return an empty map if no dcids are given
	if len(dcid) == 0 {
		return result, nil
	}

	// Perform a query to the instance table
	qStr := fmt.Sprintf("SELECT id, name, type, prov_id "+
		"FROM `%s.Instance` "+
		"WHERE id IN (%s)",
		db, util.StringList(dcid))
	log.Printf("Query: %v\n", qStr)
	q := client.Query(qStr)
	it, err := q.Read(ctx)
	if err != nil {
		return nil, err
	}
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var id, name, ntype, provID string
		for idx, cell := range row {
			v, _ := cell.(string)
			if cell == nil {
				continue
			}
			if idx == 0 {
				id = v
			} else if idx == 1 {
				name = v
			} else if idx == 2 {
				ntype = v
			} else if idx == 3 {
				provID = v
			}
		}
		if _, ok := result[id]; !ok {
			result[id] = &Node{}
		}
		result[id].Dcid = id
		result[id].Name = name
		result[id].ProvID = provID
		result[id].Types = append(result[id].Types, ntype)
	}
	return result, nil
}
