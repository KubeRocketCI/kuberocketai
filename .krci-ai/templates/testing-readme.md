# Tests Directory Structure

This directory contains all project tests written in Gherkin (BDD).

## Directory structure

```
docs/testing/
├── features/                           # Gherkin feature files (.feature)
│   ├── {Topic}.feature                 # UI/functional features
│   └── api_tests/                      # optional: API features
│       └── {Topic}.feature
└── README.md                           # This file
```

Maintenance: when adding, removing, or renaming feature files, update the Directory structure above in the same pull request so this document remains the single source of truth.

## How the QA agent works with tests

1. Analyze existing coverage under `docs/testing/features/` using filename and scenario/title/tag heuristics
2. Decide on a testing strategy using the decision matrix (Covered / Partial / Not covered)
3. Generate or extend Gherkin tests with provider parity and UI/API alignment

### Discovery and search workflow

- Build normalized keyword variants from story terms (hyphen/underscore/space/camelCase; include relevant action/result modifiers when applicable).
- Detect domain/context hints (e.g., UI vs API, module/subsystem names) and prioritize likely paths observed in the repository structure.
- Search across scenario titles, step lines, tags, and Examples; include characteristic artifacts discovered in this repository (avoid vendor-specific markers in the template).
- Present top candidates (path + snippet) for extension before proposing new files.

### Optional local Gherkin index

- You may maintain a small JSON index at `./.krci-ai/indexes/gherkin-index.json` to speed up discovery.
- Suggested fields per scenario entry: file path, feature title, scenario title, tags, step keywords, artifact constants, Examples headers.
- The agent will use the index if present; it can be regenerated at any time during onboarding or generation.

### Tagging system

Core tag families:
- Type: `@UI`, `@API`, `@E2E`, `@Integration`, `@Unit`
- Scope/suite: `@Smoke`, `@Regression`, `@ShortRegression`, `@Critical`, `@Negative`
- Non-functional: `@Performance`, `@Security`, `@Accessibility`, `@Compatibility`
- Lifecycle/maintenance: `@Cleanup`, `@DataSetup`, `@Migration`, `@Flaky`

Tag scoping rules:
- Tags above Feature apply to all scenarios
- Tags above a Scenario/Scenario Outline apply to that scenario
- Tags directly above an Examples block apply only to that Examples table

Recommended usage:
- Always include at least one type tag (`@UI` or `@API`) for filterability
- Add suite/run-scope tags to control selection windows in CI (`@Smoke`, `@ShortRegression`)
- Add subdomain tags for specific flows (e.g., create/deploy/promote) when needed
- Use lifecycle/maintenance tags for setup/cleanup and test health

## How to add your existing tests

1. Placement (choose the correct path):
   - UI/functional features: `docs/testing/features/{Topic}.feature`
   - API features (optional folder): `docs/testing/features/api_tests/{Topic}.feature`
   - Utilities/maintenance (optional): `docs/testing/features/utility/...`

2. Naming convention: {{naming_convention}}

3. Structure guidelines:
   - Prefer Scenario Outline with one or more Examples tables when variants exist
   - Place tags directly above each Examples block when provider/scenario variants differ
   - Use rich parameter tables in steps; long multi-stage flows are acceptable

4. UI/API alignment: keep core flows covered consistently across UI and API

## Integration with the QA agent

Inputs:
- Stories with acceptance criteria
- Optional scope: providers, test type (UI/API), priority/tags

Outputs:
- Coverage report with paths and tags
- Missing providers and UI/API alignment hints
- Proposed next steps (add scenarios vs create new file)



