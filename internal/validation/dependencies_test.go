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
	"testing"

	"github.com/KubeRocketCI/kuberocketai/internal/metadata"
)

func TestSimpleDependencyValidator_ValidateTaskDependencies(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create test directories
	templatesDir := filepath.Join(tmpDir, "templates")
	dataDir := filepath.Join(tmpDir, "data")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("Failed to create templates directory: %v", err)
	}
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}

	// Create test files
	storyTemplate := filepath.Join(templatesDir, "story.md")
	if err := os.WriteFile(storyTemplate, []byte("# Story Template"), 0644); err != nil {
		t.Fatalf("Failed to create story template: %v", err)
	}

	frameworkData := filepath.Join(dataDir, "sdlc-framework.md")
	if err := os.WriteFile(frameworkData, []byte("# SDLC Framework"), 0644); err != nil {
		t.Fatalf("Failed to create framework data: %v", err)
	}

	validator := NewSimpleDependencyValidator(tmpDir)

	tests := []struct {
		name         string
		taskFile     *metadata.ParsedTaskFile
		expectErrors int
	}{
		{
			name: "valid dependencies",
			taskFile: &metadata.ParsedTaskFile{
				FilePath: "test-task.md",
				Metadata: &metadata.TaskMetadata{
					Dependencies: &metadata.TaskDependencies{
						Templates: []string{"story.md"},
						Data:      []string{"sdlc-framework.md"},
					},
				},
			},
			expectErrors: 0,
		},
		{
			name: "missing template dependency",
			taskFile: &metadata.ParsedTaskFile{
				FilePath: "test-task.md",
				Metadata: &metadata.TaskMetadata{
					Dependencies: &metadata.TaskDependencies{
						Templates: []string{"missing-template.md"},
						Data:      []string{"sdlc-framework.md"},
					},
				},
			},
			expectErrors: 1,
		},
		{
			name: "missing data dependency",
			taskFile: &metadata.ParsedTaskFile{
				FilePath: "test-task.md",
				Metadata: &metadata.TaskMetadata{
					Dependencies: &metadata.TaskDependencies{
						Templates: []string{"story.md"},
						Data:      []string{"missing-data.md"},
					},
				},
			},
			expectErrors: 1,
		},
		{
			name: "no dependencies",
			taskFile: &metadata.ParsedTaskFile{
				FilePath: "test-task.md",
				Metadata: &metadata.TaskMetadata{
					Dependencies: nil,
				},
			},
			expectErrors: 0,
		},
		{
			name: "no metadata",
			taskFile: &metadata.ParsedTaskFile{
				FilePath: "test-task.md",
				Metadata: nil,
			},
			expectErrors: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := validator.ValidateTaskDependencies(tt.taskFile)

			if len(results) != tt.expectErrors {
				t.Errorf("Expected %d errors, got %d", tt.expectErrors, len(results))
				for _, result := range results {
					t.Logf("Error: %s", result.Message)
				}
			}

			// Validate error properties
			for _, result := range results {
				if result.Code != ErrorCodeMissingDependency {
					t.Errorf("Expected error code %s, got %s", ErrorCodeMissingDependency, result.Code)
				}
				if result.Severity != SeverityError {
					t.Errorf("Expected severity %s, got %s", SeverityError, result.Severity)
				}
			}
		})
	}
}

func TestSimpleDependencyValidator_ValidateDependencies(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create test directories
	templatesDir := filepath.Join(tmpDir, "templates")
	dataDir := filepath.Join(tmpDir, "data")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("Failed to create templates directory: %v", err)
	}
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}

	// Create test files
	storyTemplate := filepath.Join(templatesDir, "story.md")
	if err := os.WriteFile(storyTemplate, []byte("# Story Template"), 0644); err != nil {
		t.Fatalf("Failed to create story template: %v", err)
	}

	validator := NewSimpleDependencyValidator(tmpDir)

	tests := []struct {
		name          string
		dependencies  map[string][]string
		expectMissing int
	}{
		{
			name: "valid dependencies",
			dependencies: map[string][]string{
				"templates": {"story.md"},
			},
			expectMissing: 0,
		},
		{
			name: "missing template",
			dependencies: map[string][]string{
				"templates": {"story.md", "missing.md"},
			},
			expectMissing: 1,
		},
		{
			name:          "empty dependencies",
			dependencies:  map[string][]string{},
			expectMissing: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			missing, err := validator.ValidateDependencies(tt.dependencies)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(missing) != tt.expectMissing {
				t.Errorf("Expected %d missing dependencies, got %d: %v", tt.expectMissing, len(missing), missing)
			}
		})
	}
}

func TestSimpleDependencyValidator_ResolvePath(t *testing.T) {
	tmpDir := t.TempDir()
	validator := NewSimpleDependencyValidator(tmpDir)

	tests := []struct {
		name         string
		relativePath string
		expectedPath string
	}{
		{
			name:         "simple path",
			relativePath: "templates/story.md",
			expectedPath: filepath.Join(tmpDir, "templates/story.md"),
		},
		{
			name:         "nested path",
			relativePath: "data/common/framework.md",
			expectedPath: filepath.Join(tmpDir, "data/common/framework.md"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolved, err := validator.ResolvePath(tt.relativePath)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if resolved != tt.expectedPath {
				t.Errorf("Expected path %s, got %s", tt.expectedPath, resolved)
			}
		})
	}
}

func TestSimpleDependencyValidator_validateAsset(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test directories
	templatesDir := filepath.Join(tmpDir, "templates")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		t.Fatalf("Failed to create templates directory: %v", err)
	}

	// Create test file
	storyTemplate := filepath.Join(templatesDir, "story.md")
	if err := os.WriteFile(storyTemplate, []byte("# Story Template"), 0644); err != nil {
		t.Fatalf("Failed to create story template: %v", err)
	}

	validator := NewSimpleDependencyValidator(tmpDir)

	tests := []struct {
		name        string
		assetType   string
		assetPath   string
		expectError bool
	}{
		{
			name:        "existing asset",
			assetType:   "templates",
			assetPath:   "story.md",
			expectError: false,
		},
		{
			name:        "missing asset",
			assetType:   "templates",
			assetPath:   "missing.md",
			expectError: true,
		},
		{
			name:        "missing directory",
			assetType:   "nonexistent",
			assetPath:   "file.md",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.validateAsset(tt.assetType, tt.assetPath)
			hasError := err != nil

			if hasError != tt.expectError {
				t.Errorf("Expected error: %v, got error: %v (%v)", tt.expectError, hasError, err)
			}
		})
	}
}
