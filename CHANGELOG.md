<a name="unreleased"></a>

## [Unreleased]

### Bug Fixes

- fix markdown linting

### Chores

- **create-story:** add support for personas and assets in the site documentation

### Features

- add 'advisor' agent to the framework


<a name="v0.36.0"></a>

## [v0.36.0] - 2025-09-06

### Chores

- **create-story:** extend cursor and claudecode with commands support

### Features

- add aqa agent


<a name="v0.35.1"></a>

## [v0.35.1] - 2025-09-05

### Chores

- **xml:** add notice to the sdlc-framework general file

### Features

- **xml:** instruct agents to use xml tags correctly


<a name="v0.35.0"></a>

## [v0.35.0] - 2025-09-05

### Chores

- **migration:** add migration task to repo
- **xml:** convert update status report task
- **xml:** Add xml tags to the data sources

### Features

- **md2xml:** finalize adding xml tasks for framework components


<a name="v0.34.2"></a>

## [v0.34.2] - 2025-09-04

### Bug Fixes

- **tw:** remove reference to the non existing document


<a name="v0.34.1"></a>

## [v0.34.1] - 2025-09-04

### Chores

- update local installed framework to the latest version

### Features

- convert instructions to new format with xml tags


<a name="v0.34.0"></a>

## [v0.34.0] - 2025-09-04

### Features

- **pm:** add support for advanced project-brief creation


<a name="v0.33.1"></a>

## [v0.33.1] - 2025-09-04

### Chores

- Update sdlc framework documentation with latest changes


<a name="v0.33.0"></a>

## [v0.33.0] - 2025-09-04

### Bug Fixes

- **validation:** preserve subdirectory paths in dependency tracking

### Chores

- **ci:** fix command Injection via sonarqube-scan-action GitHub Action

### Features

- Add Technical Writer agent
- **epic-99:** add Story 99.03 Local Agent Apply Command


<a name="v0.32.0"></a>

## [v0.32.0] - 2025-09-02

### Bug Fixes

- fix Project Manager linting issues

### Chores

- Update krci-ai to version 0.31.0

### Features

- Create the PM assistant with the respective tasks, templates, data files
- Use XML tags for the most of the tasks in the framework


<a name="v0.31.0"></a>

## [v0.31.0] - 2025-09-01

### Chores

- We need to start adding documentation on how to use
- Update krci-ai to the latest version

### Features

- Use XML tags as a part of some tasks to validate efficiency


<a name="v0.30.0"></a>

## [v0.30.0] - 2025-08-27

### Bug Fixes

- Fix link to the short video
- Fix link to the short video

### Chores

- Align PRD artifacts accross the project

### Documentation

- Update status of Epic 7 to complete
- Update documentation
- Align docs to the latest changes
- Add link to quick video

### Features

- Add Product Marketing Manager persona


<a name="v0.29.1"></a>

## [v0.29.1] - 2025-08-22

### Bug Fixes

- ensure we are including go-dev assets in our bundle and advanced install

### Chores

- update go-dev agent instructions

### Code Refactoring

- remove 'config' directory that we are not using


<a name="v0.29.0"></a>

## [v0.29.0] - 2025-08-22

### Features

- **agents:** Add Go Developer agent and associated tasks and standards
- **tokens:** Add bundle token analysis command and integrate with bundle generation


<a name="v0.28.2"></a>

## [v0.28.2] - 2025-08-20

### Bug Fixes

- **validation:** warn on non-YAML agent files in agents directory ([#91](https://github.com/KubeRocketCI/kuberocketai/issues/91))

### Code Refactoring

- **install:** eliminate global variables in flag handling to improve thread safety and testability

### Documentation

- Epic 7 refinement

### Features

- **docs:** Define new epic on the MCP server topic


<a name="v0.28.1"></a>

## [v0.28.1] - 2025-08-19

### Chores

- **ci:** ensure we add all details to the release changelog

### Code Refactoring

- remove task reference from install command


<a name="v0.28.0"></a>

## [v0.28.0] - 2025-08-18

### Chores

- Align epic-9 with advanced installation
- Update krci-ai to 0.27.0

### Documentation

- simplify selective installation scope to agent-level only

### Features

- **selective:** implement agent-specific dependency validation and installation ([#65](https://github.com/KubeRocketCI/kuberocketai/issues/65))
- **stories:** add comprehensive implementation planning for Epic 8 selective installation

### Tests

- Refactor unit tests
- Refactor unit tests


<a name="v0.27.0"></a>

## [v0.27.0] - 2025-08-16


<a name="v0.26.1"></a>

## [v0.26.1] - 2025-08-14

### Bug Fixes

- **bundle:** skip additional file collection for targeted agent bundles

### Chores

- Update framework to the latest stable version


<a name="v0.26.0"></a>

## [v0.26.0] - 2025-08-14

### Bug Fixes

- Fix documentation section with TOC

### Chores

- Update changelog

### Code Refactoring

- switch to internal tools
- deduplicate helpers, centralize CLI style, and improve command exit handling

### Documentation

- Align mermaid to dark theme
- Simplify README for '/docs' directory
- **local-tasks:** make create-github-issues template-driven with dependency validation

### Features

- Implement token analysis CLI command and integrate GPT-4 tokenization
- update story template to be aligned with epic flow
- Add epic-9 to populate agents accross KubeRocketCI platform
- Add cheatsheet for krci-ai
- Create EPIC 8 for selective installation
- **epic-template:** improve outcome-focused, locally verifiable epics to reduce prescriptive testing and boost clarity


<a name="v0.25.2"></a>

## [v0.25.2] - 2025-08-06

### Bug Fixes

- Update krci-ai with correct tasks definition for PO/PM

### Chores

- Populate project with the stories
- Add artifacts to the repository
- Update bundled agents from the latest krci-ai version

### Documentation

- Update related works section
- Fix title level for video
- Add link to the youtube demo
- Update title of the repository


<a name="v0.25.1"></a>

## [v0.25.1] - 2025-07-31

### Bug Fixes

- Put updates for PO/PM into assets for proper delivery

### Chores

- Update krci-ai to the latest stable version


<a name="v0.25.0"></a>

## [v0.25.0] - 2025-07-31

### Code Refactoring

- fix unparam linter warnings by removing unused parameters
- **pm:** remove epic/story creation from PM tasks

### Documentation

- Update description and changelog

### Tests

- Add unit tests for bundling capabilities
- Add test for different packages
- add SonarCloud scan
- Increase coverage for the internal packages
- Add tests for the version package


<a name="v0.24.1"></a>

## [v0.24.1] - 2025-07-31

### Chores

- Add local task for the PO role for all IDEs

### Features

- add local task breakdown in validation framework insights


<a name="v0.24.0"></a>

## [v0.24.0] - 2025-07-30

### Bug Fixes

- Fix markdown formating

### Chores

- **bundle:** Bundle agents for WebChat usage

### Documentation

- add comprehensive user-focused documentation

### Features

- add GitHub issue creation workflow to PO agent
- implement local component override system


<a name="v0.23.0"></a>

## [v0.23.0] - 2025-07-29

### Documentation

- Put path to actual framework components
- Update README with badges
- Update README with the current state

### Features

- add single agent-task bundle generation with --task flag


<a name="v0.22.0"></a>

## [v0.22.0] - 2025-07-29

### Features

- add targeted agent bundle selection with --agent flag
- implement complete bundle generation with CLI command


<a name="v0.21.6"></a>

## [v0.21.6] - 2025-07-28

### Chores

- Update changelog

### Features

- Update task instrusction for 'update-' tasks


<a name="v0.21.5"></a>

## [v0.21.5] - 2025-07-28

### Features

- **pm:** Update tasks instructions for product manager


<a name="v0.21.4"></a>

## [v0.21.4] - 2025-07-28

### Code Refactoring

- Remove changelog from the baking


<a name="v0.21.3"></a>

## [v0.21.3] - 2025-07-27

### Chores

- Update changelog
- Update changelog

### Features

- Extend 'list agents -v' command


<a name="v0.21.2"></a>

## [v0.21.2] - 2025-07-27

### Features

- **changelog:** enhance CLI output with compact formatting and visual improvements


<a name="v0.21.1"></a>

## [v0.21.1] - 2025-07-27

### Chores

- Update changelog
- Build changelog automatically for bundling

### Code Refactoring

- Remove validate-changelog from the cli


<a name="v0.21.0"></a>

## [v0.21.0] - 2025-07-27

### Chores

- Update changelog
- Remove unused template

### Features

- **check-update:** add changelog generation, validation, and embedded display; add version/update commands


<a name="v0.20.0"></a>

## [v0.20.0] - 2025-07-27

### Chores

- Update krci-ai to the version 0.19.0

### Features

- implement enhanced framework validation with dependency analysis (Module path changed from github.com/epam/kuberocketai to github.com/KubeRocketCI/kuberocketai)

### BREAKING CHANGE


Module path changed from github.com/epam/kuberocketai to github.com/KubeRocketCI/kuberocketai


<a name="v0.19.0"></a>

## [v0.19.0] - 2025-07-27

### Bug Fixes

- **cli:** improve validate command warning display and UX

### Chores

- Update tasks for story management
- Remove review story task from the qa agent
- Add windsurf configuration
- Update business rules template
- Fix formating for templates
- **ci:** Create github issue templates


<a name="v0.18.0"></a>

## [v0.18.0] - 2025-07-26

### Chores

- Update krci-ai to the latest changes
- **claude:** Update repo instructions for Claude

### Features

- **ide:** implement Windsurf IDE integration support


<a name="v0.17.0"></a>

## [v0.17.0] - 2025-07-24

### Chores

- **ci:** run installation and validation as a part of CI

### Documentation

- **core:** standardize asset validation and reference management in agents and tasks


<a name="v0.16.0"></a>

## [v0.16.0] - 2025-07-23

### Chores

- Update krci-ai to the version 0.15.1

### Features

- **validation:** add template validation and framework link validation
- **validation:** implement task path link validation


<a name="v0.15.1"></a>

## [v0.15.1] - 2025-07-23

### Chores

- Update epic, story states
- Remove unused template from the latest version


<a name="v0.15.0"></a>

## [v0.15.0] - 2025-07-22

### Chores

- Update krci-ai to the version 0.14.0

### Features

- **tasks:** add review-story task and streamline implement-feature for multi-role story validation


<a name="v0.14.0"></a>

## [v0.14.0] - 2025-07-22

### Features

- **po:** add epic and story update tasks with change control (Product Owner agent now supports update operations with strict change control to protect completed work while enabling controlled evolution of epics and stories during development.)

### BREAKING CHANGE


Product Owner agent now supports update operations with strict change control to protect completed work while enabling controlled evolution of epics and stories during development.


<a name="v0.13.0"></a>

## [v0.13.0] - 2025-07-22

### Chores

- Update krci-ai to the version 0.12.0

### Features

- Update copilot agents tools list


<a name="v0.12.0"></a>

## [v0.12.0] - 2025-07-22

### Chores

- Update krci-ai to the version 0.11.2

### Code Refactoring

- **pm:** optimize task guidelines and enhance epic template (Content Guidelines sections restructured)

### BREAKING CHANGE


Content Guidelines sections restructured


<a name="v0.11.2"></a>

## [v0.11.2] - 2025-07-21

### Bug Fixes

- Fix commands for the developer and arch roles


<a name="v0.11.1"></a>

## [v0.11.1] - 2025-07-21

### Code Refactoring

- align VSCode chatmodes with official docs and eliminate hardcoded paths


<a name="v0.11.0"></a>

## [v0.11.0] - 2025-07-21

### Chores

- Update krci-ai version to 0.10.0

### Code Refactoring

- Remove registry.json from the instructions


<a name="v0.10.0"></a>

## [v0.10.0] - 2025-07-20

### Features

- added prompt files for BA assistant
- Add implement-feature task
- **pm:** Update tasks for PM role and update template


<a name="v0.9.0"></a>

## [v0.9.0] - 2025-07-17

### Chores

- Update krci-ai framework to the version 0.8.2

### Features

- **install:** add VS Code chat mode integration with enhanced tool support


<a name="v0.8.2"></a>

## [v0.8.2] - 2025-07-17

### Bug Fixes

- Align architect role with the list of available supported commands


<a name="v0.8.1"></a>

## [v0.8.1] - 2025-07-17

### Bug Fixes

- **installer:** correct task file validation to check for .md files

### Chores

- Update krci-ai to the version 0.8.0


<a name="v0.8.0"></a>

## [v0.8.0] - 2025-07-17

### Bug Fixes

- Fix agent description

### Chores

- Remove incorrect tasks

### Features

- **agent:** Add draft for the architect agent


<a name="v0.7.2"></a>

## [v0.7.2] - 2025-07-16

### Bug Fixes

- **validate:** enforce additionalProperties:false in agent schema validation

### Chores

- Update dogfooding version to 0.7.1


<a name="v0.7.1"></a>

## [v0.7.1] - 2025-07-11

### Chores

- Update activation prompt for agents


<a name="v0.7.0"></a>

## [v0.7.0] - 2025-07-11

### Features

- **agent:** Add customization field and activation logic to agent schema


<a name="v0.6.1"></a>

## [v0.6.1] - 2025-07-09

### Features

- **validate:** Add framework validation command with agent and task checks


<a name="v0.6.0"></a>

## [v0.6.0] - 2025-07-09

### Features

- Implement basic validation of agent yaml


<a name="v0.5.1"></a>

## [v0.5.1] - 2025-07-08

### Chores

- Update package description
- Update dogfood configuration to version 0.5.0

### Documentation

- Update README.md to reflect latest changes

### Features

- Add cursor configuration populated from the framework


<a name="v0.5.0"></a>

## [v0.5.0] - 2025-07-08

### Features

- **install:** add IDE integration support and improve install command flags


<a name="v0.4.0"></a>

## [v0.4.0] - 2025-07-08

### Bug Fixes

- Fix QA agent definition for claude code
- Fix architect agent definition for claude code

### Chores

- Add validation schema support for yaml-validator plugin
- **dogfood:** Align claude code commands

### Features

- Add create-epic/create-story task with updated templates
- **agent:** Add Product Owner agent and scope to full PO responsibilities
- **agents:** Refactor agent definitions and add markdown command files
- **schema:** Align agent schema and PO command definition


<a name="v0.3.0"></a>

## [v0.3.0] - 2025-07-07

### Code Refactoring

- move assets symlink to project root and consolidate asset embedding

### Features

- add business analyst agent framework with tasks, templates, and data
- add QA agent framework with tasks, templates, and data
- add product manager agent framework with tasks, templates, and data
- add developer agent framework with tasks, templates, and data
- First version of architect agent in framework
- **cmd:** add 'list agents' command to display installed agents


<a name="v0.2.1"></a>

## [v0.2.1] - 2025-07-07

### Bug Fixes

- **ci:** Add syft on the pipeline


<a name="v0.2.0"></a>

## [v0.2.0] - 2025-07-07

### Chores

- **ci:** Add support for signing artifacts and sbom

### Documentation

- Update installation instructions


<a name="v0.1.2"></a>

## [v0.1.2] - 2025-07-06

### Chores

- **ci:** Remove cache step


<a name="v0.1.1"></a>

## [v0.1.1] - 2025-07-06

### Chores

- **ci:** Fix GitHub Actions


<a name="v0.1.0"></a>

## v0.1.0 - 2025-07-06

### Features

- initial project setup with CLI foundation and CI/CD


[Unreleased]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.36.0...HEAD
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
