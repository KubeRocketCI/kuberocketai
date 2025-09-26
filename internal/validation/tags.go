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
package validation

import (
	"fmt"
	"regexp"
)

// xmlTag represents a parsed XML tag with its position
type xmlTag struct {
	Name        string
	IsClosing   bool
	IsSelfClose bool
	Position    int
}

// validateXMLTags validates XML-like tags in content
func (a *FrameworkAnalyzer) validateXMLTags(content string) []string {
	var issues []string
	tags := a.parseXMLTags(content)
	stack := []xmlTag{}

	for _, tag := range tags {
		if tag.IsSelfClose {
			continue
		}

		if tag.IsClosing {
			if len(stack) == 0 {
				issues = append(issues, fmt.Sprintf("Closing tag </%s> without matching opening tag", tag.Name))
				continue
			}

			// Find matching opening tag
			found := false
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i].Name == tag.Name {
					// Report unclosed nested tags before removing them
					for j := len(stack) - 1; j > i; j-- {
						issues = append(issues, fmt.Sprintf("Unclosed tag <%s>", stack[j].Name))
					}
					// Remove all tags from this position to end (handle nested tags)
					stack = stack[:i]
					found = true
					break
				}
			}

			if !found {
				issues = append(issues, fmt.Sprintf("Closing tag </%s> without matching opening tag", tag.Name))
			}
		} else {
			stack = append(stack, tag)
		}
	}

	// Check for unclosed tags
	for _, tag := range stack {
		issues = append(issues, fmt.Sprintf("Unclosed tag <%s>", tag.Name))
	}

	return issues
}

// parseXMLTags extracts XML-like tags from content, excluding those in code blocks
func (a *FrameworkAnalyzer) parseXMLTags(content string) []xmlTag {
	// First, identify code block ranges to exclude
	codeBlockRanges := a.findCodeBlockRanges(content)

	// Match XML tags including those with attributes
	// Pattern explanation:
	// <\s*           - opening bracket with optional whitespace
	// (/?)           - optional closing slash (group 1)
	// ([a-zA-Z_][\w\-:]*) - tag name (group 2): starts with letter/underscore, followed by word chars, hyphens, colons
	// [^>]*          - any attributes (non-greedy)
	// (/?)           - optional self-closing slash (group 3)
	// \s*>           - closing bracket with optional whitespace
	tagPattern := regexp.MustCompile(`<\s*(/?)([a-zA-Z_][\w\-:]*)[^>]*?(/?)\s*>`)
	matches := tagPattern.FindAllStringSubmatchIndex(content, -1)

	var tags []xmlTag
	for _, match := range matches {
		if len(match) >= 8 {
			tagStart := match[0]
			tagEnd := match[1]

			// Check if this tag is inside any code block
			if a.isPositionInCodeBlock(tagStart, tagEnd, codeBlockRanges) {
				continue // Skip tags inside code blocks
			}

			tagText := content[tagStart:tagEnd]
			nameMatch := tagPattern.FindStringSubmatch(tagText)
			if len(nameMatch) >= 4 {
				tag := xmlTag{
					Name:        nameMatch[2],
					IsClosing:   nameMatch[1] == "/",
					IsSelfClose: nameMatch[3] == "/",
					Position:    tagStart,
				}
				tags = append(tags, tag)
			}
		}
	}

	return tags
}

// CodeBlockRange represents a range of content that is inside a code block
type CodeBlockRange struct {
	Start int
	End   int
}

// findCodeBlockRanges identifies all code block ranges in the content
func (a *FrameworkAnalyzer) findCodeBlockRanges(content string) []CodeBlockRange {
	var ranges []CodeBlockRange

	// Find fenced code blocks first (``` and ~~~) - these take priority over inline
	// Handle ``` blocks
	backtickCodePattern := regexp.MustCompile(`(?m)^` + "`{3}" + `[^\n]*\n[\s\S]*?^` + "`{3}" + `\s*$`)
	backtickMatches := backtickCodePattern.FindAllStringIndex(content, -1)
	for _, match := range backtickMatches {
		ranges = append(ranges, CodeBlockRange{Start: match[0], End: match[1]})
	}

	// Handle ~~~ blocks
	tildeCodePattern := regexp.MustCompile(`(?m)^~{3}[^\n]*\n[\s\S]*?^~{3}\s*$`)
	tildeMatches := tildeCodePattern.FindAllStringIndex(content, -1)
	for _, match := range tildeMatches {
		ranges = append(ranges, CodeBlockRange{Start: match[0], End: match[1]})
	}

	// Find inline code blocks (backticks) - but skip any that overlap with fenced blocks
	inlineCodePattern := regexp.MustCompile("`[^`]*`")
	inlineMatches := inlineCodePattern.FindAllStringIndex(content, -1)
	for _, match := range inlineMatches {
		// Check if this inline code block overlaps with any fenced block
		overlaps := false
		for _, existing := range ranges {
			if match[0] >= existing.Start && match[1] <= existing.End {
				overlaps = true
				break
			}
		}
		if !overlaps {
			ranges = append(ranges, CodeBlockRange{Start: match[0], End: match[1]})
		}
	}

	return ranges
}

// isPositionInCodeBlock checks if a position range overlaps with any code block
func (a *FrameworkAnalyzer) isPositionInCodeBlock(start, end int, codeBlockRanges []CodeBlockRange) bool {
	for _, codeRange := range codeBlockRanges {
		// Check if the tag position overlaps with the code block range
		if start >= codeRange.Start && end <= codeRange.End {
			return true
		}
	}
	return false
}
