# Epic 9: Dogfooding KubeRocketAI Across KubeRocketCI Repositories

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | Planning                 |
| Priority             | High                     |
| Epic Owner           | Product Owner            |
| Timeline             | Q3 2025                  |
| Epic Key/Link        | Epic 9                   |

## Overview

### Problem Statement

KubeRocketAI is not yet consistently adopted across KubeRocketCI repositories. Teams lack a repeatable, local-first way to use agents in-day, provide feedback, and converge on shared best practices. Without dogfooding at scale, we risk slow adoption, duplicated customizations, and weak signals for which assets belong in the Framework versus repo‑local customization.

### Goal

Enable local KubeRocketAI daily usage in prioritized repositories (Wave 1) by Q3 2025; leverage Epic 8 selective installation to minimize footprint and accelerate adoption.

### Target Users

Primary: Repo Maintainers and Development Leads — integrate `.krci-ai` and use local agents day‑to‑day
Secondary: Platform Engineers — evaluate what belongs in Framework vs. local
Tertiary: Product/QA — validate usability, completeness, and quality signals

## Scope

### What's Included

1. Seed minimal `.krci-ai` with local agents in target repositories under heavy development
2. Add a short Quickstart to each repo's README (install, validate, run agents)
3. Ensure `krci-ai validate` and install flows work per repo
4. Track adoption and blockers via issues labeled `krci-ai`
5. Use Epic 8 selective installation to install only needed components per repo
6. Evaluate and document new agent patterns discovered during dogfooding:
   - Technical Writer agent for documentation repositories
   - DevOps agent (composite of Software Architect + Software Developer) for infrastructure work

### What's Not Included

1. Broad globalization of assets before feedback (we start local‑first)

### Dependencies

Epic Dependencies:

- Epic 8: Selective Installation (must be available to scope installs)
- Epic 6: Local Agent Components (baseline local agent model)
- Epic 2: Core Engine (validation/processing)

System Dependencies:

- KubeRocketAI CLI (`krci-ai`) and command framework
- GitHub access to target repositories
- Filesystem-based project layout with `.krci-ai` assets

External Dependencies:

- Team availability in target repos to review and accept PRs
- Security/compliance guidance for shared/global assets

## Solution Approach

Implementation Strategy:

1. Start local‑first: seed minimal `.krci-ai` in target repos using Framework defaults; keep customizations under `.krci-ai/local`
2. Use selective installation (Epic 8) to install only required components per repo
3. Add Quickstart README updates and a minimal usage guide
4. Open a `krci-ai` labeled issue per repo to capture blockers and usage notes

Technical Approach:

- Minimal baseline per repo: agents (PO/Dev/QA/Architect), core tasks, required templates/data; validate with CLI
- Document "How to align your repo with KubeRocketAI" in `KubeRocketCI/docs`
- Automate selective install in CI where feasible (non‑blocking for initial rollout)

## Risks & Assumptions

### Key Risks

- **Risk**: Team availability in target repositories may delay PR review and acceptance, affecting Q3 2025 timeline
  - Mitigation: Prioritize high-activity repos and coordinate with maintainers early
- **Risk**: Epic 8 (Selective Installation) delays could block dogfooding rollout
  - Mitigation: Prepare fallback with full installation if selective install isn't ready
- **Risk**: Repository-specific customizations may reveal gaps in Framework defaults
  - Mitigation: Track customization patterns via GitHub issues to inform future Framework updates

### Key Assumptions

- **Assumption**: Epic 8 (Selective Installation) will be available for scoped installs by dogfooding start
- **Assumption**: Local `.krci-ai` directory structure supports both Framework defaults and local customizations
- **Assumption**: Team feedback collection via GitHub issues will provide sufficient signal for Framework evolution
- **Assumption**: CLI validation commands work consistently across different repository structures

## Acceptance Criteria

1. **Repository agent setup and validation**
   - Scenario: Product Owner or Development Lead sets up KubeRocketAI in target repositories using minimal baseline configuration
   - Expected Behavior: Target repositories contain functional `.krci-ai` directory with appropriate agents, core tasks, and templates; CLI validation confirms setup works correctly
   - Measurement/Method: Verify directory structure and agent configuration; run CLI validation to confirm successful setup
   - Preconditions/Assumptions: Epic 8 (Selective Installation) available; target repositories identified and accessible; Framework defaults compatible with repo structures
   - Guardrails: No critical defects in agent activation flows; validation completes in reasonable time
   - Traceability: Epic 6 (Local Agent Components), Epic 2 (Core Engine)

2. **Selective installation adoption**
   - Scenario: Repository maintainer installs only required KubeRocketAI components using selective installation to minimize footprint
   - Expected Behavior: Installation process completes with appropriate subset of components relevant to repository type; documentation accurately reflects selected components
   - Measurement/Method: Compare installed components against repository needs; verify documentation matches actual installation; confirm significantly reduced footprint vs full installation
   - Preconditions/Assumptions: Epic 8 (Selective Installation) functional; component selection logic works for different repository types (frontend/backend/DevOps)
   - Guardrails: Installation time noticeably reduced vs full installation; no missing critical dependencies for selected components
   - Traceability: Epic 8 (Selective Installation)

3. **Adoption feedback and tracking**
   - Scenario: Development teams use KubeRocketAI agents daily and provide feedback on blockers, usability, and customization needs
   - Expected Behavior: Structured feedback captured via GitHub issues labeled `krci-ai`; patterns emerge indicating Framework gaps and local customization requirements
   - Measurement/Method: Review GitHub issues in target repositories; categorize feedback into Framework enhancement vs local customization; track resolution patterns
   - Preconditions/Assumptions: Teams actively using agents in daily work; GitHub issue templates available for structured feedback; maintainers responsive to feedback collection
   - Guardrails: Regular actionable feedback per target repository; reasonable response time to critical blockers
   - Deferred (non-blocking): Long-term adoption metrics and usage analytics (handled by future analytics Epic)
   - Traceability: Epic 6 (Local Agent Components)

4. **New agent pattern validation**
   - Scenario: Product team evaluates effectiveness of locally-created Technical Writer and DevOps composite agent patterns discovered during dogfooding
   - Expected Behavior: Technical Writer agent demonstrates clear value for documentation work; DevOps pattern proves effective across infrastructure repositories; framework inclusion recommendations produced with clear rationale
   - Measurement/Method: Collect usage feedback specifically for Technical Writer and DevOps patterns; analyze effectiveness compared to existing agents; document benefits and implementation recommendations
   - Preconditions/Assumptions: Technical Writer agent implemented in docs repository; DevOps pattern used across multiple infrastructure repositories; sufficient usage time for evaluation
   - Guardrails: Clear evaluation criteria established; evidence-based improvement vs existing patterns; framework impact assessment completed
   - Traceability: Epic 6 (Local Agent Components), Framework Enhancement Roadmap

## User Stories

- **Story 9.01: Frontend Developer agent setup**
  - As a Frontend Developer, I want local KubeRocketAI agents available in my frontend repository, so that I can get development assistance for UI/UX implementation tasks
  - Acceptance Criteria: `.krci-ai` directory contains Software Developer agent with frontend-specific tasks; CLI validation passes; README includes Quickstart section
  - Dependencies: Epic 8 (Selective Installation), Epic 6 (Local Agent Components)

- **Story 9.02: Backend Developer agent setup**
  - As a Backend Developer, I want local KubeRocketAI agents available in my backend repository, so that I can get development assistance and architectural guidance for API design and server-side implementation
  - Acceptance Criteria: `.krci-ai` directory contains Software Developer/Software Architect agents with backend-specific tasks; CLI validation passes; README includes Quickstart section
  - Dependencies: Epic 8 (Selective Installation), Epic 6 (Local Agent Components)

- **Story 9.03: QA Engineer agent setup**
  - As a QA Engineer, I want local KubeRocketAI agents available in my testing repositories, so that I can get quality assurance support for test planning, execution, and defect reporting
  - Acceptance Criteria: `.krci-ai` directory contains QA Engineer agent with testing-focused tasks and templates; CLI validation passes; README includes Quickstart section
  - Dependencies: Epic 8 (Selective Installation), Epic 6 (Local Agent Components)

- **Story 9.04: QA UI/UX specialist agent setup**
  - As a QA UI/UX specialist, I want local KubeRocketAI agents available for frontend testing, so that I can get specialized support for UI testing, accessibility validation, and user experience quality assurance
  - Acceptance Criteria: `.krci-ai` directory contains QA Engineer/Software Developer agents with UI/UX testing tasks; CLI validation passes; README includes Quickstart section
  - Dependencies: Epic 8 (Selective Installation), Epic 6 (Local Agent Components)

- **Story 9.05: DevOps Engineer agent setup**
  - As a DevOps Engineer, I want local KubeRocketAI agents available in my infrastructure repositories, so that I can get architectural guidance and development support for infrastructure automation and deployment workflows
  - Acceptance Criteria: `.krci-ai` directory contains Software Architect/Software Developer agents with DevOps-focused tasks; CLI validation passes; README includes Quickstart section
  - Dependencies: Epic 8 (Selective Installation), Epic 6 (Local Agent Components)

- **Story 9.06: Software Architect agent setup**
  - As a Software Architect, I want local KubeRocketAI agents available across architecture repositories, so that I can get comprehensive architectural guidance, design patterns, and system design support
  - Acceptance Criteria: `.krci-ai` directory contains Software Architect agent with architecture-focused tasks and templates; CLI validation passes; README includes Quickstart section
  - Dependencies: Epic 8 (Selective Installation), Epic 6 (Local Agent Components)

- **Story 9.07: Product Owner agent setup**
  - As a Product Owner, I want local KubeRocketAI agents available for product management work, so that I can get support for user story creation, backlog management, and requirements gathering
  - Acceptance Criteria: `.krci-ai` directory contains Product Owner agent with product management tasks; CLI validation passes; README includes Quickstart section
  - Dependencies: Epic 8 (Selective Installation), Epic 6 (Local Agent Components)

- **Story 9.08: Product Manager strategic analysis**
  - As a Product Manager, I want local KubeRocketAI agents available for strategic work, so that I can analyze dogfooding feedback patterns and create Framework optimization recommendations
  - Acceptance Criteria: `.krci-ai` directory contains Product Manager/Business Analyst agents; analysis reports identify common patterns; roadmap updated with Framework enhancement items
  - Dependencies: Epic 8 (Selective Installation), Stories 9.01-9.07 completion, minimum 4 weeks of usage feedback

- **Story 9.09: New agent pattern evaluation**
  - As a Product Manager, I want to evaluate the effectiveness of newly discovered agent patterns (Technical Writer and DevOps composite agent), so that I can recommend their inclusion in the Framework for broader adoption
  - Acceptance Criteria: Technical Writer agent effectiveness documented from docs repository usage; DevOps agent pattern validated across infrastructure repositories; framework inclusion recommendations provided with usage metrics and benefits analysis
  - Dependencies: Stories 9.01-9.08 completion, minimum 6 weeks of dogfooding feedback, active usage of local Technical Writer and DevOps patterns

## Target Repositories (initial candidate set)

KubeRocketCI GitHub repositories:

- KubeRocketCI/docs
- KubeRocketCI/krci-portal
- KubeRocketCI/krci-cache
- KubeRocketCI/gitfusion
- KubeRocketCI/terraform-aws-platform
- KubeRocketCI/tekton-custom-task
- epam/edp-install
- epam/edp-keycloak-operator
- epam/edp-sonar-operator
- epam/edp-nexus-operator
- epam/edp-tekton
- epam/edp-cluster-add-ons
- epam/edp-headlamp
- epam/edp-cd-pipeline-operator
- epam/edp-codebase-operator
