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
	"io/ioutil"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/proto"
	"github.com/datacommonsorg/mixer/server"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetPlaceStatDateWithinPlace(t *testing.T) {
	ctx := context.Background()

	memcacheData, err := loadMemcache()
	if err != nil {
		t.Fatalf("Failed to load memcache %v", err)
	}

	client, err := setup(server.NewMemcache(memcacheData))
	if err != nil {
		t.Fatalf("Failed to set up mixer and client")
	}
	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(
		path.Dir(filename), "../golden_response/staging/get_place_stat_date_within_place")

	for _, c := range []struct {
		ancestorPlace string
		placeType     string
		statVars      []string
		goldenFile    string
	}{
		{
			"geoId/06",
			"County",
			[]string{"Count_Person", "Median_Age_Person"},
			"CA_County.json",
		},
		{
			"country/USA",
			"State",
			[]string{"Count_Person", "CumulativeCount_MedicalConditionIncident_COVID_19_ConfirmedOrProbableCase"},
			"USA_State.json",
		},
	} {
		resp, err := client.GetPlaceStatDateWithinPlace(ctx, &pb.GetPlaceStatDateWithinPlaceRequest{
			AncestorPlace: c.ancestorPlace,
			PlaceType:     c.placeType,
			StatVars:      c.statVars,
		})
		if err != nil {
			t.Errorf("could not GetPlaceStatDateWithinPlace: %s", err)
			continue
		}
		goldenFile := path.Join(goldenPath, c.goldenFile)
		if generateGolden {
			updateGolden(resp, goldenFile)
			continue
		}
		var expected pb.GetPlaceStatDateWithinPlaceResponse
		file, _ := ioutil.ReadFile(goldenFile)
		err = protojson.Unmarshal(file, &expected)
		if err != nil {
			t.Errorf("Can not Unmarshal golden file")
			continue
		}

		if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
			t.Errorf("payload got diff: %v", diff)
			continue
		}
	}
}
