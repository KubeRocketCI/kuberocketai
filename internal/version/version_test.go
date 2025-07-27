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
