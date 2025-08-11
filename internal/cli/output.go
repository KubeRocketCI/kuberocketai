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

	"github.com/KubeRocketCI/kuberocketai/internal/cli/style"
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
