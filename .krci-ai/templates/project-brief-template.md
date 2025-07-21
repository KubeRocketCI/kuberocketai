# Project Brief: {{project_name}}

> **Target Length**: 2-3 pages maximum
> **Framework**: Root artifact in SDLC framework
> **File Location**: MUST be saved as `/docs/prd/project-brief.md` (exact path)

---

## Executive Summary

{{executive_summary}}

<!-- Template Guidance:
Write a compelling 3-4 sentence overview combining problem, solution approach, and expected outcome.

Example: "Our SaaS platform experiences 2,500 password-related support tickets monthly, consuming 15% of support resources and frustrating users. We will implement biometric authentication and social login options to reduce password dependency, targeting 80% reduction in support tickets and $50K annual savings. This 3-month initiative serves 10,000+ monthly active users and requires Auth0 integration with a $25K budget."

Key Elements:
- What problem are we solving? (specific and quantified)
- How will we solve it? (high-level approach)
- What's the expected outcome? (business value)
- What's the scope? (timeline, users, constraints)
-->

---

## Problem Statement

{{problem_statement}}

<!-- Template Guidance:
Define the specific pain point driving this project with clear scope boundaries.

Example: "Users frequently forget passwords leading to 15% of support tickets and account lockouts (2,500/month). Focus on authentication workflow only, excluding password policy management or user registration processes."

Best Practices:
- Start with user scenarios, not business needs
- Use concrete numbers and evidence
- Define what's included and excluded
- Avoid solution-oriented language
- Focus on pain points and their impact

Evidence to Include:
- Support ticket volumes
- User research findings
- Productivity impact data
- Cost of current workarounds
-->

---

## Opportunity

{{opportunity}}

<!-- Template Guidance:
Quantified business value plus high-level solution approach.

Example: "Reducing password-related support tickets by 80% would save $50K annually and improve user satisfaction scores by 25%. Implement biometric authentication and social login options to reduce password dependency."

Key Elements:
- Business value (cost savings, revenue, efficiency)
- User value (time savings, satisfaction, productivity)
- Market opportunity (competitive advantage, growth)
- High-level solution direction (not detailed implementation)

Quantification Examples:
- Cost reduction: "$50K annual savings"
- Time savings: "15 minutes per user per month"
- Satisfaction: "25% improvement in user satisfaction"
- Efficiency: "80% reduction in support tickets"
-->

---

## Target Users

{{target_users}}

<!-- Template Guidance:
Specific user segments who have this problem with usage patterns and demographics.

Example: "SaaS platform users (10,000+ monthly active users) who access the platform 3+ times per week. Primary segment: business professionals aged 25-45 accessing from mobile devices (60%) and desktop (40%)."

Include:
- User volume and growth trends
- Demographics (age, role, industry)
- Usage patterns (frequency, device, context)
- Segment prioritization (primary vs secondary)
- Geographic distribution if relevant

User Segment Examples:
- "10,000+ monthly active users"
- "Business professionals aged 25-45"
- "Mobile-first users (60% mobile, 40% desktop)"
- "Access platform 3+ times weekly"
- "Located primarily in North America and Europe"
-->

---

## Success Metrics

{{success_metrics}}

<!-- Template Guidance:
How we'll measure if we've solved the problem with specific timelines.

Example: "Reduce password-related support tickets by 80% within 3 months, maintain 99.9% uptime, achieve 70% user adoption of new auth methods within 6 months, improve login success rate from 85% to 95%."

Success Criteria Format:
- Specific: Exactly what will be measured
- Measurable: Numbers, percentages, timelines
- Achievable: Realistic given constraints
- Relevant: Directly tied to problem and opportunity
- Time-bound: Clear deadlines

Metric Categories:
- Problem Resolution: "80% reduction in support tickets"
- User Adoption: "70% user adoption within 6 months"
- Quality: "99.9% uptime maintained"
- User Experience: "Login success rate 85% → 95%"
- Business Impact: "$50K annual cost savings"
-->

---

## Constraints

{{constraints}}

<!-- Template Guidance:
Resource, technical, and assumption factors that limit the solution.

Example: "Must integrate with existing Auth0 setup, 3-month timeline, $25K budget, maximum 2 developers assigned. Assumes current mobile app architecture supports biometric APIs and users have compatible devices."

Constraint Categories:

### Resource Constraints:
- Budget: "$25K maximum budget"
- Timeline: "3-month delivery deadline"
- Team: "Maximum 2 developers available"
- Skills: "No iOS development expertise on team"

### Technical Constraints:
- Integration: "Must integrate with existing Auth0"
- Architecture: "Cannot modify core database schema"
- Performance: "Must maintain current response times"
- Security: "Must meet enterprise security standards"

### Business Constraints:
- Compliance: "Must maintain SOC 2 compliance"
- User Impact: "Zero downtime deployment required"
- Support: "Cannot increase support complexity"
- Branding: "Must align with current UI/UX standards"

### Key Assumptions:
- "Users have biometric-capable devices"
- "Auth0 API will remain stable"
- "No major iOS/Android changes during development"
-->

---

## Key Risks

{{key_risks}}

<!-- Template Guidance:
Major risks that could derail the project with impact assessment.

Example: "User adoption resistance (HIGH): Users may prefer familiar passwords. Auth0 API changes (MEDIUM): Potential breaking changes during integration. Biometric compatibility (MEDIUM): Older devices may not support all features. Timeline risk (HIGH): Integration complexity may exceed estimates."

Risk Assessment Format:
[Risk Name] ([Impact Level]): [Description and potential impact]

Impact Levels:
- HIGH: Could significantly delay or derail project
- MEDIUM: Could cause delays or require scope changes
- LOW: Minor impact, manageable workarounds available

Risk Categories:

### User Adoption Risks:
- "User resistance to change (HIGH)"
- "Learning curve for new features (MEDIUM)"
- "Device compatibility issues (MEDIUM)"

### Technical Risks:
- "Integration complexity (HIGH)"
- "Third-party API changes (MEDIUM)"
- "Performance impact (LOW)"

### Business Risks:
- "Timeline overrun (HIGH)"
- "Budget overrun (MEDIUM)"
- "Resource unavailability (MEDIUM)"

### Market Risks:
- "Competitive response (LOW)"
- "Regulatory changes (MEDIUM)"
- "Technology shifts (LOW)"
-->

---

## SDLC Framework Information

**Dependencies**: None (root artifact)
**Output Location**: This Project Brief MUST be saved as `/docs/prd/project-brief.md`
**Downstream Enablement**: Enables PRD creation at `/docs/prd/prd.md`

<!-- SDLC Framework Integration:
This Project Brief serves as the foundation for:
- PRD Problem/Opportunity section
- PRD Target Users & Use Cases
- PRD Goals/Measurable Outcomes
- PRD scope and constraint definition

Directory Structure:
/docs/
├── prd/                          # Product vision & requirements
│   ├── project-brief.md          # Project vision & strategy (THIS FILE)
│   └── prd.md                    # Product requirements (ENABLED BY THIS)
├── epics/                        # High-level features
├── stories/                      # User stories
├── architecture/                 # System design
└── tests/                        # Quality validation
-->

---

<!-- QUALITY CHECKLIST
✅ Document is 2-3 pages maximum
✅ Executive summary captures complete project essence
✅ Problem statement is specific and evidence-based
✅ Opportunity is quantified with business value
✅ Target users are specific with usage patterns
✅ Success metrics are measurable with timelines
✅ Constraints are realistic and comprehensive
✅ Key risks identified with impact assessment
✅ File saved exactly as /docs/prd/project-brief.md
✅ Ready to enable PRD creation
-->