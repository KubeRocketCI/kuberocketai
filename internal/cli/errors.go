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
	"os"

	"github.com/KubeRocketCI/kuberocketai/internal/cli/style"
)

// ErrorHandler provides colorized error handling for CLI commands
type ErrorHandler struct{}

// NewErrorHandler creates a new error handler with color support
func NewErrorHandler() *ErrorHandler { return &ErrorHandler{} }

// HandleError prints a colorized error message and exits with code 1
func (e *ErrorHandler) HandleError(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s: %v\n", style.Error("❌"), message, err)
		os.Exit(1)
	}
}

// HandleErrorWithCode prints a colorized error message and exits with specified code
func (e *ErrorHandler) HandleErrorWithCode(err error, message string, code int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s: %v\n", style.Error("❌"), message, err)
		os.Exit(code)
	}
}

// PrintWarning prints a colorized warning message
func (e *ErrorHandler) PrintWarning(message string) {
	fmt.Printf("%s %s\n", style.Warn("⚠️"), message)
}

// PrintError prints a colorized error message without exiting
func (e *ErrorHandler) PrintError(message string) {
	fmt.Fprintf(os.Stderr, "%s %s\n", style.Error("❌"), message)
}

// PrintErrorWithGuidance prints a colorized error message with recovery guidance
func (e *ErrorHandler) PrintErrorWithGuidance(message string, guidance []string) {
	fmt.Fprintf(os.Stderr, "%s %s\n", style.Error("❌"), message)
	if len(guidance) > 0 {
		fmt.Fprintf(os.Stderr, "\n%s Recovery guidance:\n", style.Warn("💡"))
		for _, guide := range guidance {
			fmt.Fprintf(os.Stderr, "  • %s\n", guide)
		}
	}
}

// PrintTaskValidationError prints enhanced task validation errors with specific guidance
func (e *ErrorHandler) PrintTaskValidationError(agentName string, invalidTasks []string, availableTasks []string) {
	if len(invalidTasks) == 1 {
		fmt.Fprintf(os.Stderr, "%s Task '%s' is not available for agent '%s'\n",
			style.Error("❌"), invalidTasks[0], agentName)
	} else {
		fmt.Fprintf(os.Stderr, "%s Tasks %v are not available for agent '%s'\n",
			style.Error("❌"), invalidTasks, agentName)
	}

	if len(availableTasks) > 0 {
		fmt.Fprintf(os.Stderr, "\n%s Available tasks for %s:\n", style.Warn("💡"), agentName)
		for _, task := range availableTasks {
			fmt.Fprintf(os.Stderr, "  • %s\n", task)
		}
		fmt.Fprintf(os.Stderr, "\n%s Example usage:\n", style.Warn("💡"))
		fmt.Fprintf(os.Stderr, "  krci-ai bundle --agent %s --task %s\n", agentName, availableTasks[0])
	}
}

// PrintAgentValidationError prints enhanced agent validation errors with recovery guidance
func (e *ErrorHandler) PrintAgentValidationError(invalidAgents []string, availableAgents []string) {
	if len(invalidAgents) == 1 {
		fmt.Fprintf(os.Stderr, "%s Agent '%s' not found\n", style.Error("❌"), invalidAgents[0])
	} else {
		fmt.Fprintf(os.Stderr, "%s Agents %v not found\n", style.Error("❌"), invalidAgents)
	}

	if len(availableAgents) > 0 {
		fmt.Fprintf(os.Stderr, "\n%s Available agents:\n", style.Warn("💡"))
		for _, agent := range availableAgents {
			fmt.Fprintf(os.Stderr, "  • %s\n", agent)
		}
		fmt.Fprintf(os.Stderr, "\n%s Example usage:\n", style.Warn("💡"))
		fmt.Fprintf(os.Stderr, "  krci-ai install --agent %s\n", availableAgents[0])
		if len(availableAgents) > 1 {
			fmt.Fprintf(os.Stderr, "  krci-ai install --agent %s,%s\n", availableAgents[0], availableAgents[1])
		}
	}
}

// PrintInstallationError prints enhanced installation errors with recovery guidance
func (e *ErrorHandler) PrintInstallationError(err error, phase string) {
	fmt.Fprintf(os.Stderr, "%s Installation failed during %s: %v\n", style.Error("❌"), phase, err)

	fmt.Fprintf(os.Stderr, "\n%s Recovery suggestions:\n", style.Warn("💡"))
	fmt.Fprintf(os.Stderr, "  • Check that you have write permissions in the current directory\n")
	fmt.Fprintf(os.Stderr, "  • Ensure sufficient disk space is available\n")
	fmt.Fprintf(os.Stderr, "  • Try running with --force flag to overwrite existing installation\n")
	fmt.Fprintf(os.Stderr, "  • Verify that the directory is not read-only\n")
}
