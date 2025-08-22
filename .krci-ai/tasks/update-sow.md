# Task: Update Scope of Work (SoW)

## Description

Update an existing Scope of Work document to reflect approved scope changes, requirement modifications, or constraint updates. Following PMBoK scope management processes, scope changes require formal change control, impact analysis, and stakeholder approval. This task ensures the SoW remains current and accurately reflects the authorized project scope.

## Prerequisites

- [ ] Existing SoW located at `/docs/project-management/scope-of-work.md`
- [ ] Approved change request or scope change order
- [ ] Impact analysis completed for proposed changes
- [ ] Stakeholder approval obtained for scope modifications
- [ ] Updated requirements or constraints documented

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/project-management-methodology.md
- ./.krci-ai/templates/sow-template.md
- /docs/project-management/scope-of-work.md (existing SoW)
- /docs/project-management/project-charter.md (approved charter)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Review existing SoW**: Analyze current scope documentation for areas requiring updates
2. **Apply change control**: Follow PMBoK scope change control processes
3. **Impact assessment**: Evaluate impact on schedule, budget, resources, and quality
4. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for artifact dependencies
5. **Maintain consistency**: Use [sow-template.md](./.krci-ai/templates/sow-template.md) structure
6. **Document changes**: Track all changes with rationale and approval

## Output Format

- **Location**: `/docs/project-management/scope-of-work.md` (UPDATE existing file)
- **Change tracking**: Document revision history and change rationale
- **Baseline update**: Update scope baseline as approved

## Success Criteria

- [ ] **Existing SoW reviewed** and areas for update identified
- [ ] **Change impact assessed** on schedule, budget, resources, and quality
- [ ] **WBS updated** to reflect scope changes
- [ ] **Deliverables modified** with updated acceptance criteria
- [ ] **Dependencies updated** based on scope changes
- [ ] **Change log maintained** with revision history
- [ ] **Stakeholder approval** obtained for significant changes
- [ ] **Scope baseline updated** to reflect approved changes
- [ ] **File updated** at `/docs/project-management/scope-of-work.md`

## Execution Checklist

### Change Analysis Phase

- [ ] **Current scope review**: Analyze existing SoW for completeness and accuracy
- [ ] **Change identification**: Identify specific scope changes and requirements
- [ ] **Impact assessment**: Evaluate impact on project constraints (time, cost, quality, scope)
- [ ] **Dependency analysis**: Assess how changes affect project dependencies

### Change Planning Phase

- [ ] **Change prioritization**: Prioritize scope changes based on business value and impact
- [ ] **Resource impact**: Assess resource requirements for scope changes
- [ ] **Timeline adjustment**: Evaluate schedule impact of scope modifications
- [ ] **Budget implications**: Calculate cost impact of scope changes

### Implementation Phase

- [ ] **SoW modification**: Update SoW using [sow-template.md](./.krci-ai/templates/sow-template.md)
- [ ] **WBS revision**: Update work breakdown structure for new scope elements
- [ ] **Deliverable updates**: Modify deliverables and acceptance criteria as needed
- [ ] **Change documentation**: Document all changes with rationale and approval

### Validation Phase

- [ ] **Content review**: Ensure all sections are complete and consistent
- [ ] **Stakeholder validation**: Review changes with affected stakeholders
- [ ] **Approval process**: Obtain required approvals for scope changes
- [ ] **Baseline update**: Update scope baseline and communicate changes

### Communication Phase

- [ ] **Team notification**: Communicate scope changes to project team
- [ ] **Stakeholder update**: Notify all stakeholders of approved changes
- [ ] **Documentation update**: Update related project documents
- [ ] **File update**: Save updated SoW to `/docs/project-management/scope-of-work.md`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change Traceability**: Clearly document what changed, why, and who approved it
- **Impact Transparency**: Provide clear analysis of change impacts on project constraints
- **Scope Control**: Maintain strict control over scope changes to prevent scope creep
- **Stakeholder Alignment**: Ensure all stakeholders understand and approve scope changes

### LLM Error Prevention Checklist

- **Avoid**: Making scope changes without proper impact analysis
- **Avoid**: Updating SoW without updating related project documents (plans, schedules)
- **Avoid**: Implementing changes without proper stakeholder approval
- **Avoid**: Losing traceability of original scope decisions
- **Reference**: Use existing SoW structure and [sow-template.md](./.krci-ai/templates/sow-template.md)

### PMBoK Integration Context

SoW updates follow PMBoK scope change control processes by requiring formal change requests, impact analysis, integrated change control, and scope baseline updates. Updated SoW may trigger updates to project management plans, schedules, budgets, and risk registers to maintain project integration. 