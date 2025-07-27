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
)

// Reader handles changelog reading operations
type Reader struct {
	assets embed.FS
}

// NewReader creates a new changelog reader
func NewReader(assets embed.FS) *Reader {
	return &Reader{
		assets: assets,
	}
}

// ReadChangelog reads the changelog content from embedded assets
func (r *Reader) ReadChangelog() (string, error) {
	content, err := ReadEmbeddedChangelog(r.assets)
	if err != nil {
		// Return fallback content if embedded changelog is not available
		fallbackContent := GetFallbackChangelog("unknown")
		return fallbackContent, fmt.Errorf("using fallback changelog: %w", err)
	}

	return string(content), nil
}

// ReadChangelogBytes reads the changelog content as bytes
func (r *Reader) ReadChangelogBytes() ([]byte, error) {
	return ReadEmbeddedChangelog(r.assets)
}

// HasChangelog checks if changelog is available
func (r *Reader) HasChangelog() bool {
	return HasEmbeddedChangelog(r.assets)
}

// ReadChangelog reads changelog content from embedded assets
func ReadChangelog(assets embed.FS) (string, error) {
	reader := NewReader(assets)
	return reader.ReadChangelog()
}
