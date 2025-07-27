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
	"embed"
	"fmt"
	"io/fs"
)

const (
	// ChangelogPath is the embedded path to the changelog
	ChangelogPath = "assets/changelog/CHANGELOG.md"
)

// ReadEmbeddedChangelog reads the changelog from embedded assets
func ReadEmbeddedChangelog(assets embed.FS) ([]byte, error) {
	content, err := assets.ReadFile(ChangelogPath)
	if err != nil {
		// Check if the error is due to file not found
		if _, isPathError := err.(*fs.PathError); isPathError {
			return nil, fmt.Errorf("changelog not found in embedded assets: %w", err)
		}
		return nil, fmt.Errorf("failed to read embedded changelog: %w", err)
	}

	return content, nil
}

// HasEmbeddedChangelog checks if the changelog exists in embedded assets
func HasEmbeddedChangelog(assets embed.FS) bool {
	_, err := assets.ReadFile(ChangelogPath)
	return err == nil
}

// GetChangelogPath returns the path to the embedded changelog
func GetChangelogPath() string {
	return ChangelogPath
}
