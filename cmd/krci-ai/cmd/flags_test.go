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
	"slices"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstallCommandFlags(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		description string
	}{
		{
			name:        "valid cursor IDE",
			args:        []string{"--ide=cursor"},
			expectError: false,
			description: "Should accept cursor as valid IDE",
		},
		{
			name:        "valid claude IDE",
			args:        []string{"--ide=claude"},
			expectError: false,
			description: "Should accept claude as valid IDE",
		},
		{
			name:        "valid all IDE",
			args:        []string{"--ide=all"},
			expectError: false,
			description: "Should accept all as valid IDE",
		},
		{
			name:        "invalid IDE",
			args:        []string{"--ide=invalid"},
			expectError: false, // Flag parsing succeeds, validation happens in command logic
			description: "Should parse invalid IDE flag",
		},
		{
			name:        "force flag",
			args:        []string{"--force"},
			expectError: false,
			description: "Should accept force flag",
		},
		{
			name:        "all flag",
			args:        []string{"--all"},
			expectError: false,
			description: "Should accept all flag",
		},
		{
			name:        "combined flags",
			args:        []string{"--ide=cursor", "--force"},
			expectError: false,
			description: "Should accept combined flags",
		},
		{
			name:        "all and force flags",
			args:        []string{"--all", "--force"},
			expectError: false,
			description: "Should accept all and force flags together",
		},
		{
			name:        "short form flags",
			args:        []string{"-a", "-f"},
			expectError: false,
			description: "Should accept short form flags",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new command for each test
			cmd := &cobra.Command{
				Use: "install",
				Run: func(cmd *cobra.Command, args []string) {
					// Empty run function for testing
				},
			}

			// Add the same flags as the real install command
			cmd.Flags().StringP("ide", "i", "", "IDE integration (cursor, claude, vscode, windsurf, all)")
			cmd.Flags().BoolP("force", "f", false, "Force installation even if framework is already installed")
			cmd.Flags().BoolP("all", "a", false, "Install core framework with all IDE integrations (equivalent to --ide=all)")

			cmd.SetArgs(tt.args)

			// Test flag parsing
			err := cmd.ParseFlags(tt.args)

			if tt.expectError {
				assert.Error(t, err, "Expected error for test '%s'", tt.name)
			} else {
				assert.NoError(t, err, "Unexpected error for test '%s'", tt.name)
			}

			// Test flag values if parsing succeeded
			if err == nil {
				// Test IDE flag
				ideFlag, flagErr := cmd.Flags().GetString("ide")
				require.NoError(t, flagErr, "Should be able to get IDE flag")

				if len(tt.args) > 0 && tt.args[0] == "--ide=cursor" {
					assert.Equal(t, "cursor", ideFlag, "IDE flag should be set to cursor")
				}

				// Test force flag
				forceFlag, flagErr := cmd.Flags().GetBool("force")
				require.NoError(t, flagErr, "Should be able to get force flag")

				if slices.Contains(tt.args, "--force") {
					assert.True(t, forceFlag, "Force flag should be true when --force is provided")
				}

				// Test all flag
				allFlag, flagErr := cmd.Flags().GetBool("all")
				require.NoError(t, flagErr, "Should be able to get all flag")

				if slices.Contains(tt.args, "--all") {
					assert.True(t, allFlag, "All flag should be true when --all is provided")
				}
			}
		})
	}
}

func TestInstallCommandConstants(t *testing.T) {
	// Test that our constants are defined correctly
	assert.Equal(t, "all", ideAll, "ideAll constant should be 'all'")
	assert.Equal(t, "cursor", ideCursor, "ideCursor constant should be 'cursor'")
	assert.Equal(t, "claude", ideClaude, "ideClaude constant should be 'claude'")
}
