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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBundle_Success(t *testing.T) {
	bundleContent := `# KubeRocketAI Framework Bundle

Some header content

==== FILE: .krci-ai/agents/pm.yaml ====
name: Product Manager
role: PM
==== END FILE ====

==== FILE: .krci-ai/tasks/feature-planning.md ====
# Feature Planning Task
This is a task file.
==== END FILE ====

Some footer content
`

	reader := strings.NewReader(bundleContent)
	files, err := ParseBundle(reader)

	require.NoError(t, err, "ParseBundle should not return an error")
	require.Len(t, files, 2, "Should parse exactly 2 files")

	// Check first file
	assert.Equal(t, ".krci-ai/agents/pm.yaml", files[0].Path, "First file path should match")
	expectedContent1 := "name: Product Manager\nrole: PM"
	assert.Equal(t, expectedContent1, files[0].Content, "First file content should match")

	// Check second file
	assert.Equal(t, ".krci-ai/tasks/feature-planning.md", files[1].Path, "Second file path should match")
	expectedContent2 := "# Feature Planning Task\nThis is a task file."
	assert.Equal(t, expectedContent2, files[1].Content, "Second file content should match")
}

func TestParseBundle_EmptyContent(t *testing.T) {
	bundleContent := `==== FILE: .krci-ai/agents/empty.yaml ====
==== END FILE ====`

	reader := strings.NewReader(bundleContent)
	files, err := ParseBundle(reader)

	require.NoError(t, err, "ParseBundle should not return an error")
	require.Len(t, files, 1, "Should parse exactly 1 file")
	assert.Empty(t, files[0].Content, "File content should be empty")
}

func TestParseBundle_MissingEndDelimiter(t *testing.T) {
	bundleContent := `==== FILE: .krci-ai/agents/pm.yaml ====
name: Product Manager
role: PM
`

	reader := strings.NewReader(bundleContent)
	_, err := ParseBundle(reader)

	require.Error(t, err, "Should return error for missing END FILE delimiter")
	assert.Contains(t, err.Error(), "unclosed file", "Error message should mention unclosed file")
}

func TestParseBundle_MissingStartDelimiter(t *testing.T) {
	bundleContent := `Some content
==== END FILE ====
`

	reader := strings.NewReader(bundleContent)
	_, err := ParseBundle(reader)

	require.Error(t, err, "Should return error for END FILE without matching FILE start")
	assert.Contains(t, err.Error(), "without matching FILE start", "Error message should mention missing FILE start")
}

func TestParseBundle_EmptyFilePath(t *testing.T) {
	bundleContent := `==== FILE:  ====
content
==== END FILE ====
`

	reader := strings.NewReader(bundleContent)
	_, err := ParseBundle(reader)

	require.Error(t, err, "Should return error for empty file path")
	assert.Contains(t, err.Error(), "empty file path", "Error message should mention empty file path")
}

func TestParseBundle_NestedFileStart(t *testing.T) {
	bundleContent := `==== FILE: .krci-ai/agents/pm.yaml ====
name: Product Manager
==== FILE: .krci-ai/agents/architect.yaml ====
`

	reader := strings.NewReader(bundleContent)
	_, err := ParseBundle(reader)

	require.Error(t, err, "Should return error for nested file start")
	assert.Contains(t, err.Error(), "found new file start before previous file ended", "Error message should mention nested file start")
}

func TestParseBundle_NoFiles(t *testing.T) {
	bundleContent := `# KubeRocketAI Framework Bundle
This is just header content with no files.
`

	reader := strings.NewReader(bundleContent)
	_, err := ParseBundle(reader)

	require.Error(t, err, "Should return error for bundle with no files")
	assert.Contains(t, err.Error(), "no files found", "Error message should mention no files found")
}

func TestParseBundle_PreservesWhitespace(t *testing.T) {
	bundleContent := `==== FILE: .krci-ai/tasks/test.md ====
Line 1
  Line 2 with indent
    Line 3 with more indent

Line 5 after blank line
==== END FILE ====
`

	reader := strings.NewReader(bundleContent)
	files, err := ParseBundle(reader)

	require.NoError(t, err, "ParseBundle should not return an error")
	expectedContent := "Line 1\n  Line 2 with indent\n    Line 3 with more indent\n\nLine 5 after blank line"
	assert.Equal(t, expectedContent, files[0].Content, "Content whitespace should be preserved")
}
