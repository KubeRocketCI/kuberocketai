package cmd

import (
	"slices"
	"testing"

	"github.com/spf13/cobra"
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

			if tt.expectError && err == nil {
				t.Errorf("Expected error for test '%s' but got none", tt.name)
			} else if !tt.expectError && err != nil {
				t.Errorf("Unexpected error for test '%s': %v", tt.name, err)
			}

			// Test flag values if parsing succeeded
			if err == nil {
				// Test IDE flag
				if ideFlag, err := cmd.Flags().GetString("ide"); err != nil {
					t.Errorf("Failed to get IDE flag: %v", err)
				} else if tt.args[0] == "--ide=cursor" && ideFlag != "cursor" {
					t.Errorf("Expected IDE flag to be 'cursor', got '%s'", ideFlag)
				}

				// Test force flag
				if forceFlag, err := cmd.Flags().GetBool("force"); err != nil {
					t.Errorf("Failed to get force flag: %v", err)
				} else if slices.Contains(tt.args, "--force") && !forceFlag {
					t.Errorf("Expected force flag to be true")
				}

				// Test all flag
				if allFlag, err := cmd.Flags().GetBool("all"); err != nil {
					t.Errorf("Failed to get all flag: %v", err)
				} else if slices.Contains(tt.args, "--all") && !allFlag {
					t.Errorf("Expected all flag to be true")
				}
			}
		})
	}
}

func TestInstallCommandConstants(t *testing.T) {
	// Test that our constants are defined correctly
	if ideAll != "all" {
		t.Errorf("Expected ideAll to be 'all', got '%s'", ideAll)
	}
	if ideCursor != "cursor" {
		t.Errorf("Expected ideCursor to be 'cursor', got '%s'", ideCursor)
	}
	if ideClaude != "claude" {
		t.Errorf("Expected ideClaude to be 'claude', got '%s'", ideClaude)
	}
}
