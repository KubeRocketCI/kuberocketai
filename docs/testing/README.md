# Tests Directory Structure

This directory contains all project tests written in Gherkin (BDD).

## Directory structure

```
docs/testing/
├── features/                           # 🧪 Gherkin feature files (.feature)
│   ├── tekton/
│   │   ├── AccountSettings.feature
│   │   ├── ConfigurationScreen.feature
│   │   ├── Marketplace.feature
│   │   ├── OverviewScreen.feature
│   │   ├── bitbucket/
│   │   │   ├── AddApplication.feature
│   │   │   ├── AddNewBranch.feature
│   │   │   ├── AutoDeploy.feature
│   │   │   ├── AutoDeployWithPromote.feature
│   │   │   ├── ComponentsMultiSelect.feature
│   │   │   ├── CreateDeployFlowWithYaml.feature
│   │   │   ├── DeployApplication.feature
│   │   │   ├── EditComponentPipelineStage.feature
│   │   │   ├── InputsValidation.feature
│   │   │   ├── JiraIntegration.feature
│   │   │   ├── LinksToTheComponents.feature
│   │   │   ├── PipelinesTable.feature
│   │   │   ├── PromoteApplication.feature
│   │   │   ├── RemoteCluster.feature
│   │   │   └── WidgetsIntegration.feature
│   │   ├── github/
│   │   │   ├── AddApplication.feature
│   │   │   ├── AddNewBranch.feature
│   │   │   ├── AutoDeploy.feature
│   │   │   ├── AutoDeployWithPromote.feature
│   │   │   ├── ComponentsMultiSelect.feature
│   │   │   ├── CreateDeployFlowWithYaml.feature
│   │   │   ├── DeployApplication.feature
│   │   │   ├── EditComponentPipelineStage.feature
│   │   │   ├── JiraIntegration.feature
│   │   │   ├── LinksToTheComponents.feature
│   │   │   ├── PipelinesTable.feature
│   │   │   ├── PromoteApplication.feature
│   │   │   ├── RemoteCluster.feature
│   │   │   └── WidgetsIntegration.feature
│   │   ├── gerrit/
│   │   │   ├── AddApplication.feature
│   │   │   ├── AddNewBranch.feature
│   │   │   ├── AutoDeploy.feature
│   │   │   ├── AutoDeployWithPromote.feature
│   │   │   ├── ComponentsMultiSelect.feature
│   │   │   ├── CreateDeployFlowWithYaml.feature
│   │   │   ├── DeployApplication.feature
│   │   │   ├── EditComponentPipelineStage.feature
│   │   │   ├── JiraIntegration.feature
│   │   │   ├── LinksToTheComponents.feature
│   │   │   ├── PipelinesTable.feature
│   │   │   ├── PromoteApplication.feature
│   │   │   ├── RemoteCluster.feature
│   │   │   └── WidgetsIntegration.feature
│   │   └── gitlab/
│   │       ├── AddApplication.feature
│   │       ├── AddNewBranch.feature
│   │       ├── AutoDeploy.feature
│   │       ├── AutoDeployWithPromote.feature
│   │       ├── ComponentsMultiSelect.feature
│   │       ├── CreateDeployFlowWithYaml.feature
│   │       ├── DeployApplication.feature
│   │       ├── EditComponentPipelineStage.feature
│   │       ├── JiraIntegration.feature
│   │       ├── LinksToTheComponents.feature
│   │       ├── PipelinesTable.feature
│   │       ├── PromoteApplication.feature
│   │       ├── RemoteCluster.feature
│   │       └── WidgetsIntegration.feature
│   └── api_tests/
│       ├── tekton/
│       │   ├── bitbucket/
│       │   │   ├── AutoDeploy.feature
│       │   │   ├── AutoDeployWithPromote.feature
│       │   │   ├── Deploy.feature
│       │   │   ├── Jira.feature
│       │   │   ├── PipelineRun.feature
│       │   │   ├── Promote.feature
│       │   │   ├── Recheck.feature
│       │   │   └── RemoteCluster.feature
│       │   ├── github/
│       │   │   ├── AutoDeploy.feature
│       │   │   ├── AutoDeployWithPromote.feature
│       │   │   ├── CommitMessageValidate.feature
│       │   │   ├── Deploy.feature
│       │   │   ├── Jira.feature
│       │   │   ├── PipelineRun.feature
│       │   │   ├── Promote.feature
│       │   │   ├── Recheck.feature
│       │   │   └── RemoteCluster.feature
│       │   ├── gerrit/
│       │   │   ├── AutoDeploy.feature
│       │   │   ├── AutoDeployWithPromote.feature
│       │   │   ├── Deploy.feature
│       │   │   ├── Jira.feature
│       │   │   ├── PipelineRun.feature
│       │   │   ├── Promote.feature
│       │   │   ├── Recheck.feature
│       │   │   └── RemoteCluster.feature
│       │   └── gitlab/
│       │       ├── AutoDeploy.feature
│       │       ├── AutoDeployWithPromote.feature
│       │       ├── CommitMessageValidate.feature
│       │       ├── Deploy.feature
│       │       ├── Jira.feature
│       │       ├── PipelineRun.feature
│       │       ├── Promote.feature
│       │       ├── Recheck.feature
│       │       └── RemoteCluster.feature
│       └── utility/
│           └── cleanup/
│               └── Clean.feature
└── README.md                           # 📖 This file
```

## How the QA agent works with tests

### 1. Analyze existing coverage
The QA agent performs a structured scan to understand what is already covered and where gaps exist.

- Inputs:
  - Stories (new or changed functionality to validate)
  - The feature files repository at `docs/testing/features/`

- What to scan:
  - UI features under `docs/testing/features/tekton/` (screens and provider-specific folders)
  - API features under `docs/testing/features/api_tests/tekton/` (per provider)
  - Utilities under `docs/testing/features/api_tests/utility/`

- How coverage is determined (heuristics):
  - Filename match to functionality keywords (e.g., Promote, AutoDeploy, RemoteCluster, JiraIntegration)
  - Scenario titles and steps containing functionality keywords
  - Tag match where available (e.g., `@smoke`, `@regression`, optional `@provider:github`)
  - UI ↔ API alignment for core flows (AutoDeploy, Promote, PipelineRun, RemoteCluster, Jira integration)

- Cross-provider parity check (for Tekton providers Bitbucket, GitHub, Gerrit, GitLab):
  - If functionality exists for one provider, verify presence for the others
  - Report missing providers for the same functionality

- Outputs (coverage report for the story):
  - Covered: list of matching paths to scenarios/files
  - Partial: existing paths plus missing providers or missing edge cases
  - Not covered: suggested new file path(s) and scenario placeholders
  - Recommended action (add scenario to existing file vs create new file)

### 2. Decide on a testing strategy

The QA agent uses the following decision matrix aligned to Tekton domains and providers. Actions are interactive: the agent reports findings and requests confirmation before creating or expanding tests.

| Situation | Action | Examples |
|----------|--------|----------|
| ✅ **Functionality already covered** | Inform the user that coverage exists and list matching tests; ask if additional scenarios/edge cases are desired | Covered by `tekton/github/PromoteApplication.feature` and `api_tests/tekton/github/Promote.feature` |
| 🧩 **Partial coverage: missing providers** | Show existing coverage and the list of missing providers; ask whether to add equivalent scenarios for those providers | `tekton/github/PromoteApplication.feature` exists → ask to add to `tekton/bitbucket/`, `tekton/gerrit/`, `tekton/gitlab/` |
| 🔗 **Partial coverage: UI vs API mismatch** | Present which side (UI or API) is covered; ask whether to add the complementary side to align flows | API `api_tests/tekton/github/Promote.feature` exists → ask to add UI `tekton/github/PromoteApplication.feature` (or vice versa) |
| ✳️ **New variation within an existing domain** | Propose adding scenarios to the closest existing feature; ask for confirmation; ensure provider parity plan | Add validation cases to `CreateDeployFlowWithYaml.feature` across providers |
| 🆕 **New functionality (no suitable file exists)** | Prefer Bitbucket first: propose creating `tekton/bitbucket/{NewFeature}.feature`; ask whether to cover GitHub/Gerrit/GitLab next | Create `tekton/bitbucket/InputsValidation.feature`, then ask to replicate for other providers |
| 🛠 **Utilities / maintenance flows** | Suggest using or extending utility features; ask before modifying cleanup flows | Update `api_tests/utility/cleanup/Clean.feature` |

### 3. Generating Gherkin tests

Tests in this repository often use advanced Gherkin patterns:

- Scenario Outline with multiple Examples tables per scenario
- Rich step data tables to pass parameters into steps
- Tags scoped at the Scenario level or per Examples block
- Long multi-stage flows with shared state (e.g., saving IMAGE_VERSION)

Recommended template:
```gherkin
Feature: Functionality name

  Scenario Outline: Concise, action-oriented scenario title
    Given User opens KubeRocketCI as default user
    When User performs an action with parameters
      | paramA | <paramA> |
      | paramB | <paramB> |
    And User saves the image version in memory as IMAGE_VERSION for <entity>
    Then User observes expected result for <entity>

    # Tags immediately above Examples apply to this Examples table only
    @UI @TektonGithubUI @TektonGithubCreateUI
    Examples:
      | paramA | paramB | entity |
      | foo    | bar    | app-1  |

    @UI @TektonBitbucketUI @TektonBitbucketAutoDeployWithPromoteUI
    Examples:
      | paramA | paramB | entity |
      | baz    | qux    | app-2  |
```

### 4. Tagging system

Current tags in this repository (do not replace existing tags; append if needed):

- Provider UI tags: `@TektonBitbucketUI`, `@TektonGithubUI`, `@TektonGerritUI`, `@TektonGitlabUI`
- Provider UI subdomains: `@TektonGithubCreateUI`, `@TektonGithubCloneUI`, `@TektonGithubImportUI`, `@TektonBitbucketAutoDeployWithPromoteUI`
- UI suite/run-scope tags: `@UI`, `@TektonBitbucketDeployUIRegression`, `@TektonBitbucketNoDeployUIRegression`
- API suite tags: `@TektonGithub`, `@TektonGithubShortRegression`
- Utilities: `@Clean`, `@CleanCr`

Tag scoping rules:
- Tags above Feature apply to all scenarios in the file
- Tags above a Scenario/Scenario Outline apply to that scenario
- Tags placed directly above an Examples block apply only to that Examples table (widely used in this repo)
- Multiple tags may be combined; order is not significant

Recommended usage:
- Always include a provider tag for filterability (e.g., `@TektonGithubUI`)
- Add suite/run-scope tags to control test selection (e.g., `@TektonGithubShortRegression`)
- Add subdomain tags for specific flows (e.g., `@TektonGithubCreateUI`)
- Use `@Clean`/`@CleanCr` for maintenance/cleanup scenarios

## How to add your existing tests

1. Placement (choose the correct path):
   - UI provider-agnostic screens: `docs/testing/features/tekton/*.feature`
   - UI provider-specific: `docs/testing/features/tekton/{bitbucket|github|gerrit|gitlab}/{Topic}.feature`
   - API provider-specific: `docs/testing/features/api_tests/tekton/{bitbucket|github|gerrit|gitlab}/{Topic}.feature`
   - Utilities/maintenance: `docs/testing/features/api_tests/utility/...`
   - New functionality: prefer Bitbucket first (`tekton/bitbucket/{NewFeature}.feature`), then mirror to other providers as needed

2. Naming convention:
   - Use PascalCase topic names, e.g., `PromoteApplication.feature`, `CreateDeployFlowWithYaml.feature`, `PipelineRun.feature`
   - Keep topic consistent across providers and UI/API (same filename where applicable)

3. Structure guidelines:
   - Prefer `Scenario Outline` with one or more `Examples` tables
   - Place tags directly above each `Examples` block when provider/scenario variants differ
   - Use rich parameter tables in steps; long multi-stage flows are acceptable
   - Shared state is allowed (e.g., saving IMAGE_VERSION), but keep steps explicit

4. Tagging requirements:
   - Always include a provider tag (e.g., `@TektonGithubUI`, `@TektonBitbucketUI`, or API suite tags like `@TektonGithub`)
   - Add suite/run-scope tags as needed (e.g., `@TektonBitbucketDeployUIRegression`, `@TektonGithubShortRegression`)
   - Add subdomain tags for specific flows (e.g., `@TektonGithubCreateUI`)
   - Utilities use `@Clean`/`@CleanCr`

5. Parity and alignment:
   - For existing functionality, check parity across providers; add only after confirmation it’s required
   - Align UI and API coverage for core flows (AutoDeploy, Promote, PipelineRun, RemoteCluster, Jira)

6. Cleanup and test data:
   - Ensure tests clean up created resources; reuse `api_tests/utility/cleanup/Clean.feature` where possible

7. PR checklist:
   - Link the Story and list the files/scenarios that cover it
   - Verify steps exist and tables are valid; ensure unique scenario titles
   - Confirm tags are correct (provider + suite) and Examples-scoped where needed
   - Note provider parity status (covered/missing by provider) and UI/API alignment

## Integration with the QA agent

This section defines how to interact with the QA agent (operational contract).

Inputs expected by the agent:
- Story: title and short description of the functionality to validate
- Optional scope: providers (bitbucket/github/gerrit/gitlab), test type (UI/API), priority/tags
- Preferences: whether to enforce provider parity; whether to add UI and/or API coverage

Agent actions (high level):
1. Scan the repository under `docs/testing/features/` (UI, API, utilities)
2. Build a coverage report (Covered / Partial / Not covered) with file paths
3. Follow the decision tree to propose actions and ask for confirmation

Outputs returned to the user:
- Coverage report with paths and tags
- Missing providers (parity matrix) and UI/API alignment hints
- Proposed next steps (add scenarios vs create new file), Bitbucket-first for brand-new functionality

Example (abbreviated coverage report):
```
Story: Promote application from dev to qa

Covered:
- UI: docs/testing/features/tekton/github/PromoteApplication.feature (@UI @TektonGithubUI)
- API: docs/testing/features/api_tests/tekton/github/Promote.feature (@TektonGithub)

Partial:
- Missing providers for UI: bitbucket, gerrit, gitlab

Proposed action:
- Add UI scenarios for missing providers? [yes/no]
```

Interaction loop:
- The agent reports findings → asks targeted questions (e.g., add missing providers?) → applies agreed changes.


