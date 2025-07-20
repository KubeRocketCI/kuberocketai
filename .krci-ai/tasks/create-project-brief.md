# Task: Create Project Brief

## Objective

Create a focused Project Brief that captures the essential foundation for product development by answering why, who, what success looks like, and what constraints shape the solution. Enables PRD creation and strategic alignment.

## Project Brief Purpose

The Project Brief is the **root artifact** in the SDLC framework that:

- Defines the essential foundation for all downstream artifacts
- Answers fundamental questions before solution development begins
- Provides strategic context for PRD creation
- Ensures stakeholder alignment on problem and opportunity

## Best Practice Structure (Keep to 2-3 pages maximum)

### 1. Executive Summary

**Purpose**: Complete picture in 3-4 sentences
**Include**: Problem + Solution approach + Expected outcome
**Key**: Compelling overview that captures project essence

### 2. Problem Statement

**Focus**: Specific pain point with scope boundaries
**Include**: Who has the problem + quantified impact
**Avoid**: Solution-oriented language - focus on the pain

### 3. Opportunity

**Focus**: Quantified business value + high-level approach
**Include**: Expected benefits + solution direction
**Key**: Why this is worth doing now

### 4. Target Users

**Focus**: Specific user segments with usage patterns
**Include**: Demographics + behavior patterns + volume
**Key**: Be specific enough to guide solution design

### 5. Success Metrics

**Focus**: How we'll know we solved the problem
**Include**: Specific timelines + measurable outcomes
**Avoid**: Vague aspirational statements

### 6. Constraints

**Focus**: Resource, technical, and assumption factors
**Include**: Budget + timeline + technical + team limitations
**Key**: Realistic boundaries that limit solution options

### 7. Key Risks

**Focus**: Major threats with impact assessment
**Include**: Risk level (HIGH/MEDIUM/LOW) + potential impact
**Key**: Awareness of what could derail the project

## SDLC Framework Integration

### File Structure Requirements

- **Output Location**: `/docs/prd/project-brief.md` (EXACT path and filename)
- **Registry Update**: Add artifact to `/docs/registry.json`
- **Downstream Enable**: Enables PRD creation at `/docs/prd/prd.md`

### Strategic Foundation Role

Your Project Brief should enable PRD creation by:

- Clear problem definition that becomes PRD Problem/Opportunity section
- Target user clarity that guides PRD user research and requirements
- Success metrics that become PRD Goals/Measurable Outcomes
- Constraints that inform PRD MVP scope and technical requirements

## Execution Steps

### Discovery

1. **Stakeholder Interviews**: Understand business context and strategic priorities
2. **Problem Validation**: Gather evidence that this problem is real and significant
3. **User Research**: Identify who has this problem and how it impacts them
4. **Opportunity Sizing**: Quantify business value and market opportunity

### Analysis

5. **Problem Definition**: Write specific problem statement with evidence
6. **User Segmentation**: Define target users with demographics and usage patterns
7. **Success Planning**: Define measurable outcomes with realistic timelines
8. **Constraint Assessment**: Identify realistic limitations and assumptions

### Documentation

9. **Brief Creation**: Use [project-brief-template.md](/.krci-ai/templates/project-brief-template.md) structure
10. **Length Check**: Ensure document is 2-3 pages maximum for focus
11. **Framework Compliance**: Save to `/docs/prd/project-brief.md` and update registry

## Quality Validation

### Pre-Release Checklist

- [ ] **File Location**: Saved exactly to `/docs/prd/project-brief.md`
- [ ] **Length**: Document is 2-3 pages maximum
- [ ] **Problem Clarity**: Problem statement is specific and evidence-based
- [ ] **User Specificity**: Target users are clearly defined with usage patterns
- [ ] **Measurable Success**: Success metrics are specific and testable
- [ ] **Realistic Constraints**: Constraints reflect actual limitations
- [ ] **Risk Assessment**: Key risks identified with impact levels
- [ ] **Registry Updated**: Artifact registered in `/docs/registry.json`

### Success Indicators

- Stakeholders can quickly understand why this project matters
- Problem statement guides solution development without being prescriptive
- Target users are specific enough to inform design decisions
- Success metrics provide clear validation criteria
- Constraints help teams make realistic scope decisions
- Document enables immediate PRD creation

## Best Practices

### Problem Definition Excellence

- Start with user scenarios and pain points
- Use quantified evidence (support tickets, user research, business metrics)
- Define scope boundaries (what's included and excluded)
- Avoid solution-oriented problem statements

### User Focus

- Be specific about user segments and demographics
- Include usage patterns and frequency
- Quantify user base size and growth
- Connect user needs to business value

### Success Measurement

- Define specific, measurable outcomes
- Include realistic timelines for achievement
- Focus on business and user value metrics
- Avoid vanity metrics that don't drive decisions

### Constraint Realism

- Include all major limitations (budget, timeline, technical, team)
- Document key assumptions that could change
- Be honest about what's achievable
- Consider external dependencies

## Common Pitfalls to Avoid

1. **Solution Creep**: Don't prescribe solutions - focus on problems and opportunities
2. **Scope Inflation**: Keep focused on core problem rather than expanding
3. **Vague Users**: Avoid generic user descriptions - be specific about segments
4. **Unmeasurable Success**: Ensure all success metrics are concrete and testable
5. **Unrealistic Constraints**: Be honest about actual limitations and assumptions
6. **Missing Evidence**: Support problem statements with data and research
7. **Framework Shortcuts**: Always use exact SDLC paths and naming conventions

## LLM Agent Guidance

### Discovery Questions

Ask stakeholders these key questions:

- "What specific evidence shows this problem exists and matters?"
- "Who exactly has this problem and how does it impact their work/life?"
- "What's the quantified business opportunity if we solve this?"
- "What are the realistic constraints on time, budget, and resources?"
- "What could go wrong that would prevent success?"

### Writing Focus

- Start with concrete user scenarios and pain points
- Use specific data and metrics wherever possible
- Quantify impact in business terms (cost, time, satisfaction)
- Define clear boundaries on what's in and out of scope
- Be realistic about what can be achieved with available resources

### PRD Enablement

- Problem definition should directly inform PRD Problem/Opportunity section
- Target users should guide PRD user research and use case definition
- Success metrics should become PRD Goals/Measurable Outcomes
- Constraints should inform PRD MVP scope and technical assumptions

### SDLC Framework Compliance

- **Mandatory Output**: Save Project Brief exactly to `/docs/prd/project-brief.md`
- **Registry Update**: Add artifact entry to `/docs/registry.json`
- **Quality Gate**: Project Brief approval required before PRD creation
- **Template Use**: Follow project-brief-template.md structure exactly

**Remember**: The Project Brief is the foundation for all downstream work. Spend time getting it right - a strong Project Brief makes PRD creation much easier and more focused.
