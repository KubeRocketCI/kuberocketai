# {{data_title}}

<instructions>
DATA FILE TEMPLATE: This template provides structure for creating framework data files. Data files provide knowledge, standards, and constraints that agents and tasks reference for behavioral guidance and technical requirements.

PURPOSE DEFINITION: Clearly define what knowledge, standards, or constraints this data file provides and how it integrates with the framework ecosystem.

CONTENT ORGANIZATION: Structure data logically for easy reference and consumption by framework components using inline `[filename](path)` references.
</instructions>

## Overview

<instructions>
Data file overview explaining purpose, scope, and intended usage within the framework. This section provides context for how the data integrates with agents and tasks.
</instructions>

**Purpose**: {{data_purpose}}

**Scope**: {{data_scope}}

**Usage Context**: {{usage_context}}

**Reference Pattern**: {{reference_pattern}}

## {{primary_data_section}}

<instructions>
Primary data content organized according to the specific type and purpose of the data file. Structure content for effective consumption by framework components.
</instructions>

{{primary_data_content}}

### {{data_subsection_1}}

<instructions>
Detailed data subsection with specific information, standards, or constraints. Use clear organization and formatting for easy reference.
</instructions>

{{subsection_content_1}}

### {{data_subsection_2}}

<instructions>
Additional data subsection providing complementary information. Maintain consistent organization and formatting patterns.
</instructions>

{{subsection_content_2}}

## {{secondary_data_section}}

<instructions>
Secondary data section for additional standards, examples, or reference information. Include comprehensive content that supports framework operation.
</instructions>

{{secondary_data_content}}

## {{validation_or_standards_section}}

<instructions>
Section for validation criteria, compliance standards, or usage guidelines. This content helps ensure proper framework integration and component quality.
</instructions>

{{validation_content}}

## {{examples_or_references_section}}

<instructions>
Examples, reference implementations, or additional resources that support the data content. Include practical applications and usage demonstrations.
</instructions>

{{examples_content}}

<instructions>
DATA FILE CREATION GUIDANCE:

1. **Data Organization Principles**:
   - Structure content logically for framework component consumption
   - Use clear headings and consistent formatting
   - Organize information for easy reference and accessibility
   - Group related concepts and standards together

2. **Content Types and Formats**:
   - **Behavioral Data**: Guidelines, principles, standards for agent reference
   - **Technical Data**: Specifications, schemas, constraints for task reference
   - **Reference Data**: Examples, implementations, configuration samples
   - **Validation Data**: Standards, checklists, compliance requirements

3. **File Format Selection**:
   - **Markdown (.md)**: Documentation, guidelines, explanatory content
   - **YAML (.yaml)**: Specifications, configurations, structured standards
   - **JSON (.json)**: Structured data, schemas, configuration formats
   - **Text (.txt)**: Simple reference data, plain text content
   - **CSV (.csv)**: Tabular data, metrics, comparative information

4. **Framework Integration Requirements**:
   - Design for reference by agents (behavioral guidance) or tasks (technical constraints)
   - Support inline `[filename](./.krci-ai/path/to/file)` reference patterns
   - Ensure content is accessible and consumable by framework components
   - Follow naming conventions and organizational patterns

5. **Quality and Effectiveness Standards**:
   - Provide valuable knowledge, standards, or constraints
   - Maintain accuracy and currency of information
   - Structure for effective framework operation support
   - Design for reusability and maintainability

DATA VALIDATION CHECKLIST:
- [ ] Data file serves clear purpose for framework operation
- [ ] Content organized logically for component consumption
- [ ] Appropriate file format selected for content type
- [ ] Information accurate, current, and comprehensive
- [ ] Framework integration patterns followed
- [ ] Reference compatibility ensured for agents and tasks

USAGE PATTERNS:
- **Agent References**: Behavioral data referenced in agent principles
- **Task References**: Technical data referenced in task instructions
- **Template References**: Structure and formatting guidance for templates
- **Validation References**: Standards and checklists for framework compliance
</instructions>
