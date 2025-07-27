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
