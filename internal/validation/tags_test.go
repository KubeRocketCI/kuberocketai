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
	"os"
	"path/filepath"
	"testing"

	"github.com/KubeRocketCI/kuberocketai/internal/assets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateXMLTags(t *testing.T) {
	discovery := &assets.Discovery{}
	analyzer := NewFrameworkAnalyzer(discovery)

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name: "valid_nested_tags",
			content: `<prerequisites>
- Framework installed: .krci-ai directory exists
</prerequisites>

<instructions>
1. Apply schema requirements
<nested>
Some nested content
</nested>
2. Include activation prompt
</instructions>`,
			expected: nil,
		},
		{
			name: "self_closing_tags",
			content: `<tag1/>
<tag2 attribute="value"/>
<tag3 />`,
			expected: nil,
		},
		{
			name: "unclosed_tag",
			content: `<prerequisites>
- Some content
<instructions>
More content
</instructions>`,
			expected: []string{"Unclosed tag <prerequisites>"},
		},
		{
			name: "mismatched_tags",
			content: `<prerequisites>
Content here
</instructions>`,
			expected: []string{
				"Closing tag </instructions> without matching opening tag",
				"Unclosed tag <prerequisites>",
			},
		},
		{
			name: "multiple_unclosed_tags",
			content: `<tag1>
Content 1
<tag2>
Content 2
<tag3>
Content 3`,
			expected: []string{
				"Unclosed tag <tag1>",
				"Unclosed tag <tag2>",
				"Unclosed tag <tag3>",
			},
		},
		{
			name: "closing_without_opening",
			content: `Some content
</prerequisites>
More content`,
			expected: []string{"Closing tag </prerequisites> without matching opening tag"},
		},
		{
			name: "mixed_valid_invalid",
			content: `<valid>
Valid content
</valid>

<invalid>
Invalid content

<another-valid>
More valid content
</another-valid>`,
			expected: []string{"Unclosed tag <invalid>"},
		},
		{
			name: "tags_with_attributes",
			content: `<tag attr="value" class="test">
Content
</tag>

<unclosed attr="value">
Content`,
			expected: []string{"Unclosed tag <unclosed>"},
		},
		{
			name: "nested_unclosed",
			content: `<outer>
<inner>
Content
</outer>`,
			expected: []string{"Unclosed tag <inner>"},
		},
		{
			name:     "empty_content",
			content:  "",
			expected: nil,
		},
		{
			name:     "no_tags",
			content:  "This is just regular markdown content without any tags.",
			expected: nil,
		},
		{
			name: "tags_with_hyphens_underscores",
			content: `<test-tag>
Content
</test-tag>

<test_tag>
Content
</test_tag>

<unclosed-tag>
Content`,
			expected: []string{"Unclosed tag <unclosed-tag>"},
		},
		{
			name: "tags_with_namespaces",
			content: `<ns:tag>
Content
</ns:tag>

<xml:unclosed>
Content`,
			expected: []string{"Unclosed tag <xml:unclosed>"},
		},
		{
			name: "deeply_nested_unclosed",
			content: `<level1>
<level2>
<level3>
<level4>
Content without any closing tags`,
			expected: []string{
				"Unclosed tag <level1>",
				"Unclosed tag <level2>",
				"Unclosed tag <level3>",
				"Unclosed tag <level4>",
			},
		},
		{
			name: "mixed_nested_some_closed",
			content: `<outer>
<middle1>
Content1
</middle1>
<middle2>
<inner>
Content2
</outer>`,
			expected: []string{
				"Unclosed tag <inner>",
				"Unclosed tag <middle2>",
			},
		},
		{
			name: "wrong_nesting_order",
			content: `<tag1>
<tag2>
</tag1>
</tag2>`,
			expected: []string{
				"Unclosed tag <tag2>",
				"Closing tag </tag2> without matching opening tag",
			},
		},
		{
			name: "self_closing_with_unclosed",
			content: `<self-closing/>
<unclosed>
<self-closing-2 attr="value"/>
Content`,
			expected: []string{"Unclosed tag <unclosed>"},
		},
		{
			name: "empty_tags",
			content: `<empty></empty>
<empty-unclosed>
<another-empty></another-empty>`,
			expected: []string{"Unclosed tag <empty-unclosed>"},
		},
		{
			name: "tags_with_complex_attributes",
			content: `<tag attr1="value1" attr2='value2' attr3=value3 data-test="test">
Content
</tag>

<unclosed-with-attrs class="test" id="myid" style="color: red;">
Content`,
			expected: []string{"Unclosed tag <unclosed-with-attrs>"},
		},
		{
			name: "multiple_closing_tags_no_opening",
			content: `Content before
</first>
</second>
</third>
Content after`,
			expected: []string{
				"Closing tag </first> without matching opening tag",
				"Closing tag </second> without matching opening tag",
				"Closing tag </third> without matching opening tag",
			},
		},
		{
			name: "alternating_valid_invalid",
			content: `<valid1>
Content
</valid1>
<invalid1>
<valid2>
Content
</valid2>
<invalid2>
More content`,
			expected: []string{
				"Unclosed tag <invalid1>",
				"Unclosed tag <invalid2>",
			},
		},
		{
			name: "case_sensitive_tags",
			content: `<Tag>
Content
</tag>
<TAG>
Content
</Tag>`,
			expected: []string{
				"Closing tag </tag> without matching opening tag",
				"Unclosed tag <TAG>",
			},
		},
		{
			name: "tags_with_numbers",
			content: `<tag1>
<tag2>
Content
</tag2>
</tag1>
<tag3>
<tag4>
Unclosed content`,
			expected: []string{
				"Unclosed tag <tag3>",
				"Unclosed tag <tag4>",
			},
		},
		{
			name: "whitespace_in_tag_names",
			content: `<valid-tag>
Content
</valid-tag>
< invalid tag >
Content`,
			expected: []string{"Unclosed tag <invalid>"},
		},
		{
			name: "nested_same_tag_names",
			content: `<section>
<section>
Inner content
</section>
Outer content
</section>

<div>
<div>
<div>
Triple nested
</div>
Missing two closes`,
			expected: []string{
				"Unclosed tag <div>",
				"Unclosed tag <div>",
			},
		},
		{
			name: "xml_declaration_and_comments_ignored",
			content: `<?xml version="1.0"?>
<!-- This is a comment -->
<root>
<child>
Content
</child>
</root>
<!-- Another comment -->
<unclosed>
Content`,
			expected: []string{"Unclosed tag <unclosed>"},
		},
		{
			name: "inline_code_blocks_skipped",
			content: `<valid>
Content here
</valid>

Example usage: ` + "`<skip-this-tag>`" + ` and ` + "`<skip-this-too>`" + `

<another-valid>
More content
</another-valid>

<unclosed>
This should be detected`,
			expected: []string{"Unclosed tag <unclosed>"},
		},
		{
			name: "fenced_code_blocks_skipped",
			content: `<instructions>
Follow these steps
</instructions>

` + "```xml" + `
<example-tag>
  <nested>This is just an example</nested>
</example-tag>
` + "```" + `

<prerequisites>
Requirements here
</prerequisites>

` + "```go" + `
func main() {
    fmt.Println("<ignore-this>")
}
` + "```" + `

<unclosed-real>
This should be caught`,
			expected: []string{"Unclosed tag <unclosed-real>"},
		},
		{
			name: "tilde_code_blocks_skipped",
			content: `<success_criteria>
Success defined
</success_criteria>

~~~html
<div class="example">
  <span>Example content</span>
</div>
~~~

<another-unclosed>
This should be detected`,
			expected: []string{"Unclosed tag <another-unclosed>"},
		},
		{
			name: "mixed_code_blocks_and_real_xml",
			content: `<prerequisites>
- Real requirement
</prerequisites>

Inline example: ` + "`<inline-example>`" + `

` + "```" + `
<code-block-tag>
This should be ignored
</code-block-tag>
` + "```" + `

<instructions>
1. Real instruction
2. Another step
</instructions>

~~~yaml
config:
  - name: "<yaml-example>"
~~~

<unclosed-in-real-content>
This is a real unclosed tag`,
			expected: []string{"Unclosed tag <unclosed-in-real-content>"},
		},
		{
			name: "complex_markdown_with_code_examples",
			content: `# Documentation

<overview>
This document explains XML tags.
</overview>

## Examples

Here's how to use ` + "`<tag>`" + ` in your code:

` + "```xml" + `
<example>
  <child attr="value">
    Content here
  </child>
</example>
` + "```" + `

The ` + "`<child>`" + ` tag supports attributes.

<implementation>
Follow these steps:
1. Create the structure
2. Add content
</implementation>

Code sample:
~~~javascript
function createTag() {
    return '<generated>content</generated>';
}
~~~

<validation-error>
Missing closing tag`,
			expected: []string{"Unclosed tag <validation-error>"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := analyzer.validateXMLTags(tt.content)
			assert.Equal(t, tt.expected, issues, "XML validation issues should match expected")
		})
	}
}

func TestParseXMLTags(t *testing.T) {
	discovery := &assets.Discovery{}
	analyzer := NewFrameworkAnalyzer(discovery)

	tests := []struct {
		name     string
		content  string
		expected []xmlTag
	}{
		{
			name:    "simple_tags",
			content: `<tag1>content</tag1>`,
			expected: []xmlTag{
				{Name: "tag1", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "tag1", IsClosing: true, IsSelfClose: false, Position: 14},
			},
		},
		{
			name:    "self_closing",
			content: `<tag/>`,
			expected: []xmlTag{
				{Name: "tag", IsClosing: false, IsSelfClose: true, Position: 0},
			},
		},
		{
			name:    "with_attributes",
			content: `<tag attr="value" class="test">content</tag>`,
			expected: []xmlTag{
				{Name: "tag", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "tag", IsClosing: true, IsSelfClose: false, Position: 37},
			},
		},
		{
			name:    "with_whitespace",
			content: `<tag >content</tag >`,
			expected: []xmlTag{
				{Name: "tag", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "tag", IsClosing: true, IsSelfClose: false, Position: 13},
			},
		},
		{
			name:    "nested_tags",
			content: `<outer><inner>content</inner></outer>`,
			expected: []xmlTag{
				{Name: "outer", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "inner", IsClosing: false, IsSelfClose: false, Position: 7},
				{Name: "inner", IsClosing: true, IsSelfClose: false, Position: 21},
				{Name: "outer", IsClosing: true, IsSelfClose: false, Position: 29},
			},
		},
		{
			name:    "tags_with_special_chars",
			content: `<test-tag_name:namespace>content</test-tag_name:namespace>`,
			expected: []xmlTag{
				{Name: "test-tag_name:namespace", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "test-tag_name:namespace", IsClosing: true, IsSelfClose: false, Position: 32},
			},
		},
		{
			name:     "no_tags",
			content:  "Just regular content",
			expected: nil,
		},
		{
			name:    "mixed_self_closing_styles",
			content: `<br/><hr /><img src="test" />`,
			expected: []xmlTag{
				{Name: "br", IsClosing: false, IsSelfClose: true, Position: 0},
				{Name: "hr", IsClosing: false, IsSelfClose: true, Position: 5},
				{Name: "img", IsClosing: false, IsSelfClose: true, Position: 12},
			},
		},
		{
			name:    "tags_with_underscores_and_numbers",
			content: `<tag_1>content</tag_1><element2>text</element2>`,
			expected: []xmlTag{
				{Name: "tag_1", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "tag_1", IsClosing: true, IsSelfClose: false, Position: 14},
				{Name: "element2", IsClosing: false, IsSelfClose: false, Position: 22},
				{Name: "element2", IsClosing: true, IsSelfClose: false, Position: 37},
			},
		},
		{
			name:    "complex_nested_structure",
			content: `<html><head><title>Test</title></head><body><div class="main"><p>Content</p></div></body></html>`,
			expected: []xmlTag{
				{Name: "html", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "head", IsClosing: false, IsSelfClose: false, Position: 6},
				{Name: "title", IsClosing: false, IsSelfClose: false, Position: 12},
				{Name: "title", IsClosing: true, IsSelfClose: false, Position: 22},
				{Name: "head", IsClosing: true, IsSelfClose: false, Position: 30},
				{Name: "body", IsClosing: false, IsSelfClose: false, Position: 37},
				{Name: "div", IsClosing: false, IsSelfClose: false, Position: 43},
				{Name: "p", IsClosing: false, IsSelfClose: false, Position: 62},
				{Name: "p", IsClosing: true, IsSelfClose: false, Position: 73},
				{Name: "div", IsClosing: true, IsSelfClose: false, Position: 77},
				{Name: "body", IsClosing: true, IsSelfClose: false, Position: 83},
				{Name: "html", IsClosing: true, IsSelfClose: false, Position: 90},
			},
		},
		{
			name:    "tags_with_various_attributes",
			content: `<input type="text" name="test" required disabled/><a href="http://example.com" target="_blank">Link</a>`,
			expected: []xmlTag{
				{Name: "input", IsClosing: false, IsSelfClose: true, Position: 0},
				{Name: "a", IsClosing: false, IsSelfClose: false, Position: 50},
				{Name: "a", IsClosing: true, IsSelfClose: false, Position: 98},
			},
		},
		{
			name:    "whitespace_variations",
			content: `< tag1 >content</ tag1 >< tag2/>`,
			expected: []xmlTag{
				{Name: "tag1", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "tag2", IsClosing: false, IsSelfClose: true, Position: 24},
			},
		},
		{
			name:    "namespace_variations",
			content: `<xml:root xmlns:xml="test"><ns:child>content</ns:child><x:y:z attr="val"/></xml:root>`,
			expected: []xmlTag{
				{Name: "xml:root", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "ns:child", IsClosing: false, IsSelfClose: false, Position: 27},
				{Name: "ns:child", IsClosing: true, IsSelfClose: false, Position: 44},
				{Name: "x:y:z", IsClosing: false, IsSelfClose: true, Position: 55},
				{Name: "xml:root", IsClosing: true, IsSelfClose: false, Position: 73},
			},
		},
		{
			name:    "single_character_tags",
			content: `<a>link</a><b>bold</b><i>italic</i>`,
			expected: []xmlTag{
				{Name: "a", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "a", IsClosing: true, IsSelfClose: false, Position: 7},
				{Name: "b", IsClosing: false, IsSelfClose: false, Position: 11},
				{Name: "b", IsClosing: true, IsSelfClose: false, Position: 18},
				{Name: "i", IsClosing: false, IsSelfClose: false, Position: 22},
				{Name: "i", IsClosing: true, IsSelfClose: false, Position: 31},
			},
		},
		{
			name:    "malformed_tags_ignored",
			content: `<valid>content</valid><<invalid>>content<</>not-a-tag<>`,
			expected: []xmlTag{
				{Name: "valid", IsClosing: false, IsSelfClose: false, Position: 0},
				{Name: "valid", IsClosing: true, IsSelfClose: false, Position: 14},
				{Name: "invalid", IsClosing: false, IsSelfClose: false, Position: 23},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tags := analyzer.parseXMLTags(tt.content)

			assert.Len(t, tags, len(tt.expected), "Number of parsed tags should match expected")

			for i, expectedTag := range tt.expected {
				if assert.Less(t, i, len(tags), "Should have enough parsed tags") {
					tag := tags[i]
					assert.Equal(t, expectedTag.Name, tag.Name, "Tag name should match")
					assert.Equal(t, expectedTag.IsClosing, tag.IsClosing, "IsClosing should match")
					assert.Equal(t, expectedTag.IsSelfClose, tag.IsSelfClose, "IsSelfClose should match")
				}
			}
		})
	}
}

func TestFormatMultipleReferences(t *testing.T) {
	discovery := &assets.Discovery{}
	analyzer := NewFrameworkAnalyzer(discovery)

	tests := []struct {
		name     string
		fileRef  *FileReference
		expected string
	}{
		{
			name: "single_agent_multiple_tasks",
			fileRef: &FileReference{
				FilePath: "/path/to/file.md",
				FileType: "data file",
				References: []AgentTaskRef{
					{AgentName: "architect", TaskName: "task1"},
					{AgentName: "architect", TaskName: "task2"},
				},
			},
			expected: "(data file, multiple agents/tasks) - Referenced by: architect(task1, task2)",
		},
		{
			name: "multiple_agents_single_tasks",
			fileRef: &FileReference{
				FilePath: "/path/to/file.md",
				FileType: "template",
				References: []AgentTaskRef{
					{AgentName: "architect", TaskName: "task1"},
					{AgentName: "advisor", TaskName: "task2"},
				},
			},
			expected: "(template, multiple agents/tasks) - Referenced by: advisor(task2), architect(task1)",
		},
		{
			name: "multiple_agents_multiple_tasks",
			fileRef: &FileReference{
				FilePath: "/path/to/file.md",
				FileType: "data file",
				References: []AgentTaskRef{
					{AgentName: "architect", TaskName: "task1"},
					{AgentName: "architect", TaskName: "task2"},
					{AgentName: "advisor", TaskName: "task3"},
					{AgentName: "advisor", TaskName: "task4"},
				},
			},
			expected: "(data file, multiple agents/tasks) - Referenced by: advisor(task3, task4), architect(task1, task2)",
		},
		{
			name: "agent_level_reference",
			fileRef: &FileReference{
				FilePath: "/path/to/file.md",
				FileType: "task",
				References: []AgentTaskRef{
					{AgentName: "architect", TaskName: ""},
					{AgentName: "advisor", TaskName: ""},
				},
			},
			expected: "(task, multiple agents/tasks) - Referenced by: advisor, architect",
		},
		{
			name: "mixed_references",
			fileRef: &FileReference{
				FilePath: "/path/to/file.md",
				FileType: "data file",
				References: []AgentTaskRef{
					{AgentName: "architect", TaskName: ""},
					{AgentName: "advisor", TaskName: "task1"},
					{AgentName: "advisor", TaskName: "task2"},
				},
			},
			expected: "(data file, multiple agents/tasks) - Referenced by: advisor(task1, task2), architect",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := analyzer.formatMultipleReferences(tt.fileRef)
			assert.Equal(t, tt.expected, result, "Multiple references format should match expected")
		})
	}
}

func TestCodeBlockSkipping(t *testing.T) {
	discovery := &assets.Discovery{}
	analyzer := NewFrameworkAnalyzer(discovery)

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name: "inline_code_tags_skipped",
			content: `<valid>
Valid content
</valid>

This is an example: ` + "`<skipped-tag>`" + `

<another-valid>
More content
</another-valid>`,
			expected: nil, // No issues because all tags are either valid or in code blocks
		},
		{
			name: "fenced_code_block_tags_skipped",
			content: `<valid>
Valid content
</valid>

Example:
` + "```go" + `
<tag-to-skip>
func main() {
	fmt.Println("<another-tag-to-skip>")
}
</tag-to-skip>
` + "```" + `

<another-valid>
Content
</another-valid>`,
			expected: nil, // All tags are valid or in code blocks
		},
		{
			name: "tilde_code_block_tags_skipped",
			content: `<valid>
Valid content
</valid>

Example:
~~~
<tag-to-skip>
Some code example
</tag-to-skip>
~~~

<unclosed>
This should be detected`,
			expected: []string{"Unclosed tag <unclosed>"},
		},
		{
			name: "mixed_code_blocks_and_real_tags",
			content: `<prerequisites>
- Some requirements
</prerequisites>

Example inline: ` + "`<inline-skip>`" + `

` + "```" + `
<fenced-skip>
Code content
</fenced-skip>
` + "```" + `

<instructions>
1. Do something
</instructions>

Another example:
~~~html
<html-skip>
<div>Example</div>
</html-skip>
~~~

<unclosed-real>
This should be caught`,
			expected: []string{"Unclosed tag <unclosed-real>"},
		},
		{
			name: "nested_backticks_handled",
			content: `<valid>
Content
</valid>

Example: ` + "`<tag>`" + ` and ` + "`<other-tag>`" + `

<unclosed>
This should be caught`,
			expected: []string{"Unclosed tag <unclosed>"},
		},
		{
			name: "code_blocks_with_attributes",
			content: `<tag attr="value">
Content
</tag>

` + "```xml" + `
<element attr="value">
  <child>Content</child>
</element>
` + "```" + `

<valid-tag>
Regular content
</valid-tag>`,
			expected: nil, // All valid
		},
		{
			name: "multiline_fenced_blocks",
			content: `<valid>
Valid content
</valid>

` + "```" + `
<multiline>
  Line 1
  Line 2
  <nested>
    Content
  </nested>
</multiline>
` + "```" + `

<unclosed>
Should be detected`,
			expected: []string{"Unclosed tag <unclosed>"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := analyzer.validateXMLTags(tt.content)
			assert.Equal(t, tt.expected, issues, "Code block skipping should work correctly")
		})
	}
}

func TestFindCodeBlockRanges(t *testing.T) {
	discovery := &assets.Discovery{}
	analyzer := NewFrameworkAnalyzer(discovery)

	tests := []struct {
		name     string
		content  string
		expected int // number of code block ranges
	}{
		{
			name:     "no_code_blocks",
			content:  "Just regular text with <tags>",
			expected: 0,
		},
		{
			name:     "single_inline_code",
			content:  "Text with `code` here",
			expected: 1,
		},
		{
			name:     "multiple_inline_code",
			content:  "Text with `code1` and `code2` here",
			expected: 2,
		},
		{
			name:     "single_fenced_block",
			content:  "Text\n```\ncode\n```\nMore text",
			expected: 1,
		},
		{
			name:     "mixed_code_blocks",
			content:  "Text `inline` and\n```\nfenced\n```\nand ~~~\ntilde\n~~~",
			expected: 2, // 1 inline + 1 fenced (~~~ block might not be detected as separate due to line endings)
		},
		{
			name:     "fenced_with_language",
			content:  "Text\n```go\nfunc main() {}\n```\nMore",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ranges := analyzer.findCodeBlockRanges(tt.content)
			assert.Len(t, ranges, tt.expected, "Should find correct number of code block ranges")
		})
	}
}

func TestYAMLFileSkipping(t *testing.T) {
	// Create temporary directory for test files
	tempDir := t.TempDir()

	// Create test YAML file with XML-like content
	yamlFile := filepath.Join(tempDir, "test.yaml")
	yamlContent := `apiVersion: v1
kind: ConfigMap
metadata:
  name: test
data:
  content: |
    <prerequisites>
    Some content
    # Missing closing tag intentionally
`
	err := os.WriteFile(yamlFile, []byte(yamlContent), 0644)
	require.NoError(t, err)

	// Create test YML file with XML-like content
	ymlFile := filepath.Join(tempDir, "test.yml")
	ymlContent := `version: "3.8"
services:
  app:
    image: nginx
    environment:
      - CONFIG=<config>unclosed tag</config><unclosed>
`
	err = os.WriteFile(ymlFile, []byte(ymlContent), 0644)
	require.NoError(t, err)

	// Create test Markdown file with XML-like content for comparison
	mdFile := filepath.Join(tempDir, "test.md")
	mdContent := `# Test File

<prerequisites>
Some content without closing tag
`
	err = os.WriteFile(mdFile, []byte(mdContent), 0644)
	require.NoError(t, err)

	// Create file usage map
	fileUsage := map[string]*FileReference{
		yamlFile: {
			FilePath: yamlFile,
			FileType: "data file",
			References: []AgentTaskRef{
				{AgentName: "test-agent", TaskName: "test-task"},
			},
		},
		ymlFile: {
			FilePath: ymlFile,
			FileType: "data file",
			References: []AgentTaskRef{
				{AgentName: "test-agent", TaskName: "test-task"},
			},
		},
		mdFile: {
			FilePath: mdFile,
			FileType: "task",
			References: []AgentTaskRef{
				{AgentName: "test-agent", TaskName: "test-task"},
			},
		},
	}

	// Create analyzer and run validation
	analyzer := &FrameworkAnalyzer{}
	issues := analyzer.deduplicateXMLValidationIssues(fileUsage)

	// Verify that only the markdown file has validation issues
	// YAML/YML files should be skipped entirely
	assert.Len(t, issues, 1, "Should only have issues for markdown file")

	if len(issues) > 0 {
		assert.Contains(t, issues[0].File, "test.md", "Issue should be for markdown file only")
		assert.Contains(t, issues[0].Message, "Unclosed tag <prerequisites>", "Should report unclosed tag in markdown")
	}
}
