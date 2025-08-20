# Task: Update Project Plan

## Description

Update an existing project management plan to reflect approved changes in scope, schedule, budget, resources, or other project parameters. Following PMBoK integrated change control processes, plan updates require formal change requests, impact analysis, and approval through the change control board. This task ensures the project management plan remains current and continues to guide effective project execution.

## Prerequisites

- [ ] Existing project plan located at `/docs/project-management/project-plan.md`
- [ ] Approved change request or change order available
- [ ] Impact analysis completed for all affected knowledge areas
- [ ] Change control board approval obtained
- [ ] Updated baseline data available

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/project-management-methodology.md
- ./.krci-ai/templates/project-plan-template.md
- /docs/project-management/project-plan.md (existing plan)
- /docs/project-management/project-charter.md (approved charter)
- /docs/project-management/scope-of-work.md (approved SoW)

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Review existing plan**: Analyze current project management plan for areas requiring updates
2. **Apply integrated change control**: Follow PMBoK integrated change control processes
3. **Impact assessment**: Evaluate changes across all knowledge areas and baselines
4. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for artifact dependencies
5. **Maintain integration**: Ensure all subsidiary plans remain integrated and consistent
6. **Document changes**: Track all changes with rationale, impact, and approval

## Output Format

- **Location**: `/docs/project-management/project-plan.md` (UPDATE existing file)
- **Change tracking**: Document revision history with change rationale
- **Baseline updates**: Update affected baselines as approved

## Success Criteria

- [ ] **Existing plan reviewed** and areas for update identified
- [ ] **Change impact assessed** across all knowledge areas
- [ ] **Schedule baseline updated** if scope or resource changes occurred
- [ ] **Cost baseline revised** if budget changes were approved
- [ ] **Quality plan updated** for new requirements or standards
- [ ] **Resource plan modified** for team or skill changes
- [ ] **Communication plan updated** for new stakeholders or processes
- [ ] **Risk register integrated** with updated risk responses
- [ ] **Change log maintained** with complete revision history
- [ ] **Integrated baselines** remain consistent and achievable
- [ ] **File updated** at `/docs/project-management/project-plan.md`

## Execution Checklist

### Change Analysis Phase

- [ ] **Current plan review**: Analyze existing plan for completeness and accuracy
- [ ] **Change identification**: Identify specific changes affecting the project management plan
- [ ] **Knowledge area impact**: Assess impact across all PMBoK knowledge areas
- [ ] **Baseline impact**: Evaluate changes to scope, schedule, and cost baselines

### Impact Assessment Phase

- [ ] **Schedule impact**: Analyze effect on project timeline and critical path
- [ ] **Cost impact**: Evaluate budget implications and funding requirements
- [ ] **Resource impact**: Assess changes to team structure or skill requirements
- [ ] **Quality impact**: Review effect on quality standards and deliverables
- [ ] **Risk impact**: Identify new risks or changes to existing risks

### Integration Planning Phase

- [ ] **Subsidiary plan updates**: Plan updates to all affected management plans
- [ ] **Baseline revisions**: Plan baseline updates maintaining integration
- [ ] **Change sequencing**: Determine proper sequence for implementing changes
- [ ] **Communication planning**: Plan stakeholder communication for changes

### Implementation Phase

- [ ] **Plan modification**: Update plan using [project-plan-template.md](./.krci-ai/templates/project-plan-template.md)
- [ ] **Baseline updates**: Revise scope, schedule, and cost baselines as approved
- [ ] **Subsidiary plan integration**: Update all affected management plans
- [ ] **Change documentation**: Document all changes with complete rationale

### Validation Phase

- [ ] **Content review**: Ensure all sections remain complete and consistent
- [ ] **Integration check**: Verify all subsidiary plans work together coherently
- [ ] **Baseline validation**: Confirm updated baselines are realistic and achievable
- [ ] **Stakeholder validation**: Review changes with affected stakeholders

### Approval Phase

- [ ] **Change control board**: Present updates to change control board
- [ ] **Formal approval**: Obtain required approvals for plan updates
- [ ] **Baseline approval**: Get approval for any baseline changes
- [ ] **Communication**: Communicate approved changes to all stakeholders

### Documentation Phase

- [ ] **Version control**: Update version number and revision history
- [ ] **File update**: Save updated plan to `/docs/project-management/project-plan.md`
- [ ] **Distribution**: Distribute updated plan to stakeholders
- [ ] **Archive**: Archive previous version for historical reference

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Change Traceability**: Clearly document what changed, why, and impact analysis
- **Integration Maintenance**: Ensure all subsidiary plans remain integrated and consistent
- **Baseline Integrity**: Maintain realistic and achievable baselines
- **Stakeholder Alignment**: Ensure all stakeholders understand and approve changes

### LLM Error Prevention Checklist

- **Avoid**: Making changes without proper integrated change control
- **Avoid**: Updating plans without assessing impact on all knowledge areas
- **Avoid**: Breaking integration between subsidiary management plans
- **Avoid**: Implementing changes without proper approval
- **Reference**: Use existing plan structure and [project-plan-template.md](./.krci-ai/templates/project-plan-template.md)

### PMBoK Integration Context

Project plan updates follow PMBoK integrated change control by requiring formal change requests, comprehensive impact analysis across all knowledge areas, change control board approval, and integrated baseline updates. Updated plans maintain integration between all subsidiary management plans and continue to provide effective guidance for project execution and control. 