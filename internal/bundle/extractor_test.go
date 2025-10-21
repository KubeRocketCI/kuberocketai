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
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractor_Extract_Success(t *testing.T) {
	// Create temp directory
	tempDir := t.TempDir()

	extractor := NewExtractor(tempDir)

	files := []ExtractedFile{
		{
			Path:    ".krci-ai/agents/pm.yaml",
			Content: "name: Product Manager\nrole: PM",
		},
		{
			Path:    ".krci-ai/tasks/feature-planning.md",
			Content: "# Feature Planning",
		},
	}

	stats, err := extractor.Extract(files)
	require.NoError(t, err, "Extract should not return an error")

	// Verify stats
	assert.Equal(t, 2, stats.FilesExtracted, "Should extract 2 files")
	assert.Equal(t, 0, stats.FilesOverwritten, "Should not overwrite any files")
	assert.Equal(t, 2, stats.DirsCreated, "Should create 2 directories")

	// Verify files exist and have correct content
	pmPath := filepath.Join(tempDir, ".krci-ai/agents/pm.yaml")
	content, err := os.ReadFile(pmPath)
	require.NoError(t, err, "Should be able to read extracted file")
	expectedContent := "name: Product Manager\nrole: PM"
	assert.Equal(t, expectedContent, string(content), "First file content should match")

	taskPath := filepath.Join(tempDir, ".krci-ai/tasks/feature-planning.md")
	content, err = os.ReadFile(taskPath)
	require.NoError(t, err, "Should be able to read extracted file")
	assert.Equal(t, "# Feature Planning", string(content), "Second file content should match")
}

func TestExtractor_Extract_Overwrite(t *testing.T) {
	// Create temp directory
	tempDir := t.TempDir()

	// Create existing file
	agentsDir := filepath.Join(tempDir, ".krci-ai/agents")
	if err := os.MkdirAll(agentsDir, DirPermissions); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	existingFile := filepath.Join(agentsDir, "pm.yaml")
	if err := os.WriteFile(existingFile, []byte("old content"), FilePermissions); err != nil {
		t.Fatalf("Failed to create existing file: %v", err)
	}

	extractor := NewExtractor(tempDir)

	files := []ExtractedFile{
		{
			Path:    ".krci-ai/agents/pm.yaml",
			Content: "new content",
		},
	}

	stats, err := extractor.Extract(files)
	require.NoError(t, err, "Extract should not return an error")

	// Verify stats - file should be counted as overwritten
	assert.Equal(t, 0, stats.FilesExtracted, "Should not extract new files")
	assert.Equal(t, 1, stats.FilesOverwritten, "Should overwrite 1 file")

	// Verify file content was updated
	content, err := os.ReadFile(existingFile)
	require.NoError(t, err, "Should be able to read file")
	assert.Equal(t, "new content", string(content), "File content should be updated")
}

func TestExtractor_Extract_NestedDirectories(t *testing.T) {
	// Create temp directory
	tempDir := t.TempDir()

	extractor := NewExtractor(tempDir)

	files := []ExtractedFile{
		{
			Path:    ".krci-ai/agents/nested/deep/pm.yaml",
			Content: "nested agent",
		},
	}

	stats, err := extractor.Extract(files)
	require.NoError(t, err, "Extract should not return an error")

	// Verify file was created
	assert.Equal(t, 1, stats.FilesExtracted, "Should extract 1 file")

	// Verify nested directories were created
	filePath := filepath.Join(tempDir, ".krci-ai/agents/nested/deep/pm.yaml")
	content, err := os.ReadFile(filePath)
	require.NoError(t, err, "Should be able to read nested file")
	assert.Equal(t, "nested agent", string(content), "Nested file content should match")
}

func TestExtractor_DryRun(t *testing.T) {
	// Create temp directory
	tempDir := t.TempDir()

	extractor := NewExtractor(tempDir)

	files := []ExtractedFile{
		{
			Path:    ".krci-ai/agents/pm.yaml",
			Content: "name: Product Manager",
		},
		{
			Path:    ".krci-ai/tasks/feature-planning.md",
			Content: "# Feature Planning",
		},
	}

	paths := extractor.DryRun(files)

	// Verify paths
	require.Len(t, paths, 2, "Should return 2 paths")

	expectedPath1 := filepath.Join(tempDir, ".krci-ai/agents/pm.yaml")
	expectedPath2 := filepath.Join(tempDir, ".krci-ai/tasks/feature-planning.md")

	assert.Equal(t, expectedPath1, paths[0], "First path should match")
	assert.Equal(t, expectedPath2, paths[1], "Second path should match")

	// Verify no files were created
	_, err := os.Stat(expectedPath1)
	assert.True(t, os.IsNotExist(err), "DryRun should not create first file")

	_, err = os.Stat(expectedPath2)
	assert.True(t, os.IsNotExist(err), "DryRun should not create second file")
}

func TestExtractor_Extract_EmptyContent(t *testing.T) {
	// Create temp directory
	tempDir := t.TempDir()

	extractor := NewExtractor(tempDir)

	files := []ExtractedFile{
		{
			Path:    ".krci-ai/agents/empty.yaml",
			Content: "",
		},
	}

	stats, err := extractor.Extract(files)
	require.NoError(t, err, "Extract should not return an error")
	assert.Equal(t, 1, stats.FilesExtracted, "Should extract 1 file")

	// Verify empty file was created
	filePath := filepath.Join(tempDir, ".krci-ai/agents/empty.yaml")
	content, err := os.ReadFile(filePath)
	require.NoError(t, err, "Should be able to read file")
	assert.Empty(t, string(content), "File content should be empty")
}

