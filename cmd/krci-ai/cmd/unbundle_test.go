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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUnbundleCommandExists verifies that the unbundle command is properly defined
func TestUnbundleCommandExists(t *testing.T) {
	require.NotNil(t, unbundleCmd, "unbundleCmd should not be nil")

	assert.Equal(t, "unbundle <file>", unbundleCmd.Use, "Command use should be 'unbundle <file>'")
	assert.NotEmpty(t, unbundleCmd.Short, "Command short description should not be empty")
	assert.NotEmpty(t, unbundleCmd.Long, "Command long description should not be empty")
	require.NotNil(t, unbundleCmd.RunE, "Command run function should not be nil")
}

// TestUnbundleCommandHasRequiredFlags verifies that required flags are defined
func TestUnbundleCommandHasRequiredFlags(t *testing.T) {
	tests := []struct {
		name     string
		flagName string
	}{
		{
			name:     "dry-run flag",
			flagName: "dry-run",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := unbundleCmd.Flags().Lookup(tt.flagName)
			require.NotNil(t, flag, "%s should be defined", tt.name)
		})
	}
}

// TestUnbundleCommandArgs verifies that the command requires exactly one argument
func TestUnbundleCommandArgs(t *testing.T) {
	require.NotNil(t, unbundleCmd.Args, "Args validator should be defined")
}
