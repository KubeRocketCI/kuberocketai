---
dependencies:
  data:
    - shared/validation-frameworks.md
  templates:
    - pm/project-brief-template-advanced.md
  tasks:
    - pm/gather-project-context.md
    - pm/validate-problem-statement.md
    - pm/validate-target-users.md
    - pm/validate-success-metrics.md
    - pm/validate-business-value.md
    - pm/refine-project-brief.md
---

# Task: Enhance Project Brief

## Description

Upgrade an existing standard project brief to the advanced validation flow, adding business framework validation, evidence collection, and assumption tracking. This task bridges standard rapid creation with comprehensive validation when project importance or risk increases.

This task uses the [advanced project brief template](./.krci-ai/templates/pm/project-brief-template-advanced.md) and [validation frameworks](./.krci-ai/data/pm/validation-frameworks.md) for comprehensive enhancement.

## Instructions

<instructions>
Confirm the existing project brief at `/docs/prd/project-brief.md` is accessible and created with standard flow. Verify the project has increased in strategic importance, budget, or risk level, stakeholders are requesting evidence-based validation, and time is available for comprehensive validation process (2-4 weeks). Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Follow this systematic enhancement process to upgrade the brief to advanced validation flow with business framework validation, evidence collection, and assumption tracking:

### Phase 1: Enhancement Assessment

Evaluate whether enhancement is appropriate and beneficial:

#### Enhancement Triggers Assessment

Strategic Importance:
- Project budget increased above $100K
- Timeline extended beyond 6 months
- Executive stakeholder interest increased
- Competitive or market pressure identified

Risk Level Changes:
- Key assumptions challenged by stakeholders
- Market uncertainty or user segment questions emerged
- Technical feasibility concerns raised
- Business case questioned or ROI disputed

Stakeholder Requirements:
- Evidence-based validation requested
- Detailed business case required for approval
- Comprehensive user research needed
- Risk assessment and mitigation planning required

### Phase 2: Simple Brief Analysis

Extract and analyze existing simple brief content:

#### Content Extraction

1. Existing Problem Statement: Extract current problem definition
2. Current User Definition: Identify stated target users
3. Stated Success Metrics: Capture existing success criteria
4. Documented Constraints: List known limitations and assumptions
5. Identified Risks: Extract mentioned risks and concerns

#### Gap Analysis

```markdown
Simple Brief Analysis: {{project_name}}

Strengths of Current Brief:
- {{strength_1}}
- {{strength_2}}
- {{strength_3}}

Areas Requiring Validation:
- {{validation_need_1}} (Evidence gap: {{evidence_gap_1}})
- {{validation_need_2}} (Evidence gap: {{evidence_gap_2}})
- {{validation_need_3}} (Evidence gap: {{evidence_gap_3}})

High-Risk Assumptions Identified:
- {{assumption_1}} (Risk level: {{risk_level_1}})
- {{assumption_2}} (Risk level: {{risk_level_2}})

Enhancement Priority Areas:
1. {{priority_area_1}} - {{rationale_1}}
2. {{priority_area_2}} - {{rationale_2}}
3. {{priority_area_3}} - {{rationale_3}}
```

### Phase 3: Enhanced Template Migration

Migrate content to enhanced template with validation checkpoints:

#### Template Enhancement Process

1. Backup Simple Brief: Create backup of original simple brief
2. Load Enhanced Template: Apply enhanced template with validation checkpoints
3. Migrate Content: Transfer existing content to enhanced structure
4. Add Validation Placeholders: Insert validation checkpoints and assumption tracking
5. Mark for Enhancement: Flag brief as enhanced flow in progress

#### Enhanced Structure Addition

```markdown
Enhanced Elements Added:
- Validation checkpoints for each section
- Assumption tracking with confidence levels
- Evidence quality assessment placeholders
- Business framework methodology references
- Stakeholder validation tracking
```

### Phase 4: Prioritized Validation Planning

Develop validation plan based on highest-risk and highest-impact areas:

#### Validation Prioritization

High Priority (Immediate validation required):
- Problem validation if core assumptions challenged
- User validation if market uncertainty exists
- Value validation if ROI questioned
- Metrics validation if success criteria disputed

Medium Priority (Planned validation):
- Areas with moderate evidence gaps
- Secondary assumptions with medium impact
- Competitive positioning questions
- Technical feasibility concerns

Low Priority (Monitor and validate if needed):
- Well-supported assumptions
- Low-impact variables
- Stable market conditions
- Proven technical approaches

#### Validation Sequence Planning

```markdown
Validation Execution Plan:

Week 1-2: Context Gathering
- Execute gather-context task
- Stakeholder interviews
- Evidence library creation
- Assumption inventory completion

Week 3-4: High-Priority Validations
- {{high_priority_validation_1}}
- {{high_priority_validation_2}}
- Brief refinement with results

Week 5-6: Medium-Priority Validations (if needed)
- {{medium_priority_validation_1}}
- {{medium_priority_validation_2}}
- Final brief enhancement

Week 7: Finalization
- Quality gate verification
- Stakeholder approval
- Enhanced brief completion
```

### Phase 5: Enhancement Execution

Execute the planned enhancement with validation:

#### Context Gathering Execution

- Run [gather-project-context](./.krci-ai/tasks/pm/gather-project-context.md) task to build evidence foundation
- Conduct stakeholder interviews for gap-filling
- Create comprehensive assumption inventory
- Establish evidence library with confidence assessments

#### Selective Validation Execution

Based on priority assessment, execute relevant validation tasks:
- [validate-problem-statement](./.krci-ai/tasks/pm/validate-problem-statement.md) for problem statement validation
- [validate-target-users](./.krci-ai/tasks/pm/validate-target-users.md) for user segment validation
- [validate-success-metrics](./.krci-ai/tasks/pm/validate-success-metrics.md) for success criteria validation
- [validate-business-value](./.krci-ai/tasks/pm/validate-business-value.md) for business case validation

#### Brief Integration and Refinement

- Run [refine-project-brief](./.krci-ai/tasks/pm/refine-project-brief.md) to integrate validation results
- Update confidence levels based on evidence quality
- Enhance assumptions tracking with validation outcomes
- Prepare enhanced brief for finalization
</instructions>

## Output Format

- Primary Output: Enhanced `/docs/prd/project-brief.md` with validation integration
- Supporting Outputs:
  - `/docs/prd/project-context.md` (if context gathering executed)
  - `/docs/prd/brief-assumptions.md` (enhanced assumption tracking)
  - `/docs/prd/brief-validation-*.md` (validation reports as executed)
  - `/docs/prd/enhancement-summary.md` (summary of changes made)
- Length: 2-3 pages enhanced brief (expanded from 1-2 page simple)

## Success Criteria

<success_criteria>

### Enhancement Planning

- Enhancement justification documented with clear triggers and rationale
- Gap analysis completed identifying areas requiring validation
- Validation plan created with prioritized sequence and timeline
- Resource requirements assessed for validation execution

### Template Migration

- Simple brief content preserved with backup created
- Enhanced template applied with validation checkpoints added
- Assumption tracking integrated with confidence level placeholders
- Flow marker updated from SIMPLE_FLOW to ENHANCED_FLOW

### Validation Execution

- High-priority validations completed based on risk and impact assessment
- Evidence integrated with appropriate confidence levels
- Assumption status updated based on validation results
- Brief quality enhanced with validated evidence and reduced assumptions

### Finalization

- Enhanced brief completed meeting enhanced flow standards
- Stakeholder approval obtained for enhanced version
- Enhancement documentation complete for future reference
- SDLC integration verified for downstream PRD creation
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Enhancement Assessment

- Enhancement triggers evaluated: Verify justification for enhancement
- Current brief analyzed: Extract content and identify gaps
- Stakeholder input gathered: Confirm enhancement requirements
- Resource planning: Assess time and effort required for enhancement

### Enhancement Planning

- Gap analysis conducted: Identify specific areas needing validation
- Validation prioritization: Rank validations by impact and urgency
- Timeline creation: Develop realistic schedule for enhancement process
- Success criteria definition: Define what enhanced brief should achieve

### Migration Execution

- Content backup: Preserve original simple brief
- Template application: Apply enhanced template structure
- Content migration: Transfer existing content to enhanced format
- Validation setup: Add checkpoints and assumption tracking

### Validation Execution

- Context gathering: Execute if comprehensive foundation needed
- Priority validations: Complete high-priority validation tasks
- Evidence integration: Incorporate validation results into brief
- Quality assessment: Evaluate enhanced brief against standards

### Enhancement Completion

- Final integration: Complete refine-project-brief process
- Quality verification: Ensure enhanced brief meets all standards
- Stakeholder review: Obtain approval for enhanced version
- Documentation: Complete enhancement summary and lessons learned
</execution_checklist>

## Content Guidelines

### Enhancement Quality Standards

- Value-Added: Enhancement should significantly improve brief quality and confidence
- Evidence-Based: All enhancements supported by validated evidence or analysis
- Stakeholder-Focused: Enhancement addresses specific stakeholder needs or concerns
- Risk-Mitigation: Enhanced brief reduces project risk through better understanding
- Professional Grade: Enhanced brief suitable for high-stakes decision making

### Validation Selection Standards

- Risk-Based Prioritization: Focus validation effort on highest-risk assumptions
- Impact-Driven: Validate areas with highest potential impact on project success
- Evidence-Gap Closure: Target validations that close most significant evidence gaps
- Stakeholder-Requested: Include validations specifically requested by stakeholders
- Resource-Appropriate: Balance validation depth with available time and resources

### Documentation Standards

- Enhancement Rationale: Clearly document why enhancement was pursued
- Change Tracking: Document all changes from simple to enhanced brief
- Evidence Integration: Show how validation evidence improved brief quality
- Future Reference: Document process and learnings for future enhancements
- Stakeholder Communication: Prepare clear summary of enhancements for stakeholders

## Framework Integration Notes

- SDLC Integration: Enhanced brief provides stronger foundation for downstream artifacts
- Flow Flexibility: Supports both planned enhancement and reactive enhancement triggers
- Quality Assurance: Built-in assessment ensures enhancement adds genuine value
- Resource Management: Prioritized approach manages validation effort efficiently
- Professional Standards: Enhanced output meets enterprise decision-making requirements
