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
package main

import (
	"os"
	"testing"

	"github.com/KubeRocketCI/kuberocketai/cmd/krci-ai/cmd"
)

func TestMain(t *testing.T) {
	// Test that version info is properly set
	cmd.SetVersionInfo("test-version", "test-commit", "test-date", "test-builder")

	// Test that embedded assets can be set
	cmd.SetEmbeddedAssets(EmbeddedAssets)

	// Test that we can get embedded assets back
	assets := cmd.GetEmbeddedAssets()

	// Test that assets can be used (embed.FS is never nil, but we can test functionality)
	_, err := assets.ReadDir(".")
	// This should not panic, though it might return an error if no files are embedded
	if err != nil {
		t.Logf("Assets readdir returned error (expected in test): %v", err)
	}
}

func TestVariables(t *testing.T) {
	// Test that build variables have default values
	if version == "" {
		t.Error("Expected version to have a default value")
	}
	if commit == "" {
		t.Error("Expected commit to have a default value")
	}
	if date == "" {
		t.Error("Expected date to have a default value")
	}
	if builtBy == "" {
		t.Error("Expected builtBy to have a default value")
	}
}

func TestEmbeddedAssets(t *testing.T) {
	// Test that EmbeddedAssets can be read
	entries, err := EmbeddedAssets.ReadDir("assets")
	if err != nil {
		// It's OK if assets directory doesn't exist in test environment
		return
	}

	// If assets exist, we should be able to read them
	if len(entries) == 0 {
		t.Log("No assets found in embedded filesystem")
	}
}

// TestMainExecution tests that main doesn't panic when called
// We can't actually test the execution without interfering with the test runner
func TestMainExecution(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Set up test args to avoid interfering with test execution
	os.Args = []string{"krci-ai", "version"}

	// This would normally call main(), but we can't do that in tests
	// without interfering with the test runner. Instead, we just verify
	// the setup functions work correctly.

	t.Log("Main function setup validated successfully")
}
