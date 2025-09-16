---
  templates:
    - testing-readme.md
---

# Task: Onboard Existing Gherkin Tests

## Description

Analyze an existing directory of Gherkin `.feature` files and generate a `src/main/resources/README.md` aligned with repository conventions. Onboarding is non-interactive: it performs analysis and writes the README, then reports how to change settings via the edit command.

## When to Use

- You already have `.feature` files but lack `src/main/resources/README.md`
- You are migrating an external Gherkin test suite into this system

## Inputs

- Source directory containing `.feature` files (default: `src/main/resources/features/`)

## Validation (HALT if missing)

- Source directory exists and contains at least one `.feature` file
- `src/main/resources/README.md` does not already exist (otherwise use `edit-testing-settings`)
- Write permissions for `src/main/resources/`

## Prechecks and Routing (HALT)

- If `src/main/resources/README.md` already exists → HALT and propose running `edit-testing-settings` to modify settings interactively.
- If no `.feature` files are present in the source directory → HALT and propose `setup-testing` to initialize the structure and README.
- If both `.feature` files and README are present → this flow is not applicable; use `edit-testing-settings` or `generate-test-cases` as needed.

## Flow (Non-interactive)

1) Scan and infer domains, UI/API presence, naming convention, tag families (with frequencies and scopes), discovery hints (preferred directories, topics, artifacts).
2) Create `src/main/resources/` if missing and write `src/main/resources/README.md` from [template](./.krci-ai/templates/testing-readme.md) using inferred values and hints.
3) Enrich the README with analysis outputs:
   - Replace the template "Directory structure" code block with the actual tree under `src/main/resources/features/` (preserve readability by limiting depth if very large).
   - Set the naming convention to the detected style (`PascalCase`, `kebab-case`, or `snake_case`).
   - Insert a concise UI vs API coverage summary (presence and high-level distribution).
   - Add a "Current tags in this repository" subsection listing top tags with counts, grouped into recommended families; include an "Additional tags" note for unmapped tags.
   - Append a "Discovery hints for this repository" subsection with preferred directories order, common topics/keywords, and characteristic artifacts.
4) Do not move or rename existing `.feature` files. Do not build local indexes during onboarding.
5) After writing, print a short summary with the README path and instruct the user to run `edit-testing-settings` to adjust any section interactively.

## Safety

- Read-only analysis of `.feature` files; no file moves or renames.
- No local index build (semantic or lexical) is performed in this flow.

## Analysis Heuristics

During onboarding, the agent analyzes the feature tree to infer structure and conventions:

- Domains: inferred from folder names and common path segments (e.g., `payments`, `core`)
- Test types: detect UI (`features/<domain>/*.feature`) and API (`features/api_tests/<domain>/*.feature`) presence
- Naming convention: sample filenames to detect `PascalCase` vs `kebab-case` vs `snake_case`
- Tags: aggregate tags across Feature, Scenario/Scenario Outline, and Examples; record frequency and scope
- Topics/keywords: extract top tokens from filenames and scenario titles (after normalization/stopword removal)
- Characteristic artifacts: detect frequent constants and step markers (e.g., UPPER_CASE tokens, URLs, IDs)
- Directory priority: rank subdirectories by scenario density and tag richness to suggest search order

### Tag Inference and Mapping

Extraction:
- Collect all tokens starting with `@` across Feature, Scenario, and Examples blocks
- Track per-tag frequency and where it appears (scope: Feature | Scenario | Examples)

Normalization:
- Normalize domain to PascalCase for domain-prefixed tags (e.g., `payments` → `Payments`)

Classification (recommended families):
- Type: `@UI`, `@API`, `@E2E`, `@Integration`, `@Unit`
- Scope/Suite: `@Smoke`, `@Regression`, `@ShortRegression`, `@Critical`, `@Negative`
- Non-functional: `@Performance`, `@Security`, `@Accessibility`, `@Compatibility`
- Lifecycle/Utilities: `@Cleanup`, `@DataSetup`, `@Migration`, `@Flaky`
- Optional domain/subdomain tags:
  - Domain: `@{{DomainPascal}}`
  - Subdomains: `@{{DomainPascal}}Create`, `@{{DomainPascal}}Deploy`, `@{{DomainPascal}}Promote` (and similar)

Unknown/unmapped tags:
- Keep a list of tags that don’t map cleanly; include them under an "Additional tags" note in the README so they are documented without enforcing new conventions.

### Discovery/Search Heuristics (repo-specific)

To help future search and extension:
- Preferred directories (ordered): directories ranked by feature density and tag concentration
- Common topics/keywords: top N normalized tokens from filenames and scenario titles
- Characteristic artifacts/steps: frequently recurring constants and step phrases worth using as anchors
- Parity patterns (optional): if parallel provider/module trees exist, note parity expectations

These hints are summarized into the README under a "Discovery hints for this repository" section when applicable.

## README Content (auto-generated)

- Directory structure: reflects the existing on-disk tree (including `api_tests/`, `utility/cleanup` if present)
- Tagging system: recommended families populated with inferred tags; include domain/subdomain tags if detected
- How to add your existing tests: guidance aligned to current placement and detected naming convention
- Discovery and search workflow: standard workflow plus repo-specific discovery hints
- Integration with the QA agent: inputs/outputs as defined by the template
- Additional tags: unmapped tags list (optional)

## Output

- `src/main/resources/README.md` reflecting existing suite structure and inferred conventions
- Console summary with the README path and next-step instruction (`edit-testing-settings`)

## Success Criteria

- README reflects actual on-disk feature layout without prompting the user
- Provider/tags guidance is accurate and actionable for subsequent generation
- Discovery hints help the agent prioritize relevant areas during search
- Further refinements are handled by `edit-testing-settings`
