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
	"fmt"
	"regexp"
	"strings"
)

// ValidationResult contains validation results
type ValidationResult struct {
	Valid    bool
	Errors   []string
	Warnings []string
}

// ValidateChangelog validates changelog content for proper format and structure
func ValidateChangelog(content string) error {
	result := ValidateChangelogDetailed(content)
	if !result.Valid {
		return fmt.Errorf("changelog validation failed: %s", strings.Join(result.Errors, "; "))
	}
	return nil
}

// ValidateChangelogDetailed performs detailed validation and returns comprehensive results
func ValidateChangelogDetailed(content string) *ValidationResult {
	result := &ValidationResult{
		Valid:    true,
		Errors:   []string{},
		Warnings: []string{},
	}

	if strings.TrimSpace(content) == "" {
		result.Valid = false
		result.Errors = append(result.Errors, "changelog content is empty")
		return result
	}

	// Check for required header
	if !hasChangelogHeader(content) {
		result.Warnings = append(result.Warnings, "missing standard changelog header")
	}

	// Check for proper version format
	if !hasValidVersionFormat(content) {
		result.Valid = false
		result.Errors = append(result.Errors, "no valid version headers found")
	}

	// Check for proper section structure
	sections := GetAvailableSections(content)
	if len(sections) == 0 {
		result.Warnings = append(result.Warnings, "no recognized sections found (Features, Bug Fixes, etc.)")
	}

	// Check for empty sections
	parser := NewChangelogParser(content)
	sectionMap, _ := parser.ParseSections()
	for sectionName, sectionContent := range sectionMap {
		if strings.TrimSpace(sectionContent) == "" {
			result.Warnings = append(result.Warnings, fmt.Sprintf("section '%s' is empty", sectionName))
		}
	}

	// Check for proper markdown formatting
	if !hasProperMarkdownFormat(content) {
		result.Valid = false
		result.Errors = append(result.Errors, "invalid markdown format detected")
	}

	return result
}

// hasChangelogHeader checks if the changelog has a proper header
func hasChangelogHeader(content string) bool {
	headerPatterns := []string{
		`(?i)^#\s+changelog`,
		`(?i)^#\s+change\s*log`,
		`(?i)all notable changes`,
	}

	for _, pattern := range headerPatterns {
		if matched, _ := regexp.MatchString(pattern, content); matched {
			return true
		}
	}

	return false
}

// hasValidVersionFormat checks if changelog has proper version headers
func hasValidVersionFormat(content string) bool {
	// Look for version patterns like "## [v1.0.0]" or "## v1.0.0"
	versionPatterns := []string{
		`(?m)^##\s*\[.*?\].*`,          // ## [v1.0.0] - 2025-07-27
		`(?m)^##\s*v?\d+\.\d+\.\d+.*`,  // ## v1.0.0 - 2025-07-27 or ## 1.0.0
		`(?m)^##\s*\[?Unreleased\]?.*`, // ## [Unreleased] or ## Unreleased
	}

	for _, pattern := range versionPatterns {
		if matched, _ := regexp.MatchString(pattern, content); matched {
			return true
		}
	}

	return false
}

// hasProperMarkdownFormat checks for basic markdown formatting issues
func hasProperMarkdownFormat(content string) bool {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Check for malformed headers (more than 6 #)
		if strings.HasPrefix(trimmedLine, "#") {
			headerLevel := 0
			for _, char := range trimmedLine {
				if char == '#' {
					headerLevel++
				} else {
					break
				}
			}

			if headerLevel > 6 {
				return false
			}

			// Check if header has proper spacing
			if headerLevel > 0 && len(trimmedLine) > headerLevel {
				if trimmedLine[headerLevel] != ' ' {
					return false
				}
			}
		}

		// Check for unbalanced brackets/parentheses in links
		if strings.Contains(trimmedLine, "[") || strings.Contains(trimmedLine, "]") ||
			strings.Contains(trimmedLine, "(") || strings.Contains(trimmedLine, ")") {
			if !hasBalancedMarkdownLinks(trimmedLine) {
				return false
			}
		}

		// Warn about lines that are too long (but don't fail validation)
		if len(line) > 120 {
			// This is just a warning condition, not an error
			continue
		}

		_ = i // Use the line number if needed for error reporting
	}

	return true
}

// hasBalancedMarkdownLinks checks if markdown links are properly formatted
func hasBalancedMarkdownLinks(line string) bool {
	// Simple check for balanced brackets and parentheses in links
	squareBrackets := 0
	parentheses := 0

	for _, char := range line {
		switch char {
		case '[':
			squareBrackets++
		case ']':
			squareBrackets--
		case '(':
			parentheses++
		case ')':
			parentheses--
		}

		// If we have negative counts, brackets/parentheses are unbalanced
		if squareBrackets < 0 || parentheses < 0 {
			return false
		}
	}

	// Final counts should be zero for balanced
	return squareBrackets == 0 && parentheses == 0
}

// ValidateChangelogStructure validates the overall structure follows best practices
func ValidateChangelogStructure(content string) *ValidationResult {
	result := ValidateChangelogDetailed(content)

	// Additional structural checks
	parser := NewChangelogParser(content)
	versions := parser.GetVersions()

	if len(versions) == 0 {
		result.Valid = false
		result.Errors = append(result.Errors, "no version sections found")
		return result
	}

	// Check if versions are in chronological order (latest first)
	for i := 0; i < len(versions)-1; i++ {
		current := versions[i]
		next := versions[i+1]

		// Skip unreleased versions
		if strings.ToLower(current.Version) == "unreleased" {
			continue
		}

		// Basic date format check
		if !isValidDateFormat(current.Date) {
			result.Warnings = append(result.Warnings, fmt.Sprintf("version %s has invalid date format: %s", current.Version, current.Date))
		}

		_ = next // Could add chronological validation here
	}

	return result
}

// isValidDateFormat checks if date is in acceptable format
func isValidDateFormat(date string) bool {
	// Accept various date formats commonly used in changelogs
	datePatterns := []string{
		`^\d{4}-\d{2}-\d{2}$`,        // 2025-07-27
		`^\d{2}-\d{2}-\d{4}$`,        // 27-07-2025
		`^\d{1,2}/\d{1,2}/\d{4}$`,    // 7/27/2025 or 27/7/2025
		`^[A-Za-z]+ \d{1,2}, \d{4}$`, // July 27, 2025
	}

	for _, pattern := range datePatterns {
		if matched, _ := regexp.MatchString(pattern, date); matched {
			return true
		}
	}

	return false
}
