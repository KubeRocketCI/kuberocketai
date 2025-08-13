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
package assets

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"

	"github.com/KubeRocketCI/kuberocketai/internal/validation"
)

const (
	// DefaultAgentIcon is the placeholder icon used when an agent doesn't define an icon
	DefaultAgentIcon = "🤖"
)

// AgentInfo represents basic information about an agent
type AgentInfo struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Role        string `yaml:"role"`
	Goal        string `yaml:"goal"`
	Icon        string `yaml:"icon"`
	FilePath    string `yaml:"-"` // Not from YAML, computed
	ShortName   string `yaml:"-"` // Not from YAML, computed
}

// AgentDependencyInfo extends AgentInfo with dependency information
type AgentDependencyInfo struct {
	AgentInfo
	Tasks     []string `json:"tasks"`
	Templates []string `json:"templates"`
	DataFiles []string `json:"data_files"`
}

// AgentIdentity represents the identity section of an agent YAML
type AgentIdentity struct {
	Name        string `yaml:"name"`
	ID          string `yaml:"id"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Role        string `yaml:"role"`
	Goal        string `yaml:"goal"`
	Icon        string `yaml:"icon"`
}

// Agent represents the structure of an agent YAML file
type Agent struct {
	Agent struct {
		Identity AgentIdentity `yaml:"identity"`
	} `yaml:"agent"`
}

// Discovery handles discovery and parsing of framework assets
type Discovery struct {
	installer *Installer
}

// NewDiscovery creates a new asset discovery service
func NewDiscovery(targetDir string, embeddedAssets embed.FS) *Discovery {
	return &Discovery{
		installer: NewInstaller(targetDir, embeddedAssets),
	}
}

// DiscoverAgents discovers and returns information about all installed agents
func (d *Discovery) DiscoverAgents() ([]AgentInfo, error) {
	if !d.installer.IsInstalled() {
		return nil, fmt.Errorf("framework not installed in current directory - run 'krci-ai install'")
	}

	agentsPath := d.installer.GetAgentsPath()
	agentFiles, err := filepath.Glob(filepath.Join(agentsPath, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to scan agents directory: %w", err)
	}

	if len(agentFiles) == 0 {
		return nil, fmt.Errorf("no agent files found in %s", agentsPath)
	}

	var agents []AgentInfo
	for _, file := range agentFiles {
		agentInfo, err := d.parseAgentFile(file)
		if err != nil {
			// Log warning but continue with other agents
			fmt.Fprintf(os.Stderr, "Warning: failed to parse agent file %s: %v\n", file, err)
			continue
		}

		agentInfo.FilePath = file
		agentInfo.ShortName = strings.TrimSuffix(filepath.Base(file), ".yaml")

		agents = append(agents, *agentInfo)
	}

	return agents, nil
}

// parseAgentFile parses an agent YAML file and extracts basic information
func (d *Discovery) parseAgentFile(filePath string) (*AgentInfo, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var agent Agent
	if err := yaml.Unmarshal(data, &agent); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Extract agent information
	identity := agent.Agent.Identity
	agentInfo := &AgentInfo{
		Name:        identity.Name,
		Description: identity.Description,
		Role:        identity.Role,
		Goal:        identity.Goal,
		Icon:        identity.Icon,
	}

	// Validate required fields
	if agentInfo.Name == "" {
		return nil, fmt.Errorf("agent name is required")
	}
	if agentInfo.Role == "" {
		return nil, fmt.Errorf("agent role is required")
	}

	return agentInfo, nil
}

// GetAgentByName returns information about a specific agent by name
func (d *Discovery) GetAgentByName(name string) (*AgentInfo, error) {
	agents, err := d.DiscoverAgents()
	if err != nil {
		return nil, err
	}

	for _, agent := range agents {
		if agent.Name == name {
			return &agent, nil
		}
	}

	return nil, fmt.Errorf("agent '%s' not found", name)
}

// GetAgentByShortName returns information about a specific agent by short name
func (d *Discovery) GetAgentByShortName(shortName string) (*AgentInfo, error) {
	agents, err := d.DiscoverAgents()
	if err != nil {
		return nil, err
	}

	for _, agent := range agents {
		if agent.ShortName == shortName {
			return &agent, nil
		}
	}

	return nil, fmt.Errorf("agent '%s' not found", shortName)
}

// ListAvailableAgents returns a simple list of agent names
func (d *Discovery) ListAvailableAgents() ([]string, error) {
	agents, err := d.DiscoverAgents()
	if err != nil {
		return nil, err
	}

	var names []string
	for _, agent := range agents {
		names = append(names, agent.Name)
	}

	return names, nil
}

// FormatAgentSummary returns a formatted string summarizing agent information
func (d *Discovery) FormatAgentSummary(agent AgentInfo) string {
	return fmt.Sprintf("%-15s | %-25s | %s", agent.ShortName, agent.Role, agent.Description)
}

// ValidateAgentStructure performs basic validation of agent file structure
func (d *Discovery) ValidateAgentStructure(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var agent Agent
	if err := yaml.Unmarshal(data, &agent); err != nil {
		return fmt.Errorf("invalid YAML structure: %w", err)
	}

	identity := agent.Agent.Identity

	// Check required fields
	if identity.Name == "" {
		return fmt.Errorf("agent.identity.name is required")
	}
	if identity.ID == "" {
		return fmt.Errorf("agent.identity.id is required")
	}
	if identity.Version == "" {
		return fmt.Errorf("agent.identity.version is required")
	}
	if identity.Description == "" {
		return fmt.Errorf("agent.identity.description is required")
	}
	if identity.Role == "" {
		return fmt.Errorf("agent.identity.role is required")
	}
	if identity.Goal == "" {
		return fmt.Errorf("agent.identity.goal is required")
	}

	return nil
}

// DiscoverAgentsWithDependencies discovers and returns agents with their complete dependency information
func (d *Discovery) DiscoverAgentsWithDependencies() ([]AgentDependencyInfo, error) {
	if !d.installer.IsInstalled() {
		return nil, fmt.Errorf("framework not installed in current directory - run 'krci-ai install'")
	}

	// Create validation analyzer for dependency resolution
	analyzer := validation.NewFrameworkAnalyzer(d.installer.targetDir)

	// Generate insights which include complete dependency relationships
	_, insights, err := analyzer.AnalyzeFramework()
	if err != nil {
		return nil, fmt.Errorf("failed to analyze dependencies: %w", err)
	}

	// Get basic agent information
	agents, err := d.DiscoverAgents()
	if err != nil {
		return nil, err
	}

	// Create map of agent relationships by agent name for quick lookup
	relationshipMap := make(map[string]validation.ComponentRelationship)
	for _, rel := range insights.Relationships {
		relationshipMap[rel.Agent] = rel
	}

	// Build dependency info for each agent
	var agentDeps []AgentDependencyInfo
	for _, agent := range agents {
		// Extract agent name from file path (remove .yaml extension)
		agentFileName := strings.TrimSuffix(filepath.Base(agent.FilePath), ".yaml")

		depInfo := AgentDependencyInfo{
			AgentInfo: agent,
			Tasks:     make([]string, 0),
			Templates: make([]string, 0),
			DataFiles: make([]string, 0),
		}

		// Look up dependency relationships
		if rel, exists := relationshipMap[agentFileName]; exists {
			depInfo.Tasks = rel.Tasks
			depInfo.Templates = rel.Templates
			depInfo.DataFiles = rel.DataFiles
		}

		agentDeps = append(agentDeps, depInfo)
	}

	return agentDeps, nil
}

func (d *Discovery) DiscoverAgentWithDependencies(shortName string) (AgentDependencyInfo, error) {
	agents, err := d.DiscoverAgentsWithDependencies()
	if err != nil {
		return AgentDependencyInfo{}, err
	}

	for _, agent := range agents {
		if agent.ShortName == shortName {
			return agent, nil
		}
	}

	return AgentDependencyInfo{}, fmt.Errorf("agent '%s' not found", shortName)
}

// FormatAgentDependencyTable formats agent dependency information as a table
func (d *Discovery) FormatAgentDependencyTable(agents []AgentDependencyInfo) string {
	if len(agents) == 0 {
		return "No agents found"
	}

	var result strings.Builder

	for i, agent := range agents {
		if i > 0 {
			result.WriteString("\n") // Add spacing between agents
		}

		// Format agent header with icon and colors
		agentTable := d.formatAgentTable(agent)
		result.WriteString(agentTable)
	}

	return result.String()
}

// formatAgentTable creates a professional table for a single agent with colors and icons
func (d *Discovery) formatAgentTable(agent AgentDependencyInfo) string {
	var result strings.Builder

	// Add agent header
	d.addAgentHeader(&result, agent)

	// If no tasks, show simple message
	if len(agent.Tasks) == 0 {
		magenta := color.New(color.FgMagenta).SprintFunc()
		result.WriteString(fmt.Sprintf("   %s No tasks defined\n\n", magenta("ℹ️")))
		return result.String()
	}

	// Add dependency table
	d.addDependencyTable(&result, agent)

	return result.String()
}

// addAgentHeader adds the agent header section to the result
func (d *Discovery) addAgentHeader(result *strings.Builder, agent AgentDependencyInfo) {
	// Create color functions
	bold := color.New(color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	// Agent header with icon and styling
	icon := agent.Icon
	if icon == "" {
		icon = DefaultAgentIcon
	}

	// Calculate dynamic width and create header
	agentLine := fmt.Sprintf("%s %s", icon, agent.Name)
	descLine := fmt.Sprintf("Description: %s", agent.Description)
	headerWidth := d.calculateHeaderWidth(agentLine, descLine)

	// Create dynamic borders
	topBorder := "┌" + strings.Repeat("─", headerWidth-2) + "┐"
	bottomBorder := "└" + strings.Repeat("─", headerWidth-2) + "┘"

	result.WriteString(fmt.Sprintf("\n%s\n", cyan(topBorder)))

	// Add agent name line
	d.addAgentNameLine(result, agentLine, headerWidth, bold, cyan)

	// Add description lines
	d.addDescriptionLines(result, descLine, headerWidth, cyan, yellow)

	result.WriteString(fmt.Sprintf("%s\n\n", cyan(bottomBorder)))
}

// calculateHeaderWidth determines the optimal header width
func (d *Discovery) calculateHeaderWidth(agentLine, descLine string) int {
	headerWidth := 60
	if len(agentLine) > headerWidth-4 {
		headerWidth = len(agentLine) + 6
	}
	if len(descLine) > headerWidth-4 {
		headerWidth = len(descLine) + 6
	}
	// Cap at reasonable maximum to avoid extremely wide tables
	if headerWidth > 120 {
		headerWidth = 120
	}
	return headerWidth
}

// addAgentNameLine adds the agent name line to the header
func (d *Discovery) addAgentNameLine(result *strings.Builder, agentLine string, headerWidth int, bold, cyan func(...interface{}) string) {
	namePadding := headerWidth - len(agentLine) - 4
	if namePadding < 0 {
		namePadding = 0
	}
	result.WriteString(fmt.Sprintf("%s %s%s %s\n",
		cyan("│"),
		bold(agentLine),
		strings.Repeat(" ", namePadding),
		cyan("│")))
}

// addDescriptionLines adds description lines with wrapping if needed
func (d *Discovery) addDescriptionLines(result *strings.Builder, descLine string, headerWidth int, cyan, yellow func(...interface{}) string) {
	if len(descLine) > headerWidth-4 {
		// Wrap long descriptions
		maxDescWidth := headerWidth - 4
		wrappedDesc := d.wrapText(descLine, maxDescWidth)
		for _, line := range strings.Split(wrappedDesc, "\n") {
			descPadding := headerWidth - len(line) - 4
			if descPadding < 0 {
				descPadding = 0
			}
			result.WriteString(fmt.Sprintf("%s %s%s %s\n",
				cyan("│"),
				yellow(line),
				strings.Repeat(" ", descPadding),
				cyan("│")))
		}
	} else {
		descPadding := headerWidth - len(descLine) - 4
		if descPadding < 0 {
			descPadding = 0
		}
		result.WriteString(fmt.Sprintf("%s %s%s %s\n",
			cyan("│"),
			yellow(descLine),
			strings.Repeat(" ", descPadding),
			cyan("│")))
	}
}

// addDependencyTable adds the dependency table section
func (d *Discovery) addDependencyTable(result *strings.Builder, agent AgentDependencyInfo) {
	// Create color functions
	bold := color.New(color.Bold).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	// Calculate column widths and dependencies
	taskDependencies := d.getTaskDependenciesFromValidation(agent)
	maxTaskWidth, maxTemplateWidth, maxDataWidth := d.calculateColumnWidths(agent, taskDependencies)

	// Create table format
	format := fmt.Sprintf("│ %%-%ds │ %%-%ds │ %%-%ds │\n",
		maxTaskWidth, maxTemplateWidth, maxDataWidth)

	// Add table header
	d.addTableHeader(result, maxTaskWidth, maxTemplateWidth, maxDataWidth, bold, green, blue, magenta)

	// Add table rows
	d.addTableRows(result, agent, taskDependencies, format, maxTaskWidth, maxTemplateWidth, maxDataWidth)

	// Add table footer
	d.addTableFooter(result, maxTaskWidth, maxTemplateWidth, maxDataWidth)
}

// calculateColumnWidths determines optimal column widths for the table
func (d *Discovery) calculateColumnWidths(agent AgentDependencyInfo, taskDependencies map[string]taskDependency) (int, int, int) {
	maxTaskWidth := len("Task")
	maxTemplateWidth := len("Templates")
	maxDataWidth := len("Data")

	for _, task := range agent.Tasks {
		if len(task) > maxTaskWidth {
			maxTaskWidth = len(task)
		}

		if deps, exists := taskDependencies[task]; exists {
			for _, template := range deps.templates {
				if len(template) > maxTemplateWidth {
					maxTemplateWidth = len(template)
				}
			}
			for _, dataFile := range deps.dataFiles {
				if len(dataFile) > maxDataWidth {
					maxDataWidth = len(dataFile)
				}
			}
		}
	}

	// Set reasonable minimum and maximum column widths
	maxTaskWidth = d.constrainWidth(maxTaskWidth, 20, 40)
	maxTemplateWidth = d.constrainWidth(maxTemplateWidth, 25, 50)
	maxDataWidth = d.constrainWidth(maxDataWidth, 25, 50)

	return maxTaskWidth, maxTemplateWidth, maxDataWidth
}

// constrainWidth ensures width is within reasonable bounds
func (d *Discovery) constrainWidth(width, min, max int) int {
	if width < min {
		return min
	}
	if width > max {
		return max
	}
	return width
}

// addTableHeader adds the table header section
func (d *Discovery) addTableHeader(result *strings.Builder, maxTaskWidth, maxTemplateWidth, maxDataWidth int, bold, green, blue, magenta func(...interface{}) string) {
	format := fmt.Sprintf("│ %%-%ds │ %%-%ds │ %%-%ds │\n",
		maxTaskWidth, maxTemplateWidth, maxDataWidth)

	result.WriteString(fmt.Sprintf("┌%s┬%s┬%s┐\n",
		strings.Repeat("─", maxTaskWidth+2),
		strings.Repeat("─", maxTemplateWidth+2),
		strings.Repeat("─", maxDataWidth+2)))

	result.WriteString(fmt.Sprintf(format,
		bold(green("Task")),
		bold(blue("Templates")),
		bold(magenta("Data"))))

	result.WriteString(fmt.Sprintf("├%s┼%s┼%s┤\n",
		strings.Repeat("─", maxTaskWidth+2),
		strings.Repeat("─", maxTemplateWidth+2),
		strings.Repeat("─", maxDataWidth+2)))
}

// addTableRows adds all table rows for tasks and dependencies
func (d *Discovery) addTableRows(result *strings.Builder, agent AgentDependencyInfo, taskDependencies map[string]taskDependency, format string, maxTaskWidth, maxTemplateWidth, maxDataWidth int) {
	for _, task := range agent.Tasks {
		var templates []string
		var dataFiles []string
		var fullPath string

		taskFileName := filepath.Base(task)
		if deps, exists := taskDependencies[taskFileName]; exists {
			templates = deps.templates
			dataFiles = deps.dataFiles
			fullPath = deps.fullPath
		}

		d.addTaskRows(result, taskFileName, fullPath, templates, dataFiles, format, maxTaskWidth, maxTemplateWidth, maxDataWidth)
	}
}

// addTaskRows adds rows for a single task and its dependencies
func (d *Discovery) addTaskRows(result *strings.Builder, task string, fullPath string, templates, dataFiles []string, format string, maxTaskWidth, maxTemplateWidth, maxDataWidth int) {
	maxRows := 1
	if len(templates) > maxRows {
		maxRows = len(templates)
	}
	if len(dataFiles) > maxRows {
		maxRows = len(dataFiles)
	}

	for row := 0; row < maxRows; row++ {
		taskDisplay := ""
		templatesDisplay := "-"
		dataDisplay := "-"

		// Only show task name on the first row
		if row == 0 {
			// Use fullPath if available, otherwise fall back to task name
			pathToUse := fullPath
			if pathToUse == "" {
				pathToUse = task
			}
			formattedTask := d.formatTaskDisplayName(pathToUse)
			taskDisplay = d.fitString(formattedTask, maxTaskWidth)
		}

		// Show template for this row
		if row < len(templates) {
			templatesDisplay = d.fitString(templates[row], maxTemplateWidth)
		}

		// Show data file for this row
		if row < len(dataFiles) {
			dataDisplay = d.fitString(dataFiles[row], maxDataWidth)
		}

		result.WriteString(fmt.Sprintf(format, taskDisplay, templatesDisplay, dataDisplay))
	}
}

// addTableFooter adds the table footer
func (d *Discovery) addTableFooter(result *strings.Builder, maxTaskWidth, maxTemplateWidth, maxDataWidth int) {
	result.WriteString(fmt.Sprintf("└%s┴%s┴%s┘\n",
		strings.Repeat("─", maxTaskWidth+2),
		strings.Repeat("─", maxTemplateWidth+2),
		strings.Repeat("─", maxDataWidth+2)))
}

// truncateString truncates a string to maxWidth with ellipsis if needed
func (d *Discovery) truncateString(str string, maxWidth int) string {
	if len(str) <= maxWidth {
		return str
	}
	if maxWidth <= 3 {
		return "..."
	}
	return str[:maxWidth-3] + "..."
}

// wrapText wraps text to fit within the specified width
func (d *Discovery) wrapText(text string, width int) string {
	if len(text) <= width {
		return text
	}

	var result strings.Builder
	words := strings.Fields(text)
	var currentLine strings.Builder

	for _, word := range words {
		// If adding this word would exceed the width, start a new line
		if currentLine.Len() > 0 && currentLine.Len()+len(word)+1 > width {
			result.WriteString(currentLine.String())
			result.WriteString("\n")
			currentLine.Reset()
		}

		if currentLine.Len() > 0 {
			currentLine.WriteString(" ")
		}
		currentLine.WriteString(word)
	}

	// Add the last line
	if currentLine.Len() > 0 {
		result.WriteString(currentLine.String())
	}

	return result.String()
}

// fitString fits a string within the specified width, using smart truncation when needed
func (d *Discovery) fitString(str string, maxWidth int) string {
	if len(str) <= maxWidth {
		return str
	}

	// For file names, try to show the most relevant part (usually the end)
	if strings.Contains(str, ".") {
		// For files with extensions, keep the extension visible
		if maxWidth <= 7 { // Not enough space for meaningful truncation
			return d.truncateString(str, maxWidth)
		}

		// Show beginning + "..." + end to preserve both context and extension
		prefixLen := (maxWidth - 3) / 2
		suffixLen := maxWidth - 3 - prefixLen

		if suffixLen > 0 && prefixLen > 0 {
			return str[:prefixLen] + "..." + str[len(str)-suffixLen:]
		}
	}

	// Default truncation for other strings
	return d.truncateString(str, maxWidth)
}

// taskDependency holds template and data dependencies for a task
type taskDependency struct {
	templates []string
	dataFiles []string
	fullPath  string // Store full path for local task detection
}

// getTaskDependenciesFromValidation uses the validation system to get task dependencies
func (d *Discovery) getTaskDependenciesFromValidation(agent AgentDependencyInfo) map[string]taskDependency {
	dependencies := make(map[string]taskDependency)

	// Simply use the templates and data that are already provided by the validation system
	// We'll create a single dependency entry that contains all templates and data for the agent
	// Since we show tasks individually, we'll distribute them evenly or show all for each task

	for _, task := range agent.Tasks {
		taskFileName := filepath.Base(task)
		dependencies[taskFileName] = taskDependency{
			templates: agent.Templates,
			dataFiles: agent.DataFiles,
			fullPath:  task,
		}
	}

	return dependencies
}

// formatTaskDisplayName formats task name with visual distinction for local tasks
func (d *Discovery) formatTaskDisplayName(taskPath string) string {
	taskName := filepath.Base(taskPath)
	// Check multiple patterns for local task paths
	if strings.Contains(taskPath, "/local/tasks/") || strings.Contains(taskPath, ".krci-ai/local/tasks/") {
		return "📁 " + taskName // Local task indicator
	}
	return taskName
}
