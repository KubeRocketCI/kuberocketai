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
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGPT4Calculator(t *testing.T) {
	calc, err := NewGPT4Calculator()
	require.NoError(t, err, "NewGPT4Calculator() should not fail")
	assert.NotNil(t, calc, "NewGPT4Calculator() should not return nil calculator")
	assert.NotNil(t, calc.encoding, "GPT4Calculator encoding should not be nil")
}

func TestGPT4Calculator_CalculateTokens(t *testing.T) {
	calc, err := NewGPT4Calculator()
	require.NoError(t, err, "Failed to create GPT4Calculator")

	tests := []struct {
		name     string
		text     string
		expected int
		wantErr  bool
	}{
		{
			name:     "empty string",
			text:     "",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "single word",
			text:     "hello",
			expected: 1,
			wantErr:  false,
		},
		{
			name:     "simple sentence",
			text:     "Hello world!",
			expected: 3,
			wantErr:  false,
		},
		{
			name:     "multi-line text",
			text:     "Line 1\nLine 2\nLine 3",
			expected: 11, // Adjusted based on actual GPT-4 tokenization
			wantErr:  false,
		},
		{
			name:     "text with special characters",
			text:     "Hello @world! #test $123",
			expected: 8,
			wantErr:  false,
		},
		{
			name:     "code snippet",
			text:     "func main() {\n\tfmt.Println(\"Hello, World!\")\n}",
			expected: 12, // Adjusted based on actual GPT-4 tokenization
			wantErr:  false,
		},
		{
			name:     "yaml content",
			text:     "name: test\nversion: 1.0\ndescription: A test configuration",
			expected: 16, // Adjusted based on actual GPT-4 tokenization
			wantErr:  false,
		},
		{
			name:     "markdown content",
			text:     "# Title\n\nThis is a **bold** text with *italic* and `code`.",
			expected: 18, // Adjusted based on actual GPT-4 tokenization
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tokens, err := calc.CalculateTokens(ctx, tt.text)

			if tt.wantErr {
				assert.Error(t, err, "Expected error but got none")
				return
			}

			assert.NoError(t, err, "Unexpected error")
			assert.Equal(t, tt.expected, tokens, "CalculateTokens() returned unexpected token count")
		})
	}
}

func TestGPT4Calculator_LargeText(t *testing.T) {
	calc, err := NewGPT4Calculator()
	require.NoError(t, err, "Failed to create GPT4Calculator")

	// Test with larger text to ensure performance
	largeText := ""
	for range 1000 {
		largeText += "This is a test sentence with multiple words. "
	}

	ctx := context.Background()
	tokens, err := calc.CalculateTokens(ctx, largeText)
	assert.NoError(t, err, "Failed to calculate tokens for large text")
	assert.Greater(t, tokens, 0, "Expected positive token count for large text")

	// Verify the token count is reasonable (should be more than number of words)
	expectedMinTokens := 8000 // Conservative estimate
	assert.GreaterOrEqual(t, tokens, expectedMinTokens, "Token count seems too low for large text")
}

func TestGPT4Calculator_SpecialUnicodeCharacters(t *testing.T) {
	calc, err := NewGPT4Calculator()
	require.NoError(t, err, "Failed to create GPT4Calculator")

	tests := []struct {
		name string
		text string
	}{
		{
			name: "emoji",
			text: "Hello 👋 World 🌍",
		},
		{
			name: "chinese characters",
			text: "你好世界",
		},
		{
			name: "mixed languages",
			text: "Hello 世界 مرحبا Bonjour",
		},
		{
			name: "special symbols",
			text: "©®™€£¥§¶†‡•…‰‹›\"\"''–—",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tokens, err := calc.CalculateTokens(ctx, tt.text)
			assert.NoError(t, err, "Failed to calculate tokens for %s", tt.name)
			assert.Greater(t, tokens, 0, "Expected positive token count for %s", tt.name)
		})
	}
}
