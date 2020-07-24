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
	"encoding/json"
	"fmt"

	mapset "github.com/deckarep/golang-set"

	"cloud.google.com/go/bigtable"
	"github.com/datacommonsorg/mixer/util"

	pb "github.com/datacommonsorg/mixer/proto"
)

// GetPropertyValuesPost implements API for Mixer.GetPropertyValuesPost.
func (s *Server) GetPropertyValuesPost(ctx context.Context,
	in *pb.GetPropertyValuesRequest) (*pb.GetPropertyValuesResponse, error) {
	return s.GetPropertyValues(ctx, in)
}

// GetPropertyValues implements API for Mixer.GetPropertyValues.
func (s *Server) GetPropertyValues(ctx context.Context,
	in *pb.GetPropertyValuesRequest) (*pb.GetPropertyValuesResponse, error) {
	dcids := in.GetDcids()
	prop := in.GetProperty()
	typ := in.GetValueType()
	direction := in.GetDirection()
	limit := int(in.GetLimit())
	useBranchCache := in.GetOption().GetCacheChoice() != pb.Option_BASE_CACHE_ONLY

	// Check arguments
	if prop == "" || len(dcids) == 0 {
		return nil, fmt.Errorf("missing required arguments")
	}
	if !util.CheckValidDCIDs(dcids) {
		return nil, fmt.Errorf("invalid DCIDs")
	}

	// Get in, out or both direction
	var (
		inArc  = true
		outArc = true
		inRes  = map[string][]*Node{}
		outRes = map[string][]*Node{}
	)
	var err error
	if direction == "in" {
		outArc = false
	} else if direction == "out" {
		inArc = false
	}

	if inArc {
		inRes, err = getPropertyValuesHelper(ctx, s.btTable, s.memcache, dcids,
			prop, false, useBranchCache)
		if err != nil {
			return nil, err
		}
	}
	if outArc {
		outRes, err = getPropertyValuesHelper(ctx, s.btTable, s.memcache, dcids,
			prop, true, useBranchCache)
		if err != nil {
			return nil, err
		}
	}

	result := make(map[string]map[string][]*Node)
	for _, dcid := range dcids {
		result[dcid] = map[string][]*Node{}
	}
	for dcid, nodes := range inRes {
		trimedNodes := trimNodes(nodes, typ, limit)
		if len(trimedNodes) > 0 {
			result[dcid]["in"] = trimedNodes

		}
	}
	for dcid, nodes := range outRes {
		trimedNodes := trimNodes(nodes, typ, limit)
		if len(trimedNodes) > 0 {
			result[dcid]["out"] = trimedNodes
		}
	}

	jsonRaw, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return &pb.GetPropertyValuesResponse{Payload: string(jsonRaw)}, nil
}

func getPropertyValuesHelper(
	ctx context.Context,
	btTable *bigtable.Table,
	memcache *Memcache,
	dcids []string,
	prop string,
	arcOut bool,
	useBranchCache bool,
) (map[string][]*Node, error) {
	rowList := buildPropertyValuesKey(dcids, prop, arcOut)
	nodeMap, err := readPropertyValues(ctx, btTable, rowList)
	if err != nil {
		return nil, err
	}

	// Add branch cache data
	if useBranchCache {
		branchNodeMap := memcache.ReadParallel(
			rowList,
			func(dcid string, jsonRaw []byte) (interface{}, error) {
				var propVals PropValueCache
				err := json.Unmarshal(jsonRaw, &propVals)
				if err != nil {
					return nil, err
				}
				return propVals.Nodes, nil
			}, nil)
		for dcid := range branchNodeMap {
			branchNodes := branchNodeMap[dcid].([]*Node)
			baseNodes, exist := nodeMap[dcid]
			if !exist {
				nodeMap[dcid] = branchNodes
			} else if len(branchNodes) > 0 {
				// Merge branch cache into base cache.
				itemKeys := mapset.NewSet()
				for _, n := range baseNodes {
					itemKeys.Add(n.Dcid + n.Value)
				}
				for _, n := range branchNodes {
					if itemKeys.Contains(n.Dcid + n.Value) {
						continue
					}
					nodeMap[dcid] = append(nodeMap[dcid], n)
				}
			}
		}
	}
	return nodeMap, nil
}

func trimNodes(nodes []*Node, typ string, limit int) []*Node {
	if limit == 0 && typ == "" {
		return nodes
	}
	result := []*Node{}
	for _, node := range nodes {
		if typ == "" {
			result = append(result, node)
		} else {
			for _, t := range node.Types {
				if t == typ {
					result = append(result, node)
					break
				}
			}
		}
		if limit > 0 && len(result) == limit {
			break
		}
	}
	return result
}

func readPropertyValues(
	ctx context.Context,
	btTable *bigtable.Table,
	rowList bigtable.RowList,
) (map[string][]*Node, error) {
	tmp, err := bigTableReadRowsParallel(ctx, btTable, rowList,
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			var propVals PropValueCache
			err := json.Unmarshal(jsonRaw, &propVals)
			if err != nil {
				return nil, err
			}
			return propVals.Nodes, nil
		}, nil)
	if err != nil {
		return nil, err
	}
	result := map[string][]*Node{}
	for dcid, data := range tmp {
		result[dcid] = data.([]*Node)
	}
	return result, nil
}
