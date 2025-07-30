# Task: Create GitHub Issues

## Description

Create GitHub issues (epics, stories, bugs) for KubeRocketCI/kuberocketai repository using existing GitHub templates and GitHub MCP server.

## Prerequisites

- [ ] **GitHub MCP Server**: Available and configured
- [ ] **Repository Access**: KubeRocketCI/kuberocketai with issue creation permissions

## Workflow

### 1. Gather Requirements

Ask user to specify:

- **Type**: Epic, Story, or Bug
- **Title**: Clear, descriptive title
- **Details**: All necessary information for the issue type

### 2. Confirm Before Creation

**CRITICAL**: Always confirm with user before creating any GitHub issue:

- Show what will be created
- Ask for explicit approval: "Should I create this issue?"
- Only proceed after user confirms "yes"

### 3. Create Issue

**Template Mapping**:

- **Epic** → Use `feature_request.yml` template
- **Story** → Use `enhancement.yml` template
- **Bug** → Use `bug_report.yml` template

**Repository**: `KubeRocketCI/kuberocketai`
**Default Assignee**: `SergK`

### GitHub Issue Creation

Use GitHub MCP server to create issues with appropriate template structure.

## Success Criteria

- [ ] **User Confirmation**: User explicitly approved issue creation
- [ ] **Issue Created**: GitHub issue successfully created
- [ ] **Correct Template**: Appropriate GitHub template used
- [ ] **Proper Labels**: Template-required labels applied
- [ ] **URL Provided**: GitHub issue URL returned to user

## Important Notes

- **Never create issues without user confirmation**
- GitHub templates handle all field requirements automatically
- Link stories to epics by including "Related to Epic #[number]" in story description
- All created issues auto-assigned to SergK as repository owner
