# Task: Update Epic

## Description

Update existing epic with new requirements, scope additions, or refinements while preserving completed work and maintaining story traceability. This task enables controlled epic evolution during implementation while protecting development team progress and ensuring PRD alignment.

## Prerequisites

- [ ] **Epic exists**: Target epic file exists in `/docs/epics/` with current implementation status
- [ ] **Change justification**: Clear business reason for epic update (new PRD requirements, scope clarification, story feedback)
- [ ] **Impact assessment**: Understanding of how changes affect in-progress or completed Stories
- [ ] **Stakeholder approval**: Product Owner and relevant stakeholders have approved the update scope

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/prioritization-frameworks.md
- ./.krci-ai/templates/epic.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

### CRITICAL: MANDATORY USER CONSULTATION FIRST

Before making ANY changes to the Epic, you MUST:

1. **Ask the user** what specific updates they want to make to the Epic
2. **Understand the trigger** for the changes (new PRD requirements, scope clarification, story feedback, etc.)
3. **Clarify scope** which sections need updating and why
4. **Get approval** for the proposed changes before implementation
5. **Wait for explicit confirmation** before proceeding with any edits

### ONLY AFTER USER CONFIRMATION

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for Epic update dependencies and downstream impact
2. **Assess current status**: Review epic status and dependent Stories before making changes
3. **Apply change control**: Use methods from [prioritization-frameworks.md](./.krci-ai/data/prioritization-frameworks.md) for update prioritization
4. **Maintain format**: Keep [epic.md](./.krci-ai/templates/epic.md) template structure and populate new variables
5. **Preserve traceability**: Maintain existing PRD references while adding new BR/NFR connections
6. **Ensure agent-optimization**: Maintain enhanced template compliance during all updates

### Update Validation Process

Before applying any update:

1. **Check Story Status**: Verify no completed Stories become invalid
2. **Validate Dependencies**: Ensure new dependencies don't conflict with completed work
3. **Assess Timeline**: Calculate realistic timeline changes for scope additions
4. **Document Rationale**: Record business justification for every change

## Update Types & Restrictions

### ALLOWED Updates (Safe Changes)

**Note**: All updates must maintain agent-optimized template compliance (Goal format, Dependencies grouping, Acceptance Criteria commands)

- **Add new Stories** to existing epic scope without changing completed work
- **Expand acceptance criteria** with additional validation requirements and commands
- **Add new dependencies** that don't conflict with completed Stories (maintain Epic/System/External grouping)
- **Clarify problem statement** with additional context or user research
- **Extend timeline** for additional scope (with stakeholder approval)
- **Add new target users** without removing existing personas
- **Enhance solution approach** with additional technical considerations

### RESTRICTED Updates (Requires Validation)

- **Modify goal metrics** - requires Story impact assessment and team validation
- **Change scope boundaries** - must verify no completed Stories become out-of-scope
- **Update dependencies** - requires dependency chain validation for affected Stories
- **Alter timeline** - needs development team impact assessment
- **Modify acceptance criteria** - must not invalidate completed Story validation

### FORBIDDEN Updates (Never Change)

- **Remove completed scope** - never remove features from completed Stories
- **Delete existing Stories** - completed or in-progress Stories cannot be removed
- **Change epic number** - epic numbering is immutable for traceability
- **Reduce problem scope** - cannot narrow problem if Stories address broader scope
- **Remove target users** - cannot remove personas if Stories serve them

## Output Format

- **Location**: Update existing `/docs/epics/{epic_number}-epic-{slug}.md` file in place
- **Version tracking**: Add update timestamp and change summary to file header
- **Change log**: Document what was updated and rationale in epic comments
- **Story impact**: Note which Stories are affected by changes

## Success Criteria

- [ ] **Epic updated** in place with version tracking and change documentation
- [ ] **Story compatibility** all existing Stories remain valid and implementable
- [ ] **PRD traceability** new changes connect to specific BR/NFR requirements
- [ ] **Change justification** clear business rationale documented for updates
- [ ] **Impact assessment** downstream Story effects identified and communicated
- [ ] **Template compliance** all template variables updated correctly
- [ ] **Agent-optimization maintained** epic preserves enhanced template structure after updates
- [ ] **Lifecycle-appropriate content** epic content matches status (Planning/In-Progress/Complete)
- [ ] **Goal format preserved** goal maintains "Enable X for Y% within Z timeline" pattern
- [ ] **Dependencies structure intact** Epic/System/External grouping maintained with exact headers
- [ ] **Acceptance criteria updated** validation commands reflect current implementation status
- [ ] **Quality gates passed** all validation commands execute successfully post-update

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- [ ] **User interview**: Ask user what specific changes they want to make to the Epic
- [ ] **Change justification**: Understand why these changes are needed (new PRD requirements, scope clarification, story feedback, etc.)
- [ ] **Scope definition**: Clarify which Epic sections need updating and what specific content changes are required
- [ ] **Impact discussion**: Explain potential impact on existing Stories to user
- [ ] **User approval**: Get explicit user confirmation before proceeding with any changes
- [ ] **Change plan agreement**: Confirm the proposed approach with user before implementation

### Pre-Update Assessment (ONLY AFTER USER APPROVAL)

- [ ] **Status review**: Check epic status (Planning, In Progress, Implementation, Testing, Complete)
- [ ] **Story analysis**: Review dependent Stories and their current implementation status
- [ ] **Change scope**: Define exactly what needs updating and why
- [ ] **Impact evaluation**: Assess how changes affect existing work and timeline
- [ ] **Template compliance check**: Verify current epic follows agent-optimized template structure
- [ ] **Content structure assessment**: Verify epic maintains agent-optimized template structure
- [ ] **Validation command review**: Check if existing validation commands need updates

### Change Planning Phase

- [ ] **Update classification**: Determine if changes are ALLOWED, RESTRICTED, or FORBIDDEN
- [ ] **Stakeholder validation**: Confirm updates with development team and product stakeholders
- [ ] **Story impact**: Identify which Stories need corresponding updates
- [ ] **Timeline adjustment**: Calculate any timeline changes from scope additions
- [ ] **Template compliance planning**: Ensure updates maintain agent-optimized structure
- [ ] **Template compliance planning**: Ensure updates maintain agent-optimized structure
- [ ] **Validation command updates**: Identify validation commands that need modification

### Update Implementation Phase

- [ ] **Version header**: Add update timestamp and change summary to epic file
- [ ] **Content updates**: Apply approved changes using [epic.md](./.krci-ai/templates/epic.md) structure
- [ ] **Template compliance maintenance**: Preserve Goal format, Dependencies grouping, Acceptance Criteria commands
- [ ] **Template structure maintenance**: Preserve Goal format, Dependencies grouping, Acceptance Criteria
- [ ] **Validation command updates**: Update validation methods and commands as needed
- [ ] **Agent-optimization verification**: Ensure consistent structure for automated processing
- [ ] **Change documentation**: Document what changed and why in epic comments
- [ ] **Story synchronization**: Update affected Stories to maintain epic alignment
- [ ] **Quality gates validation**: Run validation commands to verify template compliance post-update

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change-Controlled**: Only make approved updates that don't invalidate completed work
- **Impact-Aware**: Consider and document effects on all dependent Stories
- **Traceability-Maintained**: Preserve existing PRD connections while adding new ones
- **Story-Compatible**: Ensure all existing Stories remain valid and implementable

### LLM Error Prevention Checklist

- **NEVER**: Start making Epic changes without explicit user consultation and approval
- **NEVER**: Assume what changes the user wants - always ask for specific requirements
- **Avoid**: Removing scope that has completed Stories implementation
- **Avoid**: Changing epic fundamentals (number, core problem) that break traceability
- **Avoid**: Updates that make in-progress Stories irrelevant or incorrect
- **Always**: Wait for user confirmation before proceeding with any edits
- **Reference**: Use change control principles to validate every update decision
