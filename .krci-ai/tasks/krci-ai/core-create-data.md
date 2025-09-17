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

<prerequisites>
- Framework installed: .krci-ai directory exists with data/ subdirectory structure
- Data purpose defined: Clear understanding of what knowledge or standards data will provide
- Usage context: Knowledge of which agents or tasks will reference this data
- Content type identified: Understanding of appropriate file format for data content
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/templates/krci-ai/core-data-template.md
- ./.krci-ai/data/krci-ai/core-validation-checklist.md

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.

<instructions>
1. Determine data format: Choose appropriate format (.md, .yaml, .json, .txt, .csv) based on content type and usage
2. Organize data structure: Create logical organization following framework patterns for reference integration
3. Define reference context: Ensure data supports agent behavioral guidance or task technical constraints effectively
4. Apply naming conventions: Use descriptive names that indicate data purpose and scope clearly
5. Structure for accessibility: Organize content for easy reference and consumption by framework components
6. Format output: Use [core-data-template.md](./.krci-ai/templates/krci-ai/core-data-template.md) as structural guide
</instructions>

## Framework Context: Data Architecture and Integration Patterns

**Data Types and Organization**:

- Behavioral Data: Guidelines, principles, standards referenced by agents for decision-making
- Technical Data: Specifications, schemas, constraints referenced by tasks for implementation
- Reference Data: Examples, implementations, configuration samples for guidance
- Validation Data: Standards, checklists, compliance requirements for quality assurance

**File Format Selection Strategy**: Framework supports multiple formats based on content requirements:

- Markdown (.md): Documentation, guidelines, principles, explanatory content with rich formatting
- YAML (.yaml/.yml): Specifications, configurations, structured standards with hierarchy
- JSON (.json): Structured data, API schemas, configuration formats with strict schema
- Text (.txt): Simple reference data, plain text content for basic information
- CSV (.csv): Tabular data, metrics, lists, comparative information with structured columns

**Reference Integration Architecture**: Data files integrate with framework components through:

- Frontmatter dependency pattern: List in YAML frontmatter for component linking
- Agent behavioral integration: Data supports principle definitions and guidance references
- Task technical integration: Data provides specifications and constraint information
- Framework validation compatibility: Data structure supports automated validation processes

## Output Format

- Location: `./.krci-ai/data/**/{data-name}.{ext}` following naming conventions and organization
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

- [ ] Framework validation: Run `krci-ai validate` to ensure clean starting state
- [ ] Dependency verification: Confirm all reference assets exist at specified paths
- [ ] Context gathering: Review data requirements and intended usage scenarios
- [ ] Purpose definition: Clear understanding of what knowledge or standards data will provide
- [ ] Usage context identification: Determine which agents or tasks will reference this data

### Execution Phase

- [ ] Data format selection: Choose appropriate format (.md, .yaml, .json, .txt, .csv) based on content
- [ ] Content structure creation: Organize data logically following framework patterns
- [ ] Dependency optimization: Structure content to support frontmatter dependency references
- [ ] Naming conventions application: Use descriptive file and section names indicating purpose
- [ ] Accessibility focus: Organize information for effective agent and task consumption
- [ ] Template application: Use [core-data-template.md](./.krci-ai/templates/krci-ai/core-data-template.md) structure
- [ ] Content development: Populate data with comprehensive, accurate information

### Validation Phase

- [ ] Format validation: Verify file format syntax is correct and parseable
- [ ] Structure consistency: Ensure content organization follows established framework patterns
- [ ] Reference compatibility: Confirm data supports proper framework integration
- [ ] Content quality assurance: Verify information accuracy and completeness
- [ ] Framework validation: Test data file compatibility with validation requirements
- [ ] Integration testing: Confirm data integrates properly with intended framework components
