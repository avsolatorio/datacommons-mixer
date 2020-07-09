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
)

func TestGetPlacesIn(t *testing.T) {
	ctx := context.Background()
	client, err := setup(server.NewMemcache(map[string][]byte{}))
	if err != nil {
		t.Fatalf("Failed to set up mixer and client")
	}
	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(
		path.Dir(filename), "../golden_response/staging/get_places_in")

	for _, c := range []struct {
		goldenFile string
		dcids      []string
		typ        string
	}{
		{
			"usa-state.json",
			[]string{"country/USA"},
			"State",
		},
		{
			"state_county.json",
			[]string{"geoId/05", "geoId/06"},
			"County",
		},
		{
			"county_zip.json",
			[]string{"geoId/06085"},
			"CensusZipCodeTabulationArea",
		},
	} {
		req := &pb.GetPlacesInRequest{
			Dcids:     c.dcids,
			PlaceType: c.typ,
		}
		resp, err := client.GetPlacesIn(ctx, req)
		if err != nil {
			t.Errorf("could not GetPlacesIn: %s", err)
			continue
		}
		var result []map[string]string
		err = json.Unmarshal([]byte(resp.GetPayload()), &result)
		if err != nil {
			t.Errorf("Can not Unmarshal payload")
			continue
		}
		var expected []map[string]string
		file, _ := ioutil.ReadFile(path.Join(goldenPath, c.goldenFile))
		err = json.Unmarshal(file, &expected)
		if err != nil {
			t.Errorf("Can not Unmarshal golden file %s: %v", c.goldenFile, err)
			continue
		}
		if diff := cmp.Diff(result, expected); diff != "" {
			t.Errorf("payload got diff: %v", diff)
			continue
		}
	}
}
