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

package e2etest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/proto"
	"github.com/datacommonsorg/mixer/server"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestQuery(t *testing.T) {
	ctx := context.Background()
	client, err := setup(server.NewMemcache(map[string][]byte{}))
	if err != nil {
		t.Fatalf("Failed to set up mixer and client")
	}
	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(
		path.Dir(filename), "../golden_response/staging/query")

	for _, c := range []struct {
		sparql     string
		goldenFile string
	}{
		{
			`BASE <http://schema.org/>
			SELECT ?MeanTemp
			WHERE {
				?o typeOf WeatherObservation .
				?o measuredProperty temperature .
				?o meanValue ?MeanTemp .
				?o observationDate "2018-01" .
				?o observedNode ?place .
				?place dcid geoId/4261000
			}
			LIMIT 10`,
			"weather1.json",
		},
		{
			`
			SELECT ?date ?mean ?unit
			WHERE {
				?o typeOf WeatherObservation .
				?o observedNode geoId/0649670 .
				?o measuredProperty temperature .
				?o observationDate ?date .
				?o unit ?unit .
				?o meanValue ?mean .
			}`,
			"weather2.json",
		},
		{
			`
			BASE <http://schema.org/>
			SELECT  ?pop ?Unemployment
			WHERE {
				?pop typeOf StatisticalPopulation .
				?o typeOf Observation .
				?pop dcid ("dc/p/qep2q2lcc3rcc" "dc/p/gmw3cn8tmsnth" "dc/p/92cxc027krdcd") .
				?o observedNode ?pop .
				?o measuredValue ?Unemployment
			}
			ORDER BY DESC(?Unemployment)
			LIMIT 10`,
			"unemployment.json",
		},
		{
			`SELECT ?a WHERE {?a typeOf RaceCodeEnum}`,
			"race_code_enum.json",
		},
	} {
		req := &pb.QueryRequest{Sparql: c.sparql}
		resp, err := client.Query(ctx, req)
		if err != nil {
			t.Errorf("could not Query: %v", err)
			continue
		}
		result := resp.Rows
		var expected []*pb.QueryResponseRow
		file, _ := ioutil.ReadFile(path.Join(goldenPath, c.goldenFile))
		if err := json.Unmarshal(file, &expected); err != nil {
			t.Errorf("Can not Unmarshal golden file %s: %v", c.goldenFile, err)
			continue
		}
		if diff := cmp.Diff(result, expected, protocmp.Transform()); diff != "" {
			t.Errorf("payload got diff: %v", diff)
			continue
		}
	}
}
