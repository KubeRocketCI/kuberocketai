# Task: Edit Testing Settings (Interactive Wizard)

## Description

Interactively edit testing settings stored in `src/main/resources/README.md`. The wizard shows current content for each section and asks whether to change it, add new data, or skip. Changes are applied in-place with a backup created automatically.

## When to Use

- You already have `src/main/resources/README.md` and want to adjust providers, tags, structure, or guidance
- After onboarding or setup, to refine details without regenerating from scratch

## Prerequisites (HALT if missing)

- `src/main/resources/README.md` must exist
- Write permissions for `src/main/resources/`

If README is missing:
- If `src/main/resources/features/` exists → propose running `onboard-testing`
- If both missing → propose `setup-testing`

## Modes

Choose one of:
- Add new data: append sections/notes without altering existing content
- Edit specific section: pick a section by name and update only it
- Guided edit: iterate all known sections; for each, preview current content and ask to edit or skip

## Known Sections (detected and editable)

- Tests Directory Structure
- How the QA agent works with tests
- Generating Gherkin tests
- Discovery and search workflow
- Local Gherkin index (optional)
- Tagging system
- How to add your existing tests
- Integration with the QA agent
- Maintenance notes (e.g., update this README when structure changes)

The wizard should detect headings by `##`/`###` and operate on those blocks. If a section is not found, offer to create it.

## Interaction Flow

1) Select mode (add, edit specific, guided)
2) Build Quick Section Index from current README headings (H2/H3):
   - Show numbered list like: `1) ## Tests Directory Structure`, `2) ### Tagging system`, ...
   - Only offer sections that actually exist; unknown sections can be created on demand
3) If edit specific:
   - Show list of detected sections with indices
   - Display current content of the chosen section
   - Ask: "Edit this section? [yes/no]"; if yes, accept new markdown text (multiline)
4) If guided:
   - For each section in Quick Section Index order (actual document order):
     - Show current content (trim to reasonable length if very long)
     - Ask: "Edit this section? [yes/no]"; if yes, accept new markdown text
     - Ask: "Continue to next section? [yes/no]"
5) If add new data:
   - Ask for a new section title and markdown body
   - Insert after a chosen anchor section or at the end

## Tag Recommendations (helper)

- Offer provider-agnostic defaults first:
  - Type: `@UI`, `@API`, `@E2E`, `@Integration`, `@Unit`
  - Scope/Suite: `@Smoke`, `@Regression`, `@ShortRegression`, `@Critical`, `@Negative`
  - Non-functional: `@Performance`, `@Security`, `@Accessibility`, `@Compatibility`
  - Lifecycle/Utilities: `@Cleanup`, `@DataSetup`, `@Migration`, `@Flaky`
- Optional domain/subdomain tags:
  - Domain: `@{{DomainPascal}}`
  - Subdomains: `@{{DomainPascal}}Create`, `@{{DomainPascal}}Deploy`, `@{{DomainPascal}}Promote`
- Present suggested lists and ask to insert/update in the Tagging system section
- Present suggested lists and ask to insert/update in the Tagging system section

## Safety and Backups

- Before any edit, create `src/main/resources/README.md.bak` (overwrite if exists)
- Validate that resulting README contains the edited headings and is non-empty

## Outputs

- Updated `src/main/resources/README.md`
- Optional: backup at `src/main/resources/README.md.bak`
- Summary of edited sections (titles) for the user

## Success Criteria

- Requested sections updated exactly as confirmed by the user
- README structure and headings preserved or improved
- Tagging guidance consistent with selected providers and domain


