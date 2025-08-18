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
package github

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReleaseStruct(t *testing.T) {
	// Test JSON unmarshaling of Release struct
	jsonData := `{
		"tag_name": "v1.2.3",
		"published_at": "2025-07-27T10:30:00Z",
		"html_url": "https://github.com/test/repo/releases/tag/v1.2.3",
		"body": "Release notes here"
	}`

	var release Release
	err := json.Unmarshal([]byte(jsonData), &release)
	require.NoError(t, err)

	assert.Equal(t, "v1.2.3", release.TagName)
	assert.Equal(t, "https://github.com/test/repo/releases/tag/v1.2.3", release.HTMLURL)
	assert.Equal(t, "Release notes here", release.Body)

	expectedTime := time.Date(2025, 7, 27, 10, 30, 0, 0, time.UTC)
	assert.Equal(t, expectedTime, release.PublishedAt)
}

func TestReleaseIsStable(t *testing.T) {
	tests := []struct {
		name       string
		release    Release
		wantStable bool
	}{
		{
			name:       "stable release",
			release:    Release{Draft: false, Prerelease: false},
			wantStable: true,
		},
		{
			name:       "draft release",
			release:    Release{Draft: true, Prerelease: false},
			wantStable: false,
		},
		{
			name:       "prerelease",
			release:    Release{Draft: false, Prerelease: true},
			wantStable: false,
		},
		{
			name:       "draft prerelease",
			release:    Release{Draft: true, Prerelease: true},
			wantStable: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.release.IsStable()
			assert.Equal(t, tt.wantStable, got)
		})
	}
}

func TestReleaseGetVersion(t *testing.T) {
	tests := []struct {
		name    string
		tagName string
		want    string
	}{
		{
			name:    "version with v prefix",
			tagName: "v1.2.3",
			want:    "1.2.3",
		},
		{
			name:    "version without v prefix",
			tagName: "1.2.3",
			want:    "1.2.3",
		},
		{
			name:    "empty version",
			tagName: "",
			want:    "",
		},
		{
			name:    "only v",
			tagName: "v",
			want:    "",
		},
		{
			name:    "complex version with v",
			tagName: "v1.2.3-alpha.1",
			want:    "1.2.3-alpha.1",
		},
		{
			name:    "complex version without v",
			tagName: "1.2.3-alpha.1",
			want:    "1.2.3-alpha.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			release := Release{TagName: tt.tagName}
			got := release.GetVersion()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReleaseToReleaseInfo(t *testing.T) {
	release := Release{
		TagName:     "v1.2.3",
		Name:        "Release v1.2.3",
		Body:        "Release notes",
		HTMLURL:     "https://github.com/test/repo/releases/tag/v1.2.3",
		PublishedAt: time.Date(2025, 7, 27, 10, 30, 0, 0, time.UTC),
		Draft:       false,
		Prerelease:  false,
	}

	releaseInfo := release.ToReleaseInfo()

	assert.Equal(t, "1.2.3", releaseInfo.Version)
	assert.Equal(t, "v1.2.3", releaseInfo.TagName)
	assert.Equal(t, "Release v1.2.3", releaseInfo.Name)
	assert.Equal(t, "Release notes", releaseInfo.Body)
	assert.Equal(t, "https://github.com/test/repo/releases/tag/v1.2.3", releaseInfo.URL)
	assert.Equal(t, time.Date(2025, 7, 27, 10, 30, 0, 0, time.UTC), releaseInfo.PublishedAt)
	assert.True(t, releaseInfo.IsStable)
}

func TestReleaseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		check   func(t *testing.T, release Release)
	}{
		{
			name: "RFC3339 date format",
			json: `{
				"tag_name": "v1.0.0",
				"published_at": "2025-07-27T10:30:00Z"
			}`,
			wantErr: false,
			check: func(t *testing.T, release Release) {
				expected := time.Date(2025, 7, 27, 10, 30, 0, 0, time.UTC)
				assert.Equal(t, expected, release.PublishedAt)
			},
		},
		{
			name: "alternative date format",
			json: `{
				"tag_name": "v1.0.0",
				"published_at": "2025-07-27T10:30:00-07:00"
			}`,
			wantErr: false,
			check: func(t *testing.T, release Release) {
				assert.False(t, release.PublishedAt.IsZero())
			},
		},
		{
			name: "null published_at",
			json: `{
				"tag_name": "v1.0.0",
				"published_at": null
			}`,
			wantErr: false,
			check: func(t *testing.T, release Release) {
				assert.True(t, release.PublishedAt.IsZero())
			},
		},
		{
			name: "invalid date format",
			json: `{
				"tag_name": "v1.0.0",
				"published_at": "invalid-date"
			}`,
			wantErr: false,
			check: func(t *testing.T, release Release) {
				assert.True(t, release.PublishedAt.IsZero())
			},
		},
		{
			name:    "invalid JSON",
			json:    `{"tag_name": "v1.0.0"`,
			wantErr: true,
			check:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var release Release
			err := json.Unmarshal([]byte(tt.json), &release)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				if tt.check != nil {
					tt.check(t, release)
				}
			}
		})
	}
}
