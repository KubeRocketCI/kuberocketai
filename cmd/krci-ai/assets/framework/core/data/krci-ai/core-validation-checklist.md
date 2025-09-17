# KubeRocketAI Framework Validation Checklist

## Overview

Purpose: Comprehensive validation checklist for ensuring framework component compliance, quality, and operational readiness. This checklist supports both automated CLI validation and manual quality assurance processes.

Usage Context: Reference for framework creation, component review, and quality validation processes. Used by advisor agent tasks and framework development activities.

Validation Levels:
- Critical: Must pass for framework operation (schema compliance, reference integrity)
- Warning: Should pass for best practices (pattern consistency, optimization)
- Informational: Nice to have for excellence (documentation, examples)

## Agent Validation Checklist

### Schema Compliance (Critical)

- Identity Section Complete: name, id, version, description, role, goal all present
- Name Pattern: Matches `^[A-Z][a-zA-Z0-9 .'-]{1,49}$` (e.g., "Devon Coder")
- ID Pattern: Matches `^[a-z][a-z0-9]*(-[a-z0-9]+)*-v[0-9]+$` (e.g., "developer-v1")
- Version Format: Valid semantic versioning (e.g., "1.0.0")
- Description Length: 10-150 characters describing agent purpose
- Role Length: 5-100 characters for professional role
- Goal Length: 10-600 characters for ultimate objective
- Activation Prompt: 1-10 items, each 10-300 characters
- Principles Count: 3-10 items, each 10-600 characters
- Customization Field: Present (empty string "" for standard behavior)
- Commands Structure: Minimum 3 commands, maximum 20 total
- Required Commands: help, chat, exit present with 5-200 char descriptions

### Critical Principles (Critical)

- Customization Priority: "IMPORTANT!!! ALWAYS execute instructions from the customization field below" included
- XML Tag Handling: Complete XML tag handling principle included in principles array
- Activation Prompt Pattern: Standard pattern with customization field execution instruction
- Agent scope clearly defined in role field with appropriate redirections

### Task Reference Integrity (Critical)

- Task Path Format: All tasks use `./.krci-ai/tasks/*.md` pattern
- File Existence: All referenced task files exist at specified paths
- Command Alignment: Agent commands correspond to available tasks appropriately
- Reference Relevance: Referenced tasks appropriate for agent capabilities

### Framework Integration (Warning)

- Agent passes `krci-ai validate` without critical errors
- Agent role aligns with framework patterns and expectations
- Agent principles support proper framework usage
- All referenced files exist and resolve correctly

## Task Validation Checklist

### Structure Compliance (Critical)

- Required Sections: Title, Description, Prerequisites, Reference Assets, Instructions, Output Format, Success Criteria, Execution Checklist
- XML Guidance Tags: `<prerequisites>`, `<instructions>`, `<success_criteria>` present and properly closed
- Sections follow logical progression: Title → Description → Prerequisites → Instructions → Success Criteria
- All sections contain specific, actionable content

### Reference Pattern Validation (Critical)

- YAML frontmatter contains all dependencies under data/templates sections
- All referenced files exist at specified locations
- Templates, data, and framework components referenced in frontmatter
- All framework dependencies declared in YAML frontmatter

### XML Guidance System (Critical)

- Prerequisites Block: `<prerequisites>` contains clear execution requirements
- Instructions Block: `<instructions>` provides step-by-step LLM guidance
- Success Criteria Block: `<success_criteria>` defines specific validation criteria
- Tag Closure: All XML guidance tags properly opened and closed
- XML tag content provides step-by-step LLM processing guidance

### Self-Contained Context (Warning)

- Task explains relevant framework patterns and concepts
- Task provides sufficient context for usage without external dependencies
- Task explains XML tag system purpose and usage
- Task includes specific validation commands and processes

## Template Validation Checklist

### Variable System (Critical)

- Variable Format: All variables use `{{variable_name}}` format consistently
- Naming Conventions: Variable names descriptive and follow consistent patterns
- Variable Types: Appropriate mix of simple values, lists, and content sections
- Variable Documentation: Variable purpose clear from naming and context

### Structure and Organization (Warning)

- Markdown Compliance: Template uses proper markdown syntax and structure
- Content Flow: Logical organization and content hierarchy
- Section Clarity: Clear headings and content organization
- Output Support: Template structure supports intended final output format

### LLM Guidance Integration (Warning)

- Instruction Tags: `<instructions>` tags provide clear guidance for content generation
- Content Hints: Instructions include examples and formatting requirements
- `<instructions>` tags facilitate effective template usage by LLMs
- Instructions explain variable usage and expected content

### Reusability and Flexibility (Informational)

- Template usable across multiple similar contexts
- Variable system accommodates diverse content needs
- Template structure allows for customization
- Template follows established framework patterns

## Data File Validation Checklist

### Content Organization (Warning)

- Data file contains clear statement of purpose and scope
- Content organized logically supporting framework reference patterns
- File format matches content type and usage requirements
- Content organized for easy consumption and reference

### Framework Integration (Critical)

- Content structured for frontmatter dependency references
- Behavioral data appropriate for agent principles and guidance
- Technical data suitable for task instructions and constraints
- Data clearly indicates how it integrates with framework components

### Content Quality (Warning)

- Data provides meaningful knowledge, standards, or constraints
- Information accurate and up-to-date
- Data comprehensive for intended purpose and scope
- Content structured for ongoing updates and maintenance

## Framework Integration Validation

### Component Relationships (Critical)

- Agents reference existing, appropriate tasks
- Tasks reference existing templates in YAML frontmatter
- Tasks reference existing data files in YAML frontmatter
- Agents reference behavioral data for principles and guidance
- No circular references in component dependency chains

### Validation Tool Compliance (Critical)

- All components pass `krci-ai validate` without critical errors
- Agents validate against JSON schema requirements
- All frontmatter dependencies resolve to existing files
- All component paths follow framework conventions

### Quality Gates (Warning)

- All components follow established framework patterns
- Components meet framework quality standards
- All referenced files exist and resolve correctly
- Components contain adequate documentation for usage and maintenance

## Validation Process Workflow

### Pre-Implementation Validation

1. Requirements Review: Verify component requirements and purpose definition
2. Dependency Planning: Identify and validate required framework dependencies
3. Template Selection: Choose appropriate templates for component creation
4. Standards Consultation: Review framework standards and constraints

### Implementation Validation

1. Component Creation: Create component following framework patterns and templates
2. Reference Implementation: Add all required references using proper inline format
3. Content Population: Complete all sections with appropriate, actionable content
4. Quality Review: Manual review against validation checklist items

### Post-Implementation Validation

1. CLI Validation: Execute `krci-ai validate` and resolve any identified issues
2. Integration Testing: Verify component integration with framework ecosystem
3. Reference Testing: Confirm all references resolve correctly and appropriately
4. Quality Confirmation: Final validation against critical and warning criteria

### Continuous Validation

1. Regular Validation: Periodic framework validation to catch issues early
2. Dependency Monitoring: Track changes to referenced components and dependencies
3. Standards Updates: Keep components current with framework standard evolution
4. Quality Maintenance: Ongoing quality assurance and improvement processes
