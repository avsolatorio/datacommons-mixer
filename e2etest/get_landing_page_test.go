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

// TestGetLandingPage tests GetLandingPage.
func TestGetLandingPage(t *testing.T) {
	ctx := context.Background()
	client, err := setup(server.NewMemcache(map[string][]byte{}))
	if err != nil {
		t.Fatalf("Failed to set up mixer and client")
	}
	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(
		path.Dir(filename), "../golden_response/staging/get_landing_page")

	for _, c := range []struct {
		goldenFile   string
		dcids        []string
		statVarDcids []string
	}{
		{
			"county.json",
			[]string{"geoId/06085"},
			[]string{},
		},
		{
			"state.json",
			[]string{"geoId/06", "geoId/08"},
			[]string{"Count_Person_Male"},
		},
	} {
		req := &pb.GetLandingPageRequest{
			Dcids:        c.dcids,
			StatVarDcids: c.statVarDcids,
		}
		resp, err := client.GetLandingPage(ctx, req)
		if err != nil {
			t.Errorf("could not GetLandingPage: %s", err)
			continue
		}

		goldenFile := path.Join(goldenPath, c.goldenFile)
		if generateGolden {
			marshaller := protojson.MarshalOptions{Indent: " "}
			jsonStr := marshaller.Format(resp)
			err := ioutil.WriteFile(goldenFile, []byte(jsonStr), 0644)
			if err != nil {
				t.Errorf("could not write golden files to %s", c.goldenFile)
			}
			continue
		}

		var expected pb.GetLandingPageResponse
		file, _ := ioutil.ReadFile(goldenFile)
		err = protojson.Unmarshal(file, &expected)
		if err != nil {
			t.Errorf("Can not Unmarshal golden file %s: %v", c.goldenFile, err)
			continue
		}
		if diff := cmp.Diff(&resp, &expected, protocmp.Transform()); diff != "" {
			t.Errorf("payload got diff: %v", diff)
			continue
		}
	}
}
