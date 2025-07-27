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
	"regexp"
	"strings"
)

// ChangelogParser handles parsing of changelog content
type ChangelogParser struct {
	content string
}

// NewChangelogParser creates a new changelog parser
func NewChangelogParser(content string) *ChangelogParser {
	return &ChangelogParser{
		content: content,
	}
}

// ParseSections extracts sections from changelog content
func (p *ChangelogParser) ParseSections() (map[string]string, error) {
	sections := make(map[string]string)

	// Define section patterns
	sectionPatterns := []string{
		"Features",
		"Bug Fixes",
		"Breaking Changes",
		"Documentation",
		"Chores",
		"Performance Improvements",
		"Code Refactoring",
		"Tests",
		"Styles",
	}

	for _, sectionName := range sectionPatterns {
		content := p.extractSection(sectionName)
		if content != "" {
			sections[sectionName] = content
		}
	}

	return sections, nil
}

// extractSection extracts content for a specific section
func (p *ChangelogParser) extractSection(sectionName string) string {
	// Create regex pattern to match section headers
	pattern := `(?i)^###\s*` + regexp.QuoteMeta(sectionName) + `\s*$`
	sectionRegex := regexp.MustCompile(pattern)

	lines := strings.Split(p.content, "\n")
	var sectionContent []string
	inSection := false

	for _, line := range lines {
		// Check if this line starts a new section
		if strings.HasPrefix(strings.TrimSpace(line), "###") {
			if sectionRegex.MatchString(line) {
				inSection = true
				continue
			} else {
				inSection = false
				continue
			}
		}

		// Check if we've hit a new version section
		if strings.HasPrefix(strings.TrimSpace(line), "##") && !strings.HasPrefix(strings.TrimSpace(line), "###") {
			inSection = false
			continue
		}

		// Collect content if we're in the target section
		if inSection {
			sectionContent = append(sectionContent, line)
		}
	}

	return strings.TrimSpace(strings.Join(sectionContent, "\n"))
}

// ExtractSection extracts content for a specific section name
func ExtractSection(content, sectionName string) (string, error) {
	parser := NewChangelogParser(content)
	sections, err := parser.ParseSections()
	if err != nil {
		return "", err
	}

	return sections[sectionName], nil
}

// GetAvailableSections returns all available section names in the changelog
func GetAvailableSections(content string) []string {
	parser := NewChangelogParser(content)
	sections, _ := parser.ParseSections()

	var availableSections []string
	for sectionName := range sections {
		availableSections = append(availableSections, sectionName)
	}

	return availableSections
}

// GetVersions extracts version information from changelog
func (p *ChangelogParser) GetVersions() []VersionSection {
	var versions []VersionSection

	// Regex to match version headers like "## [v1.0.0] - 2025-07-27" or "## v1.0.0 - 2025-07-27"
	versionRegex := regexp.MustCompile(`^##\s*(?:\[(.+?)\]|(.+?))\s*-\s*(.+)$`)

	lines := strings.Split(p.content, "\n")
	var currentVersion *VersionSection

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Check if this is a version header
		if matches := versionRegex.FindStringSubmatch(trimmedLine); matches != nil {
			// Save previous version if exists
			if currentVersion != nil {
				versions = append(versions, *currentVersion)
			}

			// Start new version section
			version := matches[1]
			if version == "" {
				version = matches[2]
			}

			currentVersion = &VersionSection{
				Version: strings.TrimSpace(version),
				Date:    strings.TrimSpace(matches[3]),
				Content: "",
			}
		} else if currentVersion != nil {
			// Add content to current version
			currentVersion.Content += line + "\n"
		}
	}

	// Add last version
	if currentVersion != nil {
		versions = append(versions, *currentVersion)
	}

	return versions
}

// VersionSection represents a version section in the changelog
type VersionSection struct {
	Version string
	Date    string
	Content string
}
