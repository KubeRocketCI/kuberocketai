# Task: Core Create Data

## Description

Guide user through creating framework data files that provide knowledge, standards, and constraints for agent behavior and task execution. This task covers data organization, reference patterns, and content structuring for effective framework integration.

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

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Determine data format: Choose appropriate format (.md, .yaml, .json, .txt, .csv) based on content type
2. Organize data structure: Create logical organization following framework patterns
3. Define reference context: Ensure data supports agent behavioral guidance or task technical constraints
4. Apply naming conventions: Use descriptive names that indicate data purpose and scope
5. Structure for accessibility: Organize content for easy reference and consumption
6. Format output: Use [core-data-template.md](./.krci-ai/templates/krci-ai/core-data-template.md) as structural guide
</instructions>

## Framework Context: Data Architecture

Data Types and Organization:
- Behavioral Data: Guidelines, principles, standards referenced by agents
- Technical Data: Specifications, schemas, constraints referenced by tasks
- Reference Data: Examples, implementations, configuration samples
- Validation Data: Standards, checklists, compliance requirements

File Format Flexibility: Framework supports any format appropriate for content:
- Markdown (.md): Documentation, guidelines, principles, explanatory content
- YAML (.yaml/.yml): Specifications, configurations, structured standards
- JSON (.json): Structured data, API schemas, configuration formats
- Text (.txt): Simple reference data, plain text content
- CSV (.csv): Tabular data, metrics, lists, comparative information

Reference Integration: Data files are referenced by agents and tasks using inline `[filename](path)` links.

## Output Format

- Location: `./.krci-ai/data/**/{data-name}.{ext}` following naming conventions and organization
- Format: Appropriate file format based on content type and usage requirements
- Structure: Logical organization supporting framework reference patterns
- Accessibility: Content organized for easy consumption and reference

<success_criteria>
- Data file created with appropriate format and naming convention
- Content structure supports intended usage by agents or tasks
- Data organization follows framework patterns for reference integration
- File format matches content type and accessibility requirements
- Data provides valuable knowledge or constraints for framework operation
- Framework validation accommodates data file without issues
</success_criteria>

## Execution Checklist

### Data Planning and Organization

- [ ] Purpose definition: Clear understanding of what knowledge or standards data provides
- [ ] Usage context: Identification of which agents or tasks will reference this data
- [ ] Content type: Appropriate file format selected based on data structure and usage
- [ ] Organization strategy: Logical content organization supporting framework patterns

### Content Development

- [ ] Data structuring: Content organized logically for easy reference and consumption
- [ ] Naming conventions: Descriptive file and section names indicating purpose and scope
- [ ] Reference optimization: Content structured to support inline `[filename](path)` references
- [ ] Accessibility focus: Information organized for effective agent and task usage

### Format and Structure Implementation

- [ ] Format selection: File format (md, yaml, json, txt, csv) appropriate for content type
- [ ] Structure consistency: Content organization follows established framework patterns
- [ ] Reference compatibility: Data structured for effective integration with agents and tasks
- [ ] Documentation quality: Content clear, comprehensive, and actionable

### Framework Integration

- [ ] Reference patterns: Data supports proper `[filename](path)` reference usage
- [ ] Agent compatibility: Behavioral data appropriate for agent principle and guidance references
- [ ] Task compatibility: Technical data suitable for task instruction and constraint references
- [ ] Validation readiness: Data file compatible with framework validation requirements

### Quality and Effectiveness

- [ ] Content value: Data provides meaningful knowledge, standards, or constraints
- [ ] Accuracy verification: Information accurate and up-to-date for framework usage
- [ ] Completeness check: Data comprehensive for intended purpose and scope
- [ ] Usability assessment: Content organized and accessible for effective framework operation
