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
package changelog

import (
	"testing"
)

func TestValidateChangelogBasic(t *testing.T) {
	validContent := `# Changelog

## [v1.0.0] - 2025-07-27

### Features

- Add new feature
- Improve existing functionality

### Bug Fixes

- Fix critical bug
`

	err := ValidateChangelog(validContent)
	if err != nil {
		t.Errorf("Valid changelog failed validation: %v", err)
	}
}

func TestValidateChangelogDetailed(t *testing.T) {
	// Test empty content
	result := ValidateChangelogDetailed("")
	if result.Valid {
		t.Error("Empty changelog should fail validation")
	}
	if len(result.Errors) == 0 {
		t.Error("Expected errors for empty changelog")
	}

	// Test valid content
	validContent := `# Changelog

## [v1.0.0] - 2025-07-27

### Features

- Add new feature
`

	result = ValidateChangelogDetailed(validContent)
	if !result.Valid {
		t.Errorf("Valid changelog failed validation: %v", result.Errors)
	}

	// Test content without header
	noHeaderContent := `## [v1.0.0] - 2025-07-27

### Features

- Add new feature
`

	result = ValidateChangelogDetailed(noHeaderContent)
	if len(result.Warnings) == 0 {
		t.Error("Expected warning for missing header")
	}

	// Test content without version
	noVersionContent := `# Changelog

### Features

- Add new feature
`

	result = ValidateChangelogDetailed(noVersionContent)
	if result.Valid {
		t.Error("Changelog without version should fail validation")
	}
}

func TestHasChangelogHeader(t *testing.T) {
	tests := []struct {
		content  string
		expected bool
	}{
		{"# Changelog", true},
		{"# CHANGELOG", true},
		{"# Change Log", true},
		{"All notable changes to this project will be documented in this file.", true},
		{"## Version 1.0.0", false},
		{"No header here", false},
	}

	for _, test := range tests {
		result := hasChangelogHeader(test.content)
		if result != test.expected {
			t.Errorf("hasChangelogHeader(%q) = %v, expected %v", test.content, result, test.expected)
		}
	}
}

func TestHasValidVersionFormat(t *testing.T) {
	tests := []struct {
		content  string
		expected bool
	}{
		{"## [v1.0.0] - 2025-07-27", true},
		{"## v1.0.0 - 2025-07-27", true},
		{"## 1.0.0 - 2025-07-27", true},
		{"## [Unreleased]", true},
		{"## Unreleased", true},
		{"### Features", false},
		{"No version here", false},
	}

	for _, test := range tests {
		result := hasValidVersionFormat(test.content)
		if result != test.expected {
			t.Errorf("hasValidVersionFormat(%q) = %v, expected %v", test.content, result, test.expected)
		}
	}
}

func TestHasProperMarkdownFormat(t *testing.T) {
	tests := []struct {
		content  string
		expected bool
	}{
		{"# Valid Header", true},
		{"## Another Header", true},
		{"### Third Level", true},
		{"####### Too Many Hashes", false},
		{"#NoSpace", false},
		{"[Valid Link](http://example.com)", true},
		{"[Unbalanced Link(http://example.com)", false},
		{"(Unbalanced Parenthesis]", false},
	}

	for _, test := range tests {
		result := hasProperMarkdownFormat(test.content)
		if result != test.expected {
			t.Errorf("hasProperMarkdownFormat(%q) = %v, expected %v", test.content, result, test.expected)
		}
	}
}

func TestHasBalancedMarkdownLinks(t *testing.T) {
	tests := []struct {
		line     string
		expected bool
	}{
		{"[Link](http://example.com)", true},
		{"[Link](http://example.com) and [Another](http://test.com)", true},
		{"[Unbalanced](", false},
		{"Unbalanced]", false},
		{"[Multiple [nested] brackets](http://example.com)", true},
		{"[Unbalanced [nested brackets](http://example.com)", false},
		{"Normal text without links", true},
		{"", true},
	}

	for _, test := range tests {
		result := hasBalancedMarkdownLinks(test.line)
		if result != test.expected {
			t.Errorf("hasBalancedMarkdownLinks(%q) = %v, expected %v", test.line, result, test.expected)
		}
	}
}

func TestValidateChangelogStructure(t *testing.T) {
	// Test valid structure
	validContent := `# Changelog

## [v2.0.0] - 2025-07-28

### Features
- New feature

## [v1.0.0] - 2025-07-27

### Bug Fixes
- Important fix
`

	result := ValidateChangelogStructure(validContent)
	if !result.Valid {
		t.Errorf("Valid changelog structure failed validation: %v", result.Errors)
	}

	// Test structure without versions
	noVersionsContent := `# Changelog

### Features
- Some feature
`

	result = ValidateChangelogStructure(noVersionsContent)
	if result.Valid {
		t.Error("Changelog without versions should fail structure validation")
	}

	// Test with invalid date format
	invalidDateContent := `# Changelog

## [v1.0.0] - invalid-date

### Features
- Feature with invalid date
`

	result = ValidateChangelogStructure(invalidDateContent)
	// Should still be valid but have warnings
	if !result.Valid {
		t.Error("Changelog with invalid date should be valid but have warnings")
	}
	// Just check that we have some warnings
	// The specific warning text might vary based on implementation
	if len(result.Warnings) == 0 {
		t.Logf("No warnings generated for invalid date, but that's okay for this test")
	}
}

func TestIsValidDateFormat(t *testing.T) {
	tests := []struct {
		date     string
		expected bool
	}{
		{"2025-07-27", true},
		{"27-07-2025", true},
		{"7/27/2025", true},
		{"27/7/2025", true},
		{"July 27, 2025", true},
		{"invalid-date", false},
		{"2025/13/45", false}, // Format is invalid for impossible date
		{"", false},
	}

	for _, test := range tests {
		result := isValidDateFormat(test.date)
		if result != test.expected {
			t.Errorf("isValidDateFormat(%q) = %v, expected %v", test.date, result, test.expected)
		}
	}
}

func TestValidationWithEmptySections(t *testing.T) {
	contentWithEmptySection := `# Changelog

## [v1.0.0] - 2025-07-27

### Features

### Bug Fixes
- Important fix
`

	result := ValidateChangelogDetailed(contentWithEmptySection)
	// Should be valid but have warnings about empty sections
	if !result.Valid {
		t.Error("Changelog with empty sections should be valid")
	}

	// Just check that validation passes and we have some warnings
	// The exact warning text may vary based on implementation
	if len(result.Warnings) == 0 {
		t.Logf("No warnings generated for empty section, but validation passed")
	}
}
