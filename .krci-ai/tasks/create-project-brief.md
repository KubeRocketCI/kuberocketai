# Task: Create Project Brief

## Description

Create a comprehensive project brief defining the foundation for product development by answering why, who, what success looks like, and what constraints shape the solution. This document serves as the **root artifact** in the SDLC framework that defines the essential foundation for all downstream artifacts, answers fundamental questions before solution development begins, and provides strategic context for PRD creation.

## Prerequisites

- [ ] Business opportunity or problem identified
- [ ] Initial stakeholder discussions completed
- [ ] Market context and user insights available
- [ ] Strategic goals and constraints understood

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/data/business-frameworks.md
- ./.krci-ai/templates/project-brief-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for role responsibilities and artifact flow
2. **Use business frameworks**: Apply methodologies from [business-frameworks.md](./.krci-ai/data/business-frameworks.md)
3. **Format output**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for consistent structure

## Output Format

- **Location**: `/docs/prd/project-brief.md` (EXACT path and filename)
- **Length**: 2-3 pages maximum for executive consumption
- **Downstream Enable**: Enables PRD creation at `/docs/prd/prd.md`

## Success Criteria

- [ ] **File saved** to `/docs/prd/project-brief.md`
- [ ] **Length** is 2-3 pages maximum
- [ ] **Problem** is specific and evidence-based
- [ ] **Users** are clearly defined with usage patterns
- [ ] **Success metrics** are specific and testable
- [ ] **Constraints** reflect actual limitations
- [ ] **Risks** identified with impact levels (HIGH/MEDIUM/LOW)

## Execution Checklist

### Discovery Phase

- [ ] **Stakeholder interviews**: Understand business context and strategic priorities
- [ ] **Problem validation**: Gather evidence that this problem is real and significant
- [ ] **User research**: Identify who has this problem and how it impacts them
- [ ] **Opportunity sizing**: Quantify business value and market opportunity

### Analysis Phase

- [ ] **Problem definition**: Write specific problem statement with evidence
- [ ] **User segmentation**: Define target users with demographics and usage patterns
- [ ] **Success planning**: Define measurable outcomes with realistic timelines
- [ ] **Constraint assessment**: Identify realistic limitations and assumptions

### Documentation Phase

- [ ] **Brief creation**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) structure
- [ ] **Content validation**: Ensure all required sections are completed
- [ ] **Length verification**: Confirm document is 2-3 pages maximum
- [ ] **File placement**: Save to exact location `/docs/prd/project-brief.md`

## Content Guidelines

### Quality Principles for LLM Self-Evaluation

- **Problem Focus**: Use concrete user scenarios and quantified evidence, not solution-oriented statements
- **User Specificity**: Define target users specifically enough to guide solution design decisions
- **Measurable Success**: Create specific, testable outcomes with realistic timelines and evidence
- **Evidence-Based**: Support all statements with data, research, and quantified metrics

### LLM Error Prevention Checklist

- **Avoid**: Solution-oriented problem statements (focus on user pain, not missing features)
- **Avoid**: Vague user descriptions without usage patterns and demographics
- **Avoid**: Unmeasurable success metrics or aspirational statements without evidence
- **Reference**: Use [project-brief-template.md](./.krci-ai/templates/project-brief-template.md) for all formatting guidance and examples

### SDLC Integration Context

This Project Brief enables immediate PRD creation by providing clear problem definition for PRD Problem/Opportunity section, target user clarity for PRD user research and requirements, success metrics for PRD Goals/Measurable Outcomes, and constraints for PRD MVP scope and technical requirements.
