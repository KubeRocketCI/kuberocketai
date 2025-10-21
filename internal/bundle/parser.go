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
package bundle

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ExtractedFile represents a file extracted from a bundle
type ExtractedFile struct {
	Path    string
	Content string
}

// ParseBundle parses a bundle markdown file and extracts all files
func ParseBundle(reader io.Reader) ([]ExtractedFile, error) {
	scanner := bufio.NewScanner(reader)
	var files []ExtractedFile
	var currentFile *ExtractedFile
	var contentBuilder strings.Builder

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Check for file start delimiter
		if strings.HasPrefix(line, FileStartDelimiter) && strings.HasSuffix(line, " ====") {
			// Save previous file if exists
			if currentFile != nil {
				return nil, fmt.Errorf("line %d: found new file start before previous file ended (missing END FILE delimiter)", lineNum)
			}

			// Extract file path
			filePath := strings.TrimSuffix(strings.TrimPrefix(line, FileStartDelimiter), " ====")
			if filePath == "" {
				return nil, fmt.Errorf("line %d: empty file path in FILE delimiter", lineNum)
			}

			currentFile = &ExtractedFile{Path: filePath}
			contentBuilder.Reset()
			continue
		}

		// Check for file end delimiter
		if strings.TrimSpace(line) == FileEndDelimiter {
			if currentFile == nil {
				return nil, fmt.Errorf("line %d: found END FILE delimiter without matching FILE start", lineNum)
			}

			// Save the file
			currentFile.Content = contentBuilder.String()
			files = append(files, *currentFile)
			currentFile = nil
			contentBuilder.Reset()
			continue
		}

		// Collect file content
		if currentFile != nil {
			if contentBuilder.Len() > 0 {
				contentBuilder.WriteString("\n")
			}
			contentBuilder.WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading bundle: %w", err)
	}

	// Check for unclosed file
	if currentFile != nil {
		return nil, fmt.Errorf("unclosed file at end of bundle: %s (missing END FILE delimiter)", currentFile.Path)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no files found in bundle")
	}

	return files, nil
}
