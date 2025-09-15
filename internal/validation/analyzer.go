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
package validation

import (
	"fmt"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
)

// Severity levels for validation issues
type Severity int

const (
	SeverityInfo Severity = iota
	SeverityWarning
	SeverityError
	SeverityCritical
)

func (s Severity) String() string {
	switch s {
	case SeverityInfo:
		return "INFO"
	case SeverityWarning:
		return "WARNING"
	case SeverityCritical:
		return "CRITICAL"
	default:
		return "ERROR"
	}
}

// ValidationIssue represents a single validation issue
type ValidationIssue struct {
	Type        string
	Severity    Severity
	File        string
	Line        int
	Message     string
	FixGuidance string
}

// FrameworkAnalyzer provides comprehensive framework validation
type FrameworkAnalyzer struct {
	discovery *assets.Discovery
}

// NewFrameworkAnalyzer creates a new framework analyzer
func NewFrameworkAnalyzer(discovery *assets.Discovery) *FrameworkAnalyzer {
	return &FrameworkAnalyzer{
		discovery: discovery,
	}
}

// AnalyzeFramework performs comprehensive framework analysis
func (a *FrameworkAnalyzer) AnalyzeFramework() ([]ValidationIssue, *FrameworkInsights, error) {
	var issues []ValidationIssue

	// Phase 3: Generate framework insights
	insights, err := a.generateFrameworkInsights()
	if err != nil {
		return nil, nil, fmt.Errorf("framework insights generation failed: %w", err)
	}

	return issues, insights, nil
}
