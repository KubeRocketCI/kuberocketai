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

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/epam/kuberocketai/internal/engine/processor"
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

// Validation flags
var (
	verboseOutput bool
	quietOutput   bool
)

func init() {
	rootCmd.AddCommand(validateCmd)

	// Add flags
	validateCmd.Flags().BoolVarP(&verboseOutput, "verbose", "v", false, "verbose output with detailed validation results")
	validateCmd.Flags().BoolVarP(&quietOutput, "quiet", "q", false, "quiet output, only show summary")
}

// runValidate executes the validation command
func runValidate(cmd *cobra.Command, args []string) error {
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	if !quietOutput {
		fmt.Printf("🔍 Validating framework components in: %s\n\n", currentDir)
	}

	// Initialize validation system
	validator, err := NewFrameworkValidator(currentDir)
	if err != nil {
		return fmt.Errorf("failed to initialize validator: %w", err)
	}

	// Run validation
	results, err := validator.ValidateFramework()
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Display results
	displayValidationResults(results)

	// Exit with error code if validation failed
	if !results.IsValid() {
		os.Exit(1)
	}

	return nil
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

// validateTasks validates task files (basic existence check)
func (v *FrameworkValidator) validateTasks(frameworkDir string, results *ValidationResults) error {
	tasksDir := filepath.Join(frameworkDir, "tasks")

	// Check if tasks directory exists
	if _, err := os.Stat(tasksDir); os.IsNotExist(err) {
		// No tasks directory is not an error - might be a minimal setup
		return nil
	}

	// Find all task files
	taskFiles, err := filepath.Glob(filepath.Join(tasksDir, "*.md"))
	if err != nil {
		return fmt.Errorf("failed to find task files: %w", err)
	}

	// Validate each task file (basic check for now)
	for _, taskFile := range taskFiles {
		result := v.validateTaskFile(taskFile)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateTaskFile validates a single task file
func (v *FrameworkValidator) validateTaskFile(filePath string) ValidationResult {
	// Convert absolute path to relative path from baseDir
	relPath, err := filepath.Rel(v.baseDir, filePath)
	if err != nil {
		relPath = filePath // fallback to absolute path if conversion fails
	}

	result := ValidationResult{
		Type:     "task",
		File:     relPath,
		IsValid:  true,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}

	// Check if file exists and is readable
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		result.IsValid = false
		result.Errors = append(result.Errors, "Task file does not exist")
		return result
	}

	// Check if file has .md extension
	if !strings.HasSuffix(filePath, ".md") {
		result.IsValid = false
		result.Errors = append(result.Errors, "Task file must have .md extension")
		return result
	}

	// Try to read the file
	if _, err := os.ReadFile(filePath); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Cannot read task file: %s", err.Error()))
		return result
	}

	return result
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
	// Convert absolute path to relative path from baseDir
	relPath, err := filepath.Rel(v.baseDir, filePath)
	if err != nil {
		relPath = filePath // fallback to absolute path if conversion fails
	}

	result := ValidationResult{
		Type:     "template",
		File:     relPath,
		IsValid:  true,
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}

	// Check if file exists and is readable
	if _, statErr := os.Stat(filePath); os.IsNotExist(statErr) {
		result.IsValid = false
		result.Errors = append(result.Errors, "Template file does not exist")
		return result
	}

	// Check if file has .md extension
	if !strings.HasSuffix(filePath, ".md") {
		result.IsValid = false
		result.Errors = append(result.Errors, "Template file must have .md extension")
		return result
	}

	// Try to read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, fmt.Sprintf("Cannot read template file: %s", err.Error()))
		return result
	}

	// Validate template structure
	v.validateTemplateStructure(filePath, string(content), &result)

	return result
}

// validateTemplateStructure validates the structure and content of a template
func (v *FrameworkValidator) validateTemplateStructure(filePath string, content string, result *ValidationResult) {
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

	// Check agents, tasks, templates, and data directories
	dirs := []string{"agents", "tasks", "templates", "data"}

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
		result := v.validateInternalLinksInFile(markdownFile, frameworkDir)
		results.Results = append(results.Results, result)
	}

	return nil
}

// validateInternalLinksInFile validates internal framework links in a single markdown file
func (v *FrameworkValidator) validateInternalLinksInFile(filePath string, frameworkDir string) ValidationResult {
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
	v.checkInternalFrameworkLinks(string(content), filePath, frameworkDir, &result)

	return result
}

// checkInternalFrameworkLinks checks for internal framework references and validates they exist
func (v *FrameworkValidator) checkInternalFrameworkLinks(content string, filePath string, frameworkDir string, result *ValidationResult) {
	lines := strings.Split(content, "\n")

	for lineNum, line := range lines {
		// Only validate markdown links that reference internal framework paths
		// Format: [text](./.krci-ai/path/to/file.md)
		if strings.Contains(line, "](./.krci-ai/") {
			v.validateMarkdownFrameworkLinks(line, lineNum+1, filePath, frameworkDir, result)
		}
	}
}

// validateMarkdownFrameworkLinks validates markdown links that reference internal framework paths
func (v *FrameworkValidator) validateMarkdownFrameworkLinks(line string, lineNum int, filePath string, frameworkDir string, result *ValidationResult) {
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

// displayValidationResults displays the validation results
func displayFileResult(result ValidationResult) {
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)

	displayPath := result.File

	if result.IsValid {
		if verboseOutput {
			_, _ = green.Printf("✅ %s: %s\n", strings.ToUpper(result.Type), displayPath)

			// Display warnings for valid files in verbose mode
			for _, warning := range result.Warnings {
				_, _ = yellow.Printf("   ⚠️  %s\n", warning)
			}
		}
	} else {
		_, _ = red.Printf("❌ %s: %s\n", strings.ToUpper(result.Type), displayPath)

		// Display errors
		for _, errorMsg := range result.Errors {
			fmt.Printf("   🔸 %s\n", errorMsg)
		}

		// Display warnings
		for _, warning := range result.Warnings {
			_, _ = yellow.Printf("   ⚠️  %s\n", warning)
		}
	}
}

func displayValidationResults(results *ValidationResults) {
	// Color setup
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	bold := color.New(color.Bold)

	// Display individual results if verbose or if there are errors
	if verboseOutput || results.InvalidFiles > 0 {
		fmt.Println("📋 Validation Results:")
		fmt.Println(strings.Repeat("-", 60))

		for _, result := range results.Results {
			displayFileResult(result)
		}
		fmt.Println()
	}

	// Display summary
	_, _ = bold.Println("📊 Validation Summary:")
	fmt.Println(strings.Repeat("-", 40))

	fmt.Printf("Total Files:     %d\n", results.TotalFiles)

	if results.ValidFiles > 0 {
		_, _ = green.Printf("Valid Files:     %d\n", results.ValidFiles)
	} else {
		fmt.Printf("Valid Files:     %d\n", results.ValidFiles)
	}

	if results.InvalidFiles > 0 {
		_, _ = red.Printf("Invalid Files:   %d\n", results.InvalidFiles)
	} else {
		fmt.Printf("Invalid Files:   %d\n", results.InvalidFiles)
	}

	if results.TotalErrors > 0 {
		_, _ = red.Printf("Total Errors:    %d\n", results.TotalErrors)
	} else {
		fmt.Printf("Total Errors:    %d\n", results.TotalErrors)
	}

	if results.TotalWarnings > 0 {
		_, _ = yellow.Printf("Total Warnings:  %d\n", results.TotalWarnings)
		if !verboseOutput {
			fmt.Printf("💡 Use 'krci-ai validate -v' to see warning details\n")
		}
	} else {
		fmt.Printf("Total Warnings:  %d\n", results.TotalWarnings)
	}

	fmt.Println()

	// Display validation scope information
	fmt.Printf("Framework Link Scope: Validates markdown links [text](./.krci-ai/...) only\n")
	fmt.Println()

	// Display final status
	if results.IsValid() {
		_, _ = green.Println("🎉 Framework validation completed successfully!")
	} else {
		_, _ = red.Println("❌ Framework validation failed!")
		if !quietOutput {
			fmt.Println("   Please fix the errors above and run validation again.")
		}
	}
}
