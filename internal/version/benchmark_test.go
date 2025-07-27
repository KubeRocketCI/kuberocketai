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
	"testing"
)

func BenchmarkParseVersion(b *testing.B) {
	versions := []string{
		"1.0.0",
		"v1.0.0",
		"1.2.3",
		"v1.2.3",
		"1.0.0-alpha.1",
		"v2.0.0-beta.2",
		"1.0.0-rc.1",
		"v1.0.0+build.1",
		"2.1.0-alpha.1+build.2",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		version := versions[i%len(versions)]
		_, _ = ParseVersion(version)
	}
}

func BenchmarkCompareVersions(b *testing.B) {
	versionPairs := [][2]string{
		{"1.0.0", "1.0.1"},
		{"v1.0.0", "v2.0.0"},
		{"1.0.0-alpha.1", "1.0.0"},
		{"v2.0.0", "v1.9.9"},
		{"1.0.0+build.1", "1.0.0+build.2"},
		{"1.0.0-beta.1", "1.0.0-alpha.1"},
		{"2.0.0", "1.9.9"},
		{"1.0.0", "1.0.0"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pair := versionPairs[i%len(versionPairs)]
		_, _ = CompareVersions(pair[0], pair[1])
	}
}

func BenchmarkGetCurrentVersion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetCurrentVersion()
	}
}

func BenchmarkGetVersionInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetVersionInfo()
	}
}

func BenchmarkVersionInfoString(b *testing.B) {
	info := GetVersionInfo()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = info.String()
	}
}

// Benchmark memory allocations
func BenchmarkParseVersionMemory(b *testing.B) {
	version := "v1.2.3-alpha.1+build.123"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v, _ := ParseVersion(version)
		_ = v
	}
}

func BenchmarkCompareVersionsMemory(b *testing.B) {
	current := "v1.0.0"
	latest := "v2.0.0"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result, _ := CompareVersions(current, latest)
		_ = result
	}
}

func BenchmarkVersionInfoMemory(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		info := GetVersionInfo()
		_ = info
	}
}

// Concurrent version operations
func BenchmarkConcurrentParseVersion(b *testing.B) {
	version := "v1.2.3-alpha.1"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = ParseVersion(version)
		}
	})
}

func BenchmarkConcurrentCompareVersions(b *testing.B) {
	current := "v1.0.0"
	latest := "v2.0.0"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = CompareVersions(current, latest)
		}
	})
}

func BenchmarkConcurrentVersionInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GetVersionInfo()
		}
	})
}

// Edge case version parsing
func BenchmarkComplexVersions(b *testing.B) {
	complexVersions := []string{
		"1.0.0-alpha+dev.build.123",
		"v2.0.0-beta.2+exp.sha.5114f85",
		"1.0.0-x.7.z.92+alpha.1",
		"v1.0.0-alpha0.valid+beta",
		"1.0.0+0.build.1-rc.10000aaa-kk-0.1",
		"v99999999999999999999999.999999999999999999.99999999999999999",
		"1.0.0-0A.is.legal",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		version := complexVersions[i%len(complexVersions)]
		_, _ = ParseVersion(version)
	}
}

// Benchmark invalid version handling
func BenchmarkInvalidVersions(b *testing.B) {
	invalidVersions := []string{
		"invalid",
		"1.2",
		"1.2.3.4",
		"",
		"v",
		"1.2.3-",
		"1.2.3+",
		"1.2.3-+",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		version := invalidVersions[i%len(invalidVersions)]
		_, _ = ParseVersion(version)
	}
}
