<a name="unreleased"></a>

## [Unreleased]

### Chores

- Remove unused template


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


[Unreleased]: https://github.com/KubeRocketCI/kuberocketai/compare/v0.20.0...HEAD
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
