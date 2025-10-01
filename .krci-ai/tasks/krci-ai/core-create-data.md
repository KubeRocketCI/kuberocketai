---
dependencies:
  data:
    - krci-ai/core-framework-standards.yaml
    - krci-ai/core-validation-checklist.md
  templates:
    - krci-ai/core-data-template.md
---

# Task: Core Create Data

## Description

Guide user through creating framework data files that provide knowledge, standards, and constraints for agent behavior and task execution. This task covers data organization, reference patterns, and content structuring for effective framework integration and component operation.

<instructions>
Define data file specification. Ask user for exact data file name, purpose, and content scope (example: technical-standards.yaml, api-guidelines.md). Clarify which agents or tasks will reference this data and what knowledge or standards it will provide.

Determine appropriate data format. Choose format based on content type and usage requirements. Use markdown (.md) for documentation, guidelines, principles, and explanatory content with rich formatting. Use YAML (.yaml/.yml) for specifications, configurations, and structured standards with hierarchy. Use JSON (.json) for structured data, API schemas, and configuration formats with strict schema. Use text (.txt) for simple reference data and plain text content. Use CSV (.csv) for tabular data, metrics, lists, and comparative information.

Organize data structure logically. Create organization following framework patterns for reference integration. Structure content to support frontmatter dependency patterns. Organize for agent behavioral integration or task technical integration as appropriate. Design structure compatible with framework validation processes.

Define reference context and integration. Ensure data supports agent principles and behavioral guidance if referenced by agents. Structure data to provide task specifications and technical constraints if referenced by tasks. Include examples, implementations, or configuration samples as reference data. Provide standards, checklists, or compliance requirements for validation purposes.

Apply naming conventions and scope indicators. Use descriptive file names clearly indicating data purpose and scope. Follow framework naming patterns for consistency. Ensure names differentiate between behavioral data (for agents) and technical data (for tasks).

Structure for accessibility and consumption. Organize information for easy reference by framework components. Create clear sections with appropriate headings. Provide comprehensive, accurate information for intended purpose. Enable straightforward consumption by both agents and tasks.

Format output following data organization principles. Save to ./.krci-ai/data/{category}/{data-name}.{ext} following path and naming conventions. Use core-data-template.md as structural guide where applicable. Ensure format syntax is correct and parseable.

Validate data file integration. Verify file format syntax correctness. Confirm content organization follows framework patterns. Test data supports proper framework component integration. Ensure information accuracy and completeness for intended purpose.
</instructions>

## Framework Context: Data Architecture and Integration Patterns

Data Types and Organization:

- Behavioral Data: Guidelines, principles, standards referenced by agents for decision-making
- Technical Data: Specifications, schemas, constraints referenced by tasks for implementation
- Reference Data: Examples, implementations, configuration samples for guidance
- Validation Data: Standards, checklists, compliance requirements for quality assurance

File Format Selection Strategy: Framework supports multiple formats based on content requirements:

- Markdown (.md): Documentation, guidelines, principles, explanatory content with rich formatting
- YAML (.yaml/.yml): Specifications, configurations, structured standards with hierarchy
- JSON (.json): Structured data, API schemas, configuration formats with strict schema
- Text (.txt): Simple reference data, plain text content for basic information
- CSV (.csv): Tabular data, metrics, lists, comparative information with structured columns

Reference Integration Architecture: Data files integrate with framework components through:

- Frontmatter dependency pattern: List in YAML frontmatter for component linking
- Agent behavioral integration: Data supports principle definitions and guidance references
- Task technical integration: Data provides specifications and constraint information
- Framework validation compatibility: Data structure supports automated validation processes

## Output Format

- Location: `./.krci-ai/data//{data-name}.{ext}` following naming conventions and organization
- Format: Appropriate file format based on content type and usage requirements
- Structure: Logical organization supporting framework reference patterns
- Accessibility: Content organized for easy consumption and reference by components

<success_criteria>
- Framework compliance verified: Data file passes all automated validation checks without errors
- Pattern adherence confirmed: Data follows established framework conventions exactly
- Reference integrity validated: All references resolve correctly and appropriately
- Quality standards met: Data meets completeness, clarity, and maintainability requirements
- Integration readiness achieved: Data ready for framework operation and usage
- Documentation completeness confirmed: All required sections populated with actionable content
</success_criteria>

## Execution Checklist

### Preparation Phase

- Framework validation: Run `krci-ai validate` to ensure clean starting state
- Dependency verification: Confirm all reference assets exist at specified paths
- Context gathering: Review data requirements and intended usage scenarios
- Purpose definition: Clear understanding of what knowledge or standards data will provide
- Usage context identification: Determine which agents or tasks will reference this data

### Execution Phase

- Data format selection: Choose appropriate format (.md, .yaml, .json, .txt, .csv) based on content
- Content structure creation: Organize data logically following framework patterns
- Dependency optimization: Structure content to support frontmatter dependency references
- Naming conventions application: Use descriptive file and section names indicating purpose
- Accessibility focus: Organize information for effective agent and task consumption
- Template application: Use [core-data-template.md](./.krci-ai/templates/krci-ai/core-data-template.md) structure
- Content development: Populate data with comprehensive, accurate information

### Validation Phase

- Format validation: Verify file format syntax is correct and parseable
- Structure consistency: Ensure content organization follows established framework patterns
- Reference compatibility: Confirm data supports proper framework integration
- Content quality assurance: Verify information accuracy and completeness
- Framework validation: Test data file compatibility with validation requirements
- Integration testing: Confirm data integrates properly with intended framework components
