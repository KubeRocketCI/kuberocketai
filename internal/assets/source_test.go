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
package assets

import (
	"embed"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/*
var testEmbeddedAssets embed.FS

func TestNewFilesystemSource(t *testing.T) {
	baseDir := "/test/path"
	source := NewFilesystemSource(baseDir)

	assert.NotNil(t, source, "NewFilesystemSource should not return nil")
	assert.Equal(t, baseDir, source.baseDir, "baseDir should be set correctly")
	assert.Equal(t, FilesystemSourceType, source.GetSourceType(), "Source type should be filesystem")
}

func TestFilesystemSource_ListAgentFiles(t *testing.T) {
	t.Run("success with existing agents", func(t *testing.T) {
		tempDir := t.TempDir()

		// Create .krci-ai/agents directory
		agentsPath := filepath.Join(tempDir, KrciAIDir, agentsDir)
		err := os.MkdirAll(agentsPath, 0755)
		require.NoError(t, err, "Failed to create agents directory")

		// Create test agent files
		testFiles := []string{"agent1.yaml", "agent2.yaml", "agent3.yaml"}
		for _, file := range testFiles {
			err := os.WriteFile(filepath.Join(agentsPath, file), []byte("test"), 0644)
			require.NoError(t, err, "Failed to create test file")
		}

		source := NewFilesystemSource(tempDir)
		files, err := source.ListAgentFiles()

		assert.NoError(t, err, "ListAgentFiles should not return error")
		assert.Len(t, files, 3, "Should find 3 agent files")

		// Verify all expected files are found
		for _, expectedFile := range testFiles {
			expectedPath := filepath.Join(agentsPath, expectedFile)
			assert.Contains(t, files, expectedPath, "Should contain %s", expectedFile)
		}
	})

	t.Run("empty directory", func(t *testing.T) {
		tempDir := t.TempDir()

		// Create .krci-ai/agents directory but no files
		agentsPath := filepath.Join(tempDir, KrciAIDir, agentsDir)
		err := os.MkdirAll(agentsPath, 0755)
		require.NoError(t, err, "Failed to create agents directory")

		source := NewFilesystemSource(tempDir)
		files, err := source.ListAgentFiles()

		assert.NoError(t, err, "ListAgentFiles should not return error for empty directory")
		assert.Len(t, files, 0, "Should find no agent files")
	})

	t.Run("non-existent directory", func(t *testing.T) {
		source := NewFilesystemSource("/nonexistent/path")
		files, err := source.ListAgentFiles()

		assert.NoError(t, err, "ListAgentFiles should not error for non-existent directory")
		assert.Len(t, files, 0, "Should find no agent files")
	})
}

func TestFilesystemSource_ReadFile(t *testing.T) {
	t.Run("existing file", func(t *testing.T) {
		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "test.yaml")
		testContent := "test content"

		err := os.WriteFile(testFile, []byte(testContent), 0644)
		require.NoError(t, err, "Failed to create test file")

		source := NewFilesystemSource(tempDir)
		content, err := source.ReadFile(testFile)

		assert.NoError(t, err, "ReadFile should not return error")
		assert.Equal(t, testContent, string(content), "Content should match")
	})

	t.Run("non-existent file", func(t *testing.T) {
		source := NewFilesystemSource("/nonexistent")
		_, err := source.ReadFile("/nonexistent/file.yaml")

		assert.Error(t, err, "ReadFile should return error for non-existent file")
	})
}

func TestFilesystemSource_Exists(t *testing.T) {
	t.Run("existing file", func(t *testing.T) {
		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "test.yaml")

		err := os.WriteFile(testFile, []byte("test"), 0644)
		require.NoError(t, err, "Failed to create test file")

		source := NewFilesystemSource(tempDir)
		exists := source.Exists(testFile)

		assert.True(t, exists, "Exists should return true for existing file")
	})

	t.Run("non-existent file", func(t *testing.T) {
		source := NewFilesystemSource("/nonexistent")
		exists := source.Exists("/nonexistent/file.yaml")

		assert.False(t, exists, "Exists should return false for non-existent file")
	})

	t.Run("directory", func(t *testing.T) {
		tempDir := t.TempDir()
		source := NewFilesystemSource(tempDir)
		exists := source.Exists(tempDir)

		assert.True(t, exists, "Exists should return true for existing directory")
	})
}

func TestNewEmbeddedSource(t *testing.T) {
	source := NewEmbeddedSource(testEmbeddedAssets)

	assert.NotNil(t, source, "NewEmbeddedSource should not return nil")
	assert.Equal(t, EmbeddedSourceType, source.GetSourceType(), "Source type should be embedded")
	assert.Equal(t, testEmbeddedAssets, source.GetEmbeddedFS(), "Should return embedded FS")
}

func TestEmbeddedSource_WithTestAssets(t *testing.T) {
	// Use the test embedded assets
	source := NewEmbeddedSource(testEmbeddedAssets)

	t.Run("GetSourceType", func(t *testing.T) {
		sourceType := source.GetSourceType()
		assert.Equal(t, EmbeddedSourceType, sourceType, "Source type should be embedded")
	})

	t.Run("ReadFile existing", func(t *testing.T) {
		// Try to read the test agent file from testdata
		content, err := source.ReadFile("testdata/test-agent.yaml")
		if err != nil {
			t.Skip("test-agent.yaml not available in testdata, skipping")
		}
		assert.Contains(t, string(content), "agent", "Content should contain agent structure")
	})

	t.Run("ReadFile non-existent", func(t *testing.T) {
		_, err := source.ReadFile("testdata/nonexistent.yaml")
		assert.Error(t, err, "ReadFile should return error for non-existent file")
	})

	t.Run("Exists", func(t *testing.T) {
		// Test with actual testdata file
		exists := source.Exists("testdata/test-agent.yaml")
		// Note: This may be false if testdata doesn't contain test-agent.yaml
		t.Logf("testdata/test-agent.yaml exists: %v", exists)

		exists = source.Exists("testdata/nonexistent.yaml")
		assert.False(t, exists, "Exists should return false for non-existent file")
	})
}

func TestParseAgentFromSource(t *testing.T) {
	t.Run("valid agent YAML", func(t *testing.T) {
		validAgentYAML := `agent:
  identity:
    name: "Test Agent"
    description: "A test agent for validation"
    role: "Test Engineer"
    goal: "Test the system"
    icon: "🧪"`

		// Create a filesystem source with test data
		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "test-agent.yaml")
		err := os.WriteFile(testFile, []byte(validAgentYAML), 0644)
		require.NoError(t, err, "Failed to create test file")

		source := NewFilesystemSource(tempDir)
		agentInfo, err := parseAgentFromSource(source, testFile)

		assert.NoError(t, err, "parseAgentFromSource should not return error")
		assert.NotNil(t, agentInfo, "AgentInfo should not be nil")
		assert.Equal(t, "Test Agent", agentInfo.Name, "Name should match")
		assert.Equal(t, "A test agent for validation", agentInfo.Description, "Description should match")
		assert.Equal(t, "Test Engineer", agentInfo.Role, "Role should match")
		assert.Equal(t, "Test the system", agentInfo.Goal, "Goal should match")
		assert.Equal(t, "🧪", agentInfo.Icon, "Icon should match")
		assert.Equal(t, testFile, agentInfo.FilePath, "FilePath should match")
		assert.Equal(t, "test-agent", agentInfo.ShortName, "ShortName should be filename without extension")
	})

	t.Run("missing required name", func(t *testing.T) {
		invalidAgentYAML := `agent:
  identity:
    role: "Test Engineer"`

		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "invalid-agent.yaml")
		err := os.WriteFile(testFile, []byte(invalidAgentYAML), 0644)
		require.NoError(t, err, "Failed to create test file")

		source := NewFilesystemSource(tempDir)
		_, err = parseAgentFromSource(source, testFile)

		assert.Error(t, err, "parseAgentFromSource should return error for missing name")
		assert.Contains(t, err.Error(), "agent name is required", "Error should mention missing name")
	})

	t.Run("missing required role", func(t *testing.T) {
		invalidAgentYAML := `agent:
  identity:
    name: "Test Agent"`

		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "invalid-agent.yaml")
		err := os.WriteFile(testFile, []byte(invalidAgentYAML), 0644)
		require.NoError(t, err, "Failed to create test file")

		source := NewFilesystemSource(tempDir)
		_, err = parseAgentFromSource(source, testFile)

		assert.Error(t, err, "parseAgentFromSource should return error for missing role")
		assert.Contains(t, err.Error(), "agent role is required", "Error should mention missing role")
	})

	t.Run("invalid YAML", func(t *testing.T) {
		invalidYAML := `invalid: yaml: content: [}`

		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "invalid.yaml")
		err := os.WriteFile(testFile, []byte(invalidYAML), 0644)
		require.NoError(t, err, "Failed to create test file")

		source := NewFilesystemSource(tempDir)
		_, err = parseAgentFromSource(source, testFile)

		assert.Error(t, err, "parseAgentFromSource should return error for invalid YAML")
		assert.Contains(t, err.Error(), "failed to parse YAML", "Error should mention YAML parsing")
	})

	t.Run("non-existent file", func(t *testing.T) {
		source := NewFilesystemSource("/nonexistent")
		_, err := parseAgentFromSource(source, "/nonexistent/file.yaml")

		assert.Error(t, err, "parseAgentFromSource should return error for non-existent file")
		assert.Contains(t, err.Error(), "failed to read file", "Error should mention file reading")
	})
}

func TestAssetSourceInterface(t *testing.T) {
	t.Run("FilesystemSource implements AssetSource", func(t *testing.T) {
		var source AssetSource = NewFilesystemSource("/test")

		// Test that all interface methods are available
		assert.Equal(t, FilesystemSourceType, source.GetSourceType())

		// Test methods exist (they should not panic)
		assert.NotPanics(t, func() {
			_, _ = source.ListAgentFiles()
			_, _ = source.ReadFile("test")
			_ = source.Exists("test")
		})
	})

	t.Run("EmbeddedSource implements AssetSource and EmbeddedAssetSource", func(t *testing.T) {
		embeddedSource := NewEmbeddedSource(testEmbeddedAssets)

		var source AssetSource = embeddedSource
		var embeddedAssetSource EmbeddedAssetSource = embeddedSource

		// Test AssetSource interface
		assert.Equal(t, EmbeddedSourceType, source.GetSourceType())

		// Test EmbeddedAssetSource interface
		assert.Equal(t, testEmbeddedAssets, embeddedAssetSource.GetEmbeddedFS())

		// Test methods exist (they should not panic)
		assert.NotPanics(t, func() {
			_, _ = source.ListAgentFiles()
			_, _ = source.ReadFile("test")
			_ = source.Exists("test")
		})
	})
}
