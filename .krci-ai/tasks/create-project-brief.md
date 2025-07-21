# Task: Create Project Brief

## Description

Create a comprehensive project brief defining the foundation for product development by answering why, who, what success looks like, and what constraints shape the solution. This document serves as the **root artifact** in the SDLC framework that defines the essential foundation for all downstream artifacts, answers fundamental questions before solution development begins, and provides strategic context for PRD creation.

## Framework Context

**Reference**: [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) - See role responsibilities and artifact flow

## Prerequisites

- [ ] Business opportunity or problem identified
- [ ] Initial stakeholder discussions completed
- [ ] Market context and user insights available
- [ ] Strategic goals and constraints understood

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

### 📋 **7 Required Sections:**

1. **Executive Summary**: Complete picture in 3-4 sentences (Problem + Solution approach + Expected outcome)
2. **Problem Statement**: Specific pain point with scope boundaries (Who has problem + quantified impact)
3. **Opportunity**: Quantified business value + high-level approach (Expected benefits + solution direction)
4. **Target Users**: Specific user segments with demographics + behavior patterns + volume
5. **Success Metrics**: Measurable outcomes with specific timelines (How we'll know we solved the problem)
6. **Constraints**: Resource, technical, and assumption factors (Budget + timeline + technical + team limitations)
7. **Key Risks**: Major threats with impact assessment (Risk level + potential impact)

### ✅ **Quality Standards:**

- **Problem Focus**: Use concrete user scenarios and quantified evidence
- **User Specificity**: Be specific enough to guide solution design decisions
- **Measurable Success**: Specific, testable outcomes with realistic timelines
- **Realistic Constraints**: Honest about actual limitations and dependencies
- **Evidence-Based**: Support statements with data, research, and metrics

### ❌ **Common Pitfalls to Avoid:**

- Solution-oriented problem statements (focus on pain, not solutions)
- Vague user descriptions (be specific about segments and usage)
- Unmeasurable success metrics (avoid aspirational statements)
- Unrealistic constraints (be honest about limitations)
- Missing evidence (support with data and research)

### 🎯 **PRD Enablement:**

This Project Brief should enable immediate PRD creation by providing:

- Clear problem definition → PRD Problem/Opportunity section
- Target user clarity → PRD user research and requirements
- Success metrics → PRD Goals/Measurable Outcomes
- Constraints → PRD MVP scope and technical requirements
