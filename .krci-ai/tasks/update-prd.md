# Task: Update Product Requirements Document

## Description

Update an existing PRD with new requirements, scope changes, or refined business needs while maintaining traceability to Project Brief. Focus on change impact assessment and clear documentation to ensure requirements remain aligned with strategic objectives while defining epic-level features within the PRD.

## Prerequisites

- [ ] **Existing PRD**: `/docs/prd/prd.md` exists and is properly accessible
- [ ] **Change trigger**: Clear reason for update (Project Brief changes, user research, business priorities, technical constraints, stakeholder feedback)
- [ ] **Stakeholder input**: Understanding of what specifically needs to change and why
- [ ] **Epic/Story review**: Current understanding of feature groupings and requirements structure

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/prd-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

### CRITICAL: MANDATORY USER CONSULTATION FIRST

Before making ANY changes to the PRD, you MUST:

1. **Ask the user** what specific updates they want to make to the PRD
2. **Understand the trigger** for the changes (new requirements, stakeholder feedback, market changes, etc.)
3. **Clarify scope** which sections need updating and why
4. **Get approval** for the proposed changes before implementation
5. **Wait for explicit confirmation** before proceeding with any edits

### ONLY AFTER USER CONFIRMATION

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for change management process
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Maintain [prd-template.md](./.krci-ai/templates/prd-template.md) structure
4. **Maintain traceability**: Update BR/NFR numbering and include epic-level feature definitions

## Output Format

- **Location**: Updates existing `/docs/prd/prd.md` (EXACT path and filename)
- **Length**: Maintain 6-8 pages maximum
- **Requirements Format**: Maintain BR1, BR2, BR3... and NFR1, NFR2, NFR3... numbering with P0/P1/P2 priority indicators and epic-level feature definitions
- **Impact Documentation**: Clear notes on what changed and feature impact
- **Downstream Updates**: List of feature areas requiring updates

## Success Criteria

- [ ] **File updated** at `/docs/prd/prd.md` reflects all changes
- [ ] **Requirements numbered** BR/NFR structure maintained with priority indicators and epic-level features
- [ ] **Change documented** clear record of what changed and why
- [ ] **Feature impact** identified which feature areas need updates
- [ ] **Quality maintained** document remains 6-8 pages maximum
- [ ] **Project Brief alignment** changes align with Project Brief updates (if any)
- [ ] **Stakeholder approval** key stakeholders have approved requirement changes

## Execution Checklist

### User Consultation Phase (MANDATORY FIRST STEP)

- [ ] **User interview**: Ask user what specific changes they want to make to the PRD
- [ ] **Change justification**: Understand why these changes are needed (stakeholder feedback, new requirements, market changes, etc.)
- [ ] **Scope definition**: Clarify which PRD sections need updating and what specific content changes are required
- [ ] **Impact discussion**: Explain potential impact on existing features to user
- [ ] **User approval**: Get explicit user confirmation before proceeding with any changes
- [ ] **Change plan agreement**: Confirm the proposed approach with user before implementation

### Assessment Phase (ONLY AFTER USER APPROVAL)

- [ ] **Change scope**: Identify which sections need updating based on user requirements
- [ ] **Impact analysis**: Evaluate how changes affect existing feature definitions and requirements structure
- [ ] **Stakeholder review**: Confirm who needs to approve these changes before implementation
- [ ] **Requirements mapping**: Understand which BR/NFR numbers and priorities are affected

### Requirements Phase

- [ ] **Business requirements**: Update BR1, BR2, BR3... with new business functionality needs
- [ ] **Non-functional requirements**: Update NFR1, NFR2, NFR3... with new system behavior/performance needs
- [ ] **Priority assessment**: Review and update P0/P1/P2 priority indicators as needed
- [ ] **Epic groupings**: Ensure updated requirements can be organized into logical epic-level features within the PRD

### Update Phase

- [ ] **Section updates**: Modify specific sections using [prd-template.md](./.krci-ai/templates/prd-template.md) structure
- [ ] **Content integration**: Ensure changes are properly integrated without breaking flow
- [ ] **Length verification**: Confirm document remains 6-8 pages maximum
- [ ] **Quality validation**: Verify all changes maintain PRD quality standards

### Change Management Phase

- [ ] **Feature impact assessment**: Determine which feature areas need updating based on requirement changes
- [ ] **Team communication**: Notify development teams of requirement changes
- [ ] **Documentation**: Record change rationale and feature impact plan

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change Impact Focused**: Always assess feature impact before implementing PRD changes
- **Requirement Versioning**: Maintain BR/NFR numbering and priority consistency with epic-level feature definitions
- **Stakeholder Aligned**: Ensure all requirement changes have proper approval before implementation
- **Quality Preserved**: Keep updates within 6-8 page limit while maintaining user-centered focus

### LLM Error Prevention Checklist

- **NEVER**: Start making PRD changes without explicit user consultation and approval
- **NEVER**: Assume what changes the user wants - always ask for specific requirements
- **Avoid**: Breaking existing BR/NFR numbering that features depend on
- **Avoid**: Making changes without assessing feature impact
- **Avoid**: Updating requirements without proper stakeholder approval process
- **Always**: Wait for user confirmation before proceeding with any edits
- **Reference**: Use [prd-template.md](./.krci-ai/templates/prd-template.md) for all formatting consistency

### SDLC Integration Context

This update enables continued development by maintaining requirement traceability, preserving BR/NFR numbering with epic-level features, and communicating changes to development teams with clear impact assessment and timeline considerations.
