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
package validation

import (
	"os"
	"path/filepath"
)

// isFileModified checks if a file has been modified since the last cache
func (a *FrameworkAnalyzer) isFileModified(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return true // File doesn't exist or can't be accessed, treat as modified
	}

	lastModTime := info.ModTime()
	cachedTime, exists := a.cache[filePath]

	if !exists || lastModTime.After(cachedTime) {
		a.cache[filePath] = lastModTime
		return true
	}

	return false
}

// areAnyFilesModified checks if any files in the framework have been modified
func (a *FrameworkAnalyzer) areAnyFilesModified(frameworkDir string) bool {
	dirs := []string{"agents", "tasks", "templates", "data"}

	for _, dir := range dirs {
		dirPath := filepath.Join(frameworkDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			continue
		}

		// Check all files in this directory
		if a.checkDirectoryForModifications(dirPath) {
			return true
		}
	}

	return false
}

// checkDirectoryForModifications checks if any files in a directory were modified
func (a *FrameworkAnalyzer) checkDirectoryForModifications(dirPath string) bool {
	// Check all yaml/yml files
	yamlFiles, _ := filepath.Glob(filepath.Join(dirPath, "*.y*ml"))
	for _, file := range yamlFiles {
		if a.isFileModified(file) {
			return true
		}
	}

	// Check all markdown files
	mdFiles, _ := filepath.Glob(filepath.Join(dirPath, "*.md"))
	for _, file := range mdFiles {
		if a.isFileModified(file) {
			return true
		}
	}

	// Check other files in data directory
	if filepath.Base(dirPath) == "data" {
		otherFiles, _ := filepath.Glob(filepath.Join(dirPath, "*.*"))
		for _, file := range otherFiles {
			if a.isFileModified(file) {
				return true
			}
		}
	}

	return false
}

// OptimizedAnalyzeFramework uses caching for better performance
func (a *FrameworkAnalyzer) OptimizedAnalyzeFramework() ([]ValidationIssue, *FrameworkInsights, error) {
	frameworkDir := filepath.Join(a.baseDir, ".krci-ai")
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		return nil, nil, os.ErrNotExist
	}

	// Check if framework cache is still valid
	cacheKey := "framework_validation"
	if cachedResults, exists := a.resultCache[cacheKey]; exists {
		if !a.areAnyFilesModified(frameworkDir) {
			// Cache is still valid, generate fresh insights but reuse validation results
			insights, err := a.generateFrameworkInsights(frameworkDir)
			if err != nil {
				return cachedResults, nil, nil // Return cached results even if insights fail
			}
			return cachedResults, insights, nil
		}
	}

	// Cache is invalid or doesn't exist, run full analysis
	issues, insights, err := a.AnalyzeFramework()
	if err != nil {
		return nil, nil, err
	}

	// Cache the validation results
	a.resultCache[cacheKey] = issues

	return issues, insights, nil
}
