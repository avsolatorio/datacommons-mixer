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

// Package v2
package v2

import "fmt"

// split a string by "->" and "<-"
func split(s string) ([]string, error) {
	if len(s) < 2 {
		return nil, fmt.Errorf("invalid query string: %s", s)
	}
	if s[0:2] != "->" && s[0:2] != "<-" {
		return nil, fmt.Errorf("query string should start with arrow, %s", s)
	}
	pos := []int{}
	for i := 0; i < len(s)-2; i++ {
		if s[i:i+2] == "->" || s[i:i+2] == "<-" {
			pos = append(pos, i)
		}
	}
	parts := []string{}
	for i := 0; i < len(pos)-1; i++ {
		parts = append(parts, s[pos[i]:pos[i+1]])
	}
	if len(pos) > 0 {
		parts = append(parts, s[pos[len(pos)-1]:])
	}
	return parts, nil
}
