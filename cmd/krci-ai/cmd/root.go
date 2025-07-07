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
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Global embedded assets
var embeddedAssets embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "krci-ai",
	Short: "KubeRocketAI CLI - AI-as-Code framework for Kubernetes development",
	Long: `KubeRocketAI CLI is a command-line tool that provides an AI-as-Code framework
for Kubernetes development. It allows you to manage AI agents, tasks, templates,
and data as version-controlled assets.

With KubeRocketAI, you can:
- Install and manage AI framework components
- Work with AI agents for specific development roles
- Use templates for consistent output formatting
- Integrate with popular IDEs and editors

Get started by running 'krci-ai install' to set up your first framework.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// SetVersionInfo sets the version information for the CLI
func SetVersionInfo(version, commit, date, builtBy string) {
	versionInfo := fmt.Sprintf("%s (commit: %s, built: %s, by: %s)", version, commit, date, builtBy)
	rootCmd.Version = versionInfo
}

// SetEmbeddedAssets sets the embedded assets for use by commands
func SetEmbeddedAssets(assets embed.FS) {
	embeddedAssets = assets
}

// GetEmbeddedAssets returns the embedded assets
func GetEmbeddedAssets() embed.FS {
	return embeddedAssets
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.krci-ai.yaml)")

	// Version flag is automatically handled by Cobra when Version is set
}
