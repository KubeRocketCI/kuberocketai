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
	"slices"
	"strings"
)

// DeduplicateStrings returns a new slice with unique values in stable order, sorted alphabetically
// The function also returns values sorted for deterministic output when order is not important.
func DeduplicateStrings(values []string) []string {
	slices.Sort(values)

	return slices.Compact(values)
}

// EqualFoldInSlice returns true if the candidate matches any element in slice, case-insensitively
func EqualFoldInSlice(slice []string, candidate string) bool {
	return slices.ContainsFunc(slice, func(s string) bool {
		return strings.EqualFold(s, candidate)
	})
}
