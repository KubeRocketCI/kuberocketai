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

// MapSliceToSet applies a mapping function to each element of a slice and returns a set (map with empty struct values).
// This is useful for converting slices to sets while applying a transformation to each element.
func MapSliceToSet[T comparable](slice []T, mapf func(T) T) map[T]struct{} {
	m := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		m[mapf(v)] = struct{}{}
	}

	return m
}
