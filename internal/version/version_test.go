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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
		wantErr bool
	}{
		{
			name:    "valid version with v prefix",
			version: "v1.0.0",
			want:    "1.0.0",
			wantErr: false,
		},
		{
			name:    "valid version without prefix",
			version: "1.0.0",
			want:    "1.0.0",
			wantErr: false,
		},
		{
			name:    "valid version with pre-release",
			version: "v1.2.3-alpha.1",
			want:    "1.2.3-alpha.1",
			wantErr: false,
		},
		{
			name:    "invalid version format",
			version: "invalid",
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty version",
			version: "",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseVersion(tt.version)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got.String())
			}
		})
	}
}

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name    string
		current string
		latest  string
		want    bool
		wantErr bool
	}{
		{
			name:    "same versions",
			current: "1.0.0",
			latest:  "1.0.0",
			want:    false,
			wantErr: false,
		},
		{
			name:    "latest is newer patch",
			current: "1.0.0",
			latest:  "1.0.1",
			want:    true,
			wantErr: false,
		},
		{
			name:    "current is newer",
			current: "1.0.1",
			latest:  "1.0.0",
			want:    false,
			wantErr: false,
		},
		{
			name:    "latest is newer major",
			current: "1.0.0",
			latest:  "2.0.0",
			want:    true,
			wantErr: false,
		},
		{
			name:    "pre-release vs release",
			current: "1.0.0-alpha.1",
			latest:  "1.0.0",
			want:    true,
			wantErr: false,
		},
		{
			name:    "invalid current version",
			current: "invalid",
			latest:  "1.0.0",
			want:    false,
			wantErr: true,
		},
		{
			name:    "invalid latest version",
			current: "1.0.0",
			latest:  "invalid",
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareVersions(tt.current, tt.latest)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetCurrentVersion(t *testing.T) {
	// Test with default values
	version := GetCurrentVersion()
	assert.NotEmpty(t, version, "GetCurrentVersion should not return empty string")
}

func TestGetVersionInfo(t *testing.T) {
	info := GetVersionInfo()

	// Test that all fields are populated
	assert.NotEmpty(t, info.GoVersion, "GoVersion should not be empty")
	assert.NotEmpty(t, info.Platform, "Platform should not be empty")

	// Check that platform format is correct
	expectedPlatform := runtime.GOOS + "/" + runtime.GOARCH
	assert.Equal(t, expectedPlatform, info.Platform)

	// Check Go version format
	assert.Equal(t, runtime.Version(), info.GoVersion)

	// Version should match GetCurrentVersion
	assert.Equal(t, GetCurrentVersion(), info.Version)
}

func TestVersionInfoString(t *testing.T) {
	tests := []struct {
		name     string
		info     VersionInfo
		expected string
	}{
		{
			name: "complete version info",
			info: VersionInfo{
				Version:   "v1.2.3",
				Commit:    "abc123def",
				Date:      "2025-07-30",
				BuiltBy:   "goreleaser",
				Framework: "v1.0.0",
				GoVersion: "go1.21.5",
				Platform:  "linux/amd64",
			},
			expected: "krci-ai version v1.2.3\nCommit: abc123def\nBuilt: 2025-07-30 by goreleaser\nFramework: v1.0.0\nGo: go1.21.5\nPlatform: linux/amd64",
		},
		{
			name:     "empty version info",
			info:     VersionInfo{},
			expected: "krci-ai version \nCommit: \nBuilt:  by \nFramework: \nGo: \nPlatform: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.info.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}
