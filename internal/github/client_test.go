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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	assert.Equal(t, "https://api.github.com", client.BaseURL)
	assert.NotEmpty(t, client.UserAgent)
	assert.NotNil(t, client.HTTPClient)
}

func TestClientWithCustomTimeout(t *testing.T) {
	client := NewClient()

	// Test that timeout is reasonable (should be set)
	assert.NotZero(t, client.HTTPClient.Timeout)
}

func TestGetLatestRelease(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/repos/test/repo/releases/latest", r.URL.Path)
		assert.NotEmpty(t, r.Header.Get("User-Agent"))

		release := Release{
			TagName:     "v1.0.0",
			PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
			HTMLURL:     "https://github.com/test/repo/releases/tag/v1.0.0",
			Body:        "Release notes",
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(release)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	release, err := client.GetLatestRelease("test", "repo")

	require.NoError(t, err)
	assert.Equal(t, "v1.0.0", release.TagName)
	assert.Equal(t, "https://github.com/test/repo/releases/tag/v1.0.0", release.HTMLURL)
	assert.Equal(t, "Release notes", release.Body)
}

func TestGetLatestReleaseHTTPError(t *testing.T) {
	// Create test server that returns HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("Not found"))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	release, err := client.GetLatestRelease("test", "repo")

	assert.Error(t, err)
	assert.Nil(t, release)
}

func TestGetLatestReleaseJSONError(t *testing.T) {
	// Create test server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("invalid json"))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	_, err := client.GetLatestRelease("test", "repo")
	assert.Error(t, err)
}

func TestGetReleases(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/repos/test/repo/releases", r.URL.Path)
		assert.NotEmpty(t, r.Header.Get("User-Agent"))

		releases := []Release{
			{
				TagName:     "v2.0.0",
				PublishedAt: time.Date(2025, 7, 28, 10, 0, 0, 0, time.UTC),
				HTMLURL:     "https://github.com/test/repo/releases/tag/v2.0.0",
				Body:        "Latest release",
			},
			{
				TagName:     "v1.0.0",
				PublishedAt: time.Date(2025, 7, 27, 10, 0, 0, 0, time.UTC),
				HTMLURL:     "https://github.com/test/repo/releases/tag/v1.0.0",
				Body:        "First release",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(releases)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	releases, err := client.GetReleases("test", "repo", 1, 30)

	require.NoError(t, err)
	assert.Len(t, releases, 2)
	assert.Equal(t, "v2.0.0", releases[0].TagName)
	assert.Equal(t, "v1.0.0", releases[1].TagName)
}

func TestGetReleasesError(t *testing.T) {
	// Create test server that returns HTTP error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	releases, err := client.GetReleases("test", "repo", 1, 30)

	assert.Error(t, err)
	assert.Nil(t, releases)
}

func TestGetReleasesJSONError(t *testing.T) {
	// Create test server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("invalid json"))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient()
	client.BaseURL = server.URL

	releases, err := client.GetReleases("test", "repo", 1, 30)

	assert.Error(t, err)
	assert.Nil(t, releases)
}

func TestGetLatestReleaseRequestError(t *testing.T) {
	client := NewClient()
	client.BaseURL = "invalid-url"

	_, err := client.GetLatestRelease("test", "repo")
	assert.Error(t, err)
}

func TestGetReleasesRequestError(t *testing.T) {
	client := NewClient()
	client.BaseURL = "invalid-url"

	_, err := client.GetReleases("test", "repo", 1, 30)
	assert.Error(t, err)
}
