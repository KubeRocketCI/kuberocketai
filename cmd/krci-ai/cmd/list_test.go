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
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListCommandExists verifies that the list command is properly defined
func TestListCommandExists(t *testing.T) {
	require.NotNil(t, listCmd, "listCmd should not be nil")
	assert.Equal(t, "list", listCmd.Use, "Command name should be 'list'")
	assert.NotEmpty(t, listCmd.Short, "Command short description should not be empty")
	assert.NotEmpty(t, listCmd.Long, "Command long description should not be empty")
}

// TestListAgentsCommandExists verifies that the list agents command is properly defined
func TestListAgentsCommandExists(t *testing.T) {
	require.NotNil(t, listAgentsCmd, "listAgentsCmd should not be nil")
	assert.Equal(t, "agents", listAgentsCmd.Use, "Command name should be 'agents'")
	assert.NotEmpty(t, listAgentsCmd.Short, "Command short description should not be empty")
	assert.NotEmpty(t, listAgentsCmd.Long, "Command long description should not be empty")
	require.NotNil(t, listAgentsCmd.Run, "Command Run function should not be nil")
}

// TestListCommandHasSubcommands verifies that list command has required subcommands
func TestListCommandHasSubcommands(t *testing.T) {
	subcommands := listCmd.Commands()
	require.NotEmpty(t, subcommands, "List command should have subcommands")

	// Check that agents subcommand exists
	var agentsCmd *cobra.Command
	for _, cmd := range subcommands {
		if cmd.Use == "agents" {
			agentsCmd = cmd
			break
		}
	}
	require.NotNil(t, agentsCmd, "List command should have 'agents' subcommand")
}

// TestListAgentsCommandFlags verifies that list agents command has required flags
func TestListAgentsCommandFlags(t *testing.T) {
	tests := []struct {
		name      string
		flagName  string
		shorthand string
	}{
		{
			name:      "verbose flag",
			flagName:  "verbose",
			shorthand: "v",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := listAgentsCmd.Flags().Lookup(tt.flagName)
			require.NotNil(t, flag, "%s should be defined", tt.name)
			assert.Equal(t, tt.shorthand, flag.Shorthand, "%s shorthand should match expected", tt.name)
		})
	}
}

// TestListCommandRegistration verifies that list command is registered with root command
func TestListCommandRegistration(t *testing.T) {
	// Check that list command is registered
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "list" {
			found = true
			break
		}
	}
	assert.True(t, found, "List command should be registered with root command")
}

// TestListAgentsCommandStructure tests the command structure without execution
func TestListAgentsCommandStructure(t *testing.T) {
	// Test that verbose flag can be read
	err := listAgentsCmd.Flags().Set("verbose", "true")
	require.NoError(t, err, "Should be able to set verbose flag")

	verboseFlag, err := listAgentsCmd.Flags().GetBool("verbose")
	require.NoError(t, err, "Should be able to read verbose flag")
	assert.True(t, verboseFlag, "Verbose flag should be true when set")

	// Reset flag
	err = listAgentsCmd.Flags().Set("verbose", "false")
	require.NoError(t, err, "Should be able to reset verbose flag")
}

// TestListCommandHelp tests that help text is properly structured
func TestListCommandHelp(t *testing.T) {
	helpText := listCmd.Long

	// Check for key sections in help text
	expectedSections := []string{
		"Available subcommands:",
		"agents",
		"Examples:",
		"krci-ai list agents",
	}

	for _, section := range expectedSections {
		assert.Contains(t, helpText, section, "Help text should contain section: %s", section)
	}
}

// TestListAgentsCommandHelp tests that agents help text is properly structured
func TestListAgentsCommandHelp(t *testing.T) {
	helpText := listAgentsCmd.Long

	// Check for key sections in help text
	expectedSections := []string{
		"scans the .krci-ai/agents/",
		"Examples:",
		"krci-ai list agents",
		"krci-ai list agents -v",
	}

	for _, section := range expectedSections {
		assert.Contains(t, helpText, section, "Help text should contain section: %s", section)
	}
}
