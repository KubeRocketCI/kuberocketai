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
package changelog

import (
	"strings"
	"testing"
)

// Benchmark data - simulate large changelog
var largeChangelog = generateLargeChangelog()

func generateLargeChangelog() string {
	var sb strings.Builder
	sb.WriteString("# Changelog\n\n")

	// Generate 100 versions with multiple sections each
	for i := 100; i >= 1; i-- {
		sb.WriteString("## [v")
		sb.WriteString(string(rune('0' + i/100)))
		sb.WriteString(".")
		sb.WriteString(string(rune('0' + (i%100)/10)))
		sb.WriteString(".")
		sb.WriteString(string(rune('0' + i%10)))
		sb.WriteString("] - 2025-01-")
		if i < 10 {
			sb.WriteString("0")
		}
		sb.WriteString(string(rune('0' + i%32)))
		sb.WriteString("\n\n")

		// Add features section
		sb.WriteString("### Features\n\n")
		for j := 1; j <= 5; j++ {
			sb.WriteString("- Feature ")
			sb.WriteString(string(rune('0' + i)))
			sb.WriteString(".")
			sb.WriteString(string(rune('0' + j)))
			sb.WriteString(" - Some feature description\n")
		}
		sb.WriteString("\n")

		// Add bug fixes section
		sb.WriteString("### Bug Fixes\n\n")
		for j := 1; j <= 3; j++ {
			sb.WriteString("- Fix ")
			sb.WriteString(string(rune('0' + i)))
			sb.WriteString(".")
			sb.WriteString(string(rune('0' + j)))
			sb.WriteString(" - Some bug fix description\n")
		}
		sb.WriteString("\n")

		// Add documentation section
		if i%5 == 0 {
			sb.WriteString("### Documentation\n\n")
			sb.WriteString("- Update documentation for version ")
			sb.WriteString(string(rune('0' + i)))
			sb.WriteString("\n\n")
		}
	}

	return sb.String()
}

func BenchmarkChangelogParser(b *testing.B) {
	content := largeChangelog

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser := NewChangelogParser(content)
		_, _ = parser.ParseSections()
	}
}

func BenchmarkVersionParsing(b *testing.B) {
	content := largeChangelog

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser := NewChangelogParser(content)
		_ = parser.GetVersions()
	}
}

func BenchmarkSectionExtraction(b *testing.B) {
	content := largeChangelog

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		extractor := NewSectionExtractor(content)
		_, _ = extractor.ExtractAllSections()
	}
}

func BenchmarkValidation(b *testing.B) {
	content := largeChangelog

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ValidateChangelogDetailed(content)
	}
}

func BenchmarkFormatterLarge(b *testing.B) {
	content := largeChangelog

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FormatChangelog(content)
	}
}

func BenchmarkSectionFilter(b *testing.B) {
	extractor := NewSectionExtractor(largeChangelog)
	sections, _ := extractor.ExtractAllSections()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FilterSectionsByPattern(sections, "(?i)feature")
	}
}

func BenchmarkDisplayOptions(b *testing.B) {
	content := largeChangelog
	options := DefaultDisplayOptions()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DisplayChangelog(content, options)
	}
}

// Benchmark memory allocations
func BenchmarkParsingMemory(b *testing.B) {
	content := largeChangelog

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		parser := NewChangelogParser(content)
		sections, _ := parser.ParseSections()
		_ = sections
	}
}

func BenchmarkValidationMemory(b *testing.B) {
	content := largeChangelog

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := ValidateChangelogDetailed(content)
		_ = result
	}
}

// Cross-platform benchmarks for different input sizes
func BenchmarkSmallChangelog(b *testing.B) {
	content := `# Changelog

## [v1.0.0] - 2025-07-27

### Features
- Small feature

### Bug Fixes
- Small fix
`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser := NewChangelogParser(content)
		_, _ = parser.ParseSections()
	}
}

func BenchmarkMediumChangelog(b *testing.B) {
	var sb strings.Builder
	sb.WriteString("# Changelog\n\n")

	// Generate 10 versions
	for i := 10; i >= 1; i-- {
		sb.WriteString("## [v1.")
		sb.WriteString(string(rune('0' + i)))
		sb.WriteString(".0] - 2025-07-")
		if i < 10 {
			sb.WriteString("0")
		}
		sb.WriteString(string(rune('0' + i)))
		sb.WriteString("\n\n")

		sb.WriteString("### Features\n\n")
		for j := 1; j <= 3; j++ {
			sb.WriteString("- Feature ")
			sb.WriteString(string(rune('0' + j)))
			sb.WriteString("\n")
		}
		sb.WriteString("\n### Bug Fixes\n\n")
		for j := 1; j <= 2; j++ {
			sb.WriteString("- Fix ")
			sb.WriteString(string(rune('0' + j)))
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	content := sb.String()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parser := NewChangelogParser(content)
		_, _ = parser.ParseSections()
	}
}

// Benchmark concurrent access (for thread safety testing)
func BenchmarkConcurrentParsing(b *testing.B) {
	content := largeChangelog

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			parser := NewChangelogParser(content)
			_, _ = parser.ParseSections()
		}
	})
}

func BenchmarkConcurrentValidation(b *testing.B) {
	content := largeChangelog

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = ValidateChangelogDetailed(content)
		}
	})
}
