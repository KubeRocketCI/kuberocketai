# Task: Create Product Requirements Document (PRD)

## Objective

Create a streamlined PRD that drives team alignment on what to build and why, following the proven 6-8 page structure focused on user needs and business value rather than technical specifications.

## Best Practice Structure (Keep to 6-8 pages maximum)

### 1. Problem/Opportunity

**Focus**: Be crisp and clear about what user or business problem you're solving
**Avoid**: "User can't use [solution]" - this is NOT a problem statement
**Key**: Look deeper into what issues are caused when functionality is missing

### 2. Target Users & Use Cases

**Focus**: Always prioritize the user - this aligns product building with go-to-market
**Key**: Be specific about users and use cases, ensure team alignment on definitions
**Result**: Helps with prioritization and strategic focus

### 3. Current Journeys/Landscape (Optional)

**Purpose**: Give context on what users do today or how competitors solve this
**Format**: Quick summary + links to detailed materials (user journey flows, competitive analysis)
**Value**: Shows why the problem is painful and what solutions have been tried

### 4. Proposed Solution/Elevator Pitch

**Format**: Standard 2-3 liner in plain English
**Include**: Top 3 MVP value props + conceptual model diagram
**Key**: Focus on user value, not technical implementation

### 5. Goals/Measurable Outcomes

**Length**: Literally 2-3 bullets, no more
**Focus**: Measurable outcomes defining success or non-failure
**Avoid**: Vague aspirational statements

### 6. MVP/Functional Requirements

**Critical**: Focus on required functionality, save the rest for appendix
**Question**: What's the "min-viable" set of functionality for target user adoption?

#### Requirements Best Practices

**✅ DO:**

- Focus on functionality: "First-time user must accept privacy policy to use product"
- Include telemetry: "Product team can monitor and visualize user engagement"
- Link to appropriate UX sketches for quick visualization
- Include priorities: [P0] [P1] [P2] where P0 = truly required for MVP
- Bucket by use case/user journey (controversial but effective)
- Consider all critical user journeys (CUJs) - create, maintain, retire, navigate
- Limit to 3 phases/milestones maximum

**❌ DON'T:**

- Include performance metrics: "99.99% uptime" unless users actually require it for adoption
- Include design details: "blue 'Continue' button if no database entry"
- Include technical details: specific implementation approaches
- Write PRD without user needs clarity - write strategy doc first instead

### 7. Appendix (Links to Other Docs)

**Purpose**: Link to materials people will ask for
**Include**: Product decisions, detailed UX, product specs, go-to-market, pricing, pre-mortem
**Why Separate**: These don't matter if you can't align on problem/solution first

## SDLC Framework Integration

### File Structure Requirements

- **Input Required**: Project Brief at `/docs/prd/project-brief.md` (EXACT path)
- **Output Location**: `/docs/prd/prd.md` (EXACT path and filename)
- **Directory Structure**: Follow SDLC framework exactly

### Downstream Enablement

Your PRD should enable Epic creation by:

- Clear functional requirements that can be grouped into Epics (following `{epic_number}-epic-{slug}.md` format)
- Use case buckets that map to Epic themes
- Priority framework (P0/P1/P2) for Epic sequencing
- Success criteria that become Epic completion criteria

### SDLC Directory Structure Compliance

```bash
/docs/
├── prd/                          # Product vision & requirements
│   ├── project-brief.md          # Project vision & strategy (INPUT)
│   └── prd.md                    # Product requirements (OUTPUT)
├── epics/                        # High-level features ({epic_number}-epic-{slug}.md)
├── stories/                      # User stories ({epic_number}.{story_number}.story.md)
├── architecture/                 # System design
└── tests/                        # Quality validation
```

## Execution Steps

### Foundation

1. **Validate Project Brief**: Ensure `/docs/prd/project-brief.md` exists and is complete
2. **Stakeholder Input**: Interview key stakeholders about user needs and business context
3. **User Research Review**: Gather existing user research and pain point data

### Core Content

4. **Problem Definition**: Write crisp problem statement with supporting evidence
5. **User Analysis**: Define specific target users and primary use cases
6. **Solution Design**: Create 2-3 line elevator pitch + top 3 value props
7. **Goals Setting**: Define 2-3 measurable success outcomes

### Requirements

8. **MVP Scope**: Define minimum viable functionality for user adoption
9. **Priority Framework**: Apply P0/P1/P2 to all requirements
10. **Use Case Buckets**: Group requirements by user journey for Epic mapping

### Finalization

11. **Length Check**: Ensure document is 6-8 pages maximum
12. **Epic Readiness**: Validate requirements can be grouped into logical Epics
13. **Framework Compliance**: Save to `/docs/prd/prd.md`

## Quality Validation

### Pre-Release Checklist

- [ ] **File Location**: Saved exactly to `/docs/prd/prd.md`
- [ ] **Length**: Document is 6-8 pages maximum
- [ ] **User Focus**: Every requirement maps to user value
- [ ] **Problem Clarity**: Problem statement avoids solution-oriented language
- [ ] **Measurable Goals**: Success criteria are specific and testable
- [ ] **MVP Boundaries**: Clear distinction between P0 (required) and P1/P2 (nice-to-have)
- [ ] **Epic Ready**: Requirements can be logically grouped for Epic creation following naming convention

### Success Indicators

- Stakeholders can quickly understand what to build and why
- Engineering can estimate effort without additional requirements gathering
- Product team can create Epics directly from PRD functional requirements using SDLC naming conventions
- Document remains relevant throughout development without constant updates
- Team debates focus on priorities rather than missing requirements

## Common Pitfalls to Avoid

1. **Comprehensive Documentation Trap**: Don't try to document everything - focus on alignment
2. **Technical Implementation Details**: Avoid "how" details - focus on "what" and "why"
3. **Perfectionist Requirements**: Don't over-specify - leave room for design and engineering creativity
4. **Missing User Context**: Always anchor requirements in user scenarios and business value
5. **Scope Creep Prevention**: Use appendix for future considerations rather than expanding core PRD
6. **Framework Over-Engineering**: Keep SDLC compliance minimal and non-intrusive
7. **File Structure Violations**: Always use exact SDLC paths and naming conventions

## LLM Agent Guidance

### Research Phase Questions

Ask stakeholders these key questions:

- "What specific user scenarios validate this problem?"
- "How does this problem impact daily user workflows?"
- "What evidence supports the opportunity size?"
- "What solutions have users tried that didn't work?"
- "What's the minimum functionality needed for user adoption?"

### Writing Phase Focus

- Start with user problems, not product features
- Use concrete examples and user quotes when possible
- Quantify impact and success criteria with specific metrics
- Group related functionality by user journey
- Apply strict priority framework to every requirement

### Epic Enablement Approach

- Structure requirements so they naturally group into Epic themes
- Include enough detail for Epic estimation but not implementation
- Provide clear success criteria that become Epic completion criteria
- Ensure each Epic will have clear user value proposition
- Follow SDLC naming convention: `{epic_number}-epic-{slug}.md`

### SDLC Framework Compliance

- **Mandatory Input**: Validate `/docs/prd/project-brief.md` exists and is complete
- **Mandatory Output**: Save PRD exactly to `/docs/prd/prd.md`
- **Quality Gate**: PRD approval required before Epic creation
- **Downstream**: Enables `/docs/epics/` and `/docs/architecture/` creation

**Remember**: The PRD is a vehicle for alignment, not comprehensive documentation. If your team can align without this format, that's perfectly fine too.
