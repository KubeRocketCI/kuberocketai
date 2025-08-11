# Epic 9: Dogfooding KubeRocketAI Across KubeRocketCI Repositories

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | Planning                 |
| Priority             | High                     |
| Epic Owner           | Product Owner            |
| Timeline             | Q3 2025                  |

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
2. Add a short Quickstart to each repo’s README (install, validate, run agents)
3. Ensure `krci-ai validate` and install flows work per repo
4. Track adoption and blockers via issues labeled `krci-ai`
5. Use Epic 8 selective installation to install only needed components per repo

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
- Document “How to align your repo with KubeRocketAI” in `KubeRocketCI/docs`
- Automate selective install in CI where feasible (non‑blocking for initial rollout)

## Acceptance Criteria

1. Local agents bootstrapped in Wave 1 repos and validated
   - Validation: Each repo contains a `.krci-ai` directory with agents, tasks, and templates; CLI validation passes
   - Command (run per repo): `krci-ai validate | cat`

2. Selective installation is used to install only required components per repo (Epic 8 dependency)
   - Validation: Install command installs a subset; repo docs capture selected set
   - Command (representative): `krci-ai install --ide=cursor | cat` (use `--all` to install all; add `--force` if re-installing)

3. Feedback captured and tracked
   - Validation: GitHub issues labeled `krci-ai` created for each repo with findings and proposals
   - Command (representative): link to issues list filtered by label

## User Stories

Planned Stories for Implementation:

- Story 09.01: Front-end repositories rollout
  - As a Maintainer, I want local agents available in front-end repos so the team can use them daily
  - Acceptance:
    - `.krci-ai` present with local agents
    - `krci-ai validate` passes
    - README Quickstart added (install, validate, basic usage)
    - Issue labeled `krci-ai` opened to track adoption/blockers
  - Dependencies: Epic 8

- Story 09.02: Back-end repositories rollout
  - As a Maintainer, I want local agents available in back-end repos so the team can use them daily
  - Acceptance:
    - `.krci-ai` present with local agents
    - `krci-ai validate` passes
    - README Quickstart added (install, validate, basic usage)
    - Issue labeled `krci-ai` opened to track adoption/blockers
  - Dependencies: Epic 8

- Story 09.03: DevOps repositories rollout
  - As a Platform Engineer, I want local agents available in DevOps repos so the team can use them daily
  - Acceptance:
    - `.krci-ai` present with local agents
    - `krci-ai validate` passes
    - README Quickstart added (install, validate, basic usage)
    - Issue labeled `krci-ai` opened to track adoption/blockers
  - Dependencies: Epic 8

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
