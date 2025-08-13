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
func (o *OutputHandler) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Newline prints a newline character
func (o *OutputHandler) Newline() {
	fmt.Println()
}
