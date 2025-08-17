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

func TestRunVersion(t *testing.T) {
	SetVersionInfo("1.0.0", "abc123", "2025-01-01", "test-builder")

	t.Run("runVersion basic functionality", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("output", "", "Output format")

		err := runVersion(cmd, []string{})
		assert.NoError(t, err, "runVersion should not return error")
	})

	t.Run("runVersion with JSON output", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("output", "json", "Output format")

		err := cmd.Flags().Set("output", "json")
		require.NoError(t, err, "Should be able to set output flag to json")

		err = runVersion(cmd, []string{})
		assert.NoError(t, err, "runVersion with JSON output should not return error")
	})
}

func TestVersionCommandStructure(t *testing.T) {
	t.Run("version command exists", func(t *testing.T) {
		require.NotNil(t, versionCmd, "Version command should not be nil")
		assert.Equal(t, "version", versionCmd.Use, "Version command Use should be 'version'")
	})
}
