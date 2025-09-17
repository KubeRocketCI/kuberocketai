---
dependencies:
  data:
    - krci-ai/core-framework-standards.yaml
    - krci-ai/core-validation-checklist.md
  templates:
    - krci-ai/core-template-template.md
---

# Task: Core Create Template

## Description

Guide user through creating framework-compliant templates for consistent LLM output formatting. This task provides instructions for template creation using variable placeholders, XML guidance tags, and structured content organization for effective LLM processing and reusable content generation.

<prerequisites>
- Framework installed: .krci-ai directory exists with templates/ subdirectory
- Template purpose defined: Clear understanding of what output template will structure
- Variable identification: Knowledge of dynamic content areas requiring `{{variable}}` placeholders
- Usage context: Understanding of which tasks will reference this template
</prerequisites>

<instructions>
1. Apply template structure: Use markdown format with `{{variable}}` placeholders for dynamic content substitution
2. Include XML guidance tags: Add `<instructions>` sections for LLM processing guidance and content generation
3. Design variable system: Create descriptive variable names using consistent naming patterns and types
4. Provide LLM guidance: Include comments and instructions for content generation within templates
5. Ensure reusability: Design template for use across multiple similar contexts with flexible variable system
6. Format output: Use [core-template-template.md](./.krci-ai/templates/krci-ai/core-template-template.md) as structural guide
</instructions>

## Framework Context: Template Architecture and Variable System

Template Variable System: Templates use `{{variable_name}}` placeholders for LLM content substitution with specific types:

- Simple values: `{{project_name}}`, `{{description}}`, `{{author}}` for single-line content
- Content lists: `{{requirements}}`, `{{features}}`, `{{task_items}}` for structured lists
- Large sections: `{{technical_details}}`, `{{implementation_approach}}` for multi-paragraph content
- Optional content: `{{optional_section}}`, `{{additional_notes}}` may be empty based on context

XML Guidance Integration: Templates include XML tags for LLM processing instructions:

- `<instructions>` tags provide guidance for LLM content generation within sections
- Comments explain variable usage examples and formatting requirements
- Section organization guidance helps maintain consistent structure across outputs
- Processing hints facilitate effective template usage and content generation

LLM-Friendly Design Principles: Templates optimized for natural language processing with clear structure, logical content flow, and comprehensive guidance for autonomous content generation.

Template Reusability Architecture: Design patterns support use across multiple similar contexts while maintaining structural integrity and variable system consistency.

## Output Format

- Location: `./.krci-ai/templates/{template-name}.md` following naming conventions
- Structure: Markdown format with clear sections and variable placeholders
- LLM guidance: XML instruction tags and comments for processing guidance
- Variable consistency: Descriptive variable names with consistent patterns and types

<success_criteria>
- Framework compliance verified: Template passes all automated validation checks without errors
- Pattern adherence confirmed: Template follows established framework conventions exactly
- Reference integrity validated: All references resolve correctly and appropriately
- Quality standards met: Template meets completeness, clarity, and maintainability requirements
- Integration readiness achieved: Template ready for framework operation and usage
- Documentation completeness confirmed: All required sections populated with actionable content
</success_criteria>

## Execution Checklist

### Preparation Phase

- Framework validation: Run `krci-ai validate` to ensure clean starting state
- Dependency verification: Confirm all reference assets exist at specified paths
- Context gathering: Review template requirements and intended usage scenarios
- Template purpose definition: Clear understanding of what output template will structure
- Variable identification: Analyze dynamic content areas requiring placeholders

### Execution Phase

- Template structure creation: Use markdown format with clear sections and logical organization
- Variable system implementation: Create `{{variable}}` placeholders with descriptive names
- Variable type organization: Implement simple values, lists, sections, and optional content appropriately
- XML guidance integration: Add `<instructions>` tags for LLM processing guidance
- Content examples inclusion: Provide examples and hints for variable content within instruction tags
- LLM guidance documentation: Include comments explaining variable purpose and expected content
- Reusability design: Ensure template works across multiple similar contexts
- Template application: Use [core-template-template.md](./.krci-ai/templates/krci-ai/core-template-template.md) structure

### Validation Phase

- Template syntax verification: Ensure markdown syntax is correct and renders properly
- Variable consistency check: Verify all variables follow naming conventions and patterns
- XML tag validation: Confirm instruction tags are properly formatted and closed
- Framework validation: Run template validation against framework standards
- Reusability testing: Verify template supports flexible content generation scenarios
- Integration testing: Confirm template integrates properly with framework ecosystem
