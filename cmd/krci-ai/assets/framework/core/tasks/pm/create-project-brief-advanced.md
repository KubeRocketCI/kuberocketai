---
dependencies:
  data:
    - shared/validation-frameworks.md
    - shared/business-frameworks.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - pm/project-brief-template-advanced.md
    - pm/context-gathering-guide-template.md
    - pm/assumption-tracker-template.md
    - pm/validation-report-template.md
  tasks:
    - pm/gather-project-context.md
    - pm/validate-problem-statement.md
    - pm/validate-target-users.md
    - pm/validate-success-metrics.md
    - pm/validate-business-value.md
    - pm/refine-project-brief.md
    - pm/finalize-project-brief.md
    - pm/enhance-project-brief.md
---

# Task: Create Project Brief (Advanced)

## Description

Create a comprehensive project brief using advanced validation workflow with business framework methodology, evidence collection, and stakeholder validation. This advanced version extends the standard project brief creation with systematic validation, assumption tracking, and quality assurance for high-stakes projects requiring executive approval or comprehensive stakeholder buy-in.

This task leverages [validation frameworks](./.krci-ai/data/shared/validation-frameworks.md) and the [advanced project brief template](./.krci-ai/templates/pm/project-brief-template-advanced.md) for comprehensive project analysis.

Use Advanced Flow When:
- High-stakes projects (>$100K budget, >6 month timeline)
- Strategic initiatives requiring executive approval
- Market uncertainty or unvalidated assumptions
- Stakeholder validation and evidence-based decision making required
- Competitive or complex market environment

## Instructions

<instructions>
Confirm this advanced validation flow is justified (high-stakes project >$100K, executive approval needed, market uncertainty, or stakeholder validation required). Ensure 2-4 weeks are available for the multi-session validation process, stakeholders are accessible for interviews, and dependencies declared in the YAML frontmatter are readable before proceeding.

Execute comprehensive context gathering using business frameworks by running [gather-project-context](./.krci-ai/tasks/pm/gather-project-context.md) task for systematic evidence collection. Conduct structured stakeholder interviews using the [context gathering guide](./.krci-ai/templates/pm/context-gathering-guide-template.md) and build comprehensive evidence base with confidence assessments. Create detailed assumption tracking using the [assumption tracker template](./.krci-ai/templates/pm/assumption-tracker-template.md) and output results to `/docs/prd/project-context.md` with supporting evidence and assumption documentation.

Create initial brief using advanced template with validation checkpoints including built-in validation status tracking for each section, confidence levels with evidence source documentation, systematic assumption identification and risk assessment, methodology citations with validation approach, and professional standards with executive-ready formatting. Output initial `/docs/prd/project-brief.md` with validation placeholders and checkpoint system.

Apply systematic validation using established business methodologies by executing [validate-problem-statement](./.krci-ai/tasks/pm/validate-problem-statement.md) using Lean Startup Problem-Solution Fit, then execute [validate-target-users](./.krci-ai/tasks/pm/validate-target-users.md) using Jobs-to-be-Done framework, followed by [validate-success-metrics](./.krci-ai/tasks/pm/validate-success-metrics.md) using SMART criteria and OKR alignment, and complete with [validate-business-value](./.krci-ai/tasks/pm/validate-business-value.md) using Value Proposition Canvas. Output validation reports to `/docs/prd/brief-validation-*.md` using the [validation report template](./.krci-ai/templates/pm/validation-report-template.md) with evidence and confidence assessments.

Integrate validation results into enhanced project brief by running [refine-project-brief](./.krci-ai/tasks/pm/refine-project-brief.md) to integrate validation findings. Update section confidence based on validation evidence quality, modify assumption tracker with validation results, and incorporate validated evidence into brief sections. Output enhanced `/docs/prd/project-brief.md` with validated evidence and reduced assumptions.

Complete quality gates and stakeholder approval for production readiness by running [finalize-project-brief](./.krci-ai/tasks/pm/finalize-project-brief.md) for comprehensive quality review. Obtain sign-off from key decision makers, confirm all advanced flow standards are met, and verify enablement for downstream PRD development. Output production-ready project brief with comprehensive validation documentation.
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

Standard project briefs can be upgraded to advanced using [enhance-project-brief](./.krci-ai/tasks/pm/enhance-project-brief.md) command when project importance or requirements change.

## Framework Integration Notes

- SDLC Integration: Advanced brief provides stronger foundation for complex PRD development
- Business Framework Usage: Leverages established methodologies for professional validation
- Evidence Standards: Maintains quantified, multi-source validation approach
- Quality Assurance: Built-in validation ensures enterprise-grade output
- Professional Output: Investment-grade documentation suitable for executive and board presentation
