---
dependencies:
  data:
    - ../../.github/ISSUE_TEMPLATE/enhancement.yml
    - ../../.github/ISSUE_TEMPLATE/feature_request.yml
    - ../../.github/ISSUE_TEMPLATE/bug_report.yml
  mcp_servers:
    - github
---

# Task: Create GitHub Issues

## Description

Create GitHub issues (epics, stories, bugs) for KubeRocketCI/kuberocketai repository using existing GitHub templates and GitHub MCP server.

## Prerequisites

<prerequisites>
- GitHub MCP Server: Available and configured
- Repository Access: KubeRocketCI/kuberocketai with issue creation permissions
</prerequisites>

## Instructions

<instructions>

### 1. Gather Requirements

Ask user to specify:

- Type: Epic, Story (Enhancement), or Bug
- Title: Clear, descriptive title
- Repository: Default `KubeRocketCI/kuberocketai` (allow override)
- Labels: Optional extra labels
- Assignees: Default `SergK` (allow override)
- Related Epic #: Optional (adds a note at the end of the body)

Then, based on the selected Type, prompt for fields derived from the corresponding template (see Template-Driven Rendering):

- Epic (feature_request.yml):
  - Feature Summary (required)
  - Problem Statement (required)
  - Proposed Solution (required)
  - Alternative Solutions (optional)
  - Usage Examples (optional)
  - Acceptance Criteria (optional)

- Story/Enhancement (enhancement.yml):
  - Current Functionality (required)
  - Current Limitations (required)
  - Proposed Improvement (required)
  - Expected Benefits (required)
  - Implementation Examples (optional)
  - Testing Considerations (optional)

- Bug (bug_report.yml):
  - Bug Description (required)
  - Steps to Reproduce (required)
  - Expected Behavior (required)
  - Actual Behavior (required)
  - Error Logs/Output (optional)

### 2. Preview & Confirm Before Creation

CRITICAL: Always confirm with user before creating any GitHub issue:

- Show a full preview of the issue body rendered from the selected template and provided fields (H2 headers per textarea label, in template order)
- Ask for explicit approval: "Should I create this issue?"
- Only proceed after user confirms "yes"

### 3. Label Strategy

Base Labels: Use template defaults from `.github/ISSUE_TEMPLATE/*.yml`

Repository Extensions (add if conditions match):

- All Epics → +`epic`
- Breaking changes mentioned → +`breaking-change`
- Technical debt scope → +`technical-debt`
- High priority/complexity → +`critical`

### 4. Create Issue

### Template-Driven Rendering (for API/MCP creation)

When creating issues programmatically, derive the output structure from the corresponding GitHub Issue Template to keep a single source of truth.

- Locate template by type:
  - Epic → `.github/ISSUE_TEMPLATE/feature_request.yml`
  - Story/Enhancement → `.github/ISSUE_TEMPLATE/enhancement.yml`
  - Bug → `.github/ISSUE_TEMPLATE/bug_report.yml`

- Parse the template YAML and render sections in order:
  - For each `body` item with `type: textarea`, use `attributes.label` as a Markdown H2 header (e.g., `## {label}`)
  - Preserve item ordering from the template
  - Optionally include `attributes.description` as helper text under the header (plain text), if needed
  - Respect `validations.required`; HALT if any required textarea is missing in user input

- Metadata handling:
  - Title: use the template `title` prefix unless a custom title is provided by the user
  - Labels: include template default labels plus any user-specified labels
  - Assignees: include template default assignees unless overridden
  - Non-textarea fields (dropdowns, inputs, checkboxes): if user provided values, include a short "Metadata" section listing key-value pairs

- Validation:
  - CRITICAL: If the mapped template file is missing, HALT and report the exact missing path; do not create the issue
  - If any required field per template is missing, HALT and list missing fields
  - If `Type` is not one of: Epic, Story (Enhancement), Bug — HALT and show the allowed values

- Conventions:
  - Keep content concise and outcome-focused; avoid command-level testing instructions in high-level sections
  - Maintain template order and naming to match the UI form experience

  </instructions>

## Output Format

- Location: GitHub issue in the target repository (return created issue URL)
- Type: Epic → Feature Request, Story → Enhancement, Bug → Bug Report
- Title: Use template title prefix unless user provides an explicit title override
- Labels/Assignees: Apply template defaults plus any user-provided additions

## Execution Checklist

### Discovery Phase

- Validate Type is one of: Epic, Story (Enhancement), Bug
- Verify mapped template file exists for the selected Type
- Collect required fields as defined by the template (textareas with required: true)

### Planning Phase

- Render preview using template labels as H2 headers, in template order
- Validate all required fields populated; list any missing and HALT
- Apply label strategy: template defaults + repository extensions
- Confirm labels and assignees (template defaults + repository patterns)

### Creation Phase

- Create issue via GitHub MCP server
- Append `Relates to: #<n>` if provided
- Return issue URL to the user

## Success Criteria

<success_criteria>
- User Confirmation: User explicitly approved issue creation
- Issue Created: GitHub issue successfully created
- Correct Template: Appropriate GitHub template used
- Proper Labels: Template-required labels applied
- URL Provided: GitHub issue URL returned to user
</success_criteria>

## Important Notes

- Never create issues without user confirmation
- GitHub templates handle all field requirements automatically
- To link to an epic, provide the "Related Epic #"; the agent will append `Relates to: #<number>` automatically
- All created issues auto-assigned to SergK as repository owner
