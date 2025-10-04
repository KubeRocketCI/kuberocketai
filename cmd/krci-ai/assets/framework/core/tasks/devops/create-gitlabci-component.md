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

Scaffold production-ready GitLab CI/CD component library with review and build pipelines for any tech stack. Component libraries provide reusable job templates, separate review/build workflows, and standardized CI/CD patterns that accelerate pipeline development across projects.

<instructions>
Gather tech stack requirements. Ask for programming language, framework, build tool, test framework, container image, and typical build/test commands. Identify code quality tools (linters, formatters, SonarQube). Determine deployment artifacts and Helm chart requirements if applicable.

Create repository structure. Initialize Git repository with templates/ directory. Create common.yml for shared job templates, review.yml for merge request pipeline, build.yml for main branch pipeline. Add README.md, LICENSE.md, .gitlab-ci.yml, and .gitignore. Include sample source code for testing components.

Define shared job templates in common.yml. Create hidden jobs (prefixed with .) for build, test, lint, and quality checks. Define spec:inputs for stage names, container image, and tech-stack-specific parameters. Implement caching patterns for dependencies. Use $[[ inputs.name ]] interpolation throughout.

Implement review pipeline in review.yml. Include common.yml with local include pattern. Extend shared job templates for MR validation. Add jobs for code formatting, linting, unit tests, quality analysis, Dockerfile validation, and Docker build verification without publishing. Configure component inputs for customization.

Implement build pipeline in build.yml. Include common.yml and extend shared templates. Add initialization job for version/metadata. Include packaging stage for Docker image build and push. Add Git tagging job for version management. Configure registry and credential variables as inputs.

Create orchestration in main .gitlab-ci.yml. Define workflow rules for merge requests and protected branches. Include review component conditionally for merge requests. Include build component conditionally for main branch. Pass global variables as component inputs. Define all pipeline stages.

Document component library in README.md. Explain library purpose and available components. Create inputs table documenting all parameters with defaults. Provide usage examples for including components. List required CI/CD variables. Document features and integration patterns.

Test components locally. Run pipeline with components referencing $CI_COMMIT_SHA. Verify review workflow executes on merge requests. Verify build workflow executes on main branch. Validate all inputs work with various values. Confirm caching, artifacts, and dependencies function correctly.

Publish component library. Enable CI/CD Catalog project setting in repository. Create semantic version tag. Add release job using release-cli image. Test component consumption from catalog. Document versioning and upgrade guidance.
</instructions>

## Framework Context

Component Library Pattern: GitLab CI/CD component libraries organize related components (common, review, build) into cohesive units. Common templates define reusable job patterns, while review and build components extend these for specific workflow contexts. GitLab supports up to 100 components per project.

Spec Inputs: Components declare inputs in spec section with defaults and descriptions. Inputs use `$[[ inputs.name ]]` interpolation syntax. Always provide defaults for usability. Document every input in README with type, default value, and description.

Component Composition: Use local includes to compose components. Review and build components typically include common.yml to extend shared job templates. This promotes DRY principles and consistency across workflows.

Workflow Separation: Review pipelines validate merge requests with linting, testing, and verification without publishing. Build pipelines execute on main branch with packaging, artifact publishing, and tagging. Use workflow rules and component-level rules for conditional execution. GitLab supports multiple pipeline types (branch, tag, merge request) that can run simultaneously - workflow rules prevent duplicate pipeline execution.

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

Common template pattern:

```yaml
spec:
  inputs:
    stage_build:
      default: 'build'
    container_image:
      default: 'alpine:latest'
---
.build-job:
  stage: $[[ inputs.stage_build ]]
  image: $[[ inputs.container_image ]]
  script:
    - make build
```

Orchestration pattern:

```yaml
include:
  - component: $CI_SERVER_FQDN/$CI_PROJECT_PATH/review@$CI_COMMIT_SHA
    inputs: { stage_build: build, container_image: 'node:18' }
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
