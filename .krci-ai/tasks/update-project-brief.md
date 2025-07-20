# Task: Update Project Brief

## Objective

Update existing Project Brief with new strategic information while maintaining quality, stakeholder alignment, and SDLC framework compliance. Focus on change impact assessment and downstream artifact management.

## Prerequisites

- [ ] **Existing Project Brief**: `/docs/prd/project-brief.md` must exist and be properly registered
- [ ] **Change Trigger**: Clear reason for update (strategic shifts, market changes, new insights, stakeholder feedback, resource changes)
- [ ] **Impact Assessment**: Understanding of how changes affect dependent PRD and downstream artifacts

## Update Process

### Strategic Assessment

1. **Identify Change Scope**: What sections need updating? (Executive Summary, Problem, Opportunity, Users, Success Metrics, Constraints, Risks)
2. **Business Impact**: How do changes affect product strategy and business case?
3. **Downstream Impact**: How do changes affect existing PRD (`/docs/prd/prd.md`) and subsequent artifacts?

### Update Execution

4. **Section Updates**: Modify specific sections using [project-brief-template.md](/.krci-ai/templates/project-brief-template.md) structure
5. **Strategic Alignment**: Ensure updates maintain strategic coherence and business focus
6. **Quality Check**: Ensure updated Project Brief maintains 2-3 page limit and foundation quality

### Change Management

7. **PRD Impact Analysis**: Determine if PRD needs updating based on Project Brief changes
8. **Stakeholder Communication**: Notify key stakeholders of strategic changes and implications
9. **Registry Update**: Update `/docs/registry.json` with new version information

## SDLC Framework Compliance

### File Structure

- **Target File**: Update existing `/docs/prd/project-brief.md` (EXACT path)
- **Template**: Use same [project-brief-template.md](/.krci-ai/templates/project-brief-template.md) structure
- **Registry**: Update artifact version in `/docs/registry.json`

### Downstream Impact Management

- **PRD Updates**: Assess if PRD sections need updating to maintain alignment
- **Epic Impact**: Determine if existing Epics (`{epic_number}-epic-{slug}.md`) are affected
- **Strategic Cascade**: Manage strategic changes flowing through entire SDLC chain

## Quality Validation

### Pre-Update Checklist

- [ ] **Change Justification**: Clear strategic reason for update documented
- [ ] **Scope Defined**: Specific sections and changes identified with business rationale
- [ ] **Impact Assessed**: Downstream effects on PRD and subsequent artifacts evaluated
- [ ] **Stakeholders Informed**: Key business stakeholders aware of planned strategic changes

### Post-Update Checklist

- [ ] **Template Compliance**: Updated sections follow project-brief-template.md structure
- [ ] **Length Maintained**: Document remains 2-3 pages maximum
- [ ] **Strategic Coherence**: All sections maintain strategic alignment and business focus
- [ ] **File Location**: Changes saved to `/docs/prd/project-brief.md`
- [ ] **Registry Updated**: New version registered in `/docs/registry.json`
- [ ] **Downstream Assessment**: PRD alignment evaluated and update plan created if needed

## Common Update Scenarios

### Strategic Pivots

- Update Problem Statement with new market insights
- Modify Opportunity with revised business case
- Adjust Success Metrics based on new strategic priorities
- Revise Target Users based on market research

### Resource Changes

- Update Constraints with new budget, timeline, or team changes
- Modify Success Metrics to reflect realistic resource availability
- Adjust Key Risks based on new resource constraints

### Market Intelligence

- Update Problem Statement with competitive intelligence
- Modify Opportunity with new market size or timing insights
- Adjust Target Users with refined segmentation
- Update Key Risks with new market or technology threats

### Business Context Changes

- Update Executive Summary with new business priorities
- Modify Success Metrics with revised business objectives
- Adjust Constraints with new compliance or regulatory requirements

## LLM Agent Guidance

### Strategic Change Questions

- "What specific strategic changes are driving this update?"
- "How do these changes affect our core business case and value proposition?"
- "Which stakeholders need to approve these strategic changes?"
- "How do these changes impact the existing PRD and development roadmap?"

### Update Best Practices

- Maintain strategic foundation role for all downstream artifacts
- Focus on strategic changes rather than tactical adjustments
- Preserve business case coherence and value proposition clarity
- Ensure changes strengthen rather than weaken the overall foundation
- Consider long-term strategic implications beyond immediate changes

### SDLC Impact Management

- **PRD Cascade**: Determine which PRD sections need updating to maintain alignment
- **Epic Review**: Assess if strategic changes affect Epic prioritization or scope
- **Stakeholder Communication**: Ensure strategic changes are communicated across teams
- **Quality Gate**: Updated Project Brief approval required before any dependent artifact changes

### SDLC Framework Compliance

- **Mandatory File**: Update existing `/docs/prd/project-brief.md` only
- **Template Consistency**: Maintain same project-brief-template.md structure
- **Version Control**: Proper change documentation and registry updates
- **Strategic Gate**: Updated Project Brief approval required before PRD or Epic changes
- **Foundation Role**: Ensure updated brief continues to enable strong PRD creation

**Remember**: Project Brief updates affect the entire SDLC chain. Focus on strategic foundation changes that strengthen the business case and provide better guidance for all downstream artifacts.
