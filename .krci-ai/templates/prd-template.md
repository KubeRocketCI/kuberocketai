# Product Requirements Document: {{product_name}}

> **Target Length**: 6-8 pages maximum
> **Framework**: PRD following proven best practices from industry leaders
> **SDLC**: Enables Epic creation and Architecture documents
> **File Location**: MUST be saved as `/docs/prd/prd.md` (exact path)

---

## 1. Problem/Opportunity

<!-- Be crisp and clear about what user or business problem you're solving -->
<!-- AVOID: "User can't use [solution]" - this is NOT a problem statement -->
<!-- FOCUS: What issues are caused when functionality is missing? -->

{{problem_statement}}

**Evidence:**
{{supporting_evidence}}

<!-- Template Guidance:
Problem Example: "Users frequently forget passwords leading to 15% of support tickets and account lockouts (2,500/month). This creates frustration and consumes 15% of support resources."

Evidence Example: "User research shows 65% of users reset passwords monthly. Support ticket analysis reveals 2,500 password-related tickets costing $50K annually."
-->

---

## 2. Target Users & Use Cases

<!-- Always focus on the user - this aligns building with go-to-market -->
<!-- Be specific about users and use cases, ensure team alignment on definitions -->

**Primary Users:**
{{primary_users}}

**Key Use Cases:**
{{key_use_cases}}

<!-- Template Guidance:
Primary Users Example: "SaaS platform users (10,000+ monthly active) who access platform 3+ times weekly. Primary segment: business professionals aged 25-45 accessing from mobile (60%) and desktop (40%)."

Key Use Cases Example:
1. Daily login for work tasks (highest frequency)
2. Password recovery when locked out (highest pain)
3. Multi-device access synchronization (growing need)
-->

---

## 3. Current Journeys/Landscape *(Optional)*

<!-- Give context on what users do today or how competitors solve this -->
<!-- Quick summary + links to detailed materials -->

**Current User Journey:**
{{current_journey}}

**Competitive Landscape:**
{{competitive_analysis}}

<!-- Template Guidance:
Current Journey Example: "Users must remember complex passwords, leading to frequent lockouts. Recovery process takes 5-10 minutes via email verification."

Competitive Analysis: "Auth0, Okta provide enterprise solutions. Consumer apps use Face ID/Touch ID. Gap exists for SMB-focused authentication."

Links: "[Detailed user journey flow](link)" or "[Competitive analysis doc](link)"
-->

---

## 4. Proposed Solution/Elevator Pitch

<!-- Standard 2-3 liner in plain English -->
<!-- Include top 3 MVP value props + conceptual model -->

**Elevator Pitch:**
{{elevator_pitch}}

**Top 3 MVP Value Props:**

1. {{value_prop_1}}
2. {{value_prop_2}}
3. {{value_prop_3}}

**Conceptual Model:**
{{conceptual_model}}

<!-- Template Guidance:
Elevator Pitch Example: "Enable users to login using biometric authentication (fingerprint/face) and social login options, reducing password dependency by 80% while maintaining enterprise security standards."

Value Props Example:
1. 3-second biometric login eliminates password frustration
2. Social login reduces new user signup friction
3. Enterprise security maintains compliance requirements

Conceptual Model: "[Include simple diagram or description of how users will interact with the solution]"
-->

---

## 5. Goals/Measurable Outcomes

<!-- Literally 2-3 bullets, no more -->
<!-- Measurable outcomes defining success or non-failure -->

**Success Metrics:**

1. {{success_metric_1}}
2. {{success_metric_2}}
3. {{success_metric_3}}

<!-- Template Guidance:
Success Metrics Example:
1. Reduce password-related support tickets by 80% within 3 months
2. Achieve 70% user adoption of new auth methods within 6 months
3. Improve login success rate from 85% to 95%

AVOID vague statements like "improve user experience" or "increase engagement"
-->

---

## 6. MVP/Functional Requirements

<!-- Focus on required functionality, save the rest for appendix -->
<!-- What's the "min-viable" set of functionality for target user adoption? -->
<!-- Group by use case/user journey to enable Epic creation with SDLC naming: {epic_number}-epic-{slug}.md -->

### P0 Requirements (Must Have for MVP)

{{p0_requirements}}

### P1 Requirements (High Value, Important Additions)

{{p1_requirements}}

### P2 Requirements (Nice to Have)

{{p2_requirements}}

<!-- Template Guidance:
Format: Focus on functionality, not implementation
✅ DO: "First-time user must accept privacy policy to use product"
✅ DO: "Product team can monitor and visualize user engagement"
✅ DO: Link to UX sketches for quick visualization
✅ DO: Bucket by use case/user journey for Epic creation

❌ DON'T: Performance metrics unless required for adoption
❌ DON'T: Design details like "blue 'Continue' button"
❌ DON'T: Technical implementation specifics

Priority Examples:
P0: User can login using biometric authentication with <3 second response
P1: User can view login history with timestamps and device info
P2: Admin can configure password complexity requirements

Use Case Buckets for Epic Creation:
### Login & Authentication (Epic: 01-epic-login-authentication.md)
- User completes biometric setup in <2 minutes
- User authenticates using fingerprint/face recognition
- System provides clear error messages for failed authentication

### Account Management (Epic: 02-epic-account-management.md)
- User can view and manage connected devices
- User can revoke access for specific devices

Each bucket should map to an Epic following SDLC naming: {epic_number}-epic-{slug}.md
-->

---

## 7. Appendix

<!-- Links to materials people will ask for -->
<!-- These don't matter if you can't align on problem/solution first -->

**Supporting Documents:**

- {{supporting_doc_1}}
- {{supporting_doc_2}}
- {{supporting_doc_3}}

**SDLC Framework Information:**

- **Dependencies**: Project Brief at `/docs/prd/project-brief.md` (EXACT path)
- **Output Location**: This PRD MUST be saved as `/docs/prd/prd.md` (EXACT path)
- **Registry Update**: Add artifact entry to `/docs/registry.json`
- **Downstream Enablement**:
  - Epic creation in `/docs/epics/` using naming format `{epic_number}-epic-{slug}.md`
  - Architecture documents in `/docs/architecture/`
  - User stories in `/docs/stories/` using format `{epic_number}.{story_number}.story.md`

<!-- Template Guidance:
Supporting Documents Example:
- [Detailed UX flows and wireframes](link)
- [Go-to-market strategy and pricing](link)
- [Technical architecture decisions](link)
- [User research and competitive analysis](link)
- [Risk assessment and pre-mortem](link)

SDLC Framework Directory Structure:
/docs/
├── registry.json                 # Central artifact registry
├── prd/                          # Product vision & requirements
│   ├── project-brief.md          # Project vision & strategy (INPUT)
│   └── prd.md                    # Product requirements (THIS FILE)
├── epics/                        # High-level features ({epic_number}-epic-{slug}.md)
├── stories/                      # User stories ({epic_number}.{story_number}.story.md)
├── architecture/                 # System design
└── tests/                        # Quality validation

Keep SDLC framework details minimal and in appendix to maintain focus on product requirements.
-->

---

<!-- QUALITY CHECKLIST
✅ Document is 6-8 pages maximum
✅ Every requirement maps to user value
✅ Problem statement avoids solution-oriented language
✅ Success criteria are specific and testable
✅ Clear MVP boundaries (P0 vs P1/P2)
✅ Requirements grouped by use case for Epic creation
✅ SDLC framework compliance: exact file paths and naming conventions
✅ File saved exactly as /docs/prd/prd.md
✅ Registry updated at /docs/registry.json
✅ Epic naming ready: {epic_number}-epic-{slug}.md format
✅ Story naming ready: {epic_number}.{story_number}.story.md format
-->
