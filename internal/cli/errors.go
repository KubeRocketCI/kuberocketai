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

	"github.com/fatih/color"
)

// ErrorHandler provides colorized error handling for CLI commands
type ErrorHandler struct {
	red    func(a ...interface{}) string
	yellow func(a ...interface{}) string
}

// NewErrorHandler creates a new error handler with color support
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		red:    color.New(color.FgRed).SprintFunc(),
		yellow: color.New(color.FgYellow).SprintFunc(),
	}
}

// HandleError prints a colorized error message and exits with code 1
func (e *ErrorHandler) HandleError(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s: %v\n", e.red("❌"), message, err)
		os.Exit(1)
	}
}

// HandleErrorWithCode prints a colorized error message and exits with specified code
func (e *ErrorHandler) HandleErrorWithCode(err error, message string, code int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s: %v\n", e.red("❌"), message, err)
		os.Exit(code)
	}
}

// PrintWarning prints a colorized warning message
func (e *ErrorHandler) PrintWarning(message string) {
	fmt.Printf("%s %s\n", e.yellow("⚠️"), message)
}

// PrintError prints a colorized error message without exiting
func (e *ErrorHandler) PrintError(message string) {
	fmt.Fprintf(os.Stderr, "%s %s\n", e.red("❌"), message)
}
