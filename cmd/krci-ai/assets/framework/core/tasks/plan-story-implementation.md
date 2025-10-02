---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
  templates:
    - story.md
---

# Task: Plan Story Implementation

## Description

Comprehensive technical planning task for developers to analyze, validate, and enhance story implementation details before beginning development work. This task ensures complete technical understanding, detailed task/subtask planning, and implementation readiness with source code structure, libraries, patterns, schemas, and technical specifications.

## Instructions

<instructions>
Confirm the target story file exists in `/docs/stories/` requiring implementation planning, you are executing as a development team member with implementation responsibility, the story has been reviewed and approved for implementation, and you have access to architecture documentation, existing codebase, and technical standards. CRITICAL: Load all dependencies (./.krci-ai/templates/story.md and ./.krci-ai/data/krci-ai/core-sdlc-framework.md) by reading their complete content before task execution. HALT if any dependencies are missing.

Validate story completeness to ensure it has sufficient technical detail for implementation. Analyze technical requirements through deep dive into implementation needs, dependencies, and constraints. Plan implementation approach by defining specific technical approach, libraries, patterns, and structure. Enhance task/subtask sections with detailed, executable implementation steps. Validate technical understanding to ensure complete comprehension before implementation begins. Document comprehensive technical specifications and implementation approach.
</instructions>

## Output Format

- Location: Update existing story file with implementation planning enhancements
- Template: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- Content Placement: Technical details in Description section, enhanced tasks in Tasks/Subtasks section
- Implementation Ready: Story contains specific file paths, commands, and technical specifications
- Verification: Story enables autonomous development without additional technical consultation

## Success Criteria

<success_criteria>
- Technical understanding complete: Developer has full comprehension of implementation requirements
- Implementation plan detailed: All technical decisions documented and validated
- Tasks/Subtasks enhanced: Atomic, executable steps with complete specifications
- Code structure planned: Clear file structure, directories, and integration approach
- Dependencies identified: All libraries, tools, and external requirements specified
- Validation strategy defined: Complete testing approach with specific commands and criteria
- Implementation ready: Developer can begin work with confidence and clarity
</success_criteria>

## Execution Checklist

### Story Technical Validation

<story_validation>
- Story completeness check: Verify story has business requirements and acceptance criteria
- Technical gap analysis: Identify missing technical details or specifications
- Dependencies review: Validate all technical dependencies are specified
- Architecture alignment: Confirm implementation approach fits project architecture
</story_validation>

### Technical Requirements Analysis

<technical_completeness>
- Existing code analysis: Review current codebase for integration points and patterns
- Project structure mapping: Analyze existing directory structure and identify where new components fit
- Library specification: Research and document required libraries with exact versions (format: `Library: package_name v1.2.3`)
- Dependency compatibility: Validate library compatibility with existing project dependencies
- Pattern identification: Define specific design patterns and approaches following project standards
- Data structure planning: Design schemas, models, and data flow with specific formats
- Integration analysis: Plan integration with existing systems and components
- Configuration requirements: Define environment setup, configuration files, and deployment needs
- Performance considerations: Identify performance requirements and optimization needs
</technical_completeness>

### Implementation Approach Planning

<architecture_validation>
- Directory organization: Plan specific directory structure following project patterns (src/, tests/, docs/, config/)
- File structure design: Define exact file paths and names for creation/modification (`/src/component/file.ext`)
- Integration point mapping: Identify specific integration points with existing codebase and APIs
- Component architecture: Define classes, functions, and component responsibilities with interfaces
- Code reuse identification: Identify opportunities to reuse existing components and shared utilities
- Data flow design: Map input/output flow and transformation logic with specific data formats
- Error handling strategy: Plan exception handling and error recovery following project patterns
- Testing approach: Define unit, integration, and validation testing strategy with existing frameworks
- Security considerations: Identify security requirements and implementation approach per project standards
</architecture_validation>

### Task/Subtask Enhancement

<task_implementation_review>
- Enhanced task formatting: Use format "Task N: Description (AC: X, Y)" with clear acceptance criteria mapping
- Atomic task breakdown: Create single-responsibility implementation tasks with specific deliverables
- Specific file targets: Define exact file paths for creation/modification (`create file: /path/to/file.ext`)
- Command specifications: Include executable commands for each step (`run: command with args`)
- Validation command integration: Add verification commands for each task (`Command: \`test_command\``)
- Purpose specification: Document the purpose and responsibility of each file/component created
- Dependency mapping: Define dependencies between tasks using "depends on Task X completion" format
- Success criteria: Specify measurable completion criteria (file exists, tests pass, output matches)
- Error recovery planning: Define rollback steps if subtasks fail during implementation
</task_implementation_review>

### Technical Specifications Documentation

<technical_specifications>
- Libraries and versions: Document all dependencies with specific versions
- Configuration details: Specify environment setup and configuration requirements
- Database schemas: Define data models, tables, and relationships if applicable
- API specifications: Document interfaces, endpoints, and data contracts
- File formats: Specify input/output formats, validation rules, and constraints
- Command patterns: Document CLI commands, scripts, and automation tools
</technical_specifications>

### Implementation Validation Planning

<validation_strategy>
- Unit testing plan: Define specific unit tests for each component
- Integration testing: Plan testing of component interactions and data flow
- Validation commands: Create specific commands to verify implementation correctness
- Performance testing: Define performance benchmarks and testing approach
- Security validation: Plan security testing and vulnerability assessment
- End-to-end verification: Create complete workflow validation steps
</validation_strategy>

### Quality Assurance Integration

<quality_assurance>
- Code review preparation: Identify areas requiring review and validation
- Documentation requirements: Plan code documentation and technical specifications
- Compliance verification: Ensure implementation meets project standards
- Rollback planning: Define rollback procedures if implementation fails
- Monitoring setup: Plan logging, monitoring, and observability integration
- Deployment considerations: Address deployment, configuration, and environment needs
</quality_assurance>

## Content Guidelines

### Technical Planning Principles for LLM Self-Evaluation

- Implementation-Ready Planning: All technical decisions documented with specific details and rationale
- Executable Task Enhancement: Every task/subtask enhanced to be executable without additional research
- Comprehensive Technical Validation: Complete testing and verification approach planned for implementation
- Architecture Integration: All integration points, dependencies, and technical standards identified

### LLM Error Prevention Checklist

- Avoid: Generic planning without specific technical details (libraries, versions, file paths)
- Avoid: Task enhancement without validation commands and success criteria
- Avoid: Implementation planning that ignores existing project structure and patterns
- Reference: Use [story.md](./.krci-ai/templates/story.md) template for consistent enhancement formatting
