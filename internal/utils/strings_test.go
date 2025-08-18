/*
Copyright © 2025 KubeRocketAI Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicateStrings(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		want   []string
	}{
		{
			name:   "no duplicates",
			values: []string{"a", "b", "c"},
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "with duplicates",
			values: []string{"b", "a", "b", "c", "a"},
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "empty slice",
			values: nil,
			want:   nil,
		},
		{
			name:   "single element",
			values: []string{"x"},
			want:   []string{"x"},
		},
		{
			name:   "all duplicates",
			values: []string{"same", "same", "same"},
			want:   []string{"same"},
		},
		{
			name:   "empty strings",
			values: []string{"", "a", "", "b"},
			want:   []string{"", "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeduplicateStrings(tt.values)
			assert.Equal(t, tt.want, got)
		})
	}
}
