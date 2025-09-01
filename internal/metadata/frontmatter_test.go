package metadata

import (
	"os"
	"testing"
)

func TestFrontmatterParser_ParseTaskContent(t *testing.T) {
	// Test with schema from file
	parser, err := NewFrontmatterParserFromFile("../../cmd/krci-ai/assets/schemas/task-metadata.json")
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	tests := []struct {
		name          string
		content       string
		expectError   bool
		wantTemplates int
		wantData      int
	}{
		{
			name: "valid frontmatter with dependencies",
			content: `---
dependencies:
  templates:
    - story.md
    - epic.md
  data:
    - common/sdlc-framework.md
    - coding-standards.md
---

# Task Content

This is the task content.`,
			expectError:   false,
			wantTemplates: 2,
			wantData:      2,
		},
		{
			name: "no frontmatter",
			content: `# Task Content

This task has no frontmatter.`,
			expectError:   false,
			wantTemplates: 0,
			wantData:      0,
		},
		{
			name: "empty frontmatter",
			content: `---
---

# Task Content

This is the task content.`,
			expectError:   false,
			wantTemplates: 0,
			wantData:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := parser.ParseTaskContent(tt.content, "test.md")

			if tt.expectError {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if parsed.Metadata == nil {
				t.Fatal("Expected metadata, got nil")
			}

			// Check template count
			var actualTemplates int
			if parsed.Metadata.Dependencies != nil {
				actualTemplates = len(parsed.Metadata.Dependencies.Templates)
			}
			if actualTemplates != tt.wantTemplates {
				t.Errorf("Expected %d templates, got %d", tt.wantTemplates, actualTemplates)
			}

			// Check data count
			var actualData int
			if parsed.Metadata.Dependencies != nil {
				actualData = len(parsed.Metadata.Dependencies.Data)
			}
			if actualData != tt.wantData {
				t.Errorf("Expected %d data files, got %d", tt.wantData, actualData)
			}

			// Validate metadata
			if err := parser.ValidateMetadata(parsed.Metadata, "test.md"); err != nil {
				t.Errorf("Metadata validation failed: %v", err)
			}
		})
	}
}

func TestFrontmatterParser_ValidateMetadata(t *testing.T) {
	// Test with schema from file
	parser, err := NewFrontmatterParserFromFile("../../cmd/krci-ai/assets/schemas/task-metadata.json")
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	tests := []struct {
		name        string
		metadata    *TaskMetadata
		expectError bool
	}{
		{
			name: "valid metadata",
			metadata: &TaskMetadata{
				Dependencies: &TaskDependencies{
					Templates: []string{"story.md", "epic.md"},
					Data:      []string{"common/sdlc-framework.md"},
					MCP:       []string{"github-server"},
				},
			},
			expectError: false,
		},
		{
			name:        "empty metadata",
			metadata:    &TaskMetadata{},
			expectError: false,
		},
		{
			name: "invalid template extension",
			metadata: &TaskMetadata{
				Dependencies: &TaskDependencies{
					Templates: []string{"invalid.txt"}, // Should be .md
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parser.ValidateMetadata(tt.metadata, "test.md")

			if tt.expectError {
				if err == nil {
					t.Error("Expected validation error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected validation error: %v", err)
				}
			}
		})
	}
}

func TestTaskMetadata_HasDependencies(t *testing.T) {
	tests := []struct {
		name     string
		metadata *TaskMetadata
		want     bool
	}{
		{
			name: "has templates",
			metadata: &TaskMetadata{
				Dependencies: &TaskDependencies{
					Templates: []string{"story.md"},
				},
			},
			want: true,
		},
		{
			name: "has data",
			metadata: &TaskMetadata{
				Dependencies: &TaskDependencies{
					Data: []string{"sdlc-framework.md"},
				},
			},
			want: true,
		},
		{
			name: "has mcp",
			metadata: &TaskMetadata{
				Dependencies: &TaskDependencies{
					MCP: []string{"github-server"},
				},
			},
			want: true,
		},
		{
			name: "no dependencies",
			metadata: &TaskMetadata{
				Dependencies: &TaskDependencies{},
			},
			want: false,
		},
		{
			name:     "nil dependencies",
			metadata: &TaskMetadata{},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.metadata.HasDependencies(); got != tt.want {
				t.Errorf("HasDependencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestWithRealSchema tests the parser with the actual embedded schema file
func TestFrontmatterParser_WithRealSchema(t *testing.T) {
	// Check if schema file exists
	if _, err := os.Stat("../../cmd/krci-ai/assets/schemas/task-metadata.json"); os.IsNotExist(err) {
		t.Skip("Schema file not found, skipping test")
	}

	parser, err := NewFrontmatterParserFromFile("../../cmd/krci-ai/assets/schemas/task-metadata.json")
	if err != nil {
		t.Fatalf("Failed to create parser with real schema: %v", err)
	}

	// Test parsing a real migrated task file
	realTaskContent := `---
dependencies:
    templates:
        - story.md
    data:
        - common/sdlc-framework.md
        - prioritization-frameworks.md
---
# Task: Create Story

## Description

Create user-focused requirements with implementation tasks and acceptance criteria.
`

	parsed, err := parser.ParseTaskContent(realTaskContent, "create-story.md")
	if err != nil {
		t.Fatalf("Failed to parse real task content: %v", err)
	}

	if !parsed.Metadata.HasDependencies() {
		t.Error("Expected parsed metadata to have dependencies")
	}

	if err := parser.ValidateMetadata(parsed.Metadata, "create-story.md"); err != nil {
		t.Errorf("Real task metadata validation failed: %v", err)
	}
}
