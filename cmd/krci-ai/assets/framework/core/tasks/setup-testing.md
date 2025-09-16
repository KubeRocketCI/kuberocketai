---
dependencies:
  templates:
    - testing-readme.md
---

# Task: Setup Testing (First-Run Wizard)

## Description

Interactive, first-run setup that initializes your testing workspace. This wizard asks a short series of questions and then generates a testing README and directory structure for Gherkin-based tests.

## Goal

- Create `src/main/resources/README.md` from a template, filled with your choices
- Create `src/main/resources/features/` directory structure (including optional provider subfolders)
- Establish Gherkin as the single source of truth for test cases

## When to Use

- On a fresh installation where `src/main/resources/README.md` or `src/main/resources/features/` is missing
- When reconfiguring your testing strategy, providers, or tagging conventions

## Validation (HALT if missing)

- Write permissions for `src/main/resources/`

## Interactive Questions (Wizard)

Ask the following in sequence. After each answer, confirm or offer a recommended default.

1) Domain name for UI/API test features (e.g., `core`, `payments`)
   - Example: `core`, `payments`, `orders`

2) Test types to include (multi-select; default: `UI, API`)
   - Allowed values: `UI`, `API`

3) Tagging approach
   - Default: provider-agnostic families (Type, Scope/Suite, Non-functional, Lifecycle)
   - Optional: domain-specific tags (free-form), e.g., `@Payments`, `@Orders`

### Recommended default tags

Provider-agnostic families (recommended for most projects):

- Type: `@UI`, `@API`, `@E2E`, `@Integration`, `@Unit`
- Scope/Suite: `@Smoke`, `@Regression`, `@ShortRegression`, `@Critical`, `@Negative`
- Non-functional: `@Performance`, `@Security`, `@Accessibility`, `@Compatibility`
- Lifecycle/Utilities: `@Cleanup`, `@DataSetup`, `@Migration`, `@Flaky`

Optional domain/subdomain tags (enable only if relevant):
- Domain: `@{{DomainPascal}}`
- Subdomains (examples): `@{{DomainPascal}}Create`, `@{{DomainPascal}}Deploy`, `@{{DomainPascal}}Promote`

Wizard behavior:
- Offer provider-agnostic defaults first; only add domain/subdomain tags if enabled.
- Allow the user to accept defaults, enter custom tag families, or ask the agent to propose tags based on a brief description.

6) Include utilities/cleanup flows? (default: `yes`)

7) Naming convention for feature files (default: `PascalCase`, e.g., `PromoteApplication.feature`)

8) Create starter example feature file? (default: `no`)

### Wizard contract (HALT checkpoints)

- Before starting: HALT with a short intro explaining the wizard flow and expected outputs.
- For each question above: ask → HALT for the answer → echo back the interpreted value → HALT for confirmation or edits.
- Existing features: recommend the user to place `.feature` files directly into `src/main/resources/features/` (next to the future `src/main/resources/README.md`) with desired subfolders; alternatively offer to accept explicit paths or to scan typical paths (`docs/**`, `test/**`, `**/*.feature`) and HALT to confirm which to import and where they will be placed under `src/main/resources/features/`.
- Plan summary: after gathering all inputs, present a concise plan (directories to create, README sections to include/customize, files to import/move) and HALT for approval.
- Diff/preview: before writing, show a short preview of the new `src/main/resources/README.md` (first 20–30 lines and any customized sections) plus the file/folder actions list, and HALT for final confirmation.
- On any "no" at confirmation points: allow the user to revise inputs or cancel without writing.
- Optional local index: offer to build a local Gherkin index only if at least one non-starter `.feature` file is present (imported or detected). If only starter example(s) were created by the wizard, skip this prompt entirely.

## Actions

1. Create directories:
   - `src/main/resources/`
   - `src/main/resources/features/`
   - `src/main/resources/features/{{domain_name}}/`
   - For each selected provider:
     - `src/main/resources/features/{{domain_name}}/{{provider}}/` (if UI selected)
     - `src/main/resources/features/api_tests/{{domain_name}}/{{provider}}/` (if API selected)
   - If utilities enabled: `src/main/resources/features/api_tests/utility/cleanup/`

2. Generate `src/main/resources/README.md` from [template](./.krci-ai/templates/testing-readme.md) by replacing placeholders:
   - `{{domain_name}}`
   - `{{providers_csv}}` (comma-separated)
   - `{{include_ui}}` / `{{include_api}}` → `true/false`
   - `{{enforce_parity}}` → `true/false`
   - `{{provider_ui_tags}}`, `{{api_suite_tags}}`, `{{suite_tags}}`
   - `{{naming_convention}}`

   Block-by-block prompting aligned to the template:
   - Tests Directory Structure: show the template section and ask whether to accept as-is or customize notes. HALT.
   - How the QA agent works with tests: show and optionally tailor to project wording. HALT.
   - Discovery and search workflow: show and confirm or edit. HALT.
   - Optional local Gherkin index: show and confirm whether to include. HALT.
   - Tagging system: show recommended families; incorporate user-chosen domain tags. HALT.
   - How to add your existing tests: include chosen naming convention and placement guidance. HALT.
   - Integration with the QA agent: confirm inputs/outputs reflect this project. HALT.

3. Optionally create a starter example feature file if requested.

## Output

- `src/main/resources/README.md` created and aligned with your selections
- Feature directories created for the selected domain and providers
- Optional starter `.feature` file created
- Optional local Gherkin indexes (if chosen):
  - Semantic FAISS index at `./.krci-ai/indexes/gherkin-faiss.index` with metadata `gherkin-faiss.meta.json`
  - Lexical JSON `./.krci-ai/indexes/gherkin-lex.json` and SQLite `./.krci-ai/indexes/gherkin-lex.sqlite`

## Success Criteria

- README exists with your chosen providers, tags, and structure
- Feature directories exist and match the selected test types (UI/API)
- Agent can operate in BDD-only mode using the created README and features

## Notes

- This setup aligns with the QA agent contract: Gherkin is the single source of truth, README is the governing document for process and tags.
- After setup, you can create local Gherkin indexes to speed up discovery; they are optional and can be regenerated anytime.
- After setup, run the discovery/search phase (see generate task) to index existing or starter features and validate provider/tag structure.
- You can rerun this wizard later to adjust configuration; review diffs before committing changes.
