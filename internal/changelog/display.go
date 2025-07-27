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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// DisplayOptions configures how changelog content is displayed
type DisplayOptions struct {
	ShowColors    bool
	Paginate      bool
	PageSize      int
	ShowSummary   bool
	FilterSection string
}

// DefaultDisplayOptions returns default display options
func DefaultDisplayOptions() *DisplayOptions {
	return &DisplayOptions{
		ShowColors:    true,
		Paginate:      false,
		PageSize:      20,
		ShowSummary:   false,
		FilterSection: "",
	}
}

// DisplayChangelog displays changelog content with enhanced formatting and options
func DisplayChangelog(content string, options *DisplayOptions) error {
	if options == nil {
		options = DefaultDisplayOptions()
	}

	// Parse and extract sections
	extractor := NewSectionExtractor(content)
	sections, err := extractor.ExtractAllSections()
	if err != nil {
		return fmt.Errorf("failed to extract sections: %w", err)
	}

	// Filter sections if requested
	if options.FilterSection != "" {
		sections = filterSectionsByName(sections, options.FilterSection)
	}

	// Show summary if requested
	if options.ShowSummary {
		displaySummary(sections, options.ShowColors)
		return nil
	}

	// Display content
	if options.Paginate {
		return displayPaginated(sections, options)
	}

	return displayDirect(sections, options)
}

// displaySummary shows a summary of changelog sections
func displaySummary(sections []Section, useColors bool) {
	var titleColor, countColor *color.Color

	if useColors {
		titleColor = color.New(color.FgCyan, color.Bold)
		countColor = color.New(color.FgYellow)
	} else {
		titleColor = color.New()
		countColor = color.New()
	}

	_, _ = titleColor.Println("📊 Changelog Summary")
	_, _ = titleColor.Println("===================")
	fmt.Println()

	totalItems := 0
	for _, section := range sections {
		itemCount := len(section.Items)
		totalItems += itemCount

		if useColors {
			fmt.Printf("  %s: ", getSectionColor(section.Type).Sprint(section.Title))
			_, _ = countColor.Printf("%d items\n", itemCount)
		} else {
			fmt.Printf("  %s: %d items\n", section.Title, itemCount)
		}
	}

	fmt.Println()
	if useColors {
		_, _ = color.New(color.FgGreen, color.Bold).Printf("Total: %d items across %d sections\n", totalItems, len(sections))
	} else {
		fmt.Printf("Total: %d items across %d sections\n", totalItems, len(sections))
	}
}

// displayDirect displays content directly without pagination
func displayDirect(sections []Section, options *DisplayOptions) error {
	for i, section := range sections {
		if i > 0 {
			fmt.Println()
		}

		displaySection(section, options.ShowColors)
	}

	return nil
}

// displayPaginated displays content with pagination
func displayPaginated(sections []Section, options *DisplayOptions) error {
	totalSections := len(sections)
	if totalSections == 0 {
		color.Yellow("No sections to display")
		return nil
	}

	currentPage := 0
	sectionsPerPage := options.PageSize

	for {
		start := currentPage * sectionsPerPage
		end := start + sectionsPerPage

		if end > totalSections {
			end = totalSections
		}

		// Display current page
		for i := start; i < end; i++ {
			if i > start {
				fmt.Println()
			}
			displaySection(sections[i], options.ShowColors)
		}

		// Show pagination info
		if options.ShowColors {
			color.Cyan("\n--- Page %d of %d (sections %d-%d of %d) ---",
				currentPage+1,
				(totalSections+sectionsPerPage-1)/sectionsPerPage,
				start+1, end, totalSections)
		} else {
			fmt.Printf("\n--- Page %d of %d (sections %d-%d of %d) ---\n",
				currentPage+1,
				(totalSections+sectionsPerPage-1)/sectionsPerPage,
				start+1, end, totalSections)
		}

		// Check if we're at the end
		if end >= totalSections {
			break
		}

		// Ask for continuation
		fmt.Print("Press Enter to continue, 'q' to quit: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		if strings.TrimSpace(strings.ToLower(input)) == "q" {
			break
		}

		currentPage++
	}

	return nil
}

// displaySection displays a single section with formatting
func displaySection(section Section, useColors bool) {
	sectionColor := getSectionColor(section.Type)

	// Display section title
	if useColors {
		_, _ = sectionColor.Printf("### %s\n", section.Title)
	} else {
		fmt.Printf("### %s\n", section.Title)
	}

	fmt.Println()

	// Display items
	for _, item := range section.Items {
		if useColors {
			_, _ = color.New(color.FgWhite).Printf("  • %s\n", item)
		} else {
			fmt.Printf("  • %s\n", item)
		}
	}
}

// getSectionColor returns appropriate color for section type
func getSectionColor(sectionType SectionType) *color.Color {
	switch sectionType {
	case SectionFeatures:
		return color.New(color.FgGreen, color.Bold)
	case SectionBugFixes:
		return color.New(color.FgYellow, color.Bold)
	case SectionBreaking:
		return color.New(color.FgRed, color.Bold)
	case SectionDocumentation:
		return color.New(color.FgBlue, color.Bold)
	case SectionChores:
		return color.New(color.FgMagenta, color.Bold)
	case SectionPerformance:
		return color.New(color.FgCyan, color.Bold)
	case SectionRefactoring:
		return color.New(color.FgHiBlue, color.Bold)
	case SectionTests:
		return color.New(color.FgHiGreen, color.Bold)
	case SectionStyles:
		return color.New(color.FgHiMagenta, color.Bold)
	default:
		return color.New(color.FgWhite, color.Bold)
	}
}

// filterSectionsByName filters sections by name pattern
func filterSectionsByName(sections []Section, pattern string) []Section {
	var filtered []Section
	pattern = strings.ToLower(pattern)

	for _, section := range sections {
		if strings.Contains(strings.ToLower(section.Title), pattern) {
			filtered = append(filtered, section)
		}
	}

	return filtered
}

// DisplayCompactChangelog displays a compact version of the changelog
func DisplayCompactChangelog(content string) error {
	extractor := NewSectionExtractor(content)
	sections, err := extractor.ExtractAllSections()
	if err != nil {
		return fmt.Errorf("failed to extract sections: %w", err)
	}

	color.Cyan("📋 Changelog (Compact View)")
	color.Cyan("===========================")
	fmt.Println()

	for _, section := range sections {
		sectionColor := getSectionColor(section.Type)
		_, _ = sectionColor.Printf("▸ %s (%d items)\n", section.Title, len(section.Items))

		// Show only first 2 items for compact view
		for i, item := range section.Items {
			if i >= 2 {
				_, _ = color.New(color.FgHiBlack).Printf("    ... and %d more\n", len(section.Items)-2)
				break
			}
			fmt.Printf("    • %s\n", item)
		}
		fmt.Println()
	}

	return nil
}
