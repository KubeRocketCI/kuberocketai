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
package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/KubeRocketCI/kuberocketai/internal/cli/style"
	"github.com/KubeRocketCI/kuberocketai/internal/validation"
)

// Table styling constants
const (
	TableCyanColor      = "14"
	TableGrayColor      = "245"
	TableLightGrayColor = "241"
	TableCellWidth      = 20
	DescriptionMaxLen   = 50
	DescriptionTruncLen = 47
	NoneValue           = "None"
)

// OutputHandler provides colorized output functions for CLI commands
type OutputHandler struct{}

// NewOutputHandler creates a new output handler with color support
func NewOutputHandler() *OutputHandler { return &OutputHandler{} }

// PrintSuccess prints a success message with green color
func (o *OutputHandler) PrintSuccess(message string) {
	fmt.Printf("%s %s\n", style.Success("✅"), message)
}

// PrintInfo prints an info message with blue color
func (o *OutputHandler) PrintInfo(message string) {
	fmt.Printf("%s %s\n", style.Info("ℹ️"), message)
}

// PrintProgress prints a progress message with cyan color
func (o *OutputHandler) PrintProgress(message string) {
	fmt.Printf("%s %s\n", style.Progress("🔄"), message)
}

// PrintWarning prints a warning message with yellow color
func (o *OutputHandler) PrintWarning(message string) {
	fmt.Printf("%s %s\n", style.Warn("⚠️"), message)
}

// PrintError prints an error message with red color
func (o *OutputHandler) PrintError(message string) {
	fmt.Printf("%s %s\n", style.Error("❌"), message)
}

// PrintBold prints text in bold
func (o *OutputHandler) PrintBold(message string) {
	fmt.Println(style.Bold(message))
}

// Bold returns text formatted in bold
func (o *OutputHandler) Bold(text string) string {
	return style.Bold(text)
}

// PrintCyan prints text in cyan color
func (o *OutputHandler) PrintCyan(text string) string {
	return style.Cyan(text)
}

// PrintYellow prints text in yellow color
func (o *OutputHandler) PrintYellow(text string) string {
	return style.Yellow(text)
}

// PrintMagenta prints text in magenta color
func (o *OutputHandler) PrintMagenta(text string) string {
	return style.Magenta(text)
}

// PrintBlue prints text in blue color
func (o *OutputHandler) PrintBlue(text string) string {
	return style.Blue(text)
}

// PrintGreenBold prints text in green bold
func (o *OutputHandler) PrintGreenBold(text string) string {
	return style.GreenBold(text)
}

// Printf provides formatted printing with style support
func (o *OutputHandler) Printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

// Newline prints a newline character
func (o *OutputHandler) Newline() {
	fmt.Println()
}

// CreateStyledTable creates a lipgloss table with consistent styling
func CreateStyledTable() *table.Table {
	var (
		cyan      = lipgloss.Color(TableCyanColor)
		gray      = lipgloss.Color(TableGrayColor)
		lightGray = lipgloss.Color(TableLightGrayColor)

		headerStyle  = lipgloss.NewStyle().Foreground(cyan).Bold(true).Align(lipgloss.Center)
		cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(TableCellWidth)
		oddRowStyle  = cellStyle.Foreground(gray)
		evenRowStyle = cellStyle.Foreground(lightGray)
	)

	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(cyan)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		})
}

// TruncateDescription truncates description if it's too long
func TruncateDescription(description string) string {
	if len(description) > DescriptionMaxLen {
		return description[:DescriptionTruncLen] + "..."
	}
	return description
}

// PrintFrameworkInsights prints comprehensive framework insights with styled output
func (o *OutputHandler) PrintFrameworkInsights(insights *validation.FrameworkInsights, issueCount int) {
	if insights == nil {
		o.PrintError("No framework insights available")
		return
	}

	// Print overview
	o.Printf("Overview: %s, %s, %s, %s\n",
		o.PrintCyan(fmt.Sprintf("%d agents", insights.TotalAgents)),
		o.PrintCyan(fmt.Sprintf("%d tasks", insights.TotalTasks)),
		o.PrintCyan(fmt.Sprintf("%d templates", insights.TotalTemplates)),
		o.PrintCyan(fmt.Sprintf("%d data files", insights.TotalDataFiles)))

	// Print link resolution status
	if issueCount == 0 {
		o.Printf("All internal links resolved (%s checked)\n",
			o.PrintCyan(fmt.Sprintf("%d references", insights.TotalReferences)))
	} else {
		o.Printf("Found %s out of %s checked\n",
			o.PrintYellow(fmt.Sprintf("%d broken links", issueCount)),
			o.PrintCyan(fmt.Sprintf("%d references", insights.TotalReferences)))
	}

	o.Newline()

	// Print agent insights table
	o.PrintBold("FRAMEWORK INSIGHTS:")
	o.Newline()

	if len(insights.AgentStats) > 0 {
		agentTable := CreateStyledTable()
		agentTable.Headers("Agent", "Tasks", "Templates", "Data Files")

		for _, agent := range insights.AgentStats {
			agentTable.Row(
				agent.Name,
				fmt.Sprintf("%d tasks", agent.TaskCount),
				fmt.Sprintf("%d templates", agent.TemplateCount),
				fmt.Sprintf("%d data files", agent.DataFileCount),
			)
		}

		fmt.Println(agentTable.Render())
		o.Newline()
	}

	// Print most used components
	if insights.MostUsedTemplate != nil {
		o.Printf("Most used template: %s (%s)\n",
			o.PrintYellow(insights.MostUsedTemplate.Path),
			o.PrintCyan(fmt.Sprintf("%d references", insights.MostUsedTemplate.Count)))
	}

	if insights.MostUsedTask != nil {
		o.Printf("Most used task: %s (%s)\n",
			o.PrintYellow(insights.MostUsedTask.Path),
			o.PrintCyan(fmt.Sprintf("%d references", insights.MostUsedTask.Count)))
	}

	if insights.MostUsedDataFile != nil {
		o.Printf("Most used data file: %s (%s)\n",
			o.PrintYellow(insights.MostUsedDataFile.Path),
			o.PrintCyan(fmt.Sprintf("%d references", insights.MostUsedDataFile.Count)))
	}
}
