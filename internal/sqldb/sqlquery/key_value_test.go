// Copyright 2024 Google LLC
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

package sqlquery

import (
	"database/sql"
	"testing"

	pb "github.com/avsolatorio/datacommons-mixer/internal/proto"
	"github.com/go-test/deep"
)

func TestGetKeyValue(t *testing.T) {
	sqlClient, err := sql.Open("sqlite3", "../../../test/sqlquery/key_value/datacommons.db")
	if err != nil {
		t.Fatalf("Could not open test database: %v", err)
	}

	want := &pb.StatVarGroups{
		StatVarGroups: map[string]*pb.StatVarGroupNode{
			"svg1": {AbsoluteName: "SVG1"},
		},
	}

	var got pb.StatVarGroups

	found, _ := GetKeyValue(sqlClient, StatVarGroupsKey, &got)
	if !found {
		t.Errorf("Key value data not found: %s", StatVarGroupsKey)
	}
	if diff := deep.Equal(want, &got); diff != nil {
		t.Errorf("Unexpected diff %v", diff)
	}

}
