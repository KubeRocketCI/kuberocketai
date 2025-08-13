# Task: Create Story

## Description

Create user-focused requirements with implementation tasks and acceptance criteria that break down Epic features into actionable development work. This story provides specific user value and clear implementation guidance for development teams.

## Prerequisites

- [ ] **Epic exists**: Target Epic is defined and available in `/docs/epics/`
- [ ] **Epic context understood**: Epic problem, goal, and scope are clear
- [ ] **User persona identified**: Target user from Epic is specified
- [ ] **Story scope defined**: Specific functionality this Story will deliver

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/prioritization-frameworks.md
- ./.krci-ai/templates/story.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for Story dependencies and naming conventions
2. **Apply prioritization**: Use methods from [prioritization-frameworks.md](./.krci-ai/data/prioritization-frameworks.md)
3. **Format output**: Use [story.md](./.krci-ai/templates/story.md) template with proper naming
4. **Ensure Epic traceability**: Reference parent Epic and map to specific Epic deliverables

## Output Format

- **Location**: `/docs/stories/{epic_number}.{story_number}.story.md` (EXACT naming pattern)
- **Story numbering**: Sequential number within Epic (e.g., 01.01, 01.02, 01.03 for Epic 1)
- **Epic reference**: Clear connection to parent Epic in format "Epic {number}: {title}"
- **Implementation Ready**: Story contains sufficient detail for autonomous development
- **Testing Ready**: Acceptance criteria provide clear validation steps for QA

## Success Criteria

- [ ] **File saved** to `/docs/stories/{epic_number}.{story_number}.story.md` with correct naming
- [ ] **Epic traceability** clear connection to parent Epic and its goals
- [ ] **User story format** follows "As a [user], I want [goal], so that [value]" structure
- [ ] **Acceptance criteria** specific, testable conditions for completion
- [ ] **Implementation ready** provides sufficient detail for development
- [ ] **Template compliance** all template variables populated correctly

## Execution Checklist

### Discovery & Planning

- [ ] **Epic verification**: Confirm Epic exists at `/docs/epics/{epic_number}-epic-{slug}.md`
- [ ] **Story numbering**: Check existing stories for next sequential number within Epic
- [ ] **User persona**: Extract target user from Epic's user definitions
- [ ] **Story scope**: Define specific functionality this Story delivers
- [ ] **Epic reference**: Create proper Epic reference format "Epic {number}: {title}"
- [ ] **Dependencies identification**: Identify other Stories or systems this depends on

### Story Definition

- [ ] **User story creation**: Write "As a [user], I want [goal], so that [value]" aligned with Epic features
- [ ] **Story points estimation**: Estimate complexity (1, 2, 3, 5, 8, 13) using Epic context
- [ ] **Business value validation**: Ensure Story delivers measurable user value
- [ ] **Status format**: Use table format as defined in template (| Field | Value |)

### Requirements Specification

- [ ] **Acceptance criteria**: Define specific, testable conditions with measurable outcomes and file deliverables
- [ ] **Technical approach**: Define overall implementation approach and strategy with specific technologies
- [ ] **Architecture references**: Include direct links to specific architecture sections needed (format: `[filename](path)`)

### Tasks/Subtasks Development

- [ ] **Task breakdown**: Create main implementation Tasks with:
  - **Acceptance Criteria mapping** (e.g., "Task 1: Implement Bundle Export Command (AC: 1)")
  - **Specific file paths** for inputs/outputs (e.g., `create file: /path/to/file.ext`)
  - **Command patterns** LLMs can execute (e.g., `run: command with args`)
  - **Dependencies** between tasks (e.g., "depends on Task X completion")
- [ ] **Subtask checklists**: Define atomic, executable subtasks as checkboxes:
  - **Single action per checkbox** (create, edit, run, validate)
  - **Specific file/directory targets** with exact paths
  - **Success criteria** for each subtask (file exists, tests pass, output matches)
  - **Error recovery steps** if subtask fails
- [ ] **Task formatting**: Use `**Task N: Description (AC: X, Y)**` format with bullet subtasks

### Quality Assurance Planning

- [ ] **QA checklist**: Define testing requirements with clear verification steps/commands and expected outputs
- [ ] **Validation plan**: Create testing checklist with:
  - **Verification methods** LLMs can run (commands where applicable)
  - **Expected outputs** and success indicators
  - **Rollback steps** if testing fails

### Documentation & Delivery

- [ ] **Story creation**: Use [story.md](./.krci-ai/templates/story.md) template structure
- [ ] **Variable population**: Complete all template variables ({{story_number}}, {{story_title}}, etc.)
- [ ] **Content validation**: Ensure user story, acceptance criteria, and Tasks/Subtasks are complete
- [ ] **File placement**: Save to exact location `/docs/stories/{epic_number}.{story_number}.story.md`
