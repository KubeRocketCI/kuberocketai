---
dependencies:
  data:
    - validation-frameworks.md
    - business-frameworks.md
  templates:
    - validation-report-template.md
---

# Task: Validate Problem Statement

## Description

Apply Lean Startup Problem-Solution Fit Assessment framework to validate problem statement accuracy, evidence quality, and strategic alignment. This validation ensures the problem is real, urgent, and worth solving while identifying gaps that could impact project success.

This validation uses [validation frameworks](./.krci-ai/data/validation-frameworks.md) and outputs results using the [validation report template](./.krci-ai/templates/validation-report-template.md).

## Instructions

<instructions>
Confirm project brief with problem statement section exists, basic problem research or stakeholder input is available, access to potential users or customer data is accessible, and competitive landscape understanding is available. Ensure dependencies declared in the YAML frontmatter are readable before proceeding.

Apply Lean Startup Problem-Solution Fit Assessment framework from [validation-frameworks.md](./.krci-ai/data/validation-frameworks.md) to validate problem statement accuracy, evidence quality, and strategic alignment. Extract and structure the problem hypothesis from the current project brief including core problem statement, problem scope (included/excluded), affected user segments, and problem context (when, where, why). Use [validation-report-template.md](./.krci-ai/templates/validation-report-template.md) for output.
</instructions>

## Problem Hypothesis Structure

```markdown
Problem Hypothesis: {{user_segment}} experiences {{problem_description}} when {{situation_context}}, resulting in {{negative_impact}}.

Specific Claims to Validate:
- Problem occurs for {{user_segment}} with frequency of {{frequency_claim}}
- Problem causes {{impact_type}} impact measured by {{impact_metric}}
- Current solutions fail because {{solution_gap_reason}}
- Problem urgency driven by {{urgency_factors}}
```

### Phase 2: Lean Startup Validation Framework Application

#### Step 1: Problem Validation Evidence Collection

Customer Interview Validation:
- Conduct minimum 10 target user interviews
- Use structured interview guide focusing on problem experiences
- Document specific problem scenarios and pain points
- Quantify problem frequency and impact per user

Quantified Problem Metrics:
- Collect measurable indicators of problem existence
- Support ticket data, time spent on workarounds, cost of current solutions
- User behavior analytics showing problem manifestation
- Business impact data (revenue loss, efficiency reduction)

Competitive Analysis:
- Document existing solutions and their limitations
- Identify solution gaps that validate problem persistence
- Assess market demand through competitive pricing and adoption
- Analyze competitor positioning to validate problem importance

#### Step 2: Problem-Solution Fit Assessment

Problem Validation Scoring (1-10 scale):
- Problem Intensity: How painful is the problem for users?
- Problem Frequency: How often do users experience this problem?
- Problem Reach: How many users/companies experience this problem?
- Problem Urgency: How urgent is solving this problem for users?

Evidence Quality Assessment:
- Primary evidence percentage (direct user research vs secondary sources)
- Evidence recency (how recent is the supporting data)
- Evidence diversity (multiple source types and perspectives)
- Evidence reliability (credible sources and methodologies)

#### Step 3: Solution Hypothesis Validation

Solution Approach Assessment:
- Does the proposed solution address identified root causes?
- Is the solution technically and commercially feasible?
- Does the solution create meaningful differentiation vs alternatives?
- Can the solution be delivered within stated constraints?

Solution Desirability Testing:
- User reaction to proposed solution approach
- Willingness to pay or switch from current solutions
- Feature prioritization based on problem-solving value
- Solution-market fit assessment

### Phase 3: Evidence Analysis and Validation

#### Evidence Quality Framework

Primary Evidence Sources (Highest confidence):
- Direct customer interviews with problem validation
- First-party usage data showing problem manifestation
- Customer support data quantifying problem frequency
- User research studies with problem-focused findings

Secondary Evidence Sources (Moderate confidence):
- Industry reports mentioning related problems
- Competitive analysis revealing solution gaps
- Expert opinions on problem significance
- Market research highlighting related pain points

Evidence Confidence Scoring:
- High (80-100%): Multiple primary sources, quantified data, recent research
- Medium (60-79%): Mix of primary/secondary, some quantification, reasonably recent
- Low (40-59%): Primarily secondary sources, limited quantification, dated research
- Very Low (<40%): Assumptions without validation, anecdotal evidence only

#### Validation Results Synthesis

Problem-Solution Fit Score Calculation:

```txt
Problem Score = (Intensity + Frequency + Reach + Urgency) / 4
Solution Score = (Root Cause Fit + Feasibility + Differentiation + Deliverability) / 4
Overall Fit = (Problem Score + Solution Score) / 2
```

Fit Assessment Thresholds:
- Strong Fit (8-10): High confidence to proceed with current problem/solution
- Moderate Fit (6-7.9): Some refinement needed, additional validation recommended
- Weak Fit (4-5.9): Significant issues identified, major refinement required
- Poor Fit (<4): Problem or solution hypothesis needs fundamental revision

### Phase 4: Validation Results Documentation

#### Create Validation Report

Use validation report template to document:
- Framework methodology and process followed
- Evidence collected with quality assessment
- Problem-solution fit scores and rationale
- Key findings and insights
- Assumptions validated, challenged, or identified
- Recommendations for project brief updates

#### Update Project Brief

Based on validation results:
- Enhance problem statement with validated evidence
- Update confidence levels and assumption tracking
- Revise problem scope or focus based on findings
- Add validation checkpoint completion status

#### Update Assumption Tracker

Document assumption changes:
- Mark validated assumptions with evidence sources
- Update confidence levels based on validation results
- Add new assumptions identified during validation
- Prioritize remaining assumptions for future validation

## Output Format

<output_format>
- Primary Output: `/docs/prd/brief-validation-problem.md` using validation report template
- Secondary Outputs:
  - Updated `/docs/prd/project-brief.md` with validation results
  - Updated `/docs/prd/brief-assumptions.md` with assumption status
- Length: Validation report 2-3 pages, brief updates minimal
- Evidence Documentation: All sources properly cited and quality-assessed
</output_format>

## Success Criteria

<success_criteria>

### Phase 1: Problem Hypothesis

- Problem hypothesis extracted from current project brief
- Specific validation claims identified with measurable criteria
- Problem scope clearly defined with inclusion/exclusion boundaries
- Target user segments specified for validation activities

### Phase 2: Framework Application

- Customer interviews completed with minimum 10 target users
- Quantified metrics collected showing problem existence and impact
- Competitive analysis conducted validating solution gaps
- Problem-solution fit scores calculated using framework methodology

### Phase 3: Evidence Analysis

- Evidence quality assessed with confidence levels assigned
- Primary evidence prioritized over secondary sources
- Multiple source validation completed for key claims
- Evidence gaps identified with recommendations for closure

### Phase 4: Documentation

- Validation report completed using standardized template
- Project brief updated with validated evidence and confidence levels
- Assumption tracker updated with validation results
- Recommendations provided for next validation steps
</success_criteria>

## Execution Checklist

<execution_checklist>

### Pre-Validation Setup

- Current brief analysis: Review existing problem statement and evidence
- Validation planning: Identify specific claims to test and evidence needed
- Stakeholder coordination: Arrange access to customers, data, and subject matter experts
- Interview guide preparation: Develop structured questions for customer interviews

### Problem Validation Execution

- Customer interviews: Conduct minimum 10 interviews with structured approach
- Data collection: Gather quantified evidence from internal and external sources
- Competitive research: Analyze existing solutions and market positioning
- Expert consultation: Validate findings with internal subject matter experts

### Analysis and Synthesis

- Evidence analysis: Assess quality and reliability of collected evidence
- Scoring calculation: Apply framework scoring methodology consistently
- Pattern identification: Identify themes and insights across evidence sources
- Gap identification: Document areas needing additional validation

### Documentation and Communication

- Report creation: Complete validation report using standardized template
- Brief updates: Update project brief with validation results and evidence
- Assumption updates: Revise assumption tracker based on validation findings
- Stakeholder communication: Present findings to key project stakeholders

</execution_checklist>

## Content Guidelines

### Validation Quality Standards

- Evidence-Based: All conclusions supported by documented evidence with source attribution
- Quantified Focus: Prioritize measurable evidence over qualitative opinions
- Multiple Source Validation: Key findings validated across multiple evidence sources
- Bias Mitigation: Actively seek disconfirming evidence and diverse perspectives
- Transparency: Clearly document limitations, assumptions, and confidence levels

### Interview Best Practices

- Open-ended Questions: Avoid leading questions that bias toward expected answers
- Specific Scenarios: Ask for concrete examples rather than general opinions
- Quantification: Request specific numbers (frequency, time, cost) when possible
- Problem Focus: Focus on problems and pain points before discussing solutions
- Context Gathering: Understand when, where, and why problems occur

### Evidence Documentation Standards

- Source Attribution: Clearly identify all evidence sources with dates and methodology
- Quality Assessment: Rate evidence quality and reliability consistently
- Confidence Levels: Assign and document confidence levels for all key findings
- Assumptions Tracking: Explicitly identify and track all assumptions made during analysis
- Update Requirements: Document what evidence updates would change conclusions

## Framework Integration Notes

- SDLC Integration: Validation results feed into Epic creation and solution architecture
- Business Framework Usage: Leverages Lean Startup methodology from [validation frameworks](./.krci-ai/data/validation-frameworks.md) for systematic validation
- Evidence Standards: Maintains KubeRocketAI focus on quantified, evidence-based approach
- Quality Assurance: Built-in scoring and confidence assessment ensure reliable results
- Professional Output: Structured documentation suitable for stakeholder decision-making
