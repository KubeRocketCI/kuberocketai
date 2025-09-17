---
dependencies:
  data:
    - krci-ai/core-sdlc-framework.md
    - coding-standards.md
    - best-practices.md
---

# Task: Implement Feature

## Description

Implement Story requirements according to Architecture specifications and coding standards, ensuring quality and maintaining system consistency. This task enables developers to systematically transform user stories into working code while maintaining Epic alignment and architectural compliance.

## Prerequisites

<prerequisites>
- Story ready: Story has been reviewed and validated with complete Tasks/Subtasks
- Technical requirements clear: All implementation details, file paths, and commands specified
- Development environment: Project codebase access and development tools configured
- Dependencies available: Required libraries, tools, and systems accessible
</prerequisites>

## Instructions

<instructions>
1. Follow SDLC workflow: Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for implementation dependencies and handoff requirements
2. Apply coding standards: Use guidelines from [coding-standards.md](./.krci-ai/data/coding-standards.md) and [best-practices.md](./.krci-ai/data/best-practices.md)
3. Document progress: Update Story file with implementation progress and results
4. Ensure Story compliance: Implement all Story acceptance criteria following Architecture specifications
5. Maintain Epic traceability: Preserve Story's Epic alignment and contribute to Epic completion
</instructions>

### Ready to Implement

Assume story is ready with complete Tasks/Subtasks for step-by-step execution.

## Implementation Approach

### STEP-BY-STEP Implementation

<implementation_steps>
1. Update story status to "In Progress"
2. Execute Tasks/Subtasks sequentially - Follow the story implementation roadmap systematically
3. Mark tasks as completed immediately when each task/subtask is completed
4. Run tests and validation - Check project documentation (README.md, Makefile, package.json) for test/build commands
5. Update story status to "Completed" when all tasks done
6. ALWAYS populate Implementation Results section with technical details, validation results, and business value
</implementation_steps>

<critical_documentation>
- IMPORTANT!!! Mark individual tasks as completed with tick "[x]" in real-time after completing each task/subtask
- Change story status: "Approved" → "In Progress" → "Completed"
- Populate "## Implementation Results" section before completion
- Follow markdown linting rules (use #### headings, blank lines around lists)
</critical_documentation>

## Output Format

- **Location**: Working code implementation with updated Story file in `/docs/stories/`
- **Story completion**: All empty sections populated with implementation details
- **Progress tracking**: Real-time updates to Tasks/Subtasks completion status
- **Quality documentation**: Completed QA Checklist and Implementation Results

## Success Criteria

<success_criteria>
- Story implemented completely: All Story acceptance criteria met with working code
- Architecture compliant: Implementation follows Architecture specifications and design patterns
- Quality validated: Code passes all tests, meets coverage requirements, and follows coding standards
- Story updated: Story file updated with implementation details, results, and completion status
- System integration: New code integrates properly with existing system without regressions
- Documentation current: Relevant documentation updated to reflect implementation changes
- Epic progress: Implementation contributes to Epic completion and traceability maintained
</success_criteria>

## Execution Checklist

### Setup

<setup_tasks>
- Locate story: Find Story file in `/docs/stories/{epic_number}.{story_number}.story.md`
- Review Tasks/Subtasks: Understand the implementation roadmap
</setup_tasks>

### Execute Tasks/Subtasks

<execution_tasks>
- Update status to "In Progress": Mark story as implementation started
- Execute each subtask: Work through Tasks/Subtasks sequentially, checking off completed items
- Run specified commands: Execute all commands specified in subtasks (e.g., `create file: path/file.ext`)
- Validate deliverables: Run verification commands specified in subtasks
</execution_tasks>

### Complete Story

<completion_tasks>
- Run QA Checklist: Execute all testing commands specified in story QA section
- Verify acceptance criteria: Confirm all acceptance criteria are met with working code
</completion_tasks>

### Document Results

<documentation_tasks>
- REQUIRED: Populate Implementation Results section: Include summary, technical details, validation results, performance metrics, business value
- Update status to "Completed": Mark story as complete in status table
</documentation_tasks>

## Implementation Guidelines

### Simple Execution Rules

<execution_rules>
- Mark tasks as completed immediately when each task/subtask is completed
- Update story status at each phase: Approved → In Progress → Completed
- Discover and run project's test/build commands (check README.md, Makefile, package.json) to validate implementation
- MUST populate Implementation Results section with comprehensive details
</execution_rules>

### If Something Is Unclear

<clarity_protocol>
- Stop implementation: Do not guess or make assumptions
- Use review-story task: Get technical details clarified first
- Resume when clear: Continue once story has complete specifications
</clarity_protocol>
