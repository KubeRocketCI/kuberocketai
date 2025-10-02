## KubeRocketAI CLI Release

Welcome to this releases of KubeRocketAI CLI!

## Installation

### Homebrew (macOS)

```bash
brew tap KubeRocketCI/homebrew-tap
brew install krci-ai
```

<a name="unreleased"></a>

### [Unreleased]

#### Code Refactoring

* Merge prerequisites section into instructions for tasks

#### Features

* Merge prerequisites section into instruction one for tasks

<a name="v0.39.0"></a>

### [v0.39.0] - 2025-10-01

#### Chores

* cb58749 update PO task for github issue creation
* 080b652 update changelog

#### Features

* 9f4774c Implement XML validation for agent files

<a name="v0.38.5"></a>

### [v0.38.5] - 2025-09-19

#### Chores

* 391db7f update to the latest version of the framework

#### Code Refactoring

* f47680b enhance validation output and insights generation

<a name="v0.38.4"></a>

### [v0.38.4] - 2025-09-17

#### Features

* cfa9b4d Add task dependencies field to task frontmatter ([#98](https://github.com/KubeRocketCI/kuberocketai/issues/98))

<a name="v0.38.3"></a>

### [v0.38.3] - 2025-09-17

#### Bug Fixes

* fde6e84 fix xml tags for the number of tasks

<a name="v0.38.2"></a>

### [v0.38.2] - 2025-09-17

#### Code Refactoring

* f895666 simplify formatting for tasks

<a name="v0.38.1"></a>

### [v0.38.1] - 2025-09-17

#### Chores

* c42fe8a update krci-ai to the latest version

#### Code Refactoring

* f11fddb update activation prompt by introducing dependecies

<a name="v0.38.0"></a>

### [v0.38.0] - 2025-09-17

#### Bug Fixes

* 93bc6f5 remove unrelated dependecies from the tasks
* f61cc8a align tasks to the latest dependency schema

#### Chores

* 4a94f20 update dependecy for AQA tasks

#### Code Refactoring

* 5173c8f align to the latest dependency schema

#### Features

* 3647be2 Update 'advisor' persona with new dependency management
* 925d36d Refactor task dependency tracking system ([#98](https://github.com/KubeRocketCI/kuberocketai/issues/98))

<a name="v0.37.2"></a>

### [v0.37.2] - 2025-09-16

#### Bug Fixes

* 6b9bfe0 we've migrated to the new location for the framework data

#### Code Refactoring

* 6a31dc6 update sdlc-framework instructions and align with structure

<a name="v0.37.1"></a>

### [v0.37.1] - 2025-09-15

#### Chores

* ae37cea update installed framework to the latest version

#### Features

* e464d44 update instructions on dependecy loading

<a name="v0.37.0"></a>

### [v0.37.0] - 2025-09-13

#### Bug Fixes

* d1d4f02 fix markdown linting

#### Chores

* 7895445 Install the latest version of framework locally
* 6ca3379 **create-story:** add support for personas and assets in the site documentation

#### Features

* 367bcfe add 'advisor' agent to the framework

<a name="v0.36.0"></a>

### [v0.36.0] - 2025-09-06

#### Chores

* 3af08b3 **create-story:** extend cursor and claudecode with commands support

#### Features

* f33a827 add aqa agent

<a name="v0.35.1"></a>

### [v0.35.1] - 2025-09-05

#### Chores

* 3786c73 **xml:** add notice to the sdlc-framework general file

#### Features

* 8a0b77f **xml:** instruct agents to use xml tags correctly

<a name="v0.35.0"></a>

### [v0.35.0] - 2025-09-05

#### Chores

* 9c3a1fb **migration:** add migration task to repo
* 978a248 **xml:** convert update status report task
* 5a20cf2 **xml:** Add xml tags to the data sources

#### Features

* 6040c83 **md2xml:** finalize adding xml tasks for framework components

<a name="v0.34.2"></a>

### [v0.34.2] - 2025-09-04

#### Bug Fixes

* 21c8d30 **tw:** remove reference to the non existing document

<a name="v0.34.1"></a>

### [v0.34.1] - 2025-09-04

#### Chores

* 8799db7 update local installed framework to the latest version

#### Features

* e7ee4b7 convert instructions to new format with xml tags

<a name="v0.34.0"></a>

### [v0.34.0] - 2025-09-04

#### Features

* 209b203 **pm:** add support for advanced project-brief creation

<a name="v0.33.1"></a>

### [v0.33.1] - 2025-09-04

#### Chores

* 21b3199 Update sdlc framework documentation with latest changes

<a name="v0.33.0"></a>

### [v0.33.0] - 2025-09-04

#### Bug Fixes

* 43dedfb **validation:** preserve subdirectory paths in dependency tracking

#### Chores

* d3de017 **ci:** fix command Injection via sonarqube-scan-action GitHub Action

#### Features

* c3d0bf7 Add Technical Writer agent
* a4ff4c5 **epic-99:** add Story 99.03 Local Agent Apply Command

<a name="v0.32.0"></a>

### [v0.32.0] - 2025-09-02

#### Bug Fixes

* b9345f7 fix Project Manager linting issues

#### Chores

* e0b661d Update krci-ai to version 0.31.0

#### Features

* 9458075 Create the PM assistant with the respective tasks, templates, data files
* c06dd1c Use XML tags for the most of the tasks in the framework

<a name="v0.31.0"></a>

### [v0.31.0] - 2025-09-01

#### Chores

* 8035595 We need to start adding documentation on how to use
* 7bacbec Update krci-ai to the latest version

#### Features

* b67cffe Use XML tags as a part of some tasks to validate efficiency

<a name="v0.30.0"></a>

### [v0.30.0] - 2025-08-27

#### Bug Fixes

* 32bd264 Fix link to the short video
* 95ee150 Fix link to the short video

#### Chores

* 90da6d7 Align PRD artifacts accross the project

#### Documentation

* 571fa24 Update status of Epic 7 to complete
* 9ec86ca Update documentation
* 3e5c05a Align docs to the latest changes
* ced0e22 Add link to quick video

#### Features

* 5011414 Add Product Marketing Manager persona

<a name="v0.29.1"></a>

### [v0.29.1] - 2025-08-22

#### Bug Fixes

* 1f7d9da ensure we are including go-dev assets in our bundle and advanced install

#### Chores

* 2c2845e update go-dev agent instructions

#### Code Refactoring

* 00f4328 remove 'config' directory that we are not using

<a name="v0.29.0"></a>

### [v0.29.0] - 2025-08-22

#### Features

* 31a10e5 **agents:** Add Go Developer agent and associated tasks and standards
* d4bf723 **tokens:** Add bundle token analysis command and integrate with bundle generation

<a name="v0.28.2"></a>

### [v0.28.2] - 2025-08-20

#### Bug Fixes

* a357a22 **validation:** warn on non-YAML agent files in agents directory ([#91](https://github.com/KubeRocketCI/kuberocketai/issues/91))

#### Code Refactoring

* e8dcaa8 **install:** eliminate global variables in flag handling to improve thread safety and testability

#### Documentation

* 5748a7a Epic 7 refinement

#### Features

* 1e4de46 **docs:** Define new epic on the MCP server topic

<a name="v0.28.1"></a>

### [v0.28.1] - 2025-08-19

#### Chores

* fbf355f **ci:** ensure we add all details to the release changelog

#### Code Refactoring

* 297973b remove task reference from install command

<a name="v0.28.0"></a>

### [v0.28.0] - 2025-08-18

#### Chores

* 59afe34 Align epic-9 with advanced installation
* 7b9c87d Update krci-ai to 0.27.0

#### Documentation

* b7948e6 simplify selective installation scope to agent-level only

#### Features

* 5e7c3b6 **selective:** implement agent-specific dependency validation and installation ([#65](https://github.com/KubeRocketCI/kuberocketai/issues/65))
* 49e7cf3 **stories:** add comprehensive implementation planning for Epic 8 selective installation

#### Tests

* be0f16c Refactor unit tests
* bb576f7 Refactor unit tests

<a name="v0.27.0"></a>

### [v0.27.0] - 2025-08-16

<a name="v0.26.1"></a>

### [v0.26.1] - 2025-08-14

#### Bug Fixes

* 12f4fe3 **bundle:** skip additional file collection for targeted agent bundles

#### Chores

* 64aeaeb Update framework to the latest stable version

<a name="v0.26.0"></a>

### [v0.26.0] - 2025-08-14

#### Bug Fixes

* 7c3a7e2 Fix documentation section with TOC

#### Chores

* 2fa7a0d Update changelog

#### Code Refactoring

* 7712911 switch to internal tools
* a28d431 deduplicate helpers, centralize CLI style, and improve command exit handling

#### Documentation

* 76acd3f Align mermaid to dark theme
* de53fe7 Simplify README for '/docs' directory
* be0284c **local-tasks:** make create-github-issues template-driven with dependency validation

#### Features

* a20af26 Implement token analysis CLI command and integrate GPT-4 tokenization
* a961aa5 update story template to be aligned with epic flow
* c8920b6 Add epic-9 to populate agents accross KubeRocketCI platform
* 7bf1b57 Add cheatsheet for krci-ai
* 49e9e5d Create EPIC 8 for selective installation
* 5aefa4c **epic-template:** improve outcome-focused, locally verifiable epics to reduce prescriptive testing and boost clarity

<a name="v0.25.2"></a>

### [v0.25.2] - 2025-08-06

#### Bug Fixes

* ede5d38 Update krci-ai with correct tasks definition for PO/PM

#### Chores

* abd777e Populate project with the stories
* e4cc86d Add artifacts to the repository
* dd234f1 Update bundled agents from the latest krci-ai version

#### Documentation

* af76a25 Update related works section
* c5e3342 Fix title level for video
* 0e39f3f Add link to the youtube demo
* 6312a54 Update title of the repository

<a name="v0.25.1"></a>

### [v0.25.1] - 2025-07-31

#### Bug Fixes

* 5916293 Put updates for PO/PM into assets for proper delivery

#### Chores

* b493927 Update krci-ai to the latest stable version

<a name="v0.25.0"></a>

### [v0.25.0] - 2025-07-31

#### Code Refactoring

* d2b8234 fix unparam linter warnings by removing unused parameters
* 79c596f **pm:** remove epic/story creation from PM tasks

#### Documentation

* f4a7335 Update description and changelog

#### Tests

* 88f1da0 Add unit tests for bundling capabilities
* 20c5977 Add test for different packages
* b173f2c add SonarCloud scan
* 10bc8a6 Increase coverage for the internal packages
* 77e161b Add tests for the version package

<a name="v0.24.1"></a>

### [v0.24.1] - 2025-07-31

#### Chores

* 9682a2e Add local task for the PO role for all IDEs

#### Features

* ded3546 add local task breakdown in validation framework insights

<a name="v0.24.0"></a>

### [v0.24.0] - 2025-07-30

#### Bug Fixes

* 8261887 Fix markdown formating

#### Chores

* 174ead3 **bundle:** Bundle agents for WebChat usage

#### Documentation

* 0f4435c add comprehensive user-focused documentation

#### Features

* bb8e1cc add GitHub issue creation workflow to PO agent
* da46bea implement local component override system

<a name="v0.23.0"></a>

### [v0.23.0] - 2025-07-29

#### Documentation

* bd8d269 Put path to actual framework components
* 4fb567a Update README with badges
* a0f0374 Update README with the current state

#### Features

* 114d8bc add single agent-task bundle generation with --task flag

<a name="v0.22.0"></a>

### [v0.22.0] - 2025-07-29

#### Features

* 137ba42 add targeted agent bundle selection with --agent flag
* 04a0475 implement complete bundle generation with CLI command

<a name="v0.21.6"></a>

### [v0.21.6] - 2025-07-28

#### Chores

* 6910469 Update changelog

#### Features

* 7d539a9 Update task instrusction for 'update-' tasks

<a name="v0.21.5"></a>

### [v0.21.5] - 2025-07-28

#### Features

* 7ca613f **pm:** Update tasks instructions for product manager

<a name="v0.21.4"></a>

### [v0.21.4] - 2025-07-28

#### Code Refactoring

* a321aab Remove changelog from the baking

<a name="v0.21.3"></a>

### [v0.21.3] - 2025-07-27

#### Chores

* 2b7fc14 Update changelog
* 5bbe544 Update changelog

#### Features

* 08ac7ba Extend 'list agents -v' command

<a name="v0.21.2"></a>

### [v0.21.2] - 2025-07-27

#### Features

* b07ca50 **changelog:** enhance CLI output with compact formatting and visual improvements

<a name="v0.21.1"></a>

### [v0.21.1] - 2025-07-27

#### Chores

* a2b0f3e Update changelog
* 86f84bc Build changelog automatically for bundling

#### Code Refactoring

* 94d4be3 Remove validate-changelog from the cli

<a name="v0.21.0"></a>

### [v0.21.0] - 2025-07-27

#### Chores

* 8723d5c Update changelog
* f3a1bfd Remove unused template

#### Features

* eda0fd1 **check-update:** add changelog generation, validation, and embedded display; add version/update commands

<a name="v0.20.0"></a>

### [v0.20.0] - 2025-07-27

#### Chores

* 76831fc Update krci-ai to the version 0.19.0

#### Features

* a4f8a72 implement enhanced framework validation with dependency analysis (Module path changed from github.com/epam/kuberocketai to github.com/KubeRocketCI/kuberocketai)

#### BREAKING CHANGE


Module path changed from github.com/epam/kuberocketai to github.com/KubeRocketCI/kuberocketai

<a name="v0.19.0"></a>

### [v0.19.0] - 2025-07-27

#### Bug Fixes

* 6a38c52 **cli:** improve validate command warning display and UX

#### Chores

* 40f438c Update tasks for story management
* bff07c6 Remove review story task from the qa agent
* f2de0d3 Add windsurf configuration
* 6e7aec0 Update business rules template
* d8080cf Fix formating for templates
* 77c3d20 **ci:** Create github issue templates

<a name="v0.18.0"></a>

### [v0.18.0] - 2025-07-26

#### Chores

* 0ddf6fc Update krci-ai to the latest changes
* 2a58e0d **claude:** Update repo instructions for Claude

#### Features

* 36b8f13 **ide:** implement Windsurf IDE integration support

<a name="v0.17.0"></a>

### [v0.17.0] - 2025-07-24

#### Chores

* 05592ef **ci:** run installation and validation as a part of CI

#### Documentation

* 3895c46 **core:** standardize asset validation and reference management in agents and tasks

<a name="v0.16.0"></a>

### [v0.16.0] - 2025-07-23

#### Chores

* f75f05d Update krci-ai to the version 0.15.1

#### Features

* 1875396 **validation:** add template validation and framework link validation
* 6630ac5 **validation:** implement task path link validation

<a name="v0.15.1"></a>

### [v0.15.1] - 2025-07-23

#### Chores

* 5645437 Update epic, story states
* c02d2c5 Remove unused template from the latest version

<a name="v0.15.0"></a>

### [v0.15.0] - 2025-07-22

#### Chores

* 53eaee4 Update krci-ai to the version 0.14.0

#### Features

* c4d1da8 **tasks:** add review-story task and streamline implement-feature for multi-role story validation

<a name="v0.14.0"></a>

### [v0.14.0] - 2025-07-22

#### Features

* 7ffe0b8 **po:** add epic and story update tasks with change control (Product Owner agent now supports update operations with strict change control to protect completed work while enabling controlled evolution of epics and stories during development.)

#### BREAKING CHANGE


Product Owner agent now supports update operations with strict change control to protect completed work while enabling controlled evolution of epics and stories during development.

<a name="v0.13.0"></a>

### [v0.13.0] - 2025-07-22

#### Chores

* 6d9aae2 Update krci-ai to the version 0.12.0

#### Features

* 58d2af7 Update copilot agents tools list

<a name="v0.12.0"></a>

### [v0.12.0] - 2025-07-22

#### Chores

* ec7a17f Update krci-ai to the version 0.11.2

#### Code Refactoring

* ba938e2 **pm:** optimize task guidelines and enhance epic template (Content Guidelines sections restructured)

#### BREAKING CHANGE


Content Guidelines sections restructured

<a name="v0.11.2"></a>

### [v0.11.2] - 2025-07-21

#### Bug Fixes

* db801cf Fix commands for the developer and arch roles

<a name="v0.11.1"></a>

### [v0.11.1] - 2025-07-21

#### Code Refactoring

* cd40a94 align VSCode chatmodes with official docs and eliminate hardcoded paths

<a name="v0.11.0"></a>

### [v0.11.0] - 2025-07-21

#### Chores

* fdff6ad Update krci-ai version to 0.10.0

#### Code Refactoring

* 8127747 Remove registry.json from the instructions

<a name="v0.10.0"></a>

### [v0.10.0] - 2025-07-20

#### Features

* 95e15c7 added prompt files for BA assistant
* c9a5fbf Add implement-feature task
* 3669eb7 **pm:** Update tasks for PM role and update template

<a name="v0.9.0"></a>

### [v0.9.0] - 2025-07-17

#### Chores

* e098eca Update krci-ai framework to the version 0.8.2

#### Features

* 431ff7f **install:** add VS Code chat mode integration with enhanced tool support

<a name="v0.8.2"></a>

### [v0.8.2] - 2025-07-17

#### Bug Fixes

* 0a1b795 Align architect role with the list of available supported commands

<a name="v0.8.1"></a>

### [v0.8.1] - 2025-07-17

#### Bug Fixes

* 1c4c7ac **installer:** correct task file validation to check for .md files

#### Chores

* 84e0260 Update krci-ai to the version 0.8.0

<a name="v0.8.0"></a>

### [v0.8.0] - 2025-07-17

#### Bug Fixes

* f53b5f0 Fix agent description

#### Chores

* b054557 Remove incorrect tasks

#### Features

* 61e06fa **agent:** Add draft for the architect agent

<a name="v0.7.2"></a>

### [v0.7.2] - 2025-07-16

#### Bug Fixes

* 5577290 **validate:** enforce additionalProperties:false in agent schema validation

#### Chores

* 71a66e6 Update dogfooding version to 0.7.1

<a name="v0.7.1"></a>

### [v0.7.1] - 2025-07-11

#### Chores

* 9aa3fcb Update activation prompt for agents

<a name="v0.7.0"></a>

### [v0.7.0] - 2025-07-11

#### Features

* b3308af **agent:** Add customization field and activation logic to agent schema

<a name="v0.6.1"></a>

### [v0.6.1] - 2025-07-09

#### Features

* 8963a33 **validate:** Add framework validation command with agent and task checks

<a name="v0.6.0"></a>

### [v0.6.0] - 2025-07-09

#### Features

* 35e67b6 Implement basic validation of agent yaml

<a name="v0.5.1"></a>

### [v0.5.1] - 2025-07-08

#### Chores

* 52ab0ef Update package description
* 578fca5 Update dogfood configuration to version 0.5.0

#### Documentation

* a89b460 Update README.md to reflect latest changes

#### Features

* 3416e94 Add cursor configuration populated from the framework

<a name="v0.5.0"></a>

### [v0.5.0] - 2025-07-08

#### Features

* a36f581 **install:** add IDE integration support and improve install command flags

<a name="v0.4.0"></a>

### [v0.4.0] - 2025-07-08

#### Bug Fixes

* a7082d4 Fix QA agent definition for claude code
* 4932b8a Fix architect agent definition for claude code

#### Chores

* a39d9fd Add validation schema support for yaml-validator plugin
* b1f7a83 **dogfood:** Align claude code commands

#### Features

* ca95b70 Add create-epic/create-story task with updated templates
* 1b74738 **agent:** Add Product Owner agent and scope to full PO responsibilities
* d911a82 **agents:** Refactor agent definitions and add markdown command files
* 05ed88b **schema:** Align agent schema and PO command definition

<a name="v0.3.0"></a>

### [v0.3.0] - 2025-07-07

#### Code Refactoring

* 50f0499 move assets symlink to project root and consolidate asset embedding

#### Features

* 24168dc add business analyst agent framework with tasks, templates, and data
* d63cebb add QA agent framework with tasks, templates, and data
* e668067 add product manager agent framework with tasks, templates, and data
* dfa4bf8 add developer agent framework with tasks, templates, and data
* cf4d97c First version of architect agent in framework
* 8c641b9 **cmd:** add 'list agents' command to display installed agents

<a name="v0.2.1"></a>

### [v0.2.1] - 2025-07-07

#### Bug Fixes

* 6ae8bac **ci:** Add syft on the pipeline

<a name="v0.2.0"></a>

### [v0.2.0] - 2025-07-07

#### Chores

* 4140a3d **ci:** Add support for signing artifacts and sbom

#### Documentation

* 666f79d Update installation instructions

<a name="v0.1.2"></a>

### [v0.1.2] - 2025-07-06

#### Chores

* 78fdf51 **ci:** Remove cache step

<a name="v0.1.1"></a>

### [v0.1.1] - 2025-07-06

#### Chores

* e7dcf56 **ci:** Fix GitHub Actions

<a name="v0.1.0"></a>

### v0.1.0 - 2025-07-06

#### Features

* 4adcbf3 initial project setup with CLI foundation and CI/CD


[Unreleased]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.39.0...HEAD
[v0.39.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.38.5...v0.39.0
[v0.38.5]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.38.4...v0.38.5
[v0.38.4]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.38.3...v0.38.4
[v0.38.3]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.38.2...v0.38.3
[v0.38.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.38.1...v0.38.2
[v0.38.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.38.0...v0.38.1
[v0.38.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.37.2...v0.38.0
[v0.37.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.37.1...v0.37.2
[v0.37.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.37.0...v0.37.1
[v0.37.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.36.0...v0.37.0
[v0.36.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.35.1...v0.36.0
[v0.35.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.35.0...v0.35.1
[v0.35.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.34.2...v0.35.0
[v0.34.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.34.1...v0.34.2
[v0.34.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.34.0...v0.34.1
[v0.34.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.33.1...v0.34.0
[v0.33.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.33.0...v0.33.1
[v0.33.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.32.0...v0.33.0
[v0.32.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.31.0...v0.32.0
[v0.31.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.30.0...v0.31.0
[v0.30.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.29.1...v0.30.0
[v0.29.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.29.0...v0.29.1
[v0.29.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.28.2...v0.29.0
[v0.28.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.28.1...v0.28.2
[v0.28.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.28.0...v0.28.1
[v0.28.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.27.0...v0.28.0
[v0.27.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.26.1...v0.27.0
[v0.26.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.26.0...v0.26.1
[v0.26.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.25.2...v0.26.0
[v0.25.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.25.1...v0.25.2
[v0.25.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.25.0...v0.25.1
[v0.25.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.24.1...v0.25.0
[v0.24.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.24.0...v0.24.1
[v0.24.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.23.0...v0.24.0
[v0.23.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.22.0...v0.23.0
[v0.22.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.21.6...v0.22.0
[v0.21.6]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.21.5...v0.21.6
[v0.21.5]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.21.4...v0.21.5
[v0.21.4]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.21.3...v0.21.4
[v0.21.3]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.21.2...v0.21.3
[v0.21.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.21.1...v0.21.2
[v0.21.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.21.0...v0.21.1
[v0.21.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.20.0...v0.21.0
[v0.20.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.19.0...v0.20.0
[v0.19.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.18.0...v0.19.0
[v0.18.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.17.0...v0.18.0
[v0.17.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.16.0...v0.17.0
[v0.16.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.15.1...v0.16.0
[v0.15.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.15.0...v0.15.1
[v0.15.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.14.0...v0.15.0
[v0.14.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.13.0...v0.14.0
[v0.13.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.12.0...v0.13.0
[v0.12.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.11.2...v0.12.0
[v0.11.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.11.1...v0.11.2
[v0.11.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.11.0...v0.11.1
[v0.11.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.9.0...v0.10.0
[v0.9.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.8.2...v0.9.0
[v0.8.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.8.1...v0.8.2
[v0.8.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.8.0...v0.8.1
[v0.8.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.7.2...v0.8.0
[v0.7.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.7.1...v0.7.2
[v0.7.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.6.1...v0.7.0
[v0.6.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.4.0...v0.5.0
[v0.4.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.2.1...v0.3.0
[v0.2.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.1.2...v0.2.0
[v0.1.2]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.1.0...v0.1.1
