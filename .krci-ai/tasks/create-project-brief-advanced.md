---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - project-brief-template-advanced.md
    - context-gathering-guide-template.md
    - assumption-tracker-template.md
    - validation-report-template.md
  tasks:
    - gather-project-context.md
    - validate-problem-statement.md
    - validate-target-users.md
    - validate-success-metrics.md
    - validate-business-value.md
    - refine-project-brief.md
    - finalize-project-brief.md
    - enhance-project-brief.md
---

# Task: Create Project Brief (Advanced)

## Description

Create a comprehensive project brief using advanced validation workflow with business framework methodology, evidence collection, and stakeholder validation. This advanced version extends the standard project brief creation with systematic validation, assumption tracking, and quality assurance for high-stakes projects requiring executive approval or comprehensive stakeholder buy-in.

This task leverages [validation frameworks](./.krci-ai/data/validation-frameworks.md) and the [advanced project brief template](./.krci-ai/templates/project-brief-template-advanced.md) for comprehensive project analysis.

**Use Advanced Flow When:**
- High-stakes projects (>$100K budget, >6 month timeline)
- Strategic initiatives requiring executive approval
- Market uncertainty or unvalidated assumptions
- Stakeholder validation and evidence-based decision making required
- Competitive or complex market environment

## Prerequisites

<prerequisites>
- Project requires comprehensive validation and evidence-based decision making
- Time available for multi-session validation process (2-4 weeks)
- Access to stakeholders for interviews and validation
- Budget/timeline justifies enhanced validation effort

### Reference Assets

Dependencies:

- ./.krci-ai/data/validation-frameworks.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/data/krci-ai/core-sdlc-framework.md
- ./.krci-ai/templates/project-brief-template-advanced.md
- ./.krci-ai/tasks/gather-project-context.md

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.
</prerequisites>

## Instructions

<instructions>
### Phase 1: Context Gathering (Advanced)

Execute comprehensive context gathering using business frameworks:

#### Context Collection Process

1. Execute Context Gathering: Run [gather-project-context](./.krci-ai/tasks/gather-project-context.md) task for systematic evidence collection
2. Stakeholder Interview Program: Conduct structured interviews using the [context gathering guide](./.krci-ai/templates/context-gathering-guide-template.md)
3. Evidence Library Creation: Build comprehensive evidence base with confidence assessments
4. Assumption Inventory Development: Create detailed assumption tracking using the [assumption tracker template](./.krci-ai/templates/assumption-tracker-template.md)

Output: `/docs/prd/project-context.md` with supporting evidence and assumption documentation

### Phase 2: Advanced Brief Creation

Create initial brief using advanced template with validation checkpoints:

#### Advanced Template Features

- Validation Checkpoints: Built-in validation status tracking for each section
- Evidence Integration: Confidence levels and evidence source documentation
- Assumption Tracking: Systematic assumption identification and risk assessment
- Business Framework References: Methodology citations and validation approach
- Quality Assurance: Professional standards with executive-ready formatting

Output: Initial `/docs/prd/project-brief.md` with validation placeholders and checkpoint system

### Phase 3: Business Framework Validation Cycle

Apply systematic validation using established business methodologies:

#### Validation Sequence

1. Problem Validation: Execute [validate-problem-statement](./.krci-ai/tasks/validate-problem-statement.md) using Lean Startup Problem-Solution Fit
2. User Validation: Execute [validate-target-users](./.krci-ai/tasks/validate-target-users.md) using Jobs-to-be-Done framework
3. Metrics Validation: Execute [validate-success-metrics](./.krci-ai/tasks/validate-success-metrics.md) using SMART criteria and OKR alignment
4. Value Validation: Execute [validate-business-value](./.krci-ai/tasks/validate-business-value.md) using Value Proposition Canvas

Outputs: Validation reports at `/docs/prd/brief-validation-*.md` using the [validation report template](./.krci-ai/templates/validation-report-template.md) with evidence and confidence assessments

### Phase 4: Brief Refinement and Integration

Integrate validation results into enhanced project brief:

#### Refinement Process

1. Execute Brief Refinement: Run [refine-project-brief](./.krci-ai/tasks/refine-project-brief.md) to integrate validation findings
2. Confidence Level Updates: Update section confidence based on validation evidence quality
3. Assumption Status Updates: Modify assumption tracker with validation results
4. Evidence Integration: Incorporate validated evidence into brief sections

Output: Enhanced `/docs/prd/project-brief.md` with validated evidence and reduced assumptions

### Phase 5: Quality Assurance and Finalization

Complete quality gates and stakeholder approval for production readiness:

#### Finalization Process

1. Execute Finalization: Run [finalize-project-brief](./.krci-ai/tasks/finalize-project-brief.md) for comprehensive quality review
2. Stakeholder Approval: Obtain sign-off from key decision makers
3. Quality Gate Verification: Confirm all advanced flow standards met
4. SDLC Readiness: Verify enablement for downstream PRD development

Output: Production-ready project brief with comprehensive validation documentation
</instructions>

## Output Format

<output_format>
- Primary Output: `/docs/prd/project-brief.md` (same location as standard flow)
- Supporting Outputs:
  - `/docs/prd/project-context.md` (context gathering results)
  - `/docs/prd/brief-assumptions.md` (assumption tracking)
  - `/docs/prd/brief-validation-*.md` (validation reports)
  - `/docs/prd/brief-evidence.md` (evidence library)
- Length: 2-3 pages for main brief + supporting validation documentation
- Quality: Enterprise-grade with evidence-based validation
</output_format>

## Success Criteria

<success_criteria>

### Context Foundation

- Project context gathered using business framework methodologies
- Stakeholder interviews completed with structured approach
- Evidence library created with quality and confidence assessments
- Assumption inventory developed with impact and risk prioritization

### Validation Completion

- Problem validated using Lean Startup Problem-Solution Fit methodology
- Users validated using Jobs-to-be-Done framework with evidence
- Metrics validated using SMART criteria and OKR alignment assessment
- Value validated using Value Proposition Canvas and ROI analysis

### Quality Assurance

- Brief enhanced with validation evidence and confidence levels
- Assumptions tracked with validation status and evidence sources
- Stakeholder approval obtained from key decision makers
- SDLC integration verified for downstream artifact creation

### Professional Standards

- Executive readiness confirmed with 2-3 page length and professional format
- Evidence documentation complete with source attribution and quality assessment
- Risk mitigation addressed through assumption validation and stakeholder alignment
- Investment grade analysis suitable for budget approval and resource allocation
</success_criteria>

## Execution Checklist

<execution_checklist>

### Advanced Flow Preparation

- Flow justification: Confirm project meets criteria for advanced validation
- Resource planning: Allocate 2-4 weeks for comprehensive validation process
- Stakeholder coordination: Arrange access to key stakeholders and subject matter experts
- Evidence preparation: Identify available research, data, and documentation

### Context and Validation Execution

- Context gathering: Execute systematic context collection using business frameworks
- Brief creation: Generate initial brief with advanced template and validation checkpoints
- Validation cycle: Complete business framework validation for high-priority areas
- Evidence integration: Incorporate validation findings into enhanced brief

### Quality Assurance and Finalization

- Refinement completion: Integrate all validation results and evidence
- Quality verification: Confirm advanced flow standards and professional presentation
- Stakeholder approval: Obtain sign-off from key decision makers and stakeholders
- Documentation package: Complete all supporting documentation for handoff
</execution_checklist>

## Advanced Flow vs Standard Flow

### When to Use Standard Flow (`create-project-brief`)

- Low-to-medium risk projects
- Internal projects with known stakeholders
- Tight timelines requiring rapid project initiation
- Prototypes or experimental projects
- Limited budget for extensive validation

### When to Use Advanced Flow (`create-project-brief-advanced`)

- High-stakes or strategic projects
- External stakeholder validation required
- Market uncertainty or unvalidated assumptions
- Executive approval and evidence-based decision making needed
- Budget and timeline support comprehensive validation

### Upgrade Path

Standard project briefs can be upgraded to advanced using [enhance-project-brief](./.krci-ai/tasks/enhance-project-brief.md) command when project importance or requirements change.

## Framework Integration Notes

- **SDLC Integration**: Advanced brief provides stronger foundation for complex PRD development
- **Business Framework Usage**: Leverages established methodologies for professional validation
- **Evidence Standards**: Maintains quantified, multi-source validation approach
- **Quality Assurance**: Built-in validation ensures enterprise-grade output
- **Professional Output**: Investment-grade documentation suitable for executive and board presentation
