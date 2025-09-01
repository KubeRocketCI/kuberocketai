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
	"fmt"
	"os"
	"path/filepath"

	"github.com/KubeRocketCI/kuberocketai/internal/metadata"
)

// SimpleDependencyValidator handles basic dependency validation for local assets only
type SimpleDependencyValidator struct {
	baseDir string
}

// NewSimpleDependencyValidator creates a new simple dependency validator
func NewSimpleDependencyValidator(baseDir string) *SimpleDependencyValidator {
	return &SimpleDependencyValidator{baseDir: baseDir}
}

// ValidateTaskDependencies validates all dependencies for a task file
func (v *SimpleDependencyValidator) ValidateTaskDependencies(taskFile *metadata.ParsedTaskFile) []ValidationResult {
	var results []ValidationResult

	if taskFile.Metadata == nil || taskFile.Metadata.Dependencies == nil {
		return results // No dependencies to validate
	}

	deps := taskFile.Metadata.Dependencies

	// Validate template dependencies
	for _, template := range deps.Templates {
		if err := v.validateAsset("templates", template); err != nil {
			results = append(results, ValidationResult{
				Severity:    SeverityError,
				Type:        "dependency",
				Code:        ErrorCodeMissingDependency,
				File:        taskFile.FilePath,
				Message:     fmt.Sprintf("Missing template dependency: %s", template),
				Context:     map[string]interface{}{"dependency_type": "template", "dependency_path": template},
				FixGuidance: fmt.Sprintf("Create the template file at %s", filepath.Join(v.baseDir, "templates", template)),
			})
		}
	}

	// Validate data dependencies
	for _, data := range deps.Data {
		if err := v.validateAsset("data", data); err != nil {
			results = append(results, ValidationResult{
				Severity:    SeverityError,
				Type:        "dependency",
				Code:        ErrorCodeMissingDependency,
				File:        taskFile.FilePath,
				Message:     fmt.Sprintf("Missing data dependency: %s", data),
				Context:     map[string]interface{}{"dependency_type": "data", "dependency_path": data},
				FixGuidance: fmt.Sprintf("Create the data file at %s", filepath.Join(v.baseDir, "data", data)),
			})
		}
	}

	// MCP dependencies are not file-based, so we skip validation for now
	// Could be extended in the future to validate MCP server availability

	return results
}

// validateAsset checks if a dependency asset exists in the local installation
func (v *SimpleDependencyValidator) validateAsset(assetType, assetPath string) error {
	fullPath := filepath.Join(v.baseDir, assetType, assetPath)

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("dependency not found: %s", fullPath)
	}

	return nil
}

// ValidateDependencies validates a map of dependencies (legacy interface compatibility)
func (v *SimpleDependencyValidator) ValidateDependencies(dependencies map[string][]string) ([]string, error) {
	var missingDeps []string

	// Validate templates
	if templates, ok := dependencies["templates"]; ok {
		for _, template := range templates {
			if err := v.validateAsset("templates", template); err != nil {
				missingDeps = append(missingDeps, fmt.Sprintf("templates/%s", template))
			}
		}
	}

	// Validate data
	if dataFiles, ok := dependencies["data"]; ok {
		for _, data := range dataFiles {
			if err := v.validateAsset("data", data); err != nil {
				missingDeps = append(missingDeps, fmt.Sprintf("data/%s", data))
			}
		}
	}

	return missingDeps, nil
}

// ResolvePath resolves a relative path to absolute path (simple implementation)
func (v *SimpleDependencyValidator) ResolvePath(relativePath string) (string, error) {
	// For simple validation, we assume paths are relative to baseDir
	return filepath.Join(v.baseDir, relativePath), nil
}

// SetSearchPaths is a no-op for simple validator (interface compatibility)
func (v *SimpleDependencyValidator) SetSearchPaths(paths []string) {
	// Not needed for simple local-only validation
}
