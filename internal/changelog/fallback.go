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
package changelog

import (
	"fmt"
	"time"
)

// GetFallbackChangelog returns default changelog content when embedded file is missing
func GetFallbackChangelog(version string) string {
	currentDate := time.Now().Format("2006-01-02")

	fallback := fmt.Sprintf(`# Changelog

All notable changes to this project will be documented in this file.

## [%s] - %s

### Current Version

- **Version**: %s
- **Status**: No embedded changelog available
- **Generated**: %s

### Information

This is a fallback changelog displayed when the embedded changelog is not available.
The changelog is typically embedded during the release process.

For the most up-to-date changelog, please visit:
https://github.com/KubeRocketCI/kuberocketai/releases

### Features

- Core KubeRocketAI functionality
- CLI command interface
- Framework installation and management
- IDE integrations support

### Notes

Run 'krci-ai check-updates' to check for the latest version with complete changelog.
`, version, currentDate, version, currentDate)

	return fallback
}
