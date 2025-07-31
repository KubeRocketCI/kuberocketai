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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client.BaseURL != "https://api.github.com" {
		t.Errorf("Expected BaseURL to be https://api.github.com, got %s", client.BaseURL)
	}

	if client.UserAgent == "" {
		t.Error("UserAgent should not be empty")
	}

	if client.HTTPClient == nil {
		t.Error("HTTPClient should not be nil")
	}
}

func TestGetLatestRelease(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/test/repo/releases/latest" {
			t.Errorf("Expected path /repos/test/repo/releases/latest, got %s", r.URL.Path)
		}

		if r.Header.Get("User-Agent") == "" {
			t.Error("User-Agent header should be set")
		}

		release := Release{
			TagName:     "v1.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v1.0.0",
			Body:        "Release notes",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	release, err := client.GetLatestRelease("test", "repo")
	if err != nil {
		t.Fatalf("GetLatestRelease failed: %v", err)
	}

	if release.TagName != "v1.0.0" {
		t.Errorf("Expected tag name v1.0.0, got %s", release.TagName)
	}

	if release.HTMLURL != "https://github.com/test/repo/releases/tag/v1.0.0" {
		t.Errorf("Expected URL https://github.com/test/repo/releases/tag/v1.0.0, got %s", release.HTMLURL)
	}

	if release.Body != "Release notes" {
		t.Errorf("Expected body 'Release notes', got %s", release.Body)
	}
}

func TestGetLatestReleaseHTTPError(t *testing.T) {
	// Create test server that returns 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	_, err := client.GetLatestRelease("test", "repo")
	if err == nil {
		t.Error("Expected error for 404 response")
	}
}

func TestGetLatestReleaseJSONError(t *testing.T) {
	// Create test server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	_, err := client.GetLatestRelease("test", "repo")
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestClientWithCustomTimeout(t *testing.T) {
	client := NewClient()

	// Test that timeout is reasonable (should be set)
	if client.HTTPClient.Timeout == 0 {
		t.Error("HTTP client should have a timeout set")
	}
}

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
	if err != nil {
		t.Fatalf("Failed to unmarshal release JSON: %v", err)
	}

	if release.TagName != "v1.2.3" {
		t.Errorf("Expected TagName v1.2.3, got %s", release.TagName)
	}

	expectedTime := time.Date(2025, 7, 27, 10, 30, 0, 0, time.UTC)
	if !release.PublishedAt.Equal(expectedTime) {
		t.Errorf("Expected PublishedAt %v, got %v", expectedTime, release.PublishedAt)
	}

	if release.HTMLURL != "https://github.com/test/repo/releases/tag/v1.2.3" {
		t.Errorf("Expected HTMLURL https://github.com/test/repo/releases/tag/v1.2.3, got %s", release.HTMLURL)
	}

	if release.Body != "Release notes here" {
		t.Errorf("Expected Body 'Release notes here', got %s", release.Body)
	}
}

func TestGetReleases(t *testing.T) {
	// Create test server that returns paginated releases
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/repos/test/repo/releases"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Check query parameters
		page := r.URL.Query().Get("page")
		perPage := r.URL.Query().Get("per_page")
		if page != "1" || perPage != "10" {
			t.Errorf("Expected page=1&per_page=10, got page=%s&per_page=%s", page, perPage)
		}

		releases := []Release{
			{
				TagName:     "v1.2.0",
				Name:        "Release 1.2.0",
				Body:        "Latest release",
				HTMLURL:     "https://github.com/test/repo/releases/tag/v1.2.0",
				PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
				Draft:       false,
				Prerelease:  false,
			},
			{
				TagName:     "v1.1.0",
				Name:        "Release 1.1.0",
				Body:        "Previous release",
				HTMLURL:     "https://github.com/test/repo/releases/tag/v1.1.0",
				PublishedAt: time.Date(2025, 7, 20, 10, 0, 0, 0, time.UTC),
				Draft:       false,
				Prerelease:  true,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(releases)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	releases, err := client.GetReleases("test", "repo", 1, 10)
	if err != nil {
		t.Fatalf("GetReleases failed: %v", err)
	}

	if len(releases) != 2 {
		t.Errorf("Expected 2 releases, got %d", len(releases))
	}

	// Test first release
	if releases[0].TagName != "v1.2.0" {
		t.Errorf("Expected first release TagName v1.2.0, got %s", releases[0].TagName)
	}
	if releases[0].Name != "Release 1.2.0" {
		t.Errorf("Expected first release Name 'Release 1.2.0', got %s", releases[0].Name)
	}

	// Test second release
	if releases[1].TagName != "v1.1.0" {
		t.Errorf("Expected second release TagName v1.1.0, got %s", releases[1].TagName)
	}
	if releases[1].Prerelease != true {
		t.Errorf("Expected second release to be prerelease")
	}
}

func TestGetReleasesError(t *testing.T) {
	// Test HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	_, err := client.GetReleases("test", "repo", 1, 10)
	if err == nil {
		t.Error("Expected error for 500 response")
	}
}

func TestGetReleasesJSONError(t *testing.T) {
	// Test JSON decode error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("invalid json array"))
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	_, err := client.GetReleases("test", "repo", 1, 10)
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestReleaseIsStable(t *testing.T) {
	tests := []struct {
		name       string
		draft      bool
		prerelease bool
		want       bool
	}{
		{"stable release", false, false, true},
		{"draft release", true, false, false},
		{"prerelease", false, true, false},
		{"draft prerelease", true, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			release := &Release{
				Draft:      tt.draft,
				Prerelease: tt.prerelease,
			}
			if got := release.IsStable(); got != tt.want {
				t.Errorf("IsStable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReleaseGetVersion(t *testing.T) {
	tests := []struct {
		name    string
		tagName string
		want    string
	}{
		{"version with v prefix", "v1.2.3", "1.2.3"},
		{"version without v prefix", "1.2.3", "1.2.3"},
		{"empty version", "", ""},
		{"only v", "v", ""},
		{"complex version with v", "v1.2.3-alpha.1", "1.2.3-alpha.1"},
		{"complex version without v", "1.2.3-alpha.1", "1.2.3-alpha.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			release := &Release{TagName: tt.tagName}
			if got := release.GetVersion(); got != tt.want {
				t.Errorf("GetVersion() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestReleaseToReleaseInfo(t *testing.T) {
	publishedAt := time.Date(2025, 7, 27, 10, 30, 0, 0, time.UTC)
	release := &Release{
		TagName:     "v1.2.3",
		Name:        "Release 1.2.3",
		Body:        "Release notes",
		HTMLURL:     "https://github.com/test/repo/releases/tag/v1.2.3",
		PublishedAt: publishedAt,
		Draft:       false,
		Prerelease:  false,
	}

	info := release.ToReleaseInfo()

	if info.Version != "1.2.3" {
		t.Errorf("Expected Version 1.2.3, got %s", info.Version)
	}
	if info.TagName != "v1.2.3" {
		t.Errorf("Expected TagName v1.2.3, got %s", info.TagName)
	}
	if info.Name != "Release 1.2.3" {
		t.Errorf("Expected Name 'Release 1.2.3', got %s", info.Name)
	}
	if info.Body != "Release notes" {
		t.Errorf("Expected Body 'Release notes', got %s", info.Body)
	}
	if info.URL != "https://github.com/test/repo/releases/tag/v1.2.3" {
		t.Errorf("Expected URL https://github.com/test/repo/releases/tag/v1.2.3, got %s", info.URL)
	}
	if !info.PublishedAt.Equal(publishedAt) {
		t.Errorf("Expected PublishedAt %v, got %v", publishedAt, info.PublishedAt)
	}
	if !info.IsStable {
		t.Error("Expected IsStable to be true")
	}
}

func TestReleaseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
		checkFn  func(*testing.T, *Release)
	}{
		{
			name: "RFC3339 date format",
			jsonData: `{
				"tag_name": "v1.0.0",
				"published_at": "2025-07-27T10:30:00Z"
			}`,
			wantErr: false,
			checkFn: func(t *testing.T, r *Release) {
				expected := time.Date(2025, 7, 27, 10, 30, 0, 0, time.UTC)
				if !r.PublishedAt.Equal(expected) {
					t.Errorf("Expected PublishedAt %v, got %v", expected, r.PublishedAt)
				}
			},
		},
		{
			name: "alternative date format",
			jsonData: `{
				"tag_name": "v1.0.0",
				"published_at": "2025-07-27T10:30:00Z"
			}`,
			wantErr: false,
			checkFn: func(t *testing.T, r *Release) {
				if r.PublishedAt.IsZero() {
					t.Error("Expected PublishedAt to be parsed")
				}
			},
		},
		{
			name: "null published_at",
			jsonData: `{
				"tag_name": "v1.0.0",
				"published_at": null
			}`,
			wantErr: false,
			checkFn: func(t *testing.T, r *Release) {
				if !r.PublishedAt.IsZero() {
					t.Error("Expected PublishedAt to be zero for null value")
				}
			},
		},
		{
			name: "invalid date format",
			jsonData: `{
				"tag_name": "v1.0.0",
				"published_at": "invalid-date"
			}`,
			wantErr: false,
			checkFn: func(t *testing.T, r *Release) {
				// Should not error, but date should remain zero
				if !r.PublishedAt.IsZero() {
					t.Error("Expected PublishedAt to be zero for invalid date")
				}
			},
		},
		{
			name:     "invalid JSON",
			jsonData: `{"tag_name": "v1.0.0", "published_at":}`,
			wantErr:  true,
			checkFn:  func(t *testing.T, r *Release) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var release Release
			err := json.Unmarshal([]byte(tt.jsonData), &release)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				tt.checkFn(t, &release)
			}
		})
	}
}

// Test fallback functions
func TestGetReleasesURL(t *testing.T) {
	tests := []struct {
		owner string
		repo  string
		want  string
	}{
		{"owner1", "repo1", "https://github.com/owner1/repo1/releases"},
		{"KubeRocketCI", "kuberocketai", "https://github.com/KubeRocketCI/kuberocketai/releases"},
		{"", "", "https://github.com///releases"},
	}

	for _, tt := range tests {
		t.Run(tt.owner+"/"+tt.repo, func(t *testing.T) {
			if got := GetReleasesURL(tt.owner, tt.repo); got != tt.want {
				t.Errorf("GetReleasesURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFallbackMessage(t *testing.T) {
	tests := []struct {
		owner string
		repo  string
		want  string
	}{
		{
			"test", "repo",
			"Unable to check for updates. Visit: https://github.com/test/repo/releases to check manually",
		},
		{
			"KubeRocketCI", "kuberocketai",
			"Unable to check for updates. Visit: https://github.com/KubeRocketCI/kuberocketai/releases to check manually",
		},
	}

	for _, tt := range tests {
		t.Run(tt.owner+"/"+tt.repo, func(t *testing.T) {
			if got := GetFallbackMessage(tt.owner, tt.repo); got != tt.want {
				t.Errorf("GetFallbackMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKuberocketaiReleasesURL(t *testing.T) {
	expected := "https://github.com/KubeRocketCI/kuberocketai/releases"
	if got := GetKuberocketaiReleasesURL(); got != expected {
		t.Errorf("GetKuberocketaiReleasesURL() = %v, want %v", got, expected)
	}
}

func TestGetKuberocketaiFallbackMessage(t *testing.T) {
	expected := "Unable to check for updates. Visit: https://github.com/KubeRocketCI/kuberocketai/releases to check manually"
	if got := GetKuberocketaiFallbackMessage(); got != expected {
		t.Errorf("GetKuberocketaiFallbackMessage() = %v, want %v", got, expected)
	}
}

func TestGetLatestReleaseRequestError(t *testing.T) {
	client := NewClient()
	// Use an invalid URL to trigger request creation error
	client.BaseURL = "://invalid-url"

	_, err := client.GetLatestRelease("test", "repo")
	if err == nil {
		t.Error("Expected error for invalid URL")
	}
}

func TestGetReleasesRequestError(t *testing.T) {
	client := NewClient()
	// Use an invalid URL to trigger request creation error
	client.BaseURL = "://invalid-url"

	_, err := client.GetReleases("test", "repo", 1, 10)
	if err == nil {
		t.Error("Expected error for invalid URL")
	}
}
