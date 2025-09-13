# Task: Core Create Template

## Description

Guide user through creating framework-compliant templates for consistent LLM output formatting. This task provides instructions for template creation using variable placeholders, XML guidance tags, and structured content organization for effective LLM processing.

<prerequisites>
- Framework installed: .krci-ai directory exists with templates/ subdirectory
- Template purpose defined: Clear understanding of what output template will structure
- Variable identification: Knowledge of dynamic content areas requiring `{{variable}}` placeholders
- Usage context: Understanding of which tasks will reference this template
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/templates/krci-ai/core-template-template.md
- ./.krci-ai/data/krci-ai/core-validation-checklist.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Apply template structure: Use markdown format with `{{variable}}` placeholders for dynamic content
2. Include XML guidance tags: Add `<instructions>` sections for LLM processing guidance
3. Design variable system: Create descriptive variable names using consistent naming patterns
4. Provide LLM guidance: Include comments and instructions for content generation
5. Ensure reusability: Design template for use across multiple similar contexts
6. Format output: Use [core-template-template.md](./.krci-ai/templates/krci-ai/core-template-template.md) as structural guide
</instructions>

## Framework Context: Template Architecture

Template Variable System: Templates use `{{variable_name}}` placeholders for LLM content substitution:
- Simple Values: `{{project_name}}`, `{{description}}` for single values
- Lists: `{{requirements}}`, `{{features}}` for bullet points or numbered lists
- Sections: `{{technical_details}}`, `{{implementation_approach}}` for large content blocks
- Optional Content: `{{optional_section}}` may be empty based on context

XML Guidance Integration: Templates include XML tags for LLM processing instructions:
- `<instructions>` - Guidance for LLM content generation within sections
- Comments for variable usage examples and formatting requirements
- Section organization help for consistent structure

LLM-Friendly Design: Templates optimized for natural language processing and content generation.

## Output Format

- Location: `./.krci-ai/templates/{template-name}.md` following naming conventions
- Structure: Markdown format with clear sections and variable placeholders
- LLM guidance: XML instruction tags and comments for processing guidance
- Variable consistency: Descriptive variable names with consistent patterns

<success_criteria>
- Template file created with proper markdown structure and naming
- Variable placeholders use `{{variable_name}}` format consistently
- XML guidance tags included for LLM processing instructions
- Template structure supports intended output formatting
- Variable system enables flexible content generation
- Template validates against framework standards
</success_criteria>

## Execution Checklist

### Template Structure Design

- [ ] Markdown format: Template uses standard markdown syntax for structure
- [ ] Section organization: Logical content flow with clear headings and organization
- [ ] Variable placement: `{{variable}}` placeholders positioned appropriately for content substitution
- [ ] Content hierarchy: Proper heading levels and content organization

### Variable System Implementation

- [ ] Variable naming: Descriptive names using consistent patterns (snake_case recommended)
- [ ] Variable types: Appropriate mix of simple values, lists, and content sections
- [ ] Variable documentation: Comments explaining variable purpose and expected content
- [ ] Optional handling: Clear indication of optional vs required variables

### LLM Guidance Integration

- [ ] XML instructions: `<instructions>` tags provide clear guidance for content generation
- [ ] Content examples: Examples or hints for variable content within instruction tags
- [ ] Formatting guidance: Instructions for consistent formatting and style
- [ ] Processing hints: Guidance for LLM content generation and structure

### Template Optimization

- [ ] Reusability: Template design supports use across multiple similar contexts
- [ ] Flexibility: Variable system allows for diverse content while maintaining structure
- [ ] LLM compatibility: Template structure facilitates natural language processing
- [ ] Framework integration: Template follows established framework patterns

### Validation and Testing

- [ ] Template syntax: Markdown syntax correct and renders properly
- [ ] Variable consistency: All variables follow naming conventions and patterns
- [ ] XML tag validity: Instruction tags properly formatted and closed
- [ ] Framework compliance: Template meets framework standards and validation requirements
