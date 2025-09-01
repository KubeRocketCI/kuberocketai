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
	"strings"
	"time"

	"github.com/fatih/color"
)

// ValidationReport represents the complete validation report
type ValidationReport struct {
	Issues        []ValidationIssue
	Insights      *FrameworkInsights
	ProcessTime   time.Duration
	IsValid       bool
	HasCritical   bool
	HasWarnings   bool
	CriticalCount int
	WarningCount  int
	InfoCount     int
}

// NewValidationReport creates a new validation report
func NewValidationReport(issues []ValidationIssue, insights *FrameworkInsights, processTime time.Duration) *ValidationReport {
	report := &ValidationReport{
		Issues:      issues,
		Insights:    insights,
		ProcessTime: processTime,
		IsValid:     true,
		HasCritical: false,
		HasWarnings: false,
	}

	// Analyze issues
	for _, issue := range issues {
		switch issue.Severity {
		case SeverityCritical:
			report.HasCritical = true
			report.IsValid = false
			report.CriticalCount++
		case SeverityError:
			report.IsValid = false
			report.CriticalCount++ // Treat errors as critical for exit code purposes
		case SeverityWarning:
			report.HasWarnings = true
			report.WarningCount++
		case SeverityInfo:
			report.InfoCount++
		}
	}

	return report
}

// FormatReport formats the validation report for console output
func (r *ValidationReport) FormatReport(verbose bool) string {
	var result strings.Builder

	result.WriteString("🔍 Validating framework integrity...\n\n")
	result.WriteString(r.formatStatusHeader())
	result.WriteString(r.formatInsights())
	result.WriteString(r.formatCriticalIssues())
	result.WriteString(r.formatWarnings(verbose))
	result.WriteString(r.formatInfoMessages(verbose))
	result.WriteString(r.formatExitCode())

	return result.String()
}

// formatStatusHeader formats the validation status header
func (r *ValidationReport) formatStatusHeader() string {
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)

	if r.IsValid {
		if r.HasWarnings {
			return green.Sprint("✅ FRAMEWORK VALID (with warnings)") + "\n\n"
		}
		return green.Sprint("✅ FRAMEWORK VALID") + "\n\n"
	}
	return red.Sprintf("❌ FRAMEWORK INVALID (%d critical issues found)", r.CriticalCount) + "\n\n"
}

// formatInsights formats framework insights if appropriate
func (r *ValidationReport) formatInsights() string {
	if r.Insights != nil && (r.IsValid || !r.HasCritical) {
		return r.Insights.FormatInsights() + fmt.Sprintf("⚡ Validation completed in %.1fs\n\n", r.ProcessTime.Seconds())
	}
	return ""
}

// formatCriticalIssues formats critical issues section
func (r *ValidationReport) formatCriticalIssues() string {
	if !r.HasCritical {
		return ""
	}

	red := color.New(color.FgRed)
	var result strings.Builder

	result.WriteString(red.Sprint("🚨 CRITICAL ISSUES (must fix):"))
	result.WriteString("\n")

	issueNum := 1
	for _, issue := range r.Issues {
		if issue.Severity == SeverityCritical || issue.Severity == SeverityError {
			result.WriteString(fmt.Sprintf("   %d. %s", issueNum, issue.Message))
			if issue.Line > 0 {
				result.WriteString(fmt.Sprintf(" (line %d)", issue.Line))
			}
			if issue.File != "" {
				result.WriteString(fmt.Sprintf(" in %s", issue.File))
			}
			result.WriteString("\n")
			result.WriteString(fmt.Sprintf("      Fix: %s\n\n", issue.FixGuidance))
			issueNum++
		}
	}

	return result.String()
}

// formatWarnings formats warning issues section
func (r *ValidationReport) formatWarnings(verbose bool) string {
	if !r.HasWarnings {
		return ""
	}

	yellow := color.New(color.FgYellow)
	var result strings.Builder

	if r.HasCritical {
		result.WriteString(yellow.Sprintf("⚠️  WARNINGS (%d non-critical issues):", r.WarningCount))
	} else {
		result.WriteString(yellow.Sprintf("⚠️  WARNINGS (non-critical):"))
	}
	result.WriteString("\n")

	if verbose || !r.HasCritical {
		for _, issue := range r.Issues {
			if issue.Severity == SeverityWarning {
				result.WriteString(fmt.Sprintf("   • %s", issue.Message))
				if issue.File != "" {
					result.WriteString(fmt.Sprintf(" (%s)", issue.File))
				}
				result.WriteString("\n")
			}
		}
	} else {
		result.WriteString("   • Use 'krci-ai validate -v' to see warning details\n")
	}
	result.WriteString("\n")

	return result.String()
}

// formatInfoMessages formats info messages in verbose mode
func (r *ValidationReport) formatInfoMessages(verbose bool) string {
	if !verbose || r.InfoCount == 0 {
		return ""
	}

	blue := color.New(color.FgBlue)
	var result strings.Builder

	result.WriteString(blue.Sprint("💡 INFO:"))
	result.WriteString("\n")
	for _, issue := range r.Issues {
		if issue.Severity == SeverityInfo {
			result.WriteString(fmt.Sprintf("   • %s\n", issue.Message))
		}
	}
	result.WriteString("\n")

	return result.String()
}

// formatExitCode formats the exit code information
func (r *ValidationReport) formatExitCode() string {
	if r.HasCritical {
		return "Exit code: 1 (critical issues found)\n"
	}
	return "Exit code: 0 (framework functional)\n"
}
