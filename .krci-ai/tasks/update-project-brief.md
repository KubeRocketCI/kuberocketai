# Task: Update Project Brief

## Description

Update an existing project brief with new information, scope changes, or refined understanding while maintaining strategic alignment and enabling downstream SDLC artifacts. Focus on change impact assessment and downstream artifact management to ensure existing PRD and Epic artifacts remain aligned with strategic changes.

## Prerequisites

- [ ] **Existing Project Brief**: `/docs/prd/project-brief.md` exists and is properly accessible
- [ ] **Change trigger**: Clear reason for update (strategic shifts, market changes, new insights, stakeholder feedback, resource changes)
- [ ] **Impact assessment**: Understanding of how changes affect dependent PRD and downstream artifacts
- [ ] **Stakeholder buy-in**: Key stakeholders aware of planned strategic changes

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/project-brief-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

### CRITICAL: MANDATORY USER CONSULTATION FIRST

Before making ANY changes to the Project Brief, you MUST:

1. **Ask the user** what specific updates they want to make to the Project Brief
2. **Understand the trigger** for the changes (strategic shifts, market changes, stakeholder feedback, resource changes, etc.)
3. **Clarify scope** which sections need updating and why
4. **Get approval** for the proposed changes before implementation
5. **Wait for explicit confirmation** before proceeding with any edits

### ONLY AFTER USER CONFIRMATION

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for change impact assessment
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Maintain [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
4. **Assess downstream impact**: Identify which PRD artifacts need updates

## Output Format

- **Location**: Updates existing `/docs/prd/project-brief.md` (EXACT path and filename)
- **Length**: Maintain 2-3 pages maximum
- **Impact Documentation**: Clear notes on what changed and downstream impact
- **Downstream Updates**: List of PRD artifacts requiring updates

## Success Criteria

- [ ] **File updated** at `/docs/prd/project-brief.md` reflects all changes
- [ ] **Change documented** with clear record of what changed and why
- [ ] **Downstream impact** identified which PRD artifacts need updates
- [ ] **Quality maintained** document remains 2-3 pages maximum
- [ ] **Strategic alignment** changes support overall product strategy
- [ ] **Stakeholder communication** key stakeholders informed of strategic changes

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- [ ] **User interview**: Ask user what specific changes they want to make to the Project Brief
- [ ] **Change justification**: Understand why these changes are needed (strategic shifts, market changes, stakeholder feedback, resource changes, etc.)
- [ ] **Scope definition**: Clarify which Project Brief sections need updating and what specific content changes are required
- [ ] **Impact discussion**: Explain potential impact on existing PRD artifacts to user
- [ ] **User approval**: Get explicit user confirmation before proceeding with any changes
- [ ] **Change plan agreement**: Confirm the proposed approach with user before implementation

### Assessment Phase (ONLY AFTER USER APPROVAL)

- [ ] **Change scope**: Identify which sections need updating based on user requirements (Executive Summary, Problem, Opportunity, Users, Success Metrics, Constraints, Risks)
- [ ] **Business impact**: Analyze how changes affect product strategy and business case
- [ ] **Downstream impact**: Evaluate how changes affect existing PRD (`/docs/prd/prd.md`) artifacts
- [ ] **Stakeholder validation**: Confirm changes with key stakeholders

### Update Phase

- [ ] **Section updates**: Modify specific sections using [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
- [ ] **Strategic alignment**: Ensure updates maintain strategic coherence and business focus
- [ ] **Quality check**: Verify updated Project Brief maintains 2-3 page limit and foundation quality
- [ ] **Content validation**: Ensure all changes are properly integrated

### Change Management Phase

- [ ] **PRD impact analysis**: Determine if PRD needs updating based on Project Brief changes
- [ ] **Stakeholder communication**: Notify key stakeholders of strategic changes and implications
- [ ] **Documentation**: Record change rationale and downstream impact plan

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Strategic Focus**: Focus on strategic foundation changes rather than tactical adjustments
- **Foundation Strength**: Ensure changes strengthen rather than weaken the overall strategic foundation
- **Cascade Management**: Assess how strategic changes flow through PRD requirements
- **Long-term Alignment**: Consider long-term strategic implications beyond immediate tactical changes

### LLM Error Prevention Checklist

- **NEVER**: Start making Project Brief changes without explicit user consultation and approval
- **NEVER**: Assume what changes the user wants - always ask for specific requirements
- **Avoid**: Making changes without clear strategic justification and stakeholder approval
- **Avoid**: Updating without assessing downstream PRD impact
- **Avoid**: Expanding scope beyond strategic foundation changes into tactical details
- **Always**: Wait for user confirmation before proceeding with any edits
- **Reference**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for all formatting consistency

### SDLC Integration Context

This update enables continued strategic alignment by managing strategic changes flowing through PRD requirements, ensuring stakeholder approval of strategic changes, and maintaining clear documentation of strategic change rationale and downstream PRD impact.
