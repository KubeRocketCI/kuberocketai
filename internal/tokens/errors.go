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
package tokens

import (
	"fmt"
	"strings"
)

// FormatUserFriendlyError converts technical errors into user-friendly messages
func FormatUserFriendlyError(err error) string {
	if err == nil {
		return ""
	}

	errStr := err.Error()

	// Handle common error patterns with user-friendly messages
	switch {
	case strings.Contains(errStr, "agent") && strings.Contains(errStr, "not found"):
		return "❌ Agent not found. Use 'krci-ai list' to see available agents."

	case strings.Contains(errStr, "framework not installed"):
		return "❌ Framework not installed in current directory. Run 'krci-ai install' first."

	case strings.Contains(errStr, "file not found") || strings.Contains(errStr, "no such file"):
		return "❌ Required file is missing. Please ensure all framework files are present."

	case strings.Contains(errStr, "permission denied"):
		return "❌ Permission denied. Please check file permissions and try again."

	case strings.Contains(errStr, "failed to encode text"):
		return "❌ Failed to process text content. The file may contain unsupported characters."

	case strings.Contains(errStr, "context canceled"):
		return "❌ Operation was canceled. Please try again."

	case strings.Contains(errStr, "context deadline exceeded"):
		return "❌ Operation timed out. Please try again or check for large files."

	case strings.Contains(errStr, "token count") && strings.Contains(errStr, "exceeds") && strings.Contains(errStr, "limit"):
		return "⚠️  Token limit exceeded. Consider breaking down your configuration into smaller components."

	default:
		// For unknown errors, provide a generic helpful message
		return fmt.Sprintf("❌ An error occurred: %s\n\nIf this problem persists, please check your framework installation and file permissions.", errStr)
	}
}

// SuggestSolutions provides suggestions based on the error type
func SuggestSolutions(err error) []string {
	if err == nil {
		return nil
	}

	errStr := err.Error()
	var suggestions []string

	switch {
	case strings.Contains(errStr, "agent") && strings.Contains(errStr, "not found"):
		suggestions = append(suggestions,
			"• Run 'krci-ai list' to see all available agents",
			"• Check the agent name spelling",
			"• Ensure the framework is properly installed")

	case strings.Contains(errStr, "framework not installed"):
		suggestions = append(suggestions,
			"• Run 'krci-ai install' to install the framework",
			"• Navigate to your project directory first",
			"• Ensure you have write permissions in the current directory")

	case strings.Contains(errStr, "file not found") || strings.Contains(errStr, "no such file"):
		suggestions = append(suggestions,
			"• Verify all framework files are present",
			"• Try reinstalling with 'krci-ai install'",
			"• Check if files were accidentally deleted or moved")

	case strings.Contains(errStr, "permission denied"):
		suggestions = append(suggestions,
			"• Check file and directory permissions",
			"• Try running with appropriate privileges",
			"• Ensure you have read access to framework files")

	case strings.Contains(errStr, "token count") && strings.Contains(errStr, "exceeds"):
		suggestions = append(suggestions,
			"• Break down large configurations into smaller components",
			"• Remove unnecessary content from templates and tasks",
			"• Consider using shorter descriptions and documentation",
			"• Use references instead of inline content where possible")

	default:
		suggestions = append(suggestions,
			"• Check your framework installation: 'krci-ai validate'",
			"• Ensure all required files are present and accessible",
			"• Try reinstalling the framework if issues persist")
	}

	return suggestions
}
