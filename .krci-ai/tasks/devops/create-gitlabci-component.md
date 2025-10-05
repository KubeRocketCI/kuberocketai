---
dependencies:
  data:
    - krci-ai/core-framework-standards.yaml
    - devops/gitlabci-component-patterns.yaml
  templates:
    - devops/gitlabci-component-template.md
---

# Task: Create GitLab CI/CD Component Library

## Description

Scaffold production-ready GitLab CI/CD component library with standardized stage flow and critical dependencies for any tech stack. Component libraries provide reusable job templates, separate review/build workflows, and consistent architectural patterns while allowing technology-specific job implementations.

<instructions>
Gather tech stack requirements. Ask for programming language, framework, build tool, test framework, container image, and typical build/test commands. Identify code quality tools (linters, formatters, SonarQube). Determine deployment artifacts and Helm chart requirements if applicable.

Create repository structure with standardized stage flow. Initialize Git repository with templates/ directory. Create common.yml for shared job templates with MANDATORY 7-stage structure: [prepare, test, build, verify, package, publish, release]. Create review.yml for merge request pipeline, build.yml for main branch pipeline. Add README.md, LICENSE.md, .gitlab-ci.yml, and .gitignore. Include sample source code for testing components.

Define shared job templates in common.yml with critical dependencies. Create separated .test-job and .build-job templates (never combined). Create mandatory init-values job in prepare stage producing dotenv artifacts. Implement technology-specific caching patterns. Define spec:inputs for stage names, container image, and tech-stack-specific parameters. Use $[[ inputs.name ]] interpolation throughout.

Implement review pipeline in review.yml following test→build→verify flow. Include common.yml with local include pattern. Add MANDATORY jobs: init-values (prepare), test (test stage), build (build stage, depends on test), sonar (build stage, depends on init-values + test). Add technology-specific jobs as needed within appropriate stages. Configure Docker build verification without publishing in verify stage.

Implement build pipeline in build.yml following test→build→package→publish flow. Include common.yml and extend shared templates. Add MANDATORY dependency patterns: build depends on test completion, sonar depends on init-values + test artifacts, buildkit-build depends on init-values + build artifacts, git-tag depends on init-values artifacts. Add technology-specific packaging jobs as needed.

Create orchestration in main .gitlab-ci.yml with standardized stages. Define workflow rules for merge requests and protected branches. Include review component conditionally for merge requests. Include build component conditionally for main branch. Pass global variables as component inputs. Define MANDATORY stage sequence: [prepare, test, build, verify, package, publish, release].

Document component library in README.md. Explain library purpose, standardized stage flow, and available components. Create inputs table documenting all parameters with defaults. Provide usage examples for including components. List required CI/CD variables. Document mandatory vs technology-specific jobs and their stage placement.

Test components locally with critical dependency validation. Run pipeline with components referencing $CI_COMMIT_SHA. Verify init-values produces dotenv artifacts, test produces coverage, build depends on test, sonar consumes both init-values and test artifacts. Validate all inputs work with various values. Confirm artifact flow and dependencies function correctly.

Publish component library. Enable CI/CD Catalog project setting in repository. Create semantic version tag. Add release job using release-cli image. Test component consumption from catalog. Document versioning, stage flow requirements, and upgrade guidance.
</instructions>

## Framework Context

Standardized Stage Flow: ALL component libraries MUST follow the 7-stage architecture: [prepare, test, build, verify, package, publish, release]. This ensures consistent user experience across all technology stacks while allowing technology-specific job implementations within each stage.

Critical Dependencies: MANDATORY dependency patterns ensure proper artifact flow: init-values produces dotenv artifacts, test produces coverage artifacts, build depends on test completion, sonar depends on init-values + test artifacts, packaging jobs depend on build artifacts. These patterns are technology-agnostic and must be implemented consistently.

Component Library Pattern: GitLab CI/CD component libraries organize related components (common, review, build) into cohesive units. Common templates define reusable job patterns with separated .test-job and .build-job templates, while review and build components extend these for specific workflow contexts. GitLab supports up to 100 components per project.

Spec Inputs: Components declare inputs in spec section with defaults and descriptions. Inputs use `$[[ inputs.name ]]` interpolation syntax. Always provide defaults for usability. Document every input in README with type, default value, and description.

Component Composition: Use local includes to compose components. Review and build components typically include common.yml to extend shared job templates. This promotes DRY principles and consistency across workflows.

Workflow Separation: Review pipelines validate merge requests with test→build→verify flow without publishing. Build pipelines execute on main branch with test→build→package→publish flow. Use workflow rules and component-level rules for conditional execution. GitLab supports multiple pipeline types (branch, tag, merge request) that can run simultaneously - workflow rules prevent duplicate pipeline execution.

Testing Requirements: Component repositories must include sample source code for testing. Components are tested locally using `$CI_COMMIT_SHA` reference before publication. Validate all inputs work correctly with various configurations.

Security Requirements: Component projects must implement security best practices. Use protected branches for releases with merge request approvals. Enable 2FA for all maintainers. Sign commits for integrity verification. Store secrets securely using protected CI/CD variables which are only accessible to protected branch and tag pipelines, preventing credential exposure in fork merge requests. Use minimally scoped tokens. Audit component code before consumption. Pin component versions to specific commits or tags when consuming components.

## Output Format

Repository structure:

```
{project-root}/
├── templates/
│   ├── common.yml      # Shared job templates (.build-job, .test-job)
│   ├── review.yml      # MR validation pipeline
│   └── build.yml       # Main branch pipeline
├── .gitlab-ci.yml      # Component orchestration with conditional includes
├── README.md           # Documentation with inputs table and usage examples
├── LICENSE.md
└── {sample-code}/      # Tech-stack-specific source code for testing
```

Standardized stage and job template pattern:

```yaml
spec:
  inputs:
    stage_prepare:
      default: 'prepare'
    stage_test:
      default: 'test'
    stage_build:
      default: 'build'
    container_image:
      default: 'alpine:latest'
---
# MANDATORY: Separated test and build jobs
.test-job:
  stage: $[[ inputs.stage_test ]]
  image: $[[ inputs.container_image ]]
  script:
    - {{test_command}}
  artifacts:
    paths: [coverage/]

.build-job:
  stage: $[[ inputs.stage_build ]]
  image: $[[ inputs.container_image ]]
  script:
    - {{build_command}}
  artifacts:
    paths: [dist/]
```

Standardized orchestration pattern:

```yaml
# MANDATORY: 7-stage flow
stages: [prepare, test, build, verify, package, publish, release]

include:
  - component: $CI_SERVER_FQDN/$CI_PROJECT_PATH/review@$CI_COMMIT_SHA
    inputs:
      stage_prepare: prepare
      stage_test: test
      stage_build: build
      container_image: 'node:18'
    rules:
      - if: $CI_PIPELINE_SOURCE == "merge_request_event"
```

<success_criteria>
- Repository contains templates/ directory with common.yml, review.yml, and build.yml components
- Common template defines reusable hidden job templates extended by review and build workflows
- Review component executes on merge requests with validation jobs but no publishing
- Build component executes on main branch with packaging and publishing jobs
- All spec:inputs have default values and descriptions, documented in README inputs table
- Main .gitlab-ci.yml orchestrates conditional component inclusion with workflow rules
- Sample source code included for tech stack testing with working build/test commands
- README documents library purpose, available components, inputs, usage examples, and required variables
- Components tested locally using $CI_COMMIT_SHA reference with successful pipeline execution
- Component library published with semantic version tag and discoverable in CI/CD catalog
</success_criteria>

## Execution Checklist

### Setup Phase

- Gather tech stack details: language, framework, build tool, test framework, container images
- Identify quality tools: linters, formatters, SonarQube configuration, security scanners
- Determine artifacts: build outputs, coverage reports, Docker images, Helm charts
- Initialize repository: Create Git repo with clear project description, add .gitignore, LICENSE.md, README.md structure
- Configure security: Enable protected branches for main and release branches, require 2FA for maintainers, set up merge request approvals
- Create templates/ directory: Establish directory structure for component YAML files (max 100 components per project)
- Add sample code: Include minimal working source code for tech stack testing

### Implementation Phase

- Create common.yml: Define hidden job templates (.build-job, .test-job, .lint-job, .quality-job)
- Define common inputs: Add spec:inputs for stages, container images, tech-specific parameters
- Implement caching: Configure dependency caching patterns (go.sum, package-lock.json, requirements.txt)
- Create review.yml: Include common.yml, extend templates, add MR-specific jobs (lint, test, verify)
- Create build.yml: Include common.yml, extend templates, add build-specific jobs (package, publish, tag)
- Configure orchestration: Write .gitlab-ci.yml with workflow rules and conditional component includes
- Document library: Create README with inputs table, usage examples, features, and variable requirements
- Configure release job: Add create-release job using release-cli for semantic versioning

### Publication Phase

- Test locally: Run pipeline using $CI_COMMIT_SHA component references, verify all jobs execute
- Validate workflows: Confirm review runs on MRs, build runs on main, inputs work as expected
- Security review: Audit component for secrets, verify token usage, check for malicious content
- Enable catalog: Set CI/CD Catalog project setting and ensure project description is set
- Create release: Tag with signed commit using semantic version, trigger release job, verify catalog publication
- Test consumption: Include published component from catalog using pinned version (commit SHA or tag, avoid ~latest)
- Document usage: Add versioning guidance, security considerations, and upgrade instructions to README
