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
	"testing"
)

func BenchmarkOptimizedAnalyzeFramework_CacheHit(b *testing.B) {
	tempDir := b.TempDir()
	setupComplexFramework(b, tempDir)

	analyzer := NewFrameworkAnalyzer(tempDir)

	// Prime the cache
	_, _, err := analyzer.OptimizedAnalyzeFramework()
	if err != nil {
		b.Fatalf("Failed to prime cache: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := analyzer.OptimizedAnalyzeFramework()
		if err != nil {
			b.Fatalf("Cached analysis failed: %v", err)
		}
	}
}

func BenchmarkOptimizedAnalyzeFramework_CacheMiss(b *testing.B) {
	tempDir := b.TempDir()
	setupComplexFramework(b, tempDir)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		analyzer := NewFrameworkAnalyzer(tempDir)
		_, _, err := analyzer.OptimizedAnalyzeFramework()
		if err != nil {
			b.Fatalf("Fresh analysis failed: %v", err)
		}
	}
}

func BenchmarkAnalyzeFramework_NoCaching(b *testing.B) {
	tempDir := b.TempDir()
	setupComplexFramework(b, tempDir)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		analyzer := NewFrameworkAnalyzer(tempDir)
		_, _, err := analyzer.AnalyzeFramework()
		if err != nil {
			b.Fatalf("Non-cached analysis failed: %v", err)
		}
	}
}

func BenchmarkIsFileModified(b *testing.B) {
	tempDir := b.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")

	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		b.Fatalf("Failed to create test file: %v", err)
	}

	analyzer := NewFrameworkAnalyzer(tempDir)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		analyzer.isFileModified(testFile)
	}
}

func BenchmarkAreAnyFilesModified(b *testing.B) {
	tempDir := b.TempDir()
	setupComplexFramework(b, tempDir)

	analyzer := NewFrameworkAnalyzer(tempDir)
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Prime the cache
	analyzer.areAnyFilesModified(frameworkDir)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		analyzer.areAnyFilesModified(frameworkDir)
	}
}

// setupComplexFramework creates a complex framework structure for benchmarking
func setupComplexFramework(b *testing.B, tempDir string) {
	frameworkDir := filepath.Join(tempDir, ".krci-ai")

	// Create directories
	dirs := []string{"agents", "tasks", "templates", "data"}
	for _, dir := range dirs {
		dirPath := filepath.Join(frameworkDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			b.Fatalf("Failed to create directory %s: %v", dirPath, err)
		}
	}

	// Create multiple agent files
	agentContent := `agent:
  identity:
    name: "test-agent-%d"
    description: "Test agent %d"
  tasks:
    - ./.krci-ai/tasks/task-%d.md
    - ./.krci-ai/tasks/common-task.md
  commands:
    help: "Show help"
    execute: "Execute task"`

	for i := 1; i <= 10; i++ {
		agentFile := filepath.Join(frameworkDir, "agents", "agent-"+fmt.Sprintf("%d", i)+".yaml")
		content := fmt.Sprintf(agentContent, i, i, i)
		if err := os.WriteFile(agentFile, []byte(content), 0644); err != nil {
			b.Fatalf("Failed to create agent file: %v", err)
		}
	}

	// Create task files
	taskContent := `# Task %d

## Description
This is test task %d with detailed instructions.

## Parameters
- param1: Description of parameter 1
- param2: Description of parameter 2

## Examples
Example usage of task %d.

## Internal Links
[Related Template](./.krci-ai/templates/template-%d.md)
[Related Data](./.krci-ai/data/data-%d.json)`

	for i := 1; i <= 10; i++ {
		taskFile := filepath.Join(frameworkDir, "tasks", "task-"+fmt.Sprintf("%d", i)+".md")
		content := fmt.Sprintf(taskContent, i, i, i, i, i)
		if err := os.WriteFile(taskFile, []byte(content), 0644); err != nil {
			b.Fatalf("Failed to create task file: %v", err)
		}
	}

	// Create common task
	commonTaskFile := filepath.Join(frameworkDir, "tasks", "common-task.md")
	commonContent := "# Common Task\nThis task is used by multiple agents."
	if err := os.WriteFile(commonTaskFile, []byte(commonContent), 0644); err != nil {
		b.Fatalf("Failed to create common task file: %v", err)
	}

	// Create template files
	templateContent := `# Template %d

This is template %d for testing purposes.

## Variables
- ${variable1}: Description
- ${variable2}: Description

## Usage
Template %d usage instructions.`

	for i := 1; i <= 5; i++ {
		templateFile := filepath.Join(frameworkDir, "templates", "template-"+fmt.Sprintf("%d", i)+".md")
		content := fmt.Sprintf(templateContent, i, i, i)
		if err := os.WriteFile(templateFile, []byte(content), 0644); err != nil {
			b.Fatalf("Failed to create template file: %v", err)
		}
	}

	// Create data files
	dataContent := `{
  "id": %d,
  "name": "data-%d",
  "description": "Test data file %d",
  "values": [1, 2, 3, 4, 5],
  "metadata": {
    "type": "test",
    "version": "1.0.0"
  }
}`

	for i := 1; i <= 5; i++ {
		dataFile := filepath.Join(frameworkDir, "data", "data-"+fmt.Sprintf("%d", i)+".json")
		content := fmt.Sprintf(dataContent, i, i, i)
		if err := os.WriteFile(dataFile, []byte(content), 0644); err != nil {
			b.Fatalf("Failed to create data file: %v", err)
		}
	}
}
