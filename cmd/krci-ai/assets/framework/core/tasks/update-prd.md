# Task: Update Product Requirements Document

## Description

Update an existing PRD with new requirements, scope changes, or refined business needs while maintaining traceability to Project Brief and enabling Epic/Story creation. Focus on change impact assessment and downstream communication to ensure existing Epics and Stories remain aligned with updated requirements.

## Prerequisites

- [ ] **Existing PRD**: `/docs/prd/prd.md` exists and is properly accessible
- [ ] **Change trigger**: Clear reason for update (Project Brief changes, user research, business priorities, technical constraints, stakeholder feedback)
- [ ] **Stakeholder input**: Understanding of what specifically needs to change and why
- [ ] **Epic/Story review**: Current status of downstream artifacts

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/prd-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for change management process
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Maintain [prd-template.md](./.krci-ai/templates/prd-template.md) structure
4. **Maintain traceability**: Update BR/NFR numbering and assess Epic/Story impact

## Output Format

- **Location**: Updates existing `/docs/prd/prd.md` (EXACT path and filename)
- **Length**: Maintain 6-8 pages maximum
- **Requirements Format**: Maintain BR1, BR2, BR3... and NFR1, NFR2, NFR3... numbering with P0/P1/P2 priority indicators
- **Impact Documentation**: Clear notes on what changed and Epic/Story impact
- **Downstream Updates**: List of Epic/Story artifacts requiring updates

## Success Criteria

- [ ] **File updated** at `/docs/prd/prd.md` reflects all changes
- [ ] **Requirements numbered** BR/NFR structure maintained with priority indicators for Epic mapping
- [ ] **Change documented** clear record of what changed and why
- [ ] **Downstream impact** identified which Epic/Story artifacts need updates
- [ ] **Quality maintained** document remains 6-8 pages maximum
- [ ] **Project Brief alignment** changes align with Project Brief updates (if any)
- [ ] **Stakeholder approval** key stakeholders have approved requirement changes

## Execution Checklist

### Assessment Phase

- [ ] **Change scope**: Identify which sections need updating (Problem/Opportunity, Target Users & Use Cases, Current Journeys/Landscape, Proposed Solution/Elevator Pitch, Goals/Measurable Outcomes, MVP/Functional Requirements)
- [ ] **Impact analysis**: Evaluate how changes affect existing Epics (`/docs/epics/`) and Stories (`/docs/stories/`)
- [ ] **Stakeholder review**: Confirm who needs to approve these changes before implementation
- [ ] **Requirements mapping**: Understand which BR/NFR numbers and priorities are affected

### Requirements Phase

- [ ] **Business requirements**: Update BR1, BR2, BR3... with new business functionality needs
- [ ] **Non-functional requirements**: Update NFR1, NFR2, NFR3... with new system behavior/performance needs
- [ ] **Priority assessment**: Review and update P0/P1/P2 priority indicators as needed
- [ ] **Epic mapping**: Ensure updated requirements can still map to logical Epic groupings

### Update Phase

- [ ] **Section updates**: Modify specific sections using [prd-template.md](./.krci-ai/templates/prd-template.md) structure
- [ ] **Content integration**: Ensure changes are properly integrated without breaking flow
- [ ] **Length verification**: Confirm document remains 6-8 pages maximum
- [ ] **Quality validation**: Verify all changes maintain PRD quality standards

### Change Management Phase

- [ ] **Epic impact assessment**: Determine which Epics (`{epic_number}-epic-{slug}.md`) need updating
- [ ] **Story impact review**: Assess if in-progress Stories (`{epic_number}.{story_number}.story.md`) are affected
- [ ] **Team communication**: Notify Epic owners and development teams of changes
- [ ] **Documentation**: Record change rationale and downstream impact plan

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change Impact Focused**: Always assess Epic/Story impact before implementing PRD changes
- **Requirement Versioning**: Maintain BR/NFR numbering and priority consistency for Epic traceability
- **Stakeholder Aligned**: Ensure all requirement changes have proper approval before implementation
- **Quality Preserved**: Keep updates within 6-8 page limit while maintaining user-centered focus

### LLM Error Prevention Checklist

- **Avoid**: Breaking existing BR/NFR numbering that Epics depend on
- **Avoid**: Making changes without assessing downstream Epic/Story impact
- **Avoid**: Updating requirements without proper stakeholder approval process
- **Reference**: Use [prd-template.md](./.krci-ai/templates/prd-template.md) for all formatting consistency

### SDLC Integration Context

This update enables continued Epic/Story development by maintaining requirement traceability, preserving BR/NFR numbering for Epic mapping, assessing which in-progress Stories need updates, and communicating changes to development teams with clear impact assessment and timeline considerations.
