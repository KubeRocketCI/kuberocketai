---
dependencies:
  data:
    - shared/validation-frameworks.md
  templates:
    - pm/project-brief-template.md
    - pm/assumption-tracker-template.md
---

# Task: Refine Project Brief

## Description

Incorporate validation feedback and evidence into the project brief, updating sections with validated information, improved confidence levels, and refined assumptions. This task synthesizes multiple validation results to create an enhanced, evidence-based project brief.

This task uses [validation frameworks](./.krci-ai/data/shared/validation-frameworks.md) and may utilize the [assumption tracker template](./.krci-ai/templates/pm/assumption-tracker-template.md) for tracking refinements.

## Instructions

<instructions>
Confirm the original project brief exists at `/docs/prd/project-brief.md`, one or more validation reports are completed, assumption tracker is updated with validation results, and stakeholder feedback on validation results is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Collect and analyze all completed validation results from `/docs/prd/brief-validation-*.md`. Review validation reports to extract key findings that impact project brief sections. Identify confidence level changes based on validation evidence and document required brief updates. Aggregate evidence quality scores across all validations, identify strongest and weakest evidence areas, update overall brief confidence based on evidence assessment, and flag areas requiring additional validation.

Update project brief sections systematically with validation results. Enhance executive summary by updating problem description with validated evidence, refining solution approach based on validation findings, updating business value projections with validated metrics, and enhancing scope definition with validated constraints. Refine problem statement by integrating root cause analysis findings, adding quantified problem evidence from validation, updating problem scope based on validation boundaries, and enhancing impact assessment with validated metrics.

Enhance target users section by updating user segments with Jobs-to-be-Done validation, adding validated demographic and behavioral data, including opportunity scores and user priorities, and refining user context with journey validation insights. Improve success metrics by updating metrics with SMART criteria validation results, adding baseline data discovered during validation, refining targets based on evidence and benchmarks, and balancing leading and lagging indicators appropriately.

Update constraints and risks by adding constraints discovered during validation, updating risk assessment with validation insights, including evidence-based risk probability estimates, and enhancing mitigation strategies with validated approaches. For each section, update validation checkpoints by marking completed validations as verified, updating confidence levels based on evidence quality, documenting validation methods used, and including evidence sources and quality assessments.

Integrate assumption tracker updates by updating assumption status based on validation results, adding new assumptions discovered during validation, removing or modifying disproven assumptions, and prioritizing remaining assumptions for future validation. Ensure 2-3 page length limit maintained, verify executive-ready language and structure, confirm all sections have appropriate evidence support, and validate SDLC integration requirements met.

Prepare stakeholder review by creating summary of changes made based on validation, highlighting areas of increased or decreased confidence, documenting remaining uncertainties and validation needs, and creating stakeholder presentation of refined brief.
</instructions>

## Output Format

<output_format>
- Primary Output: Updated `/docs/prd/project-brief.md` with validation integration
- Secondary Outputs:
  - `/docs/prd/brief-refinement-summary.md` documenting changes made
  - Updated `/docs/prd/brief-assumptions.md` with current status
  - `/docs/prd/validation-dashboard.md` with overall validation status
- Length: Project brief remains 2-3 pages maximum
</output_format>

## Success Criteria

<success_criteria>
- All validation results integrated into appropriate brief sections
- Confidence levels updated based on evidence quality assessment
- Assumption tracker synchronized with brief content and validation results
- Brief length maintained within 2-3 page executive limit
- Evidence sources documented with quality assessments
- Remaining validation needs identified and prioritized
</success_criteria>

## Execution Checklist

<execution_checklist>

### Validation Integration

- Collect validation reports: Gather all completed validation documents
- Extract key findings: Identify findings that require brief updates
- Assess evidence quality: Evaluate strength of validation evidence
- Plan section updates: Determine which sections need refinement using [project brief template](./.krci-ai/templates/pm/project-brief-template.md) structure

### Brief Refinement

- Update executive summary: Integrate high-level validation insights
- Refine problem statement: Add validated evidence and refined scope
- Enhance user sections: Include validated user research and insights
- Improve success metrics: Update with validated baselines and targets
- Update constraints/risks: Include validation-discovered factors

### Quality Assurance

- Length verification: Ensure brief remains within 2-3 page limit
- Evidence documentation: Verify all claims have supporting evidence cited
- Assumption alignment: Ensure brief content aligns with assumption tracker
- Stakeholder readiness: Confirm brief is executive-ready and decision-enabling

</execution_checklist>

## Content Guidelines

### Integration Principles

- Evidence Primacy: Prioritize validated evidence over original assumptions
- Transparency: Clearly indicate confidence levels and evidence sources
- Conciseness: Maintain executive brevity while including essential validation insights
- Actionability: Ensure refined brief enables clear next steps and decisions

### Quality Standards

- Professional Presentation: Maintain executive-level language and structure
- Evidence Attribution: Clearly cite validation sources and evidence quality
- Balanced Perspective: Include both confirming and challenging validation results
- Future-Focused: Identify areas needing continued validation and monitoring
