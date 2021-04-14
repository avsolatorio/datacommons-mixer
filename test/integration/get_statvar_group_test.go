// Copyright 2021 Google LLC
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

package integration

import (
	"context"
	"io/ioutil"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetStatVarGroup(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client, err := setupStatVar()
	if err != nil {
		t.Fatalf("Failed to set up mixer and client")
	}
	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(
		path.Dir(filename), "golden_response/staging/get_statvar_group")

	for _, c := range []struct {
		place      string
		goldenFile string
		checkCount bool
	}{
		{
			"country/GBR",
			"uk.json",
			false,
		},
		{
			"badDcid",
			"empty.json",
			false,
		},
		{
			"",
			"",
			true,
		},
	} {
		resp, err := client.GetStatVarGroup(ctx, &pb.GetStatVarGroupRequest{
			Place: c.place,
		})
		if err != nil {
			t.Errorf("could not GetStatVarGroup: %s", err)
			continue
		}
		goldenFile := path.Join(goldenPath, c.goldenFile)

		if generateGolden {
			if !c.checkCount {
				updateProtoGolden(resp, goldenFile)
			}
			continue
		}

		var expected pb.StatVarGroups
		file, _ := ioutil.ReadFile(goldenFile)
		err = protojson.Unmarshal(file, &expected)
		if err != nil {
			t.Errorf("Can not Unmarshal golden file")
			continue
		}

		if c.checkCount {
			num := len(resp.StatVarGroups)
			if num < 100 {
				t.Errorf("Too few stat var groups: %d", num)
			}
		}

		if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
			t.Errorf("payload got diff: %v", diff)
			continue
		}
	}
}
