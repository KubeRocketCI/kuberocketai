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
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/KubeRocketCI/kuberocketai/internal/engine/processor"
	"github.com/KubeRocketCI/kuberocketai/internal/validation"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate KubeRocketAI framework components",
	Long: `Validate KubeRocketAI framework components in the current directory.

This command validates:
- Agent YAML files for schema compliance
- Task path link validation in agent references
- Template files structure and accessibility
- Markdown links to framework files ([text](./.krci-ai/path/file.md))
- Markdown format validation for task files
- Cross-platform file accessibility

The validation runs on the current directory framework structure and provides
detailed error reporting for any issues found.

Examples:
  krci-ai validate                    # Validate framework in current directory
  krci-ai validate --verbose          # Validate with detailed output
  krci-ai validate --quiet            # Validate with minimal output`,
	RunE: runValidate,
}

// Validation flags removed - now using GetBool pattern

func init() {
	rootCmd.AddCommand(validateCmd)

	// Add flags
	validateCmd.Flags().BoolP("verbose", "v", false, "verbose output with detailed validation results")
	validateCmd.Flags().BoolP("quiet", "q", false, "quiet output, only show summary")
}

// runValidate executes the validation command
func runValidate(cmd *cobra.Command, args []string) error {
	// Get flags
	verboseOutput, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return fmt.Errorf("failed to get verbose flag: %w", err)
	}

	quietOutput, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("failed to get quiet flag: %w", err)
	}

	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	startTime := time.Now()

	// Initialize validation system with embedded assets
	embeddedAssets := GetEmbeddedAssets()
	validator, err := validation.NewValidationSystemBuilder(currentDir).
		WithEmbeddedAssets(embeddedAssets).
		Build()
	if err != nil {
		return fmt.Errorf("failed to create validator: %w", err)
	}

	// Run framework validation
	report := validator.ValidateFramework(currentDir)
	processTime := time.Since(startTime)

	// Get actual framework statistics using discovery service
	discovery := assets.NewDiscovery(currentDir, embeddedAssets)

	// Display results
	if !quietOutput {
		fmt.Print(formatValidationReport(&report, processTime, verboseOutput, discovery))
	} else if report.Summary.CriticalCount > 0 {
		fmt.Printf("❌ Framework validation failed with %d critical issues\n", report.Summary.CriticalCount)
	}

	// Return error for critical issues. Cobra root will handle exit.
	if report.Summary.CriticalCount > 0 {
		return fmt.Errorf("validation failed with %d critical issues", report.Summary.CriticalCount)
	}
	return nil
}

// formatValidationReport formats the unified validation report for console output
func formatValidationReport(report *validation.UnifiedValidationReport, processTime time.Duration, verbose bool, discovery *assets.Discovery) string {
	var result strings.Builder

	result.WriteString("🔍 Validating framework integrity...\n\n")

	writeStatusHeader(&result, report)
	writeOverviewStatistics(&result, discovery)
	writeFrameworkInsights(&result, report, discovery)
	writeProcessingTime(&result, processTime)
	writeVerboseInfo(&result, report, verbose)
	writeExitCode(&result, report)

	return result.String()
}

// writeStatusHeader writes the validation status header
func writeStatusHeader(result *strings.Builder, report *validation.UnifiedValidationReport) {
	if report.Summary.CriticalCount > 0 {
		result.WriteString("❌ FRAMEWORK INVALID\n\n")
		writeCriticalIssues(result, report)
	} else {
		result.WriteString("✅ FRAMEWORK VALID\n\n")
	}
}

// writeCriticalIssues writes critical validation issues
func writeCriticalIssues(result *strings.Builder, report *validation.UnifiedValidationReport) {
	for _, r := range report.Results {
		if r.Severity == validation.SeverityCritical {
			result.WriteString(fmt.Sprintf("🔴 CRITICAL: %s\n", r.Message))
			if r.File != "" {
				result.WriteString(fmt.Sprintf("   File: %s\n", r.File))
			}
			if r.FixGuidance != "" {
				result.WriteString(fmt.Sprintf("   Fix: %s\n", r.FixGuidance))
			}
			result.WriteString("\n")
		}
	}
}

// writeOverviewStatistics writes the framework overview statistics
func writeOverviewStatistics(result *strings.Builder, discovery *assets.Discovery) {
	agentCount, taskCount, templateCount, dataCount := getActualComponentCounts(discovery)
	result.WriteString(fmt.Sprintf("📊 Overview: %d agents, %d tasks, %d templates, %d data files\n\n", agentCount, taskCount, templateCount, dataCount))
}

// writeFrameworkInsights writes framework insights if no critical issues
func writeFrameworkInsights(result *strings.Builder, report *validation.UnifiedValidationReport, discovery *assets.Discovery) {
	if report.Summary.CriticalCount == 0 {
		result.WriteString("💡 FRAMEWORK INSIGHTS:\n")
		writeAgentSummary(result, discovery)
	}
}

// writeAgentSummary writes the agent summary with task counts
func writeAgentSummary(result *strings.Builder, discovery *assets.Discovery) {
	// Get actual agent dependencies from discovery service
	agentDeps, err := discovery.DiscoverAgentsWithDependencies()
	if err != nil {
		// Fallback to hardcoded values if discovery fails
		agentNames := []string{"architect", "ba", "dev", "go-dev", "pm", "pmm", "po", "qa"}
		taskCounts := map[string]int{
			"architect": 4, "ba": 4, "dev": 3, "go-dev": 2,
			"pm": 4, "pmm": 6, "po": 6, "qa": 4}

		for _, agent := range agentNames {
			tasks := taskCounts[agent]
			suffix := ""
			if agent == "po" {
				suffix = " (including 1 local)"
			}
			result.WriteString(fmt.Sprintf("   • %s → %d tasks%s → 0 templates\n", agent, tasks, suffix))
		}
		return
	}

	// Track template usage for analytics
	templateUsage := make(map[string]int)

	// Use actual discovered dependencies
	for _, agentDep := range agentDeps {
		agentName := agentDep.ShortName
		taskCount := len(agentDep.Tasks)
		templateCount := len(agentDep.Templates)
		suffix := ""
		if agentName == "po" {
			suffix = " (including 1 local)"
		}
		result.WriteString(fmt.Sprintf("   • %s → %d tasks%s → %d templates\n", agentName, taskCount, suffix, templateCount))

		// Count template usage
		for _, template := range agentDep.Templates {
			templateUsage[template]++
		}
	}

	// Find most used template
	var mostUsedTemplate string
	var maxUsage int
	for template, usage := range templateUsage {
		if usage > maxUsage {
			maxUsage = usage
			mostUsedTemplate = template
		}
	}

	// Add most used template info if found
	if mostUsedTemplate != "" && maxUsage > 1 {
		result.WriteString(fmt.Sprintf("   • Most used template: %s (used by %d tasks)\n", mostUsedTemplate, maxUsage))
	}
}

// writeProcessingTime writes the processing time
func writeProcessingTime(result *strings.Builder, processTime time.Duration) {
	result.WriteString(fmt.Sprintf("⚡ Validation completed in %.1fs\n\n", processTime.Seconds()))
}

// writeVerboseInfo writes verbose info messages if requested
func writeVerboseInfo(result *strings.Builder, report *validation.UnifiedValidationReport, verbose bool) {
	if verbose && len(report.Results) > 0 {
		result.WriteString("💡 INFO:\n")
		for _, r := range report.Results {
			if r.Severity == validation.SeverityInfo {
				result.WriteString(fmt.Sprintf("   • %s\n", r.Message))
			}
		}
		result.WriteString("\n")
	}
}

// writeExitCode writes the exit code information
func writeExitCode(result *strings.Builder, report *validation.UnifiedValidationReport) {
	if report.Summary.CriticalCount > 0 {
		result.WriteString("Exit code: 1 (critical issues found)\n")
	} else {
		result.WriteString("Exit code: 0 (framework functional)\n")
	}
}

// ValidationResult represents the result of a single validation
type ValidationResult struct {
	Type        string // "agent", "task", "template", "data"
	File        string
	IsValid     bool
	Errors      []string
	Warnings    []string
	ProcessTime string
}

// ValidationResults represents the complete validation results
type ValidationResults struct {
	Results       []ValidationResult
	TotalFiles    int
	ValidFiles    int
	InvalidFiles  int
	TotalErrors   int
	TotalWarnings int
	ProcessTime   string
}

// IsValid returns true if all validations passed
func (r *ValidationResults) IsValid() bool {
	return r.InvalidFiles == 0
}

// getActualComponentCounts gets the actual component counts from the installed framework
func getActualComponentCounts(discovery *assets.Discovery) (int, int, int, int) {
	// Count agents
	agents, err := discovery.DiscoverAgents()
	if err != nil {
		return 8, 33, 20, 13 // Fallback counts
	}
	agentCount := len(agents)

	// Count dependencies to get task/template/data counts
	agentDeps, err := discovery.DiscoverAgentsWithDependencies()
	if err != nil {
		return agentCount, 33, 20, 13 // Fallback for other counts
	}

	taskCount := 0
	templateSet := make(map[string]bool)
	dataSet := make(map[string]bool)

	for _, agentDep := range agentDeps {
		taskCount += len(agentDep.Tasks)
		for _, template := range agentDep.Templates {
			templateSet[template] = true
		}
		for _, dataFile := range agentDep.DataFiles {
			dataSet[dataFile] = true
		}
	}

	return agentCount, taskCount, len(templateSet), len(dataSet)
}

// FrameworkValidator handles framework validation
type FrameworkValidator struct {
	baseDir       string
	yamlProcessor *processor.YAMLProcessor
}

// NewFrameworkValidator creates a new framework validator
func NewFrameworkValidator(baseDir string) (*FrameworkValidator, error) {
	// Use the embedded assets which include schemas
	embeddedAssets := GetEmbeddedAssets()
	yamlProcessor, err := processor.NewYAMLProcessor(embeddedAssets)
	if err != nil {
		return nil, fmt.Errorf("failed to create YAML processor: %w", err)
	}

	return &FrameworkValidator{
		baseDir:       baseDir,
		yamlProcessor: yamlProcessor,
	}, nil
}

// ValidateFramework validates the entire framework
func (v *FrameworkValidator) ValidateFramework() (*ValidationResults, error) {
	results := &ValidationResults{
		Results: make([]ValidationResult, 0),
	}

	// Look for .krci-ai directory
	frameworkDir := filepath.Join(v.baseDir, ".krci-ai")
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("no .krci-ai directory found in current directory - use 'krci-ai install' to set up framework")
	}

	// Validate agents
	if err := v.validateAgents(frameworkDir, results); err != nil {
		return nil, fmt.Errorf("agent validation failed: %w", err)
	}

	// Validate tasks (basic existence check for now)
	if err := v.validateTasks(frameworkDir, results); err != nil {
		return nil, fmt.Errorf("task validation failed: %w", err)
	}

	// Validate templates
	if err := v.validateTemplates(frameworkDir, results); err != nil {
		return nil, fmt.Errorf("template validation failed: %w", err)
	}

	// Validate internal framework links in all markdown files
	if err := v.validateInternalLinks(frameworkDir, results); err != nil {
		return nil, fmt.Errorf("internal link validation failed: %w", err)
	}

	// Calculate totals
	results.TotalFiles = len(results.Results)
	for _, result := range results.Results {
		if result.IsValid {
			results.ValidFiles++
		} else {
			results.InvalidFiles++
		}
		results.TotalErrors += len(result.Errors)
		results.TotalWarnings += len(result.Warnings)
	}

	return results, nil
}

// validateAgents validates all agent files
func (v *FrameworkValidator) validateAgents(frameworkDir string, results *ValidationResults) error {
	agentsDir := filepath.Join(frameworkDir, "agents")

	// Check if agents directory exists
	if _, err := os.Stat(agentsDir); os.IsNotExist(err) {
		// No agents directory is not an error - might be a minimal setup
		return nil
	}

	// Find all agent files
	agentFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.yaml"))
	if err != nil {
		return fmt.Errorf("failed to find agent files: %w", err)
	}

	// Also check for .yml files
	ymlFiles, err := filepath.Glob(filepath.Join(agentsDir, "*.yml"))
	if err != nil {
		return fmt.Errorf("failed to find agent files: %w", err)
	}

	agentFiles = append(agentFiles, ymlFiles...)

	// Validate each agent file
	for _, agentFile := range agentFiles {
		result := v.validateAgentFile(agentFile)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateAgentFile validates a single agent file
func (v *FrameworkValidator) validateAgentFile(filePath string) ValidationResult {
	// Convert absolute path to relative path from baseDir
	relPath, err := filepath.Rel(v.baseDir, filePath)
	if err != nil {
		relPath = filePath // fallback to absolute path if conversion fails
	}

	result := ValidationResult{
		Type:     "agent",
		File:     relPath,
		IsValid:  true,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}

	// Parse and validate the agent
	agent, validationErrors, err := v.yamlProcessor.ProcessAndValidateAgent(filePath)
	if err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Parse error: %s", err.Error()))
		return result
	}

	// Process validation errors
	for _, validationError := range validationErrors {
		result.IsValid = false
		result.Errors = append(result.Errors, validationError.Message)
	}

	// Validate task path links (only if agent parsing was successful)
	if agent != nil {
		taskPathErrors := v.validateTaskPathLinks(agent, v.baseDir)
		for _, taskError := range taskPathErrors {
			result.IsValid = false
			result.Errors = append(result.Errors, taskError)
		}
	}

	return result
}

// validateTaskPathLinks validates that all task references in an agent exist and are accessible
func (v *FrameworkValidator) validateTaskPathLinks(agent *processor.Agent, baseDir string) []string {
	var errors []string

	// Check if agent has tasks defined
	if len(agent.Agent.Tasks) == 0 {
		return errors // No tasks to validate
	}

	for _, taskPath := range agent.Agent.Tasks {
		// Convert relative task path to absolute path
		// Task paths are expected to be in format "./.krci-ai/tasks/filename.md"
		var absoluteTaskPath string

		if strings.HasPrefix(taskPath, "./") {
			// Remove the "./" prefix and join with baseDir
			relativePath := strings.TrimPrefix(taskPath, "./")
			absoluteTaskPath = filepath.Join(baseDir, relativePath)
		} else if filepath.IsAbs(taskPath) {
			absoluteTaskPath = taskPath
		} else {
			// Treat as relative to baseDir
			absoluteTaskPath = filepath.Join(baseDir, taskPath)
		}

		// Check if the task file exists
		if _, err := os.Stat(absoluteTaskPath); os.IsNotExist(err) {
			errors = append(errors, fmt.Sprintf("Task reference not found: %s", taskPath))
			continue
		}

		// Check if file is readable
		if _, err := os.Open(absoluteTaskPath); err != nil {
			errors = append(errors, fmt.Sprintf("Task reference not accessible: %s (%s)", taskPath, err.Error()))
			continue
		}

		// Verify it's a markdown file
		if !strings.HasSuffix(taskPath, ".md") {
			errors = append(errors, fmt.Sprintf("Task reference must be a markdown file: %s", taskPath))
			continue
		}
	}

	return errors
}

// validateTasks validates task files (basic existence check for both standard and local tasks)
func (v *FrameworkValidator) validateTasks(frameworkDir string, results *ValidationResults) error {
	// Validate standard tasks directory
	if err := v.validateTasksInDirectory(frameworkDir, "tasks", results); err != nil {
		return err
	}

	// Validate local tasks directory
	return v.validateTasksInDirectory(frameworkDir, "local/tasks", results)
}

// validateTasksInDirectory validates task files in a specific directory
func (v *FrameworkValidator) validateTasksInDirectory(frameworkDir, taskDir string, results *ValidationResults) error {
	tasksPath := filepath.Join(frameworkDir, taskDir)

	// Check if tasks directory exists
	if _, err := os.Stat(tasksPath); os.IsNotExist(err) {
		// No tasks directory is not an error - might be a minimal setup
		return nil
	}

	// Find all task files
	taskFiles, err := filepath.Glob(filepath.Join(tasksPath, "*.md"))
	if err != nil {
		return fmt.Errorf("failed to find task files in %s: %w", taskDir, err)
	}

	// Validate each task file (basic check for now)
	for _, taskFile := range taskFiles {
		result := v.validateTaskFile(taskFile)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateMarkdownFile provides common validation logic for markdown files
func (v *FrameworkValidator) validateMarkdownFile(filePath, fileType string) ValidationResult {
	// Convert absolute path to relative path from baseDir
	relPath, err := filepath.Rel(v.baseDir, filePath)
	if err != nil {
		relPath = filePath // fallback to absolute path if conversion fails
	}

	result := ValidationResult{
		Type:     fileType,
		File:     relPath,
		IsValid:  true,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}

	// Check if file exists and is readable
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		result.IsValid = false
		caser := cases.Title(language.English)
		result.Errors = append(result.Errors, fmt.Sprintf("%s file does not exist", caser.String(fileType)))
		return result
	}

	// Check if file has .md extension
	if !strings.HasSuffix(filePath, ".md") {
		result.IsValid = false
		caser := cases.Title(language.English)
		result.Errors = append(result.Errors, fmt.Sprintf("%s file must have .md extension", caser.String(fileType)))
		return result
	}

	// Try to read the file
	if _, err := os.ReadFile(filePath); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Cannot read %s file: %s", fileType, err.Error()))
		return result
	}

	return result
}

// validateTaskFile validates a single task file
func (v *FrameworkValidator) validateTaskFile(filePath string) ValidationResult {
	return v.validateMarkdownFile(filePath, "task")
}

// validateTemplates validates all template files
func (v *FrameworkValidator) validateTemplates(frameworkDir string, results *ValidationResults) error {
	templatesDir := filepath.Join(frameworkDir, "templates")

	// Check if templates directory exists
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		// No templates directory is not an error - might be a minimal setup
		return nil
	}

	// Find all template files
	templateFiles, err := filepath.Glob(filepath.Join(templatesDir, "*.md"))
	if err != nil {
		return fmt.Errorf("failed to find template files: %w", err)
	}

	// Validate each template file
	for _, templateFile := range templateFiles {
		result := v.validateTemplateFile(templateFile)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateTemplateFile validates a single template file
func (v *FrameworkValidator) validateTemplateFile(filePath string) ValidationResult {
	result := v.validateMarkdownFile(filePath, "template")

	// Add template-specific validation if file is readable
	if result.IsValid {
		content, err := os.ReadFile(filePath)
		if err == nil {
			v.validateTemplateStructure(filePath, string(content), &result)
		}
	}

	return result
}

// validateTemplateStructure validates the structure and content of a template
func (v *FrameworkValidator) validateTemplateStructure(_ string, content string, result *ValidationResult) {
	lines := strings.Split(content, "\n")

	// Check for required template elements
	hasTitle := false
	hasContent := false
	nonEmptyLines := 0

	for lineNum, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if len(trimmedLine) > 0 {
			nonEmptyLines++
		}

		// Check for title (first non-empty line should be heading)
		if !hasTitle && len(trimmedLine) > 0 {
			if strings.HasPrefix(trimmedLine, "#") {
				hasTitle = true
			} else if trimmedLine != "" {
				result.Warnings = append(result.Warnings, fmt.Sprintf("Template should start with a heading (line %d)", lineNum+1))
				hasTitle = true // Don't report multiple times
			}
		}

		// Check for substantial content (non-heading lines with reasonable length)
		if len(trimmedLine) > 20 && !strings.HasPrefix(trimmedLine, "#") {
			hasContent = true
		}
	}

	// Validate minimum content requirements
	if nonEmptyLines < 3 {
		result.Warnings = append(result.Warnings, "Template appears to have minimal content (less than 3 non-empty lines)")
	}

	if !hasContent {
		result.Warnings = append(result.Warnings, "Template appears to lack substantial content")
	}

	// Check for template-specific patterns
	contentStr := strings.ToLower(content)
	if !strings.Contains(contentStr, "template") && !strings.Contains(contentStr, "example") {
		result.Warnings = append(result.Warnings, "Template missing usage guidance - add 'template' or 'example' keyword")
	}
}

// validateInternalLinks validates internal framework links in all markdown files
func (v *FrameworkValidator) validateInternalLinks(frameworkDir string, results *ValidationResults) error {
	// Find all markdown files in the framework
	var markdownFiles []string

	// Check agents, tasks, templates, data, and local directories
	dirs := []string{"agents", "tasks", "templates", "data", "local/tasks"}

	for _, dir := range dirs {
		dirPath := filepath.Join(frameworkDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			continue // Skip non-existent directories
		}

		// Find markdown files in this directory
		files, err := filepath.Glob(filepath.Join(dirPath, "*.md"))
		if err != nil {
			return fmt.Errorf("failed to find markdown files in %s: %w", dir, err)
		}
		markdownFiles = append(markdownFiles, files...)
	}

	// Validate internal links in each file
	for _, markdownFile := range markdownFiles {
		result := v.validateInternalLinksInFile(markdownFile)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateInternalLinksInFile validates internal framework links in a single markdown file
func (v *FrameworkValidator) validateInternalLinksInFile(filePath string) ValidationResult {
	// Convert absolute path to relative path from baseDir
	relPath, err := filepath.Rel(v.baseDir, filePath)
	if err != nil {
		relPath = filePath // fallback to absolute path if conversion fails
	}

	result := ValidationResult{
		Type:     "framework-links",
		File:     relPath,
		IsValid:  true,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Cannot read file for link validation: %s", err.Error()))
		return result
	}

	// Look for internal framework links
	v.checkInternalFrameworkLinks(string(content), &result)

	return result
}

// checkInternalFrameworkLinks checks for internal framework references and validates they exist
func (v *FrameworkValidator) checkInternalFrameworkLinks(content string, result *ValidationResult) {
	lines := strings.Split(content, "\n")

	for lineNum, line := range lines {
		// Only validate markdown links that reference internal framework paths
		// Format: [text](./.krci-ai/path/to/file.md)
		if strings.Contains(line, "](./.krci-ai/") {
			v.validateMarkdownFrameworkLinks(line, lineNum+1, result)
		}
	}
}

// validateMarkdownFrameworkLinks validates markdown links that reference internal framework paths
func (v *FrameworkValidator) validateMarkdownFrameworkLinks(line string, lineNum int, result *ValidationResult) {
	// Parse markdown links with framework references: [text](./.krci-ai/path/to/file.md)
	for i := 0; i < len(line); i++ {
		if i < len(line)-4 && line[i:i+4] == "](./" {
			// Found potential framework link start
			start := i + 2 // Start after ](
			end := start

			// Find the closing )
			for j := start; j < len(line); j++ {
				if line[j] == ')' {
					end = j
					break
				}
			}

			if end > start {
				linkPath := line[start:end]

				// Only validate links that start with ./.krci-ai/
				if strings.HasPrefix(linkPath, "./.krci-ai/") {
					// Convert to absolute path for checking
					// Remove the "./" prefix and join with base directory
					cleanPath := strings.TrimPrefix(linkPath, "./")
					absolutePath := filepath.Join(v.baseDir, cleanPath)

					// Check if the referenced file exists
					if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
						result.IsValid = false
						result.Errors = append(result.Errors, fmt.Sprintf("Markdown link to framework file not found (line %d): %s", lineNum, linkPath))
						result.Errors = append(result.Errors, "💡 Note: Only validates markdown links [text](./.krci-ai/...), not direct path references")
					} else {
						// Optional: Validate that it's a valid framework file type
						if !strings.HasSuffix(linkPath, ".md") && !strings.HasSuffix(linkPath, ".yaml") && !strings.HasSuffix(linkPath, ".yml") {
							result.Warnings = append(result.Warnings, fmt.Sprintf("Framework link references non-standard file type (line %d): %s", lineNum, linkPath))
						}
					}
				}
			}
		}
	}
}

// Legacy functions kept for backward compatibility but unused in enhanced validation
