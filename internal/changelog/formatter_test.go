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
	"testing"

	"github.com/fatih/color"
)

func TestFormatChangelogCompact(t *testing.T) {
	testChangelog := `# Changelog

## v1.2.0 - 2023-01-15

### Features
  • Add new feature A
  • Implement feature B
  • Create feature C
  • Add feature D
  • Implement feature E

### Bug Fixes
  • Fix critical bug in module X
  • Resolve issue with API calls

### Chores
  • Update dependencies
  • Improve documentation

## v1.1.0 - 2023-01-01

### Features
  • Add initial feature set

### Bug Fixes
  • Fix startup issues
`

	result, err := FormatChangelogCompact(testChangelog)
	if err != nil {
		t.Fatalf("FormatChangelogCompact failed: %v", err)
	}

	// Check that result contains expected elements
	if !strings.Contains(result, "🏷️") {
		t.Error("Expected version emoji in compact format")
	}
	if !strings.Contains(result, "🚀") {
		t.Error("Expected features emoji in compact format")
	}
	if !strings.Contains(result, "🐛") {
		t.Error("Expected fixes emoji in compact format")
	}

	// Check that it limits items (should show max 4 items and "... and X more")
	if !strings.Contains(result, "... and 1 more") {
		t.Error("Expected truncation message for features section")
	}

	// Check GitHub link is present (only if more than 3 releases would trigger the link)
	// Since our test has only 2 releases, the link won't be shown
	// This is correct behavior - link only shown when truncating releases
	lines := strings.Split(testChangelog, "\n")
	releaseCount := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "## v") {
			releaseCount++
		}
	}

	if releaseCount > 3 {
		if !strings.Contains(result, "https://github.com/KubeRocketCI/kuberocketai/blob/main/CHANGELOG.md") {
			t.Error("Expected GitHub changelog link when more than 3 releases")
		}
	}
}

func TestParseVersionAndDate(t *testing.T) {
	tests := []struct {
		input        string
		expectedVer  string
		expectedDate string
	}{
		{"v1.2.0 - 2023-01-15", "v1.2.0", "2023-01-15"},
		{"Unreleased", "Unreleased", ""},
		{"v2.0.0", "v2.0.0", ""},
		{"v1.0.0 - 2022-12-01", "v1.0.0", "2022-12-01"},
	}

	for _, test := range tests {
		version, date := parseVersionAndDate(test.input)
		if version != test.expectedVer {
			t.Errorf("parseVersionAndDate(%q) version = %q, want %q", test.input, version, test.expectedVer)
		}
		if date != test.expectedDate {
			t.Errorf("parseVersionAndDate(%q) date = %q, want %q", test.input, date, test.expectedDate)
		}
	}
}

func TestGetSectionTypeFromTitle(t *testing.T) {
	tests := []struct {
		title    string
		expected string
	}{
		{"Features", sectionTypeFeatures},
		{"New Features", sectionTypeFeatures},
		{"Added", sectionTypeFeatures},
		{"Bug Fixes", sectionTypeFixes},
		{"Fixes", sectionTypeFixes},
		{"Fixed", sectionTypeFixes},
		{"Breaking Changes", sectionTypeBreaking},
		{"BREAKING CHANGE", sectionTypeBreaking},
		{"Chores", sectionTypeChores},
		{"Code Refactoring", sectionTypeRefactoring},
		{"Documentation", sectionTypeDocumentation},
		{"Other Stuff", sectionTypeOther},
	}

	for _, test := range tests {
		result := getSectionTypeFromTitle(test.title)
		if result != test.expected {
			t.Errorf("getSectionTypeFromTitle(%q) = %q, want %q", test.title, result, test.expected)
		}
	}
}

func TestFormatVersionHeader(t *testing.T) {
	config := DefaultColorConfig()

	tests := []struct {
		version  string
		date     string
		hasEmoji bool
	}{
		{"v1.2.0", "2023-01-15", true},
		{"Unreleased", "", true},
		{"v2.0.0", "", true},
	}

	for _, test := range tests {
		result := formatVersionHeader(test.version, test.date, config)

		if test.hasEmoji && !strings.Contains(result, "🏷️") && !strings.Contains(result, "🚧") {
			t.Errorf("formatVersionHeader(%q, %q) missing emoji", test.version, test.date)
		}

		if !strings.Contains(result, test.version) {
			t.Errorf("formatVersionHeader(%q, %q) missing version", test.version, test.date)
		}

		if test.date != "" && !strings.Contains(result, test.date) {
			t.Errorf("formatVersionHeader(%q, %q) missing date", test.version, test.date)
		}
	}
}

func TestPrioritizeSections(t *testing.T) {
	sections := []ReleaseSection{
		{Type: sectionTypeChores, Title: "Chores", Items: []string{"item1"}},
		{Type: sectionTypeFeatures, Title: "Features", Items: []string{"item2"}},
		{Type: sectionTypeBreaking, Title: "Breaking", Items: []string{"item3"}},
		{Type: sectionTypeFixes, Title: "Fixes", Items: []string{"item4"}},
	}

	result := prioritizeSections(sections)

	// Check order: breaking > features > fixes > chores
	expectedOrder := []string{sectionTypeBreaking, sectionTypeFeatures, sectionTypeFixes, sectionTypeChores}

	if len(result) != len(expectedOrder) {
		t.Fatalf("Expected %d sections, got %d", len(expectedOrder), len(result))
	}

	for i, expected := range expectedOrder {
		if result[i].Type != expected {
			t.Errorf("Section %d: expected type %q, got %q", i, expected, result[i].Type)
		}
	}
}

func TestGetSectionEmoji(t *testing.T) {
	tests := []struct {
		sectionType string
		expected    string
	}{
		{sectionTypeFeatures, "🚀"},
		{sectionTypeFixes, "🐛"},
		{sectionTypeBreaking, "💥"},
		{sectionTypeChores, "🔧"},
		{sectionTypeRefactoring, "♻️"},
		{sectionTypeDocumentation, "📚"},
		{sectionTypeOther, "✨"},
		{"unknown", "✨"},
	}

	for _, test := range tests {
		result := getSectionEmoji(test.sectionType)
		if result != test.expected {
			t.Errorf("getSectionEmoji(%q) = %q, want %q", test.sectionType, result, test.expected)
		}
	}
}

func TestGetSectionColorByType(t *testing.T) {
	config := DefaultColorConfig()

	tests := []struct {
		sectionType string
		expected    *color.Color
	}{
		{sectionTypeFeatures, config.Features},
		{sectionTypeFixes, config.Fixes},
		{sectionTypeBreaking, config.Breaking},
		{sectionTypeOther, config.Header},
	}

	for _, test := range tests {
		result := getSectionColorByType(test.sectionType, config)
		if result != test.expected {
			t.Errorf("getSectionColorByType(%q) returned unexpected color", test.sectionType)
		}
	}
}

func TestShortenItem(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"cli: improve command handling", "improve command handling"},
		{"cmd: add new feature", "add new feature"},
		{"internal: refactor code", "refactor code"},
		{"Update krci-ai to the version 1.2.0", "Update to version 1.2.0"},
		{"Add support for new feature", "Add new feature"},
		{"Implement new functionality", "Add new functionality"},
		{"Create new component", "Add new component"},
		{"Normal item without prefix", "Normal item without prefix"},
	}

	for _, test := range tests {
		result := shortenItem(test.input)
		if result != test.expected {
			t.Errorf("shortenItem(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestShortenItemLongText(t *testing.T) {
	longText := strings.Repeat("a", 85)
	result := shortenItem(longText)

	if len(result) > 80 {
		t.Errorf("shortenItem should limit text to 80 characters, got %d", len(result))
	}

	if !strings.HasSuffix(result, "...") {
		t.Error("shortenItem should add '...' to truncated text")
	}
}

func TestFormatCompactSection(t *testing.T) {
	config := DefaultColorConfig()

	section := ReleaseSection{
		Type:  sectionTypeFeatures,
		Title: "Features",
		Items: []string{
			"Add feature A",
			"Add feature B",
			"Add feature C",
			"Add feature D",
			"Add feature E",
		},
	}

	result := formatCompactSection(section, config)

	// Check for emoji
	if !strings.Contains(result, "🚀") {
		t.Error("Expected features emoji in section formatting")
	}

	// Check for section title
	if !strings.Contains(result, "Features") {
		t.Error("Expected section title in formatting")
	}

	// Check for truncation since we have 5 items (more than 4)
	if !strings.Contains(result, "... and 1 more") {
		t.Error("Expected truncation message for 5 items")
	}

	// Check that first 4 items are present
	for i := 0; i < 4; i++ {
		expected := string(rune('A' + i))
		if !strings.Contains(result, expected) {
			t.Errorf("Expected feature %s in formatted section", expected)
		}
	}
}

func TestParseReleases(t *testing.T) {
	testChangelog := `# Changelog

## v1.2.0 - 2023-01-15

### Features
  • Add new feature A
  • Implement feature B

### Bug Fixes
  • Fix critical bug

## v1.1.0 - 2023-01-01

### Features
  • Add initial feature
`

	releases := parseReleases(testChangelog)

	if len(releases) != 2 {
		t.Fatalf("Expected 2 releases, got %d", len(releases))
	}

	// Check first release
	if releases[0].Version != "v1.2.0" {
		t.Errorf("Expected version v1.2.0, got %q", releases[0].Version)
	}
	if releases[0].Date != "2023-01-15" {
		t.Errorf("Expected date 2023-01-15, got %q", releases[0].Date)
	}
	if len(releases[0].Sections) != 2 {
		t.Errorf("Expected 2 sections in first release, got %d", len(releases[0].Sections))
	}

	// Check sections content
	featuresSection := releases[0].Sections[0]
	if featuresSection.Type != sectionTypeFeatures {
		t.Errorf("Expected features section type, got %q", featuresSection.Type)
	}
	if len(featuresSection.Items) != 2 {
		t.Errorf("Expected 2 feature items, got %d", len(featuresSection.Items))
	}

	fixesSection := releases[0].Sections[1]
	if fixesSection.Type != sectionTypeFixes {
		t.Errorf("Expected fixes section type, got %q", fixesSection.Type)
	}
	if len(fixesSection.Items) != 1 {
		t.Errorf("Expected 1 fix item, got %d", len(fixesSection.Items))
	}
}

func TestParseReleasesEmpty(t *testing.T) {
	releases := parseReleases("")
	if len(releases) != 0 {
		t.Errorf("Expected 0 releases for empty content, got %d", len(releases))
	}
}

func TestParseReleasesUnreleased(t *testing.T) {
	testChangelog := `## Unreleased

### Features
  • New unreleased feature`

	releases := parseReleases(testChangelog)

	if len(releases) != 1 {
		t.Fatalf("Expected 1 release, got %d", len(releases))
	}

	if releases[0].Version != "Unreleased" {
		t.Errorf("Expected 'Unreleased' version, got %q", releases[0].Version)
	}
	if releases[0].Date != "" {
		t.Errorf("Expected empty date for unreleased, got %q", releases[0].Date)
	}
}

func TestReleaseParserHandleVersionHeader(t *testing.T) {
	parser := &releaseParser{
		versionRegex: regexp.MustCompile(`^##\s+(.+)$`),
	}

	tests := []struct {
		line        string
		shouldMatch bool
		version     string
	}{
		{"## v1.2.0 - 2023-01-15", true, "v1.2.0"},
		{"## Unreleased", true, "Unreleased"},
		{"### Features", false, ""},
		{"Normal text", false, ""},
	}

	for _, test := range tests {
		result := parser.handleVersionHeader(test.line)
		if result != test.shouldMatch {
			t.Errorf("handleVersionHeader(%q) = %v, want %v", test.line, result, test.shouldMatch)
		}

		if test.shouldMatch && parser.currentRelease != nil {
			if parser.currentRelease.Version != test.version {
				t.Errorf("Expected version %q, got %q", test.version, parser.currentRelease.Version)
			}
		}
	}
}

func TestReleaseParserHandleSectionHeader(t *testing.T) {
	parser := &releaseParser{
		sectionRegex: regexp.MustCompile(`^###\s+(.+)$`),
		currentRelease: &Release{
			Version:  "v1.0.0",
			Date:     "2023-01-01",
			Sections: []ReleaseSection{},
		},
	}

	tests := []struct {
		line        string
		shouldMatch bool
		sectionType string
	}{
		{"### Features", true, sectionTypeFeatures},
		{"### Bug Fixes", true, sectionTypeFixes},
		{"### Breaking Changes", true, sectionTypeBreaking},
		{"## v1.2.0", false, ""},
		{"Normal text", false, ""},
	}

	for _, test := range tests {
		result := parser.handleSectionHeader(test.line)
		if result != test.shouldMatch {
			t.Errorf("handleSectionHeader(%q) = %v, want %v", test.line, result, test.shouldMatch)
		}

		if test.shouldMatch && parser.currentSection != nil {
			if parser.currentSection.Type != test.sectionType {
				t.Errorf("Expected section type %q, got %q", test.sectionType, parser.currentSection.Type)
			}
		}
	}
}

func TestReleaseParserHandleListItem(t *testing.T) {
	parser := &releaseParser{
		itemRegex: regexp.MustCompile(`^\s*[•*-]\s+(.+)$`),
		currentSection: &ReleaseSection{
			Type:  sectionTypeFeatures,
			Title: "Features",
			Items: []string{},
		},
	}

	tests := []struct {
		line     string
		expected string
	}{
		{"  • Add new feature", "Add new feature"},
		{"  * Fix important bug", "Fix important bug"},
		{"  - Update documentation", "Update documentation"},
		{"    • Indented item", "Indented item"},
		{"Normal text", ""},
		{"### Header", ""},
	}

	initialItemCount := len(parser.currentSection.Items)
	validItemCount := 0

	for _, test := range tests {
		parser.handleListItem(test.line)
		if test.expected != "" {
			validItemCount++
		}
	}

	finalItemCount := len(parser.currentSection.Items)
	expectedFinalCount := initialItemCount + validItemCount

	if finalItemCount != expectedFinalCount {
		t.Errorf("Expected %d items after processing, got %d", expectedFinalCount, finalItemCount)
	}

	// Check that valid items were added correctly
	for i, test := range tests {
		if test.expected != "" && i < len(parser.currentSection.Items) {
			// Find the item in the list (order might vary due to multiple valid items)
			found := false
			for _, item := range parser.currentSection.Items {
				if item == test.expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected item %q not found in section items", test.expected)
			}
		}
	}
}

func TestReleaseParserCompleteFlow(t *testing.T) {
	parser := &releaseParser{
		versionRegex: regexp.MustCompile(`^##\s+(.+)$`),
		sectionRegex: regexp.MustCompile(`^###\s+(.+)$`),
		itemRegex:    regexp.MustCompile(`^\s*[•*-]\s+(.+)$`),
	}

	lines := []string{
		"## v1.2.0 - 2023-01-15",
		"### Features",
		"  • Add feature A",
		"  • Add feature B",
		"### Bug Fixes",
		"  • Fix bug X",
		"## v1.1.0 - 2023-01-01",
		"### Features",
		"  • Initial feature",
	}

	for _, line := range lines {
		parser.processLine(line)
	}
	parser.finalizeCurrentRelease()

	if len(parser.releases) != 2 {
		t.Fatalf("Expected 2 releases, got %d", len(parser.releases))
	}

	// Verify first release
	release1 := parser.releases[0]
	if release1.Version != "v1.2.0" {
		t.Errorf("Expected version v1.2.0, got %q", release1.Version)
	}
	if len(release1.Sections) != 2 {
		t.Errorf("Expected 2 sections in first release, got %d", len(release1.Sections))
	}

	// Verify second release
	release2 := parser.releases[1]
	if release2.Version != "v1.1.0" {
		t.Errorf("Expected version v1.1.0, got %q", release2.Version)
	}
	if len(release2.Sections) != 1 {
		t.Errorf("Expected 1 section in second release, got %d", len(release2.Sections))
	}
}

func TestFormatFallback(t *testing.T) {
	content := "# Simple changelog\n\n## v1.0.0\n- Simple item"
	config := DefaultColorConfig()

	result, err := formatFallback(content, config)
	if err != nil {
		t.Fatalf("formatFallback failed: %v", err)
	}

	if result == "" {
		t.Error("formatFallback should return non-empty result")
	}
}
