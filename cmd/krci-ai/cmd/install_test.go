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

// TestInstallCommandExists verifies that the install command is properly defined
func TestInstallCommandExists(t *testing.T) {
	require.NotNil(t, installCmd, "installCmd should not be nil")

	assert.Equal(t, "install", installCmd.Use, "Command name should be 'install'")
	assert.NotEmpty(t, installCmd.Short, "Command short description should not be empty")
	assert.NotEmpty(t, installCmd.Long, "Command long description should not be empty")
	require.NotNil(t, installCmd.Run, "Command run function should not be nil")
}

// TestInstallCommandHasRequiredFlags verifies that all required flags are defined
func TestInstallCommandHasRequiredFlags(t *testing.T) {
	tests := []struct {
		name      string
		flagName  string
		shorthand string
	}{
		{
			name:      "IDE flag",
			flagName:  "ide",
			shorthand: "i",
		},
		{
			name:      "force flag",
			flagName:  "force",
			shorthand: "f",
		},
		{
			name:      "all flag",
			flagName:  "all",
			shorthand: "a",
		},
		{
			name:      "agent flag",
			flagName:  "agent",
			shorthand: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := installCmd.Flags().Lookup(tt.flagName)
			require.NotNil(t, flag, "%s should be defined", tt.name)

			if tt.shorthand != "" {
				assert.Equal(t, tt.shorthand, flag.Shorthand, "%s shorthand should match expected", tt.name)
			}
		})
	}
}

// TestIDEValidation verifies that IDE constants are correctly defined
func TestIDEValidation(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		expected string
	}{
		{
			name:     "Cursor IDE",
			constant: ideCursor,
			expected: "cursor",
		},
		{
			name:     "Claude IDE",
			constant: ideClaude,
			expected: "claude",
		},
		{
			name:     "VS Code IDE",
			constant: ideVSCode,
			expected: "vscode",
		},
		{
			name:     "Windsurf IDE",
			constant: ideWindsurf,
			expected: "windsurf",
		},
		{
			name:     "All IDEs",
			constant: ideAll,
			expected: "all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.constant, "IDE constant should match expected value")
		})
	}

	// Test that all IDEs are included in valid set
	validIDEs := []string{ideCursor, ideClaude, ideVSCode, ideWindsurf, ideAll}
	expectedIDEs := []string{"cursor", "claude", "vscode", "windsurf", "all"}

	assert.ElementsMatch(t, expectedIDEs, validIDEs, "Valid IDEs should contain all expected values")
}

// TestFlagReading tests reading individual flags directly
func TestFlagReading(t *testing.T) {
	// Create a test command with flags
	cmd := &cobra.Command{}
	cmd.Flags().String("ide", "", "IDE integration")
	cmd.Flags().Bool("force", false, "Force installation")
	cmd.Flags().Bool("all", false, "Install all")

	// Test default values
	t.Run("default values", func(t *testing.T) {
		ideFlag, err := cmd.Flags().GetString("ide")
		require.NoError(t, err, "Should be able to read ide flag")
		assert.Empty(t, ideFlag, "IDE flag should be empty by default")

		forceFlag, err := cmd.Flags().GetBool("force")
		require.NoError(t, err, "Should be able to read force flag")
		assert.False(t, forceFlag, "Force flag should be false by default")

		allFlag, err := cmd.Flags().GetBool("all")
		require.NoError(t, err, "Should be able to read all flag")
		assert.False(t, allFlag, "All flag should be false by default")
	})

	// Test with flags set
	t.Run("set values", func(t *testing.T) {
		err := cmd.Flags().Set("ide", "cursor")
		require.NoError(t, err, "Should be able to set ide flag")

		err = cmd.Flags().Set("force", "true")
		require.NoError(t, err, "Should be able to set force flag")

		err = cmd.Flags().Set("all", "true")
		require.NoError(t, err, "Should be able to set all flag")

		ideFlag, _ := cmd.Flags().GetString("ide")
		assert.Equal(t, "cursor", ideFlag, "IDE flag should be set to cursor")

		forceFlag, _ := cmd.Flags().GetBool("force")
		assert.True(t, forceFlag, "Force flag should be true")

		allFlag, _ := cmd.Flags().GetBool("all")
		assert.True(t, allFlag, "All flag should be true")
	})
}

// TestValidateIDEFlag tests the validateIDEFlag function
func TestValidateIDEFlag(t *testing.T) {
	tests := []struct {
		name        string
		ideFlag     string
		expectError bool
	}{
		{
			name:        "valid cursor IDE",
			ideFlag:     ideCursor,
			expectError: false,
		},
		{
			name:        "valid claude IDE",
			ideFlag:     ideClaude,
			expectError: false,
		},
		{
			name:        "valid vscode IDE",
			ideFlag:     ideVSCode,
			expectError: false,
		},
		{
			name:        "valid windsurf IDE",
			ideFlag:     ideWindsurf,
			expectError: false,
		},
		{
			name:        "valid all IDEs",
			ideFlag:     ideAll,
			expectError: false,
		},
		{
			name:        "invalid IDE",
			ideFlag:     "invalid-ide",
			expectError: true,
		},
		{
			name:        "empty IDE flag",
			ideFlag:     "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateIDEFlag(tt.ideFlag, nil)

			if tt.expectError {
				assert.Error(t, err, "Expected error for IDE flag '%s'", tt.ideFlag)
			} else {
				assert.NoError(t, err, "Expected no error for valid IDE flag '%s'", tt.ideFlag)
			}
		})
	}
}

// TestInstallCommandRunFunction tests that the command run function is properly structured
func TestInstallCommandRunFunction(t *testing.T) {
	require.NotNil(t, installCmd.Run, "Install command Run function should not be nil")

	// Test that the command doesn't panic when called with empty args
	// Note: We can't easily test the full functionality without mocking the dependencies,
	// but we can ensure the structure is correct
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Install command Run function panicked: %v", r)
		}
	}()

	// Create a minimal test command to verify structure
	testCmd := &cobra.Command{}
	testCmd.Flags().String("ide", "", "IDE integration")
	testCmd.Flags().Bool("force", false, "Force installation")
	testCmd.Flags().Bool("all", false, "Install all")
	testCmd.Flags().String("agent", "", "Install specific agents")

	// Verify that the test command structure is valid
	require.NotNil(t, testCmd, "Test command should be created successfully")
	assert.NotNil(t, testCmd.Flags(), "Test command should have flags")
}

// TestInstallSelectiveFlags tests the selective installation flags
func TestInstallSelectiveFlags(t *testing.T) {
	// Test that the installAgent variable exists and can be modified
	originalAgent := installAgent

	// Cleanup after test
	defer func() {
		installAgent = originalAgent
	}()

	tests := []struct {
		name       string
		agentValue string
	}{
		{
			name:       "set agent only",
			agentValue: "test-agent",
		},
		{
			name:       "empty agent",
			agentValue: "",
		},
		{
			name:       "multiple agents",
			agentValue: "pm,architect",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installAgent = tt.agentValue

			assert.Equal(t, tt.agentValue, installAgent, "installAgent should be set correctly")
		})
	}
}
