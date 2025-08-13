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
package tokens

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatUserFriendlyError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedOutput string
	}{
		{
			name:           "nil error",
			err:            nil,
			expectedOutput: "",
		},
		{
			name:           "agent not found error",
			err:            errors.New("agent 'test' not found in framework"),
			expectedOutput: "❌ Agent not found. Use 'krci-ai list' to see available agents.",
		},
		{
			name:           "framework not installed error",
			err:            errors.New("framework not installed in current directory"),
			expectedOutput: "❌ Framework not installed in current directory. Run 'krci-ai install' first.",
		},
		{
			name:           "file not found error",
			err:            errors.New("file not found: /path/to/missing/file.yaml"),
			expectedOutput: "❌ Required file is missing. Please ensure all framework files are present.",
		},
		{
			name:           "no such file error",
			err:            errors.New("open /path/file.yaml: no such file or directory"),
			expectedOutput: "❌ Required file is missing. Please ensure all framework files are present.",
		},
		{
			name:           "permission denied error",
			err:            errors.New("permission denied accessing /path/to/file"),
			expectedOutput: "❌ Permission denied. Please check file permissions and try again.",
		},
		{
			name:           "failed to encode text error",
			err:            errors.New("failed to encode text: invalid UTF-8 sequence"),
			expectedOutput: "❌ Failed to process text content. The file may contain unsupported characters.",
		},
		{
			name:           "context canceled error",
			err:            errors.New("context canceled"),
			expectedOutput: "❌ Operation was canceled. Please try again.",
		},
		{
			name:           "context deadline exceeded error",
			err:            errors.New("context deadline exceeded"),
			expectedOutput: "❌ Operation timed out. Please try again or check for large files.",
		},
		{
			name:           "token limit exceeded error",
			err:            errors.New("token count 150000 exceeds maximum limit of 128000"),
			expectedOutput: "⚠️  Token limit exceeded. Consider breaking down your configuration into smaller components.",
		},
		{
			name:           "unknown error",
			err:            errors.New("unexpected database connection error"),
			expectedOutput: "❌ An error occurred: unexpected database connection error\n\nIf this problem persists, please check your framework installation and file permissions.",
		},
		{
			name:           "complex agent error",
			err:            errors.New("failed to discover agent dependencies: agent 'unknown' not found"),
			expectedOutput: "❌ Agent not found. Use 'krci-ai list' to see available agents.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatUserFriendlyError(tt.err)
			assert.Equal(t, tt.expectedOutput, result, "FormatUserFriendlyError() output mismatch")
		})
	}
}

func TestSuggestSolutions(t *testing.T) {
	tests := []struct {
		name                string
		err                 error
		expectedSuggestions []string
	}{
		{
			name:                "nil error",
			err:                 nil,
			expectedSuggestions: nil,
		},
		{
			name: "agent not found error",
			err:  errors.New("agent 'test' not found in framework"),
			expectedSuggestions: []string{
				"• Run 'krci-ai list' to see all available agents",
				"• Check the agent name spelling",
				"• Ensure the framework is properly installed",
			},
		},
		{
			name: "framework not installed error",
			err:  errors.New("framework not installed in current directory"),
			expectedSuggestions: []string{
				"• Run 'krci-ai install' to install the framework",
				"• Navigate to your project directory first",
				"• Ensure you have write permissions in the current directory",
			},
		},
		{
			name: "file not found error",
			err:  errors.New("file not found: /path/to/missing/file.yaml"),
			expectedSuggestions: []string{
				"• Verify all framework files are present",
				"• Try reinstalling with 'krci-ai install'",
				"• Check if files were accidentally deleted or moved",
			},
		},
		{
			name: "no such file error",
			err:  errors.New("open /path/file.yaml: no such file or directory"),
			expectedSuggestions: []string{
				"• Verify all framework files are present",
				"• Try reinstalling with 'krci-ai install'",
				"• Check if files were accidentally deleted or moved",
			},
		},
		{
			name: "permission denied error",
			err:  errors.New("permission denied accessing /path/to/file"),
			expectedSuggestions: []string{
				"• Check file and directory permissions",
				"• Try running with appropriate privileges",
				"• Ensure you have read access to framework files",
			},
		},
		{
			name: "token limit exceeded error",
			err:  errors.New("token count 150000 exceeds maximum limit of 128000"),
			expectedSuggestions: []string{
				"• Break down large configurations into smaller components",
				"• Remove unnecessary content from templates and tasks",
				"• Consider using shorter descriptions and documentation",
				"• Use references instead of inline content where possible",
			},
		},
		{
			name: "unknown error",
			err:  errors.New("unexpected database connection error"),
			expectedSuggestions: []string{
				"• Check your framework installation: 'krci-ai validate'",
				"• Ensure all required files are present and accessible",
				"• Try reinstalling the framework if issues persist",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SuggestSolutions(tt.err)

			require.Len(t, result, len(tt.expectedSuggestions), "SuggestSolutions() returned unexpected number of suggestions")

			for i, suggestion := range result {
				assert.Equal(t, tt.expectedSuggestions[i], suggestion, "Suggestion %d mismatch", i)
			}
		})
	}
}

func TestFormatUserFriendlyError_EdgeCases(t *testing.T) {
	tests := []struct {
		name             string
		err              error
		shouldContain    []string
		shouldNotContain []string
	}{
		{
			name:          "mixed case agent error",
			err:           errors.New("Agent 'TEST' Not Found in framework"),
			shouldContain: []string{"❌", "An error occurred"}, // This will fall through to default case
		},
		{
			name:          "partial token limit error",
			err:           errors.New("token count exceeds limit but not full message"),
			shouldContain: []string{"⚠️", "Token limit exceeded"},
		},
		{
			name:             "very long error message",
			err:              errors.New("this is a very long error message that contains many words and should still be handled properly by the error formatter without any issues"),
			shouldContain:    []string{"❌", "An error occurred"},
			shouldNotContain: []string{"Agent not found", "Framework not installed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatUserFriendlyError(tt.err)

			for _, should := range tt.shouldContain {
				assert.Contains(t, result, should, "Expected result to contain %q", should)
			}

			for _, shouldNot := range tt.shouldNotContain {
				assert.NotContains(t, result, shouldNot, "Expected result to NOT contain %q", shouldNot)
			}
		})
	}
}

func TestSuggestSolutions_EdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		minSuggestions int
		maxSuggestions int
	}{
		{
			name:           "complex error with multiple keywords",
			err:            errors.New("agent not found and file not found in framework"),
			minSuggestions: 3, // Should match agent not found pattern
			maxSuggestions: 5,
		},
		{
			name:           "error with no matching patterns",
			err:            errors.New("completely unrelated error message"),
			minSuggestions: 3, // Should fall back to default suggestions
			maxSuggestions: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SuggestSolutions(tt.err)

			assert.GreaterOrEqual(t, len(result), tt.minSuggestions, "Expected at least %d suggestions", tt.minSuggestions)
			assert.LessOrEqual(t, len(result), tt.maxSuggestions, "Expected at most %d suggestions", tt.maxSuggestions)

			// Verify all suggestions start with bullet point
			for i, suggestion := range result {
				assert.True(t, strings.HasPrefix(suggestion, "• "), "Suggestion %d should start with '• ', got: %s", i, suggestion)
			}
		})
	}
}

// Test error message formatting consistency
func TestErrorMessageFormatting(t *testing.T) {
	testErrors := []error{
		errors.New("agent not found"),
		errors.New("framework not installed"),
		errors.New("file not found"),
		errors.New("permission denied"),
		errors.New("context canceled"),
		errors.New("token count exceeds limit"),
	}

	for _, err := range testErrors {
		t.Run(err.Error(), func(t *testing.T) {
			formatted := FormatUserFriendlyError(err)

			// All error messages should start with an emoji
			assert.True(t, strings.HasPrefix(formatted, "❌") || strings.HasPrefix(formatted, "⚠️"),
				"Error message should start with emoji, got: %s", formatted)

			// Error messages should not be empty
			assert.NotEmpty(t, strings.TrimSpace(formatted), "Error message should not be empty")

			// Error messages should not contain the raw error text for known patterns
			if strings.Contains(err.Error(), "unexpected") {
				// Skip this check for "unexpected" errors as they should contain raw text
				return
			}
			assert.False(t, strings.Contains(formatted, err.Error()),
				"Formatted error should not contain raw error text for known patterns: %s", formatted)
		})
	}
}
