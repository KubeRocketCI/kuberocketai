---
dependencies:
  data:
    - architecture-principles.md
    - krci-ai/core-sdlc-framework.md
  templates:
    - sad-template.md
    - architecture-review.md
---

# Task: Review Architecture Documentation

## Description

Conduct comprehensive review of multi-file architecture documentation to ensure technical quality, PRD requirement compliance, and readiness for development implementation. This review validates that architecture meets enterprise standards and enables successful Epic/Story development across all architecture sections.

## Prerequisites

<prerequisites>
- Completed architecture: Architecture documentation exists in `/docs/architecture/` with sections following [sad-template.md](./.krci-ai/templates/sad-template.md) structure
- Reference documents: Access to PRD (`/docs/prd/prd.md`) and Epics (`/docs/epics/`) for validation
- Architecture standards: Understanding of organizational architecture principles and guidelines from [architecture-principles.md](./.krci-ai/data/architecture-principles.md)
- Review criteria: Clear understanding of quality gates and acceptance criteria
</prerequisites>

## Instructions

<instructions>
1. Follow SDLC workflow: Reference [sdlc-framework.md](./.krci-ai/data/krci-ai/core-sdlc-framework.md) for quality gate requirements and review criteria
2. Apply review standards: Use [architecture-principles.md](./.krci-ai/data/architecture-principles.md) for quality assessment
3. Format output: Use [architecture-review.md](./.krci-ai/templates/architecture-review.md) for review documentation
4. Validate traceability: Ensure all PRD requirements and Epic features are addressed across architecture sections
</instructions>

## Output Format

<output_format>
- Location: `/docs/architecture/architecture-review-{date}.md` (EXACT path and filename)
- Review outcome: Clear PASS/FAIL determination with detailed findings for each architecture section
- Issue documentation: Specific issues found with actionable remediation guidance
- Quality gate status: Formal approval or rejection for implementation phase
</output_format>

## Success Criteria

<success_criteria>
- Review completed: Comprehensive assessment of all architecture sections documented per [sad-template.md](./.krci-ai/templates/sad-template.md) structure
- Quality determination: Clear PASS/FAIL decision with detailed rationale for each section
- Issues documented: Specific findings with actionable remediation steps
- Traceability validated: All PRD requirements verified as addressed across architecture sections
- Standards compliance: Architecture meets organizational standards and best practices
- Implementation readiness: Architecture provides sufficient guidance for development teams through 07-transition-migration.md
</success_criteria>

## Execution Checklist

### Document Review Phase

<document_review>
- Completeness check: Verify all 11 sections are present and no template variables remain
- Content quality: Assess technical accuracy, clarity, and professional presentation
- Internal consistency: Validate consistency between sections and architectural decisions
- Standards compliance: Ensure architecture follows organizational principles and guidelines
</document_review>

### Requirements Validation Phase

<requirements_validation>
- PRD traceability: Verify every BR/NFR requirement is addressed in architecture
- Epic alignment: Confirm architecture supports all Epic implementations
- Quality attributes: Validate NFR requirements have specific architectural approaches
- Constraint compliance: Ensure architecture respects stated constraints and limitations
</requirements_validation>

### Technical Assessment Phase

<technical_assessment>
- Architecture feasibility: Assess technical viability of proposed solutions
- Technology decisions: Evaluate technology choices against requirements and standards
- Risk assessment: Identify architectural risks and validate mitigation strategies
- Implementation guidance: Confirm architecture provides clear development direction
</technical_assessment>

### Quality Gate Phase

<quality_gate>
- Review documentation: Complete [architecture-review.md](./.krci-ai/templates/architecture-review.md) template
- Decision rationale: Document clear reasoning for PASS/FAIL determination
- Issue prioritization: Categorize findings by severity and implementation impact
- Next steps: Define clear action items for architecture improvement or approval
</quality_gate>

## Content Guidelines

### 📋 Review Focus Areas:

<review_focus_areas>
1. Section Completeness: All 11 sections populated with relevant, project-specific content
2. Requirements Coverage: Every PRD BR/NFR requirement mapped to architectural components
3. Epic Enablement: Architecture provides clear implementation guidance for all Epics
4. Quality Attributes: NFR requirements addressed with specific architectural approaches
5. Decision Quality: Architectural decisions have clear rationale and consider alternatives
6. Professional Standards: Document meets enterprise architecture documentation standards
</review_focus_areas>

### ✅ PASS Criteria:

<pass_criteria>
- Complete Documentation: All 11 sections fully populated without template variables
- Requirements Compliance: 100% of PRD BR/NFR requirements addressed in architecture
- Epic Support: Architecture enables all Epic implementations with clear guidance
- Quality Standards: Document meets professional architecture documentation standards
- Technical Feasibility: Proposed architecture is technically sound and implementable
- Decision Quality: Architectural decisions are well-reasoned with clear alternatives
</pass_criteria>

### ❌ FAIL Criteria:

<fail_criteria>
- Missing or incomplete sections in SAD document
- PRD requirements not addressed or poorly mapped to architecture
- Insufficient implementation guidance for Epic/Story development
- Architectural decisions without clear rationale or alternatives
- Technical solutions that are not feasible or violate constraints
- Documentation quality below professional standards
</fail_criteria>

### 🔍 Common Review Issues:

<common_review_issues>

#### Completeness Issues

- Template variables ({{variable}}) not replaced with project-specific content
- Sections missing or containing placeholder text
- Architectural diagrams missing or insufficient detail

#### Requirements Issues

- PRD BR/NFR requirements not traced to architectural components
- New requirements introduced without PRD justification
- Quality attributes without specific implementation approaches

#### Technical Issues

- Technology choices without clear rationale or trade-off analysis
- Architectural patterns that don't align with organizational standards
- Solutions that don't address stated constraints or limitations
</common_review_issues>

### 🎯 Review Questions:

<review_questions>
Key questions to evaluate during review:

- "Are all PRD BR/NFR requirements clearly addressed in the architecture?"
- "Can development teams create Epics and Stories from this architecture guidance?"
- "Are architectural decisions well-reasoned with clear alternatives considered?"
- "Does the architecture meet organizational standards and best practices?"
- "Is the proposed solution technically feasible within stated constraints?"
</review_questions>

### 📋 Quality Gate Checklist:

<quality_gate_checklist>
- Documentation Quality: Professional presentation suitable for stakeholder review
- Requirements Compliance: All PRD requirements addressed with architectural solutions
- Epic Enablement: Clear implementation guidance for all Epic features
- Technical Soundness: Proposed solutions are feasible and well-architected
- Standards Alignment: Architecture follows organizational principles and guidelines
- Decision Quality: Major decisions have clear rationale and alternatives considered
</quality_gate_checklist>
