# Task: Update Story

## Description

Update existing user story with new requirements, scope refinements, or implementation changes while preserving completed tasks/subtasks and maintaining epic traceability. This task enables controlled story evolution during development while protecting team progress and ensuring Epic alignment.

## Prerequisites

- [ ] **Story exists**: Target story file exists in `/docs/stories/` with current implementation status
- [ ] **Change justification**: Clear business reason for story update (new requirements, scope clarification, task feedback)
- [ ] **Impact assessment**: Understanding of how changes affect in-progress or completed Tasks/Subtasks
- [ ] **Epic alignment**: Product Owner confirms updates maintain Epic goals and traceability

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/prioritization-frameworks.md
- ./.krci-ai/templates/story.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

### CRITICAL: MANDATORY USER CONSULTATION FIRST

Before making ANY changes to the Story, you MUST:

1. **Ask the user** what specific updates they want to make to the Story
2. **Understand the trigger** for the changes (new requirements, scope clarification, task feedback, etc.)
3. **Clarify scope** which sections need updating and why
4. **Get approval** for the proposed changes before implementation
5. **Wait for explicit confirmation** before proceeding with any edits

### ONLY AFTER USER CONFIRMATION

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for Story update dependencies and downstream impact
2. **Assess current status**: Review story status and Tasks/Subtasks before making changes
3. **Apply change control**: Use methods from [prioritization-frameworks.md](./.krci-ai/data/prioritization-frameworks.md) for update prioritization
4. **Maintain format**: Keep [story.md](./.krci-ai/templates/story.md) template structure and populate new variables
5. **Preserve Epic traceability**: Maintain existing Epic references while ensuring new changes align with Epic goals

### Update Validation Process

Before applying any update:

1. **Check Task Status**: Verify no completed Tasks/Subtasks become invalid
2. **Validate Epic Alignment**: Ensure changes maintain Epic traceability and goals
3. **Assess Acceptance Criteria**: Calculate impact on existing acceptance criteria and validation
4. **Document Rationale**: Record business justification for every change

## Update Types & Restrictions

### ALLOWED Updates (Safe Changes)

- **Add new Tasks/Subtasks** to existing story scope without changing completed work
- **Expand acceptance criteria** with additional validation requirements
- **Add new dependencies** that don't conflict with completed Tasks
- **Clarify story description** with additional context or user research
- **Extend story points** for additional scope (with development team validation)
- **Enhance QA checklist** with additional testing requirements
- **Update implementation results** with actual deliverables as work progresses

### RESTRICTED Updates (Requires Validation)

- **Modify story goal** - requires Epic alignment check and development team validation
- **Change acceptance criteria** - must verify no completed Tasks become invalid
- **Update dependencies** - requires dependency chain validation for affected Tasks
- **Alter story points** - needs development team estimation review
- **Modify task structure** - must not invalidate completed subtask validation

### FORBIDDEN Updates (Never Change)

- **Remove completed Tasks/Subtasks** - never remove work that has been completed
- **Delete completed acceptance criteria** - completed validation cannot be removed
- **Change story number** - story numbering is immutable for Epic traceability
- **Reduce story scope** - cannot narrow scope if Tasks address broader functionality
- **Remove Epic reference** - Epic traceability must always be maintained

## Output Format

- **Location**: Update existing `/docs/stories/{epic_number}.{story_number}.story.md` file in place
- **Template**: Maintain [story.md](./.krci-ai/templates/story.md) structure (8 sections only)
- **Change Documentation**: Add timestamp to Status table, document changes in story comments
- **Content Updates**: Modify appropriate sections based on change type (AC, Tasks, Description)
- **Verification**: File maintains valid template structure with documented change history

## Success Criteria

- [ ] **Story updated** in place with version tracking and change documentation
- [ ] **Task compatibility** all existing Tasks/Subtasks remain valid and implementable
- [ ] **Epic traceability** story maintains alignment with parent Epic goals
- [ ] **Change justification** clear business rationale documented for updates
- [ ] **Impact assessment** downstream Task effects identified and communicated
- [ ] **Template compliance** all template variables updated correctly

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- [ ] **User interview**: Ask user what specific changes they want to make to the Story
- [ ] **Change justification**: Understand why these changes are needed (new requirements, scope clarification, task feedback, etc.)
- [ ] **Scope definition**: Clarify which Story sections need updating and what specific content changes are required
- [ ] **Impact discussion**: Explain potential impact on existing Tasks/Subtasks to user
- [ ] **User approval**: Get explicit user confirmation before proceeding with any changes
- [ ] **Change plan agreement**: Confirm the proposed approach with user before implementation

### Pre-Update Assessment (ONLY AFTER USER APPROVAL)

- [ ] **Status review**: Check story status (Pending, In Progress, Approved, Completed)
- [ ] **Task analysis**: Review Tasks/Subtasks and their current implementation status
- [ ] **Change scope**: Define exactly what needs updating and why
- [ ] **Epic validation**: Confirm changes maintain Epic alignment and traceability

### Change Planning Phase

- [ ] **Update classification**: Determine if changes are ALLOWED, RESTRICTED, or FORBIDDEN
- [ ] **Team validation**: Confirm updates with development team and Epic stakeholders
- [ ] **Task impact**: Identify which Tasks/Subtasks need corresponding updates
- [ ] **Story points adjustment**: Recalculate story complexity if scope changes

### Update Implementation Phase

- [ ] **Version header**: Add update timestamp and change summary to story file
- [ ] **Content updates**: Apply approved changes using [story.md](./.krci-ai/templates/story.md) structure
- [ ] **Change documentation**: Document what changed and why in story comments
- [ ] **Task synchronization**: Update affected Tasks/Subtasks to maintain story alignment

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change-Controlled**: Only make approved updates that don't invalidate completed work
- **Epic-Aligned**: Ensure all changes maintain alignment with parent Epic goals
- **Task-Compatible**: Preserve all existing Tasks/Subtasks that have been completed
- **Traceability-Maintained**: Keep Epic references and story numbering intact

### LLM Error Prevention Checklist

- **NEVER**: Start making Story changes without explicit user consultation and approval
- **NEVER**: Assume what changes the user wants - always ask for specific requirements
- **Avoid**: Removing Tasks/Subtasks that have completed implementation
- **Avoid**: Changing story fundamentals (number, Epic reference) that break traceability
- **Avoid**: Updates that make completed acceptance criteria irrelevant
- **Always**: Wait for user confirmation before proceeding with any edits
- **Reference**: Use change control principles to validate every update decision
