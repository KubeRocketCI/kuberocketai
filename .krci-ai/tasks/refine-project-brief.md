# Task: Refine Project Brief

## Description

Incorporate validation feedback and evidence into the project brief, updating sections with validated information, improved confidence levels, and refined assumptions. This task synthesizes multiple validation results to create an enhanced, evidence-based project brief.

This task uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and may utilize the [assumption tracker template](./.krci-ai/templates/assumption-tracker-template.md) for tracking refinements.

## Prerequisites

<prerequisites>
- Original project brief exists at `/docs/prd/project-brief.md`
- One or more validation reports completed
- Assumption tracker updated with validation results
- Stakeholder feedback on validation results available
</prerequisites>

### Reference Assets

Dependencies:

- /docs/prd/project-brief.md (source document)
- /docs/prd/brief-validation-*.md (validation reports)
- /docs/prd/brief-assumptions.md (assumption tracker)
- ./.krci-ai/templates/project-brief-template.md

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.

## Instructions

<instructions>
Incorporate validation feedback and evidence into the project brief, updating sections with validated information, improved confidence levels, and refined assumptions.

### Phase 1: Validation Results Integration

Collect and analyze all completed validation results:

#### Validation Report Analysis

1. **Review all validation reports** in `/docs/prd/brief-validation-*.md`
2. **Extract key findings** that impact project brief sections
3. **Identify confidence level changes** based on validation evidence
4. **Document required brief updates** based on validation results

#### Evidence Quality Assessment

1. **Aggregate evidence quality scores** across all validations
2. **Identify strongest and weakest evidence** areas
3. **Update overall brief confidence** based on evidence assessment
4. **Flag areas requiring additional validation**

### Phase 2: Project Brief Section Updates

#### Executive Summary Enhancement

- Update problem description with validated evidence
- Refine solution approach based on validation findings
- Update business value projections with validated metrics
- Enhance scope definition with validated constraints

#### Problem Statement Refinement

- Integrate root cause analysis findings
- Add quantified problem evidence from validation
- Update problem scope based on validation boundaries
- Enhance impact assessment with validated metrics

#### Target Users Enhancement

- Update user segments with Jobs-to-be-Done validation
- Add validated demographic and behavioral data
- Include opportunity scores and user priorities
- Refine user context with journey validation insights

#### Success Metrics Improvement

- Update metrics with SMART criteria validation results
- Add baseline data discovered during validation
- Refine targets based on evidence and benchmarks
- Balance leading and lagging indicators appropriately

#### Constraints and Risks Updates

- Add constraints discovered during validation
- Update risk assessment with validation insights
- Include evidence-based risk probability estimates
- Enhance mitigation strategies with validated approaches

### Phase 3: Confidence and Assumption Updates

#### Validation Checkpoint Updates

For each section, update validation checkpoints:
- Mark completed validations as verified
- Update confidence levels based on evidence quality
- Document validation methods used
- Include evidence sources and quality assessments

#### Assumption Tracker Integration

- Update assumption status based on validation results
- Add new assumptions discovered during validation
- Remove or modify disproven assumptions
- Prioritize remaining assumptions for future validation

### Phase 4: Quality Assurance and Finalization

#### Brief Quality Assessment

- Ensure 2-3 page length limit maintained
- Verify executive-ready language and structure
- Confirm all sections have appropriate evidence support
- Validate SDLC integration requirements met

#### Stakeholder Review Preparation

- Prepare summary of changes made based on validation
- Highlight areas of increased/decreased confidence
- Document remaining uncertainties and validation needs
- Create stakeholder presentation of refined brief
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

- [ ] **Collect validation reports**: Gather all completed validation documents
- [ ] **Extract key findings**: Identify findings that require brief updates
- [ ] **Assess evidence quality**: Evaluate strength of validation evidence
- [ ] **Plan section updates**: Determine which sections need refinement using [project brief template](./.krci-ai/templates/project-brief-template.md) structure

### Brief Refinement

- [ ] **Update executive summary**: Integrate high-level validation insights
- [ ] **Refine problem statement**: Add validated evidence and refined scope
- [ ] **Enhance user sections**: Include validated user research and insights
- [ ] **Improve success metrics**: Update with validated baselines and targets
- [ ] **Update constraints/risks**: Include validation-discovered factors

### Quality Assurance

- [ ] **Length verification**: Ensure brief remains within 2-3 page limit
- [ ] **Evidence documentation**: Verify all claims have supporting evidence cited
- [ ] **Assumption alignment**: Ensure brief content aligns with assumption tracker
- [ ] **Stakeholder readiness**: Confirm brief is executive-ready and decision-enabling

</execution_checklist>

## Content Guidelines

### Integration Principles

- **Evidence Primacy**: Prioritize validated evidence over original assumptions
- **Transparency**: Clearly indicate confidence levels and evidence sources
- **Conciseness**: Maintain executive brevity while including essential validation insights
- **Actionability**: Ensure refined brief enables clear next steps and decisions

### Quality Standards

- **Professional Presentation**: Maintain executive-level language and structure
- **Evidence Attribution**: Clearly cite validation sources and evidence quality
- **Balanced Perspective**: Include both confirming and challenging validation results
- **Future-Focused**: Identify areas needing continued validation and monitoring
