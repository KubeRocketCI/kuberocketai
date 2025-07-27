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

// SectionType represents different types of changelog sections
type SectionType string

const (
	SectionFeatures      SectionType = "Features"
	SectionBugFixes      SectionType = "Bug Fixes"
	SectionBreaking      SectionType = "Breaking Changes"
	SectionDocumentation SectionType = "Documentation"
	SectionChores        SectionType = "Chores"
	SectionPerformance   SectionType = "Performance Improvements"
	SectionRefactoring   SectionType = "Code Refactoring"
	SectionTests         SectionType = "Tests"
	SectionStyles        SectionType = "Styles"
)

// Section represents a changelog section with its content
type Section struct {
	Type    SectionType
	Title   string
	Content string
	Items   []string
}

// SectionExtractor handles extraction and processing of changelog sections
type SectionExtractor struct {
	content string
}

// NewSectionExtractor creates a new section extractor
func NewSectionExtractor(content string) *SectionExtractor {
	return &SectionExtractor{
		content: content,
	}
}

// ExtractAllSections extracts all sections with their detailed information
func (se *SectionExtractor) ExtractAllSections() ([]Section, error) {
	var sections []Section

	sectionTypes := []SectionType{
		SectionFeatures,
		SectionBugFixes,
		SectionBreaking,
		SectionDocumentation,
		SectionChores,
		SectionPerformance,
		SectionRefactoring,
		SectionTests,
		SectionStyles,
	}

	for _, sectionType := range sectionTypes {
		section, err := se.extractSectionDetailed(string(sectionType))
		if err == nil && section.Content != "" {
			section.Type = sectionType
			sections = append(sections, section)
		}
	}

	return sections, nil
}

// extractSectionDetailed extracts detailed section information
func (se *SectionExtractor) extractSectionDetailed(sectionName string) (Section, error) {
	section := Section{
		Title: sectionName,
	}

	// Extract raw content
	content, err := ExtractSection(se.content, sectionName)
	if err != nil {
		return section, err
	}

	section.Content = content

	// Extract individual items from the section
	items := se.extractSectionItems(content)
	section.Items = items

	return section, nil
}

// extractSectionItems extracts individual items from section content
func (se *SectionExtractor) extractSectionItems(content string) []string {
	var items []string

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Look for list items (lines starting with -, *, or •)
		if strings.HasPrefix(trimmedLine, "- ") ||
			strings.HasPrefix(trimmedLine, "* ") ||
			strings.HasPrefix(trimmedLine, "• ") {

			// Remove the list marker and clean up
			item := strings.TrimSpace(trimmedLine[2:])
			if item != "" {
				items = append(items, item)
			}
		}
	}

	return items
}

// GetSectionsByVersion extracts sections for a specific version
func (se *SectionExtractor) GetSectionsByVersion(version string) ([]Section, error) {
	versionContent, err := se.extractVersionContent(version)
	if err != nil {
		return nil, err
	}

	versionExtractor := NewSectionExtractor(versionContent)
	return versionExtractor.ExtractAllSections()
}

// extractVersionContent extracts content for a specific version
func (se *SectionExtractor) extractVersionContent(version string) (string, error) {
	parser := NewChangelogParser(se.content)
	versions := parser.GetVersions()

	for _, v := range versions {
		if strings.EqualFold(v.Version, version) ||
			strings.Contains(strings.ToLower(v.Version), strings.ToLower(version)) {
			return v.Content, nil
		}
	}

	return "", fmt.Errorf("version %s not found in changelog", version)
}

// GetNestedSections handles nested subsections within a section
func GetNestedSections(content string) map[string][]string {
	nestedSections := make(map[string][]string)

	lines := strings.Split(content, "\n")
	currentSubsection := ""

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Check for subsection headers (#### level)
		if strings.HasPrefix(trimmedLine, "#### ") {
			currentSubsection = strings.TrimSpace(trimmedLine[4:])
			nestedSections[currentSubsection] = []string{}
		} else if currentSubsection != "" && (strings.HasPrefix(trimmedLine, "- ") ||
			strings.HasPrefix(trimmedLine, "* ") || strings.HasPrefix(trimmedLine, "• ")) {

			// Add item to current subsection
			item := strings.TrimSpace(trimmedLine[2:])
			if item != "" {
				nestedSections[currentSubsection] = append(nestedSections[currentSubsection], item)
			}
		}
	}

	return nestedSections
}

// FilterSectionsByPattern filters sections based on a regex pattern
func FilterSectionsByPattern(sections []Section, pattern string) ([]Section, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	var filtered []Section
	for _, section := range sections {
		if regex.MatchString(section.Content) || regex.MatchString(section.Title) {
			filtered = append(filtered, section)
		}
	}

	return filtered, nil
}

// GetSectionSummary creates a summary of all sections
func GetSectionSummary(sections []Section) map[string]int {
	summary := make(map[string]int)

	for _, section := range sections {
		summary[section.Title] = len(section.Items)
	}

	return summary
}

// FormatSectionForDisplay formats a section for terminal display
func FormatSectionForDisplay(section Section) string {
	var result strings.Builder

	// Add section title
	result.WriteString(fmt.Sprintf("### %s\n\n", section.Title))

	// Add items
	for _, item := range section.Items {
		result.WriteString(fmt.Sprintf("- %s\n", item))
	}

	result.WriteString("\n")
	return result.String()
}
