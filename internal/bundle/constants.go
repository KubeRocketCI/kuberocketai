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
package bundle

// Bundle format delimiters
const (
	// FileStartDelimiter marks the beginning of a file section in a bundle
	FileStartDelimiter = "==== FILE: "

	// FileEndDelimiter marks the end of a file section in a bundle
	FileEndDelimiter = "==== END FILE ===="
)

// File system permissions
const (
	// FilePermissions defines the default permission mode for created files
	FilePermissions = 0644

	// DirPermissions defines the default permission mode for created directories
	DirPermissions = 0755
)
