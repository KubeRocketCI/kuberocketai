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
package style

import "github.com/fatih/color"

// Cache sprint funcs once to avoid per-call allocations
var (
	successFn   = color.New(color.FgGreen).SprintFunc()
	infoFn      = color.New(color.FgBlue).SprintFunc()
	progressFn  = color.New(color.FgCyan).SprintFunc()
	warnFn      = color.New(color.FgYellow).SprintFunc()
	errorFn     = color.New(color.FgRed).SprintFunc()
	boldFn      = color.New(color.Bold).SprintFunc()
	cyanFn      = color.New(color.FgCyan).SprintFunc()
	yellowFn    = color.New(color.FgYellow).SprintFunc()
	magentaFn   = color.New(color.FgMagenta).SprintFunc()
	blueFn      = color.New(color.FgBlue).SprintFunc()
	greenBoldFn = color.New(color.FgGreen, color.Bold).SprintFunc()
)

// Accessors return preconfigured Sprint functions for consistent styling.
func Success(a ...interface{}) string   { return successFn(a...) }
func Info(a ...interface{}) string      { return infoFn(a...) }
func Progress(a ...interface{}) string  { return progressFn(a...) }
func Warn(a ...interface{}) string      { return warnFn(a...) }
func Error(a ...interface{}) string     { return errorFn(a...) }
func Bold(a ...interface{}) string      { return boldFn(a...) }
func Cyan(a ...interface{}) string      { return cyanFn(a...) }
func Yellow(a ...interface{}) string    { return yellowFn(a...) }
func Magenta(a ...interface{}) string   { return magentaFn(a...) }
func Blue(a ...interface{}) string      { return blueFn(a...) }
func GreenBold(a ...interface{}) string { return greenBoldFn(a...) }
