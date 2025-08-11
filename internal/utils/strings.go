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
	"sort"
	"strings"
)

// ContainsString returns true when target is present in the slice (case-sensitive)
func ContainsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

// DeduplicateStrings returns a new slice with unique values in stable order, sorted alphabetically
// The function also returns values sorted for deterministic output when order is not important.
func DeduplicateStrings(values []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(values))
	for _, v := range values {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	// Keep deterministic order by sorting alphabetically. If caller needs stable
	// input order, they should avoid relying on sort and pass already-ordered input.
	// Using strings.Compare for clarity.
	sort.Strings(result)
	return result
}

// EqualFoldInSlice returns true if the candidate matches any element in slice, case-insensitively
func EqualFoldInSlice(slice []string, candidate string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, candidate) {
			return true
		}
	}
	return false
}
