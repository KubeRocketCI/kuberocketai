# Task: Update Product Requirements Document (PRD)

## Objective

Update existing PRD with new information while maintaining quality, stakeholder alignment, and SDLC framework compliance. Focus on change impact assessment and downstream communication.

## Prerequisites

- [ ] **Existing PRD**: `/docs/prd/prd.md` must exist and be properly accessible
- [ ] **Change Trigger**: Clear reason for update (Project Brief changes, user research, business priorities, technical constraints, stakeholder feedback)
- [ ] **Stakeholder Input**: Understanding of what specifically needs to change and why

## Update Process

### Change Assessment

1. **Identify Change Scope**: What sections need updating? (Problem, Users, Solution, Goals, Requirements)
2. **Impact Analysis**: How do changes affect existing Epics (`/docs/epics/`) and Stories (`/docs/stories/`)?
3. **Stakeholder Review**: Who needs to approve these changes before implementation?

### Update Execution

4. **Section Updates**: Modify specific PRD sections using [prd-template.md](/.krci-ai/templates/prd-template.md) structure
5. **Version Control**: Document changes in version history and maintain change log
6. **Quality Check**: Ensure updated PRD maintains 6-8 page limit and best practices compliance

### Change Communication

7. **Downstream Notification**: Communicate changes to Epic owners and development teams
8. **Approval Process**: Obtain stakeholder sign-off on updated PRD

## SDLC Framework Compliance

### File Structure

- **Target File**: Update existing `/docs/prd/prd.md` (EXACT path)
- **Template**: Use same [prd-template.md](/.krci-ai/templates/prd-template.md) structure

### Downstream Impact Management

- **Epic Updates**: Identify which Epics (`{epic_number}-epic-{slug}.md`) need updating
- **Story Impact**: Assess if in-progress Stories (`{epic_number}.{story_number}.story.md`) are affected
- **Architecture Review**: Determine if changes require Architecture document updates

## Quality Validation

### Pre-Update Checklist

- [ ] **Change Justification**: Clear reason for update documented
- [ ] **Scope Defined**: Specific sections and changes identified
- [ ] **Impact Assessed**: Downstream effects on Epics/Stories evaluated
- [ ] **Stakeholders Informed**: Key stakeholders aware of planned changes

### Post-Update Checklist

- [ ] **Template Compliance**: Updated sections follow prd-template.md structure
- [ ] **Length Maintained**: Document remains 6-8 pages maximum
- [ ] **Best Practices**: User focus, P0/P1/P2 priorities, and measurable goals maintained
- [ ] **File Location**: Changes saved to `/docs/prd/prd.md`
- [ ] **Change Communication**: Downstream teams notified of changes

## Common Update Scenarios

### Business Priority Changes

- Update Goals/Measurable Outcomes (Section 5)
- Adjust P0/P1/P2 priorities in Requirements (Section 6)
- Modify timeline and implementation sequence

### New User Research

- Refine Target Users & Use Cases (Section 2)
- Update Problem/Opportunity with new evidence (Section 1)
- Adjust functional requirements based on user feedback

### Technical Constraint Discovery

- Update Solution approach (Section 4)
- Modify functional requirements for technical feasibility
- Adjust MVP scope and priorities

### Scope Changes

- Update Problem/Opportunity scope boundaries
- Modify functional requirements and priorities
- Adjust success metrics and goals

## LLM Agent Guidance

### Change Impact Questions

- "What specific changes are needed and why?"
- "Which existing Epics and Stories are affected by these changes?"
- "How do these changes affect our success metrics and timeline?"
- "Who are the key stakeholders that need to approve these changes?"

### Update Best Practices

- Maintain original PRD structure and quality standards
- Focus on specific changes rather than comprehensive rewrite
- Preserve user focus and business value orientation
- Keep Epic enablement and SDLC framework compliance
- Document all changes with clear rationale

### SDLC Framework Compliance

- **Mandatory File**: Update existing `/docs/prd/prd.md` only
- **Template Consistency**: Maintain same prd-template.md structure
- **Version Control**: Proper change documentation
- **Quality Gate**: Updated PRD approval required before Epic/Story changes
- **Communication**: Notify downstream teams of changes affecting their work

**Remember**: PRD updates should be focused, well-justified, and minimize disruption to ongoing development while maintaining alignment on what to build and why.
