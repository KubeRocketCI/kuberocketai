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
	case "features", "feature":
		titleColor = config.Features
	case "fixes", "fix", "bug fixes":
		titleColor = config.Fixes
	case "breaking", "breaking changes":
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
