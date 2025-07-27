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
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// Constants for section types
const (
	sectionTypeFeatures      = "features"
	sectionTypeFixes         = "fixes"
	sectionTypeBreaking      = "breaking"
	sectionTypeChores        = "chores"
	sectionTypeRefactoring   = "refactoring"
	sectionTypeDocumentation = "documentation"
	sectionTypeOther         = "other"
)

// ColorConfig defines colors for different section types
type ColorConfig struct {
	Header   *color.Color
	Features *color.Color
	Fixes    *color.Color
	Breaking *color.Color
	Body     *color.Color
	List     *color.Color
	Bold     *color.Color
}

// DefaultColorConfig returns the default color configuration
func DefaultColorConfig() *ColorConfig {
	return &ColorConfig{
		Header:   color.New(color.FgBlue, color.Bold),
		Features: color.New(color.FgGreen),
		Fixes:    color.New(color.FgYellow),
		Breaking: color.New(color.FgRed, color.Bold),
		Body:     color.New(color.FgWhite),
		List:     color.New(color.FgCyan),
		Bold:     color.New(color.Bold),
	}
}

// FormatChangelog converts markdown content to colored terminal output
func FormatChangelog(content string) (string, error) {
	return FormatChangelogCompact(content)
}

// FormatChangelogCompact creates a compact, user-friendly view
func FormatChangelogCompact(content string) (string, error) {
	config := DefaultColorConfig()

	// Parse content into releases
	releases := parseReleases(content)
	if len(releases) == 0 {
		return formatFallback(content, config)
	}

	var result strings.Builder

	// Show only the most recent releases (up to 3)
	maxReleases := 3
	if len(releases) > maxReleases {
		releases = releases[:maxReleases]
	}

	for _, release := range releases {
		// Format release header with version and date
		versionLine := formatVersionHeader(release.Version, release.Date, config)
		result.WriteString(versionLine)
		result.WriteString("\n")

		// Group and format sections with priority
		sections := prioritizeSections(release.Sections)
		for _, section := range sections {
			if len(section.Items) == 0 {
				continue
			}

			// Format section with emoji and compact items
			sectionText := formatCompactSection(section, config)
			result.WriteString(sectionText)
		}
	}

	// Add footer hint if more releases available
	if len(releases) == maxReleases {
		result.WriteString(config.Body.Sprint("💡 Showing latest 3 releases. Full history available in "))
		result.WriteString(config.Header.Sprint("https://github.com/KubeRocketCI/kuberocketai/blob/main/CHANGELOG.md"))
		result.WriteString("\n")
	}

	return result.String(), nil
}

// FormatChangelogVerbose provides the original detailed view
func FormatChangelogVerbose(content string) (string, error) {
	config := DefaultColorConfig()

	// Parse markdown
	md := goldmark.New()
	doc := md.Parser().Parse(text.NewReader([]byte(content)))

	var result strings.Builder

	err := ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch node := n.(type) {
			case *ast.Heading:
				level := node.Level
				text := extractText(node, content)
				formatted := formatHeading(text, level, config)
				result.WriteString(formatted)
				result.WriteString("\n\n")

			case *ast.List:
				// Lists are handled by their items

			case *ast.ListItem:
				text := extractText(node, content)
				formatted := formatListItem(text, config)
				result.WriteString(formatted)
				result.WriteString("\n")

			case *ast.Paragraph:
				text := extractText(node, content)
				formatted := formatParagraph(text, config)
				result.WriteString(formatted)
				result.WriteString("\n\n")
			}
		}

		return ast.WalkContinue, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to walk markdown AST: %w", err)
	}

	return result.String(), nil
}

// formatHeading formats a heading with appropriate colors
func formatHeading(text string, level int, config *ColorConfig) string {
	switch level {
	case 1:
		return config.Header.Sprintf("# %s", text)
	case 2:
		if isFeatureHeader(text) {
			return config.Features.Sprintf("## %s", text)
		} else if isFixHeader(text) {
			return config.Fixes.Sprintf("## %s", text)
		} else if isBreakingHeader(text) {
			return config.Breaking.Sprintf("## %s", text)
		}
		return config.Header.Sprintf("## %s", text)
	case 3:
		return config.Header.Sprintf("### %s", text)
	default:
		return config.Body.Sprintf("%s%s %s", strings.Repeat("#", level), " ", text)
	}
}

// formatListItem formats a list item
func formatListItem(text string, config *ColorConfig) string {
	if strings.TrimSpace(text) == "" {
		return ""
	}
	return config.List.Sprintf("  • %s", strings.TrimSpace(text))
}

// formatParagraph formats a paragraph with bold text handling
func formatParagraph(text string, config *ColorConfig) string {
	if strings.TrimSpace(text) == "" {
		return ""
	}

	// Handle **bold** text
	boldRegex := regexp.MustCompile(`\*\*(.*?)\*\*`)
	formatted := boldRegex.ReplaceAllStringFunc(text, func(match string) string {
		content := strings.Trim(match, "*")
		return config.Bold.Sprint(content)
	})

	return config.Body.Sprint(formatted)
}

// FormatSection formats a specific section with color
func FormatSection(title, content string, sectionType string) string {
	config := DefaultColorConfig()

	var titleColor *color.Color
	switch strings.ToLower(sectionType) {
	case sectionTypeFeatures, "feature":
		titleColor = config.Features
	case sectionTypeFixes, "fix", "bug fixes":
		titleColor = config.Fixes
	case sectionTypeBreaking, "breaking changes":
		titleColor = config.Breaking
	default:
		titleColor = config.Header
	}

	result := titleColor.Sprintf("## %s", title) + "\n\n"
	result += config.Body.Sprint(content) + "\n"

	return result
}

// Helper functions

func extractText(node ast.Node, source string) string {
	var buf bytes.Buffer

	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch n := n.(type) {
			case *ast.Text:
				buf.Write(n.Segment.Value([]byte(source)))
			case *ast.String:
				buf.WriteString(string(n.Value))
			}
		}
		return ast.WalkContinue, nil
	})

	return buf.String()
}

func isFeatureHeader(text string) bool {
	lower := strings.ToLower(text)
	return strings.Contains(lower, "feature") || strings.Contains(lower, "added") || strings.Contains(lower, "new")
}

func isFixHeader(text string) bool {
	lower := strings.ToLower(text)
	return strings.Contains(lower, "fix") || strings.Contains(lower, "bug") || strings.Contains(lower, "patch")
}

func isBreakingHeader(text string) bool {
	lower := strings.ToLower(text)
	return strings.Contains(lower, "breaking") || strings.Contains(lower, "breaking changes") || strings.Contains(lower, "incompatible")
}

// Release represents a single release version with its sections
type Release struct {
	Version  string
	Date     string
	Sections []ReleaseSection
}

// ReleaseSection represents a section within a release
type ReleaseSection struct {
	Type  string
	Title string
	Items []string
}

// parseReleases extracts release information from changelog content
func parseReleases(content string) []Release {
	parser := &releaseParser{
		versionRegex: regexp.MustCompile(`^##\s+(.+)$`),
		sectionRegex: regexp.MustCompile(`^###\s+(.+)$`),
		itemRegex:    regexp.MustCompile(`^\s*[•*-]\s+(.+)$`),
	}

	return parser.parse(content)
}

// releaseParser handles parsing of changelog content
type releaseParser struct {
	versionRegex *regexp.Regexp
	sectionRegex *regexp.Regexp
	itemRegex    *regexp.Regexp

	releases       []Release
	currentRelease *Release
	currentSection *ReleaseSection
}

// parse processes the changelog content and returns releases
func (p *releaseParser) parse(content string) []Release {
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		p.processLine(line)
	}

	p.finalizeCurrentRelease()
	return p.releases
}

// processLine handles a single line of changelog content
func (p *releaseParser) processLine(line string) {
	if p.handleVersionHeader(line) {
		return
	}
	if p.handleSectionHeader(line) {
		return
	}
	p.handleListItem(line)
}

// handleVersionHeader processes version header lines
func (p *releaseParser) handleVersionHeader(line string) bool {
	match := p.versionRegex.FindStringSubmatch(line)
	if match == nil {
		return false
	}

	p.finalizeCurrentRelease()

	versionText := match[1]
	version, date := parseVersionAndDate(versionText)
	p.currentRelease = &Release{
		Version:  version,
		Date:     date,
		Sections: []ReleaseSection{},
	}
	p.currentSection = nil
	return true
}

// handleSectionHeader processes section header lines
func (p *releaseParser) handleSectionHeader(line string) bool {
	match := p.sectionRegex.FindStringSubmatch(line)
	if match == nil || p.currentRelease == nil {
		return false
	}

	p.finalizeCurrentSection()

	sectionTitle := match[1]
	p.currentSection = &ReleaseSection{
		Type:  getSectionTypeFromTitle(sectionTitle),
		Title: sectionTitle,
		Items: []string{},
	}
	return true
}

// handleListItem processes list item lines
func (p *releaseParser) handleListItem(line string) {
	match := p.itemRegex.FindStringSubmatch(line)
	if match == nil || p.currentSection == nil {
		return
	}

	item := strings.TrimSpace(match[1])
	if item != "" {
		p.currentSection.Items = append(p.currentSection.Items, item)
	}
}

// finalizeCurrentSection saves the current section to the current release
func (p *releaseParser) finalizeCurrentSection() {
	if p.currentSection != nil && len(p.currentSection.Items) > 0 && p.currentRelease != nil {
		p.currentRelease.Sections = append(p.currentRelease.Sections, *p.currentSection)
	}
}

// finalizeCurrentRelease saves the current release to the releases list
func (p *releaseParser) finalizeCurrentRelease() {
	if p.currentRelease != nil {
		p.finalizeCurrentSection()
		p.releases = append(p.releases, *p.currentRelease)
	}
}

// parseVersionAndDate extracts version and date from version header text
func parseVersionAndDate(versionText string) (string, string) {
	// Handle formats like "v1.2.3 - 2023-01-01" or "Unreleased"
	parts := strings.Split(versionText, " - ")
	if len(parts) >= 2 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(versionText), ""
}

// getSectionTypeFromTitle determines section type from title
func getSectionTypeFromTitle(title string) string {
	lower := strings.ToLower(title)
	if strings.Contains(lower, "feature") || strings.Contains(lower, "added") {
		return sectionTypeFeatures
	}
	if strings.Contains(lower, "fix") || strings.Contains(lower, "bug") {
		return sectionTypeFixes
	}
	if strings.Contains(lower, sectionTypeBreaking) {
		return sectionTypeBreaking
	}
	if strings.Contains(lower, "chore") {
		return sectionTypeChores
	}
	if strings.Contains(lower, "refactor") {
		return sectionTypeRefactoring
	}
	if strings.Contains(lower, "doc") {
		return sectionTypeDocumentation
	}
	return sectionTypeOther
}

// formatVersionHeader creates a formatted version header with emoji and colors
func formatVersionHeader(version, date string, config *ColorConfig) string {
	emoji := "🏷️"
	if strings.Contains(strings.ToLower(version), "unreleased") {
		emoji = "🚧"
	}

	header := fmt.Sprintf("%s %s", emoji, config.Header.Sprint(version))
	if date != "" {
		header += config.Body.Sprintf(" (%s)", date)
	}

	return header
}

// prioritizeSections orders sections by importance for compact display
func prioritizeSections(sections []ReleaseSection) []ReleaseSection {
	priority := map[string]int{
		sectionTypeBreaking:      1,
		sectionTypeFeatures:      2,
		sectionTypeFixes:         3,
		sectionTypeRefactoring:   4,
		sectionTypeDocumentation: 5,
		sectionTypeChores:        6,
		sectionTypeOther:         7,
	}

	// Sort sections by priority
	sorted := make([]ReleaseSection, len(sections))
	copy(sorted, sections)

	// Simple bubble sort by priority
	for i := range len(sorted) {
		for j := i + 1; j < len(sorted); j++ {
			iPriority := priority[sorted[i].Type]
			jPriority := priority[sorted[j].Type]
			if iPriority > jPriority {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	return sorted
}

// formatCompactSection formats a section with emoji and grouped items
func formatCompactSection(section ReleaseSection, config *ColorConfig) string {
	var result strings.Builder

	// Get emoji and color for section type
	emoji := getSectionEmoji(section.Type)
	sectionColor := getSectionColorByType(section.Type, config)

	// Section header with emoji
	result.WriteString(sectionColor.Sprintf("%s %s\n", emoji, section.Title))

	// Format items (limit to most important ones in compact view)
	maxItems := 4
	items := section.Items
	if len(items) > maxItems {
		items = items[:maxItems]
	}

	for _, item := range items {
		// Group similar items and shorten descriptions
		shortItem := shortenItem(item)
		result.WriteString(config.List.Sprintf("  • %s\n", shortItem))
	}

	// Show count if items were truncated
	if len(section.Items) > maxItems {
		remaining := len(section.Items) - maxItems
		result.WriteString(config.Body.Sprintf("  ... and %d more\n", remaining))
	}

	return result.String()
}

// getSectionEmoji returns appropriate emoji for section type
func getSectionEmoji(sectionType string) string {
	switch sectionType {
	case sectionTypeFeatures:
		return "🚀"
	case sectionTypeFixes:
		return "🐛"
	case sectionTypeBreaking:
		return "💥"
	case sectionTypeChores:
		return "🔧"
	case sectionTypeRefactoring:
		return "♻️"
	case sectionTypeDocumentation:
		return "📚"
	default:
		return "✨"
	}
}

// getSectionColorByType returns color based on section type
func getSectionColorByType(sectionType string, config *ColorConfig) *color.Color {
	switch sectionType {
	case sectionTypeFeatures:
		return config.Features
	case sectionTypeFixes:
		return config.Fixes
	case sectionTypeBreaking:
		return config.Breaking
	default:
		return config.Header
	}
}

// shortenItem creates shorter, more readable item descriptions
func shortenItem(item string) string {
	// Remove redundant prefixes
	item = strings.TrimSpace(item)

	// Common patterns to clean up
	patterns := map[string]string{
		"cli: ":                  "",
		"cmd: ":                  "",
		"internal: ":             "",
		"Update krci-ai to the ": "Update to ",
		"version ":               "v",
		"Add support for ":       "Add ",
		"Implement ":             "Add ",
		"Create ":                "Add ",
	}

	for pattern, replacement := range patterns {
		if strings.HasPrefix(item, pattern) {
			item = replacement + item[len(pattern):]
			break
		}
	}

	// Limit length for readability
	if len(item) > 80 {
		item = item[:77] + "..."
	}

	return item
}

// formatFallback provides fallback formatting when parsing fails
func formatFallback(content string, _ *ColorConfig) (string, error) {
	// Just apply basic formatting to raw content
	return FormatChangelogVerbose(content)
}
