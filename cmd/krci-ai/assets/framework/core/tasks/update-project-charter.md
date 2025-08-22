# Task: Update Project Charter

## Description

Update an existing project charter to reflect approved changes in project scope, objectives, stakeholders, or constraints. Following PMBoK standards, charter updates require formal change control processes and sponsor approval. This task ensures the project charter remains current and accurately reflects the authorized project parameters.

## Prerequisites

- [ ] Existing project charter located at `/docs/project-management/project-charter.md`
- [ ] Change request or approved scope change available
- [ ] Impact analysis completed
- [ ] Stakeholder approval for changes obtained
- [ ] Updated requirements or constraints documented

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/project-management-methodology.md
- ./.krci-ai/templates/project-charter-template.md
- /docs/project-management/project-charter.md (existing charter)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Review existing charter**: Analyze current project charter for areas requiring updates
2. **Apply change control**: Follow PMBoK change control processes
3. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for artifact dependencies
4. **Maintain consistency**: Use [project-charter-template.md](./.krci-ai/templates/project-charter-template.md) structure
5. **Document changes**: Clearly track what changed and why

## Output Format

- **Location**: `/docs/project-management/project-charter.md` (UPDATE existing file)
- **Change tracking**: Document revision history and change rationale
- **Approval tracking**: Update approval signatures and dates

## Success Criteria

- [ ] **Existing charter reviewed** and areas for update identified
- [ ] **Changes documented** with clear rationale and impact analysis
- [ ] **Updated objectives** remain SMART and measurable
- [ ] **Stakeholder list updated** with any new or changed roles
- [ ] **Risk register updated** to reflect new or changed risks
- [ ] **Budget and timeline adjusted** if scope changes occurred
- [ ] **Change log updated** with revision history
- [ ] **Approval process completed** for significant changes
- [ ] **File updated** at `/docs/project-management/project-charter.md`

## Execution Checklist

### Change Analysis Phase

- [ ] **Current state review**: Analyze existing charter for completeness and accuracy
- [ ] **Change identification**: Identify specific areas requiring updates
- [ ] **Impact assessment**: Evaluate impact of changes on scope, schedule, budget, and risks
- [ ] **Stakeholder impact**: Assess how changes affect stakeholder roles and responsibilities

### Update Planning Phase

- [ ] **Change prioritization**: Prioritize updates based on project impact and urgency
- [ ] **Approval requirements**: Determine what level of approval is needed for changes
- [ ] **Communication planning**: Plan how to communicate changes to stakeholders
- [ ] **Timeline impact**: Assess if charter changes affect project timeline or deliverables

### Implementation Phase

- [ ] **Charter modification**: Update charter using [project-charter-template.md](./.krci-ai/templates/project-charter-template.md)
- [ ] **Change documentation**: Document all changes with rationale and approval
- [ ] **Version control**: Update version number and revision history
- [ ] **Stakeholder notification**: Communicate changes to affected stakeholders

### Validation Phase

- [ ] **Content review**: Ensure all sections are complete and consistent
- [ ] **Approval process**: Obtain required approvals for charter changes
- [ ] **File update**: Save updated charter to `/docs/project-management/project-charter.md`
- [ ] **Change communication**: Notify team of charter updates

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change Traceability**: Clearly document what changed, why, and who approved it
- **Impact Clarity**: Explain how changes affect project scope, schedule, budget, and risks
- **Stakeholder Alignment**: Ensure all stakeholders understand and approve changes
- **Version Control**: Maintain clear revision history and version tracking

### LLM Error Prevention Checklist

- **Avoid**: Making changes without proper impact analysis or approval
- **Avoid**: Losing original charter content without proper change documentation
- **Avoid**: Updating charter without updating dependent artifacts (project plans, risk registers)
- **Reference**: Use existing charter structure and [project-charter-template.md](./.krci-ai/templates/project-charter-template.md)

### PMBoK Integration Context

Charter updates follow PMBoK change control processes by requiring formal change requests, impact analysis, stakeholder approval, and documentation. Updated charters may trigger updates to project management plans, scope of work documents, and risk registers to maintain artifact consistency. 