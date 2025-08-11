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
package version

import (
	"runtime"
	"testing"

	"github.com/KubeRocketCI/kuberocketai/internal/testutil"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		version string
		want    string
		wantErr bool
	}{
		{"v1.0.0", "1.0.0", false},
		{"1.0.0", "1.0.0", false},
		{"v1.2.3-alpha.1", "1.2.3-alpha.1", false},
		{"invalid", "", true},
		{"", "", true},
	}

	for _, tt := range tests {
		got, err := ParseVersion(tt.version)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseVersion(%q) error = %v, wantErr %v", tt.version, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && got.String() != tt.want {
			t.Errorf("ParseVersion(%q) = %q, want %q", tt.version, got.String(), tt.want)
		}
	}
}

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		current string
		latest  string
		want    bool
		wantErr bool
	}{
		{"1.0.0", "1.0.0", false, false},
		{"1.0.0", "1.0.1", true, false},
		{"1.0.1", "1.0.0", false, false},
		{"1.0.0", "2.0.0", true, false},
		{"2.0.0", "1.0.0", false, false},
		{"1.0.0-alpha.1", "1.0.0", true, false},
		{"1.0.0", "1.0.0-alpha.1", false, false},
		{"invalid", "1.0.0", false, true},
		{"1.0.0", "invalid", false, true},
	}

	for _, tt := range tests {
		got, err := CompareVersions(tt.current, tt.latest)
		if (err != nil) != tt.wantErr {
			t.Errorf("CompareVersions(%q, %q) error = %v, wantErr %v", tt.current, tt.latest, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("CompareVersions(%q, %q) = %t, want %t", tt.current, tt.latest, got, tt.want)
		}
	}
}

func TestGetCurrentVersion(t *testing.T) {
	// Test with default values
	version := GetCurrentVersion()
	if version == "" {
		t.Error("GetCurrentVersion() should not return empty string")
	}
}

func TestGetVersionInfo(t *testing.T) {
	info := GetVersionInfo()

	// Test that all fields are populated
	if info.GoVersion == "" {
		t.Error("GoVersion should not be empty")
	}

	if info.Platform == "" {
		t.Error("Platform should not be empty")
	}

	// Check that platform format is correct
	expectedPlatform := runtime.GOOS + "/" + runtime.GOARCH
	if info.Platform != expectedPlatform {
		t.Errorf("Expected platform %s, got %s", expectedPlatform, info.Platform)
	}

	// Check Go version format
	if info.GoVersion != runtime.Version() {
		t.Errorf("Expected Go version %s, got %s", runtime.Version(), info.GoVersion)
	}

	// Version should match GetCurrentVersion
	if info.Version != GetCurrentVersion() {
		t.Errorf("VersionInfo.Version (%s) should match GetCurrentVersion() (%s)", info.Version, GetCurrentVersion())
	}
}

func TestVersionInfoStruct(t *testing.T) {
	info := VersionInfo{
		Version:   "v1.0.0",
		Commit:    "abc123",
		Date:      "2025-07-27",
		BuiltBy:   "test",
		Framework: "v0.20.0",
		GoVersion: "go1.21.0",
		Platform:  "linux/amd64",
	}

	if info.Version != "v1.0.0" {
		t.Errorf("Expected Version v1.0.0, got %s", info.Version)
	}

	if info.Commit != "abc123" {
		t.Errorf("Expected Commit abc123, got %s", info.Commit)
	}

	if info.Date != "2025-07-27" {
		t.Errorf("Expected Date 2025-07-27, got %s", info.Date)
	}

	if info.BuiltBy != "test" {
		t.Errorf("Expected BuiltBy test, got %s", info.BuiltBy)
	}

	if info.Framework != "v0.20.0" {
		t.Errorf("Expected Framework v0.20.0, got %s", info.Framework)
	}

	if info.GoVersion != "go1.21.0" {
		t.Errorf("Expected GoVersion go1.21.0, got %s", info.GoVersion)
	}

	if info.Platform != "linux/amd64" {
		t.Errorf("Expected Platform linux/amd64, got %s", info.Platform)
	}
}

func TestVersionComparisonEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		current string
		latest  string
		want    bool
		wantErr bool
	}{
		{
			name:    "pre-release vs release",
			current: "1.0.0-rc.1",
			latest:  "1.0.0",
			want:    true,
			wantErr: false,
		},
		{
			name:    "build metadata ignored",
			current: "1.0.0+build1",
			latest:  "1.0.0+build2",
			want:    false,
			wantErr: false,
		},
		{
			name:    "major version difference",
			current: "1.9.9",
			latest:  "2.0.0",
			want:    true,
			wantErr: false,
		},
		{
			name:    "patch version difference",
			current: "1.0.9",
			latest:  "1.0.10",
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareVersions(tt.current, tt.latest)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareVersions() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestVersionInfoString(t *testing.T) {
	info := VersionInfo{
		Version:   "v1.2.3",
		Commit:    "abc123def",
		Date:      "2025-07-30",
		BuiltBy:   "goreleaser",
		Framework: "v1.0.0",
		GoVersion: "go1.21.5",
		Platform:  "linux/amd64",
	}

	result := info.String()
	expected := "krci-ai version v1.2.3\nCommit: abc123def\nBuilt: 2025-07-30 by goreleaser\nFramework: v1.0.0\nGo: go1.21.5\nPlatform: linux/amd64"

	if result != expected {
		t.Errorf("VersionInfo.String() = %q, want %q", result, expected)
	}

	// Test with empty values
	emptyInfo := VersionInfo{}
	emptyResult := emptyInfo.String()
	expectedEmpty := "krci-ai version \nCommit: \nBuilt:  by \nFramework: \nGo: \nPlatform: "

	if emptyResult != expectedEmpty {
		t.Errorf("VersionInfo.String() with empty values = %q, want %q", emptyResult, expectedEmpty)
	}
}

func TestValidateCompatibility(t *testing.T) {
	tests := []struct {
		name             string
		cliVersion       string
		frameworkVersion string
		wantErr          bool
		errorContains    string
	}{
		{
			name:             "compatible versions - exact match",
			cliVersion:       "1.0.0",
			frameworkVersion: "1.0.0",
			wantErr:          false,
		},
		{
			name:             "compatible versions - within matrix",
			cliVersion:       "1.0.0",
			frameworkVersion: "1.0.1",
			wantErr:          false,
		},
		{
			name:             "compatible versions - 1.1.0 with 1.0.0",
			cliVersion:       "1.1.0",
			frameworkVersion: "1.0.0",
			wantErr:          false,
		},
		{
			name:             "incompatible versions",
			cliVersion:       "1.0.0",
			frameworkVersion: "2.0.0",
			wantErr:          true,
			errorContains:    "not compatible",
		},
		{
			name:             "invalid CLI version",
			cliVersion:       "invalid",
			frameworkVersion: "1.0.0",
			wantErr:          true,
			errorContains:    "invalid CLI version",
		},
		{
			name:             "invalid framework version",
			cliVersion:       "1.0.0",
			frameworkVersion: "invalid",
			wantErr:          true,
			errorContains:    "invalid framework version",
		},
		{
			name:             "CLI version not in compatibility matrix",
			cliVersion:       "3.0.0",
			frameworkVersion: "1.0.0",
			wantErr:          true,
			errorContains:    "no compatibility information",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCompatibility(tt.cliVersion, tt.frameworkVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCompatibility() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errorContains != "" {
				if err == nil || !testutil.ContainsSubstr(err.Error(), tt.errorContains) {
					t.Errorf("ValidateCompatibility() error = %v, want error containing %q", err, tt.errorContains)
				}
			}
		})
	}
}

func TestGetCompatibilityMatrix(t *testing.T) {
	matrix := GetCompatibilityMatrix()

	// Should not be empty
	if len(matrix) == 0 {
		t.Error("GetCompatibilityMatrix() should not return empty matrix")
	}

	// Should contain expected CLI versions
	expectedVersions := []string{"1.0.0", "1.1.0", "1.2.0", "2.0.0"}
	for _, version := range expectedVersions {
		if _, exists := matrix[version]; !exists {
			t.Errorf("GetCompatibilityMatrix() should contain version %s", version)
		}
	}

	// Should return a copy (modifying returned matrix should not affect original)
	originalLen := len(matrix["1.0.0"])
	matrix["1.0.0"] = append(matrix["1.0.0"], "test-version")

	// Get a fresh copy
	freshMatrix := GetCompatibilityMatrix()
	if len(freshMatrix["1.0.0"]) != originalLen {
		t.Error("GetCompatibilityMatrix() should return a copy, not the original matrix")
	}

	// Test modification of returned matrix doesn't affect original
	delete(matrix, "1.0.0")
	freshMatrix2 := GetCompatibilityMatrix()
	if _, exists := freshMatrix2["1.0.0"]; !exists {
		t.Error("Modifying returned matrix should not affect the original")
	}
}

func TestGetCompatibleFrameworkVersions(t *testing.T) {
	tests := []struct {
		name       string
		cliVersion string
		wantLen    int
		wantErr    bool
		wantFirst  string
	}{
		{
			name:       "valid CLI version 1.0.0",
			cliVersion: "1.0.0",
			wantLen:    3,
			wantErr:    false,
			wantFirst:  "1.0.0",
		},
		{
			name:       "valid CLI version 1.1.0",
			cliVersion: "1.1.0",
			wantLen:    5,
			wantErr:    false,
			wantFirst:  "1.0.0",
		},
		{
			name:       "valid CLI version 2.0.0",
			cliVersion: "2.0.0",
			wantLen:    3,
			wantErr:    false,
			wantFirst:  "2.0.0",
		},
		{
			name:       "invalid CLI version",
			cliVersion: "invalid",
			wantLen:    0,
			wantErr:    true,
		},
		{
			name:       "unknown CLI version",
			cliVersion: "5.0.0",
			wantLen:    0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCompatibleFrameworkVersions(tt.cliVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompatibleFrameworkVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != tt.wantLen {
					t.Errorf("GetCompatibleFrameworkVersions() returned %d versions, want %d", len(got), tt.wantLen)
				}
				if len(got) > 0 && got[0] != tt.wantFirst {
					t.Errorf("GetCompatibleFrameworkVersions() first version = %s, want %s", got[0], tt.wantFirst)
				}

				// Test that modifying returned slice doesn't affect original
				if len(got) > 0 {
					originalFirst := got[0]
					got[0] = "modified"
					fresh, _ := GetCompatibleFrameworkVersions(tt.cliVersion)
					if fresh[0] != originalFirst {
						t.Error("GetCompatibleFrameworkVersions() should return a copy")
					}
				}
			}
		})
	}
}

func TestIsFrameworkVersionCompatible(t *testing.T) {
	tests := []struct {
		name             string
		cliVersion       string
		frameworkVersion string
		want             bool
	}{
		{
			name:             "compatible versions",
			cliVersion:       "1.0.0",
			frameworkVersion: "1.0.1",
			want:             true,
		},
		{
			name:             "incompatible versions",
			cliVersion:       "1.0.0",
			frameworkVersion: "2.0.0",
			want:             false,
		},
		{
			name:             "invalid CLI version",
			cliVersion:       "invalid",
			frameworkVersion: "1.0.0",
			want:             false,
		},
		{
			name:             "invalid framework version",
			cliVersion:       "1.0.0",
			frameworkVersion: "invalid",
			want:             false,
		},
		{
			name:             "unknown CLI version",
			cliVersion:       "99.0.0",
			frameworkVersion: "1.0.0",
			want:             false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFrameworkVersionCompatible(tt.cliVersion, tt.frameworkVersion)
			if got != tt.want {
				t.Errorf("IsFrameworkVersionCompatible() = %t, want %t", got, tt.want)
			}
		})
	}
}
