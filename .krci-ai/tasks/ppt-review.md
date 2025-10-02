---
dependencies:
  data:
  mcp_servers:
    - office-powerpoint
---

# Task: PowerPoint Review

## Description

Provide guidance/review/improve for a PowerPoint presentation requested by the user.

## Instructions

<instructions>
Confirm the presentation for review is presented as a file path or attached to the chat, and the file has read permissions. Ensure dependencies declared in the YAML frontmatter (Office-PowerPoint-MCP-Server MCP server) are available before proceeding. If not installed, notify the user about the dependency requirement and install if accepted (GitHub reference: `https://github.com/GongRzhe/Office-PowerPoint-MCP-Server`).

Read the presentation user provided by file path or attachment. Ask user for it if nothing is specified. Create a Python virtual environment named `venv` if not already taken and notify user. Fulfill the user's request by creating scripts in the `presentation-processing` folder in the root. Create a copy of the presentation in `presentation-processing` named `presentation-name-edited.pptx`. Clean up the virtual environment after completion and notify the user that the edited version is located in the `presentation-name-edited.pptx` file.
</instructions>

## Output Format

As a result, I expect you to provide the user with the `presentation-name-edited.pptx` file of the `presentation-processing` folder in the root of the project

## Success Criteria

<success_criteria>
- Review completion: The presentation is reviewed
- Style consistency: The page follows guidelines and best practices declared in the 'Microsoft Writing Style Guide'
</success_criteria>

## Execution Checklist

### Execution Phase

<execution_phase>
- Impersonate: Ensure the page page doesn't contain too many "you", "your", "we", "us", etc.
- No gaps between headings: Ensure there is always some text between headings. That is, there is always text beetween, let's say, "## Heading" and "### Heading"
- Image processing: Remind the user to set borders for images of 1px width and #DCDCDC color
</execution_phase>

### Validation Phase

<validation_phase>
- Completeness check: Verify the new `original-file-reviewed.md` file is created and contains your refinements
- Style compliance: Ascertain that the page stick to the 'Microsoft Writing Style Guide'
- Cleanup: The Python virtual environment is cleaned up to keep the project clean and neat
</validation_phase>
