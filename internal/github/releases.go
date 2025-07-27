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
	"time"
)

// Release represents a GitHub release
type Release struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	Body        string    `json:"body"`
	HTMLURL     string    `json:"html_url"`
	PublishedAt time.Time `json:"published_at"`
	Draft       bool      `json:"draft"`
	Prerelease  bool      `json:"prerelease"`
	TarballURL  string    `json:"tarball_url"`
	ZipballURL  string    `json:"zipball_url"`
}

// IsStable returns true if the release is stable (not draft or prerelease)
func (r *Release) IsStable() bool {
	return !r.Draft && !r.Prerelease
}

// GetVersion returns the version string, removing 'v' prefix if present
func (r *Release) GetVersion() string {
	version := r.TagName
	if len(version) > 0 && version[0] == 'v' {
		return version[1:]
	}
	return version
}

// ReleaseInfo contains processed release information
type ReleaseInfo struct {
	Version     string    `json:"version"`
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	Body        string    `json:"body"`
	URL         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
	IsStable    bool      `json:"is_stable"`
}

// ToReleaseInfo converts a Release to ReleaseInfo
func (r *Release) ToReleaseInfo() *ReleaseInfo {
	return &ReleaseInfo{
		Version:     r.GetVersion(),
		TagName:     r.TagName,
		Name:        r.Name,
		Body:        r.Body,
		URL:         r.HTMLURL,
		PublishedAt: r.PublishedAt,
		IsStable:    r.IsStable(),
	}
}

// UnmarshalJSON custom unmarshaling to handle various date formats
func (r *Release) UnmarshalJSON(data []byte) error {
	type Alias Release
	aux := &struct {
		PublishedAt interface{} `json:"published_at"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Handle different date formats
	switch v := aux.PublishedAt.(type) {
	case string:
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			r.PublishedAt = t
		} else if t, err := time.Parse("2006-01-02T15:04:05Z", v); err == nil {
			r.PublishedAt = t
		}
	case nil:
		r.PublishedAt = time.Time{}
	}

	return nil
}
