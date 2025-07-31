package cmd

import (
	"github.com/spf13/cobra"
	"testing"
)

func TestRunVersion(t *testing.T) {
	SetVersionInfo("1.0.0", "abc123", "2025-01-01", "test-builder")

	t.Run("runVersion basic functionality", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("output", "", "Output format")

		err := runVersion(cmd, []string{})
		if err != nil {
			t.Errorf("runVersion failed: %v", err)
		}
	})

	t.Run("runVersion with JSON output", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("output", "json", "Output format")
		cmd.Flags().Set("output", "json")

		err := runVersion(cmd, []string{})
		if err != nil {
			t.Errorf("runVersion with JSON failed: %v", err)
		}
	})
}

func TestVersionCommandStructure(t *testing.T) {
	t.Run("version command exists", func(t *testing.T) {
		if versionCmd == nil {
			t.Error("Version command should not be nil")
		}

		if versionCmd.Use != "version" {
			t.Errorf("Expected version command Use to be version, got %s", versionCmd.Use)
		}
	})
}
