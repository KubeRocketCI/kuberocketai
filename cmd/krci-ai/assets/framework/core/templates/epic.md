# Epic {{epic_number}}: {{epic_title}}

## Status

<status>
| Field                | Value                    |
|----------------------|--------------------------|
| Status               | {{status}}               |
| Priority             | {{priority}}            |
| Epic Owner           | {{owner}}               |
| Timeline             | {{timeline}}            |
| Epic Key/Link        | {{epic_key}}            |

<instructions>
Epic tracking and SDLC integration. Enables progress monitoring and PRD traceability validation.

Status Options: Planning -> Approved -> In Progress -> Done -> Completed
Priority Example: Critical, High, Medium, Low (align with PRD requirement priorities)
Epic Owner Example: Product Owner, Tech Lead, Development Team
Timeline Example: Sprint 1-3 (6 weeks), Q1 2025, March-April 2025

CRITICAL: Status section contains ONLY these fields. Do NOT add Dependencies or other fields here.
</instructions>
</status>

## Overview

### Problem Statement

<problem_statement>
{{problem_statement}}

<instructions>
Clear, specific user or business problem this Epic solves from PRD analysis. Must trace back to PRD business requirements and user pain points.

Problem Statement Example: Users struggle to activate and use agents consistently across supported IDEs, causing rework and reduced productivity. Inconsistent activation flows generate support tickets and block onboarding.

Format Structure:
- Start with specific user pain point from PRD
- Include quantifiable impact or evidence (if already available)
- Connect to PRD BR/NFR requirements
- Avoid solution-oriented language

DO:
- Tie the problem to concrete PRD BR/NFR references
- Use user-centric language and observable effects
- Include evidence or impact where available

DONT:
- Prescribe solutions or technical designs here
- Use vague terms without context
- Omit explicit traceability to PRD
</instructions>
</problem_statement>

### Goal

<goal>
{{goal}}

<instructions>
Capability delivered statement that defines Epic completion. Be outcome-focused and verifiable immediately after implementation in a controlled environment.

Goal Examples:
- Deliver a consistent agent activation capability across target IDEs with a single, predictable user flow
- Provide single sign-on capability across portal and tools using an internal IDP/test stub for validation

Notes:
- You may optionally link to OKRs if they are already active and verifiable during/after the release
- Avoid speculative future metrics and long-horizon targets that cannot be checked immediately after delivery

DO:
- State the capability in plain language (outcome, not implementation)
- Include near-term indicators only if verifiable post-implementation
- Align with the Problem Statement and PRD scope

DONT:
- Force %/timeline when measurement isn't feasible yet
- Embed tool-specific testing details
- Drift into story-level scope
</instructions>
</goal>

### Target Users

<target_users>
{{target_users}}

<instructions>
Specific user personas from PRD who benefit from this Epic. Must align with PRD user segments and enable Story As a user scenarios.

Target Users Example: Primary: Software Engineers and Architects – installing and activating agents in supported IDEs, Secondary: QA Engineers – validating agent behavior across environments, Tertiary: Product Managers – tracking readiness for release

Format Structure:
- List primary, secondary, tertiary users and their context
- Connect to PRD user segments and persona definitions

DO:
- Use PRD-defined personas and segments
- Provide enough context for story persona scenarios
- Keep the list focused on actual beneficiaries

DONT:
- Invent new roles not present in PRD
- Use generic labels like users or developers without context
- Omit the user context and responsibilities
</instructions>
</target_users>

## Scope

### What's Included

<in_scope>
{{in_scope}}

<instructions>
Specific features and functionality this Epic delivers. Must map to PRD requirements and enable Story breakdown.

Whats Included Example:
1. Unified activation flow across Cursor, VS Code, and JetBrains IDEs (BR2, BR3)
2. Session continuity within a workday across integrated tools (BR5)
3. Basic error handling and user messaging for failed activations (NFR-UX1)
4. Local test IDP/stubs or mocks where external dependencies would otherwise be required (NFR-TEST1)

Format Structure:
- Number items for clear tracking
- Reference specific PRD requirements in parentheses
- Focus on user-facing functionality and immediately testable outcomes

DO:
- Keep items independently verifiable after implementation
- Reference BR/NFR for each included capability
- Prefer outcomes over implementation detail

DONT:
- Include low-level design or non-essential technical tasks
- Add third-party dependencies that block local validation
- Blur boundaries with Whats Not Included
</instructions>
</in_scope>

### What's Not Included

<out_of_scope>
{{out_of_scope}}

<instructions>
Clear boundaries of what this Epic excludes to prevent scope creep.

Whats Not Included Example:
1. Production analytics dashboards for adoption metrics (handled in separate analytics Epic)
2. Advanced SSO federation across multiple enterprise providers (future roadmap)
3. Legacy IDE support beyond current LTS versions (deferred)

DO:
- Explain rationale (deferred, out of scope, future epic)
- Point to related epics when applicable
- Protect the MVP boundary

DONT:
- List ambiguous exclusions without justification
- Duplicate items that already appear in scope
- Use exclusions to mask undecided scope
</instructions>
</out_of_scope>

### Dependencies

<dependencies>
{{dependencies}}

<instructions>
Other Epics, systems, or external requirements this Epic needs. Design acceptance criteria so they can be verified without reliance on third-party services; if unavoidable, specify stubs/mocks.

Dependencies Example:
Epic Dependencies:
- Epic 1: Baseline infrastructure
- Epic 2: Core engine enhancements

System/Test Dependencies:
- Local IDP stub for SSO validation
- Supported IDE versions (latest stable) available in CI/staging

External Dependencies (if any):
- Security review sign-off (if policy requires before release)

DO:
- Call out stubs/mocks to avoid third-party blockers
- Specify versions/constraints that affect validation
- Separate epic/system/external dependencies for clarity

DONT:
- Depend on production-only services for acceptance
- Leave approvals/integration needs implicit
- Use vague placeholders
</instructions>
</dependencies>

## Solution Approach

<solution_approach>
{{solution_approach}}

<instructions>
High-level implementation strategy and architectural direction. Guides Story creation without prescribing detailed implementation.

Keep this section at architectural and capability level:
- Key integration points and boundaries
- Feature flags/toggles for safe rollout
- Use of stubs or test doubles to decouple from third-party services during validation

DO:
- Describe integration boundaries and toggles for safe rollout
- Prefer approaches enabling immediate post-implementation verification
- Note key risks and how validation addresses them

DONT:
- Specify tool commands or low-level implementation detail
- Overconstrain design choices prematurely
- Ignore rollout/operational considerations
</instructions>
</solution_approach>

## Risks & Assumptions

<risks_and_assumptions>
{{risks_and_assumptions}}

<instructions>
Key uncertainties that may impact delivery or validation; distinct from Solution Approach. Capture assumptions that, if false, would invalidate scope or acceptance.

Risks & Assumptions Examples:
- Risk: Supported IDE API changes could break activation flow near release
- Risk: Security approval needed before enabling cross-tool session persistence
- Assumption: Local IDP stub is available for SSO validation
- Assumption: Feature flags can be toggled per environment

DO:
- State risks with potential impact and a mitigation idea
- Make assumptions explicit and testable
- Keep items tied to epic scope and acceptance

DONT:
- Duplicate dependencies; reference them if needed
- List generic project risks unrelated to acceptance
- Leave critical assumptions implicit
</instructions>
</risks_and_assumptions>

## Acceptance Criteria

<acceptance_criteria>
{{acceptance_criteria}}

<instructions>
Epic-level, outcome-focused, and immediately verifiable post-implementation. Do NOT prescribe verification commands or specific tools. Avoid dependence on third-party services; if required, define local stubs/mocks.

Required Structure for each criterion:
1. Scenario: Brief user/system flow to validate
2. Expected Behavior: Clear, observable outcome
3. Measurement/Method: How we confirm now (manual steps, UI/API observation, logs), without prescribing tools/commands
4. Preconditions/Assumptions: Feature flags, environment, stubs/mocks, test users
5. Guardrails: Quality constraints (e.g., no P0/P1 defects in core flow; basic performance/security thresholds)
6. Deferred (non-blocking): Optional indicator that cannot be immediately verified (use sparingly)

Note: Deferred (non-blocking) items do not gate epic closure. Track them in analytics/ops epics or release notes.

Acceptance Criteria Example:
1. Cross-IDE agent activation consistency
   - Scenario: User activates a selected agent in Cursor, VS Code, and JetBrains using the unified flow
   - Expected Behavior: Activation completes successfully with consistent steps and end state across IDEs
   - Measurement/Method: Perform activation in staging across supported IDEs; verify final activated state and absence of additional prompts
   - Preconditions/Assumptions: Supported IDE versions installed; feature flag unifiedActivation enabled
   - Guardrails: No P0/P1 defects in activation core path; added startup latency <= 200ms
   - Traceability: BR2, BR3

DO:
- Number criteria for tracking and story mapping
- Keep outcomes measurable/observable now, without requiring external analytics
- Use stubs/mocks to validate flows when real third-party services are unavailable

DONT:
- Include command-level verification
- Rely on long-horizon metrics or OKRs that cannot be verified immediately after delivery
- Use subjective language like works well or user-friendly without observable outcomes
</instructions>
</acceptance_criteria>

## User Stories

<user_stories>
{{planned_stories}}

<instructions>
Planned Stories that implement this Epic with clear breakdown and traceability.

Recommended Structure:
- Group stories by phases or slices that deliver incremental user value
- Use consistent numbering within the Epic
- Each story should include persona, goal, and minimal acceptance criteria

DO:
- Keep stories INVEST and traceable to epic criteria
- Include dependencies only when necessary and explicit
- Ensure each story yields observable value

DONT:
- Mix multiple personas/goals in one story
- Leave stories without acceptance criteria
- Skip numbering or phase grouping without reason

Example:
Phase 1: Foundation
- Story epic_number.01: Unified activation flow – Cursor
  - As a Software Engineer, I want to activate an agent in Cursor, so that I can use it without extra setup
- Story epic_number.02: Unified activation flow – VS Code

Phase 2: Parity
- Story epic_number.03: Unified activation flow – JetBrains
- Story epic_number.04: Session continuity (intra-day)

CRITICAL: All sections must appear in exact order - Status, Overview, Scope, Solution Approach, Risks & Assumptions, Acceptance Criteria, User Stories.
</instructions>
</user_stories>
