# GitLab CI/CD Component Library Scaffold

<instructions>
This template scaffolds a complete GitLab CI/CD component library with standardized 7-stage flow and critical dependencies.
Replace all {{variable_name}} placeholders with tech-stack-specific values.
The library enforces mandatory stage sequence [prepare, test, build, verify, package, publish, release] and critical dependency patterns while allowing technology-specific job implementations.
After generating, customize job scripts, caching strategies, and quality tools for your tech stack within the standardized architecture.
</instructions>

## templates/common.yml

```yaml
spec:
  inputs:
    stage_prepare:
      default: 'prepare'
      description: 'Preparation stage name'
    stage_test:
      default: 'test'
      description: 'Test stage name'
    stage_build:
      default: 'build'
      description: 'Build stage name'
    codebase_name:
      default: '{{project_name}}'
      description: 'Project name'
    container_image:
      default: '{{default_container_image}}'
      description: '{{tech_stack}} container image with version'
    chart_dir:
      default: 'deploy-templates'
      description: 'Helm chart directory (if applicable)'
---

# Common variables
.common-variables:
  variables:
    CODEBASE_NAME: "$[[ inputs.codebase_name ]]"
    CHART_DIR: "$[[ inputs.chart_dir ]]"

# Technology-specific dependency caching template
.dependency-cache:
  variables:
    CACHE_DIR: "${CI_PROJECT_DIR}/.cache"
  before_script:
    - mkdir -p ${CACHE_DIR}
  cache:
    key:
      files:
        - {{dependency_lock_file}}  # e.g., go.sum, package-lock.json, requirements.txt
    paths:
      - ${CACHE_DIR}

# MANDATORY: Separated test job template (NO building)
.test-job:
  stage: $[[ inputs.stage_test ]]
  image: $[[ inputs.container_image ]]
  extends: .dependency-cache
  script:
    - {{test_command}}   # e.g., npm test, go test, pytest
  artifacts:
    paths:
      - coverage/
    reports:
      coverage_report:
        coverage_format: {{coverage_format}}  # e.g., cobertura, jacoco
        path: {{coverage_file}}  # e.g., coverage.xml
    expire_in: 1 week
    when: always

# MANDATORY: Separated build job template (NO testing)
.build-job:
  stage: $[[ inputs.stage_build ]]
  image: $[[ inputs.container_image ]]
  extends: .dependency-cache
  script:
    - {{build_command}}  # e.g., npm run build, make build, python setup.py build
  artifacts:
    paths:
      - {{build_output_dir}}  # e.g., dist/, build/, target/
    expire_in: 1 week
    when: always

# Linting job template
.lint-job:
  stage: $[[ inputs.stage_test ]]
  image: $[[ inputs.container_image ]]
  script:
    - {{lint_command}}    # e.g., golangci-lint run, npm run lint, flake8
    - {{format_command}}  # e.g., go fmt, prettier --check, black --check

# MANDATORY: SonarQube analysis base (runs in build stage after test)
.sonar-base:
  stage: $[[ inputs.stage_build ]]
  image:
    name: sonarsource/sonar-scanner-cli:latest
    entrypoint: [""]
```

## templates/review.yml

```yaml
spec:
  inputs:
    stage_prepare:
      default: 'prepare'
    stage_build:
      default: 'build'
    stage_test:
      default: 'test'
    stage_verify:
      default: 'verify'
    codebase_name:
      default: '{{project_name}}'
    container_image:
      default: '{{default_container_image}}'
    chart_dir:
      default: 'deploy-templates'
---

# Include common templates
include:
  - local: 'templates/common.yml'
    inputs:
      stage_prepare: $[[ inputs.stage_prepare ]]
      stage_build: $[[ inputs.stage_build ]]
      stage_test: $[[ inputs.stage_test ]]
      codebase_name: $[[ inputs.codebase_name ]]
      container_image: $[[ inputs.container_image ]]
      chart_dir: $[[ inputs.chart_dir ]]

# MANDATORY: Standardized 7-stage pipeline
stages:
  - $[[ inputs.stage_prepare ]]
  - $[[ inputs.stage_test ]]
  - $[[ inputs.stage_build ]]
  - $[[ inputs.stage_verify ]]

# MANDATORY: Initialize values (produces dotenv artifacts)
init-values:
  stage: $[[ inputs.stage_prepare ]]
  image: alpine:latest
  script:
    - echo "BRANCH_NAME=${CI_COMMIT_REF_NAME}" > branch.env
    - echo "PROJECT_NAME=${CODEBASE_NAME}" >> branch.env
  artifacts:
    reports:
      dotenv: branch.env

# MANDATORY: Run tests with coverage (NO building)
test:
  extends: .test-job
  needs: []
  dependencies: []

# MANDATORY: Build application (depends on test completion)
build:
  extends: .build-job
  needs: [test]
  dependencies: []

# Technology-specific jobs (customize as needed)
lint:
  extends: .lint-job
  needs: []

# MANDATORY: SonarQube analysis (depends on init-values + test)
sonar:
  extends: .sonar-base
  needs: [init-values, test]
  dependencies: [init-values, test]
  script:
    - >-
      if [ -n "$SONAR_ORG" ]; then
        SONAR_PROJECT_KEY="${SONAR_ORG}_${CODEBASE_NAME}";
      else
        SONAR_PROJECT_KEY="${CODEBASE_NAME}";
      fi

      sonar-scanner
      -Dsonar.projectKey=${SONAR_PROJECT_KEY}
      -Dsonar.projectName=${CODEBASE_NAME}
      -Dsonar.host.url=${SONAR_HOST_URL}
      -Dsonar.token=${SONAR_TOKEN}
      -Dsonar.pullrequest.key=${CI_MERGE_REQUEST_IID}
      -Dsonar.pullrequest.branch=${CI_MERGE_REQUEST_SOURCE_BRANCH_NAME}
      -Dsonar.pullrequest.base=${CI_MERGE_REQUEST_TARGET_BRANCH_NAME}

# Docker build verification (no push)
docker-verify:
  stage: $[[ inputs.stage_verify ]]
  image:
    name: moby/buildkit:rootless
    entrypoint: [""]
  script:
    - buildctl-daemonless.sh build
      --frontend dockerfile.v0
      --local context=.
      --local dockerfile=.
      --output type=image,name=verify:latest,push=false
```

## templates/build.yml

```yaml
spec:
  inputs:
    stage_prepare:
      default: 'prepare'
    stage_build:
      default: 'build'
    stage_test:
      default: 'test'
    stage_package:
      default: 'package'
    stage_publish:
      default: 'publish'
    codebase_name:
      default: '{{project_name}}'
    container_image:
      default: '{{default_container_image}}'
    chart_dir:
      default: 'deploy-templates'
    image_registry:
      default: 'docker.io/${DOCKERHUB_USERNAME}'
---

include:
  - local: 'templates/common.yml'
    inputs:
      stage_prepare: $[[ inputs.stage_prepare ]]
      stage_build: $[[ inputs.stage_build ]]
      stage_test: $[[ inputs.stage_test ]]
      codebase_name: $[[ inputs.codebase_name ]]
      container_image: $[[ inputs.container_image ]]
      chart_dir: $[[ inputs.chart_dir ]]

variables:
  IMAGE_REGISTRY: $[[ inputs.image_registry ]]

# MANDATORY: Standardized 7-stage pipeline
stages:
  - $[[ inputs.stage_prepare ]]
  - $[[ inputs.stage_test ]]
  - $[[ inputs.stage_build ]]
  - $[[ inputs.stage_package ]]
  - $[[ inputs.stage_publish ]]

# MANDATORY: Initialize values (produces dotenv artifacts)
init-values:
  stage: $[[ inputs.stage_prepare ]]
  image: alpine:latest
  script:
    - echo "BRANCH_NAME=${CI_MERGE_REQUEST_TARGET_BRANCH_NAME:-${CI_COMMIT_REF_NAME}}" > build.env
    - echo "PROJECT_NAME=${CODEBASE_NAME}" >> build.env
    - echo "BUILD_VERSION=${CI_COMMIT_SHORT_SHA}" >> build.env
    - echo "VCS_TAG=${CI_COMMIT_TAG:-v0.1.0-${CI_COMMIT_SHORT_SHA}}" >> build.env
    - echo "IMAGE_TAG=${CI_COMMIT_TAG:-${CI_COMMIT_SHORT_SHA}}" >> build.env
  artifacts:
    reports:
      dotenv: build.env

# MANDATORY: Run tests with coverage (NO building)
test:
  extends: .test-job
  needs: []
  dependencies: []

# MANDATORY: Build application (depends on test completion)
build:
  extends: .build-job
  needs: [test]
  dependencies: []

# Technology-specific jobs (customize as needed)
lint:
  extends: .lint-job
  needs: []

# MANDATORY: SonarQube analysis (depends on init-values + test)
sonar:
  extends: .sonar-base
  needs: [init-values, test]
  dependencies: [init-values, test]
  script:
    - >-
      if [ -n "$SONAR_ORG" ]; then
        SONAR_PROJECT_KEY="${SONAR_ORG}_${CODEBASE_NAME}";
      else
        SONAR_PROJECT_KEY="${CODEBASE_NAME}";
      fi

      sonar-scanner
      -Dsonar.projectKey=${SONAR_PROJECT_KEY}
      -Dsonar.projectName=${CODEBASE_NAME}
      -Dsonar.host.url=${SONAR_HOST_URL}
      -Dsonar.token=${SONAR_TOKEN}
      -Dsonar.branch.name=${CI_COMMIT_REF_NAME}

# MANDATORY: Container building (depends on init-values + build artifacts)
buildkit-build:
  stage: $[[ inputs.stage_package ]]
  image:
    name: moby/buildkit:rootless
    entrypoint: [""]
  script:
    - echo "{\"auths\":{\"https://index.docker.io/v1/\":{\"username\":\"${DOCKERHUB_USERNAME}\",\"password\":\"${DOCKERHUB_PASSWORD}\"}}}" > ~/.docker/config.json
    - buildctl-daemonless.sh build
      --frontend dockerfile.v0
      --local context=.
      --local dockerfile=.
      --output type=image,name=${IMAGE_REGISTRY}/${CODEBASE_NAME}:${IMAGE_TAG},push=true
  needs:
    - init-values
    - build
    - sonar
  dependencies:
    - init-values
    - build

# Technology-specific package publishing (customize as needed)
{{package_publish_job}}:
  stage: $[[ inputs.stage_package ]]
  image: $[[ inputs.container_image ]]
  script:
    - {{package_publish_command}}  # e.g., npm publish, mvn deploy
  needs:
    - init-values
    - build
  dependencies: [init-values, build]

# MANDATORY: Git tagging (depends on init-values artifacts)
git-tag:
  stage: $[[ inputs.stage_publish ]]
  image: alpine/git:latest
  script:
    - git config --global user.email "${GITLAB_USER_EMAIL:-ci@gitlab.com}"
    - git config --global user.name "${GITLAB_USER_NAME:-GitLab CI}"
    - git remote set-url origin "https://gitlab-ci:${GITLAB_ACCESS_TOKEN}@${CI_SERVER_HOST}/${CI_PROJECT_PATH}.git"
    - git tag -a "${VCS_TAG}" -m "Automated release tag"
    - git push origin "${VCS_TAG}"
  needs:
    - buildkit-build
    - init-values
  dependencies:
    - init-values
```

## .gitlab-ci.yml

```yaml
workflow:
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
    - if: $CI_COMMIT_REF_PROTECTED == "true"
    - if: $CI_COMMIT_TAG =~ /^\d+\.\d+\.\d+$/

variables:
  CODEBASE_NAME: "{{project_name}}"
  CONTAINER_IMAGE: "{{default_container_image}}"
  IMAGE_REGISTRY: "docker.io/{{docker_registry_user}}"

include:
  - component: $CI_SERVER_FQDN/$CI_PROJECT_PATH/review@$CI_COMMIT_SHA
    inputs:
      stage_build: build
      stage_test: test
      codebase_name: ${CODEBASE_NAME}
      container_image: ${CONTAINER_IMAGE}
    rules:
      - if: $CI_PIPELINE_SOURCE == "merge_request_event"

  - component: $CI_SERVER_FQDN/$CI_PROJECT_PATH/build@$CI_COMMIT_SHA
    inputs:
      stage_build: build
      stage_test: test
      stage_package: package
      stage_publish: publish
      codebase_name: ${CODEBASE_NAME}
      container_image: ${CONTAINER_IMAGE}
      image_registry: ${IMAGE_REGISTRY}
    rules:
      - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH || $CI_COMMIT_REF_PROTECTED == "true"

# MANDATORY: Standardized 7-stage flow
stages: [prepare, test, build, verify, package, publish, release]

create-release:
  stage: release
  image: registry.gitlab.com/gitlab-org/release-cli:latest
  rules:
    - if: $CI_COMMIT_TAG =~ /^\d+\.\d+\.\d+$/
  script: echo "Creating release $CI_COMMIT_TAG"
  release:
    tag_name: $CI_COMMIT_TAG
    description: "Release $CI_COMMIT_TAG"
```

## README.md Structure

```markdown
# {{project_name}} CI/CD Component Library

Reusable GitLab CI/CD components for {{tech_stack}} projects providing standardized review and build pipelines.

## Components

### Common (templates/common.yml)
Base job templates extended by review and build workflows.

### Review (templates/review.yml)
Merge request validation pipeline with linting, testing, and verification.

### Build (templates/build.yml)
Main branch pipeline with building, packaging, publishing, and tagging.

## Inputs

| Name | Description | Default |
|------|-------------|---------|
| stage_prepare | Preparation stage | `prepare` |
| stage_build | Build stage | `build` |
| stage_test | Test stage | `test` |
| stage_package | Package stage | `package` |
| stage_publish | Publish stage | `publish` |
| codebase_name | Project name | `{{project_name}}` |
| container_image | {{tech_stack}} image | `{{default_container_image}}` |
| image_registry | Docker registry | `docker.io/${DOCKERHUB_USERNAME}` |
| chart_dir | Helm chart directory | `deploy-templates` |

## Usage

```yaml
include:
  - component: gitlab.example.com/your-org/{{project_name}}-ci/review@1.0.0
    inputs:
      codebase_name: 'my-app'
      container_image: '{{alternative_image}}'
```

## Required CI/CD Variables

- `SONAR_HOST_URL`: SonarQube server URL
- `SONAR_TOKEN`: SonarQube authentication token
- `DOCKERHUB_USERNAME`: Docker Hub username
- `DOCKERHUB_PASSWORD`: Docker Hub password
- `GITLAB_ACCESS_TOKEN`: GitLab access token for tagging

## Features

- {{tech_stack}} build and test with caching
- Code coverage reporting
- SonarQube integration
- Docker image building and publishing
- Helm chart validation
- Git tagging and versioning
