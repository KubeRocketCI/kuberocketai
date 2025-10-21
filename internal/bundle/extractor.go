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

import (
	"fmt"
	"os"
	"path/filepath"
)

// ExtractorStats holds statistics about the extraction process
type ExtractorStats struct {
	FilesExtracted   int
	FilesOverwritten int
	DirsCreated      int
}

// Extractor handles extracting files from a bundle to the filesystem
type Extractor struct {
	projectRoot string
}

// NewExtractor creates a new bundle extractor
func NewExtractor(projectRoot string) *Extractor {
	return &Extractor{
		projectRoot: projectRoot,
	}
}

// Extract writes the extracted files to the filesystem
func (e *Extractor) Extract(files []ExtractedFile) (*ExtractorStats, error) {
	stats := &ExtractorStats{}
	createdDirs := make(map[string]bool)

	for _, file := range files {
		// Normalize the file path - handle both absolute and relative paths
		fullPath := e.normalizePath(file.Path)

		// Create parent directory if needed
		dir := filepath.Dir(fullPath)
		if !createdDirs[dir] {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, DirPermissions); err != nil {
					return stats, fmt.Errorf("failed to create directory %s: %w", dir, err)
				}
				stats.DirsCreated++
			}
			createdDirs[dir] = true
		}

		// Check if file already exists
		fileExists := false
		if _, err := os.Stat(fullPath); err == nil {
			fileExists = true
			stats.FilesOverwritten++
		}

		// Write file
		if err := os.WriteFile(fullPath, []byte(file.Content), FilePermissions); err != nil {
			return stats, fmt.Errorf("failed to write file %s: %w", fullPath, err)
		}

		if !fileExists {
			stats.FilesExtracted++
		}
	}

	return stats, nil
}

// DryRun returns information about what would be extracted without actually writing files
func (e *Extractor) DryRun(files []ExtractedFile) []string {
	var paths []string

	for _, file := range files {
		fullPath := e.normalizePath(file.Path)
		paths = append(paths, fullPath)
	}

	return paths
}

// normalizePath converts relative paths from bundles to full paths in the project
// Bundle files always contain relative paths (e.g., ".krci-ai/agents/dev.yaml")
func (e *Extractor) normalizePath(filePath string) string {
	// Join the relative path with the project root
	return filepath.Join(e.projectRoot, filePath)
}
