# Task: PowerPoint Review

## Description

Provide guidance/review/improve for a PowerPoint presentation requested by the user.

## Prerequisites

- [ ] **Presentation exists**: Presentation for review is presented as a file path or attached to the chat
- [ ] **Page is allowed to read**: The file has read permissions for everyone

## Instructions

1. **Check dependencies**: Check if the "Office-PowerPoint-MCP-Server" MCP server is installed. If it is not installed, notify the user that you need to install a dependency (Office-PowerPoint-MCP-Server). If user accepts, install it. GitHub reference with package and instructions: `https://github.com/GongRzhe/Office-PowerPoint-MCP-Server`
2. **Read presentation**: Read the presentation user provided you with by specifying a file path or attaching it to the chat. If nothing is specified, ask him for it
3. **Create venv**: Create a Python virtual environment. Call it `venv` if this name is not already taken. Notify user about creating a virtual environment
4. **Perform commands**: Fulfill the user's request. Notify the user that to you need to create a number of scripts in the `presentation-processing` folder in the root folder of the project. If user provided you with a PowerPoint presentation, create create a copy of this presentation in the `presentation-processing`, name the copy as `<presentation-processing>-edited.pptx`
5. **Clean up environment**: Clean up the virtual environment
6. **Notify about completion**: Tell the user that all the edited version is located in the `<presentation-processing>-edited.pptx` file of the `presentation-processing` folder in the root of the project

## Output Format

As a result, I expect you to provide the user with the `<presentation-processing>-edited.pptx` file of the `presentation-processing` folder in the root of the project

## Success criteria

- [ ] **Review completion**: The presentation is reviewed
- [ ] **Style consistency**: The page follows guidelines and best practices declared in the 'Microsoft Writing Style Guide'.

## Execution Checklist

### Execution Phase

- [ ] **Impersonate**: Ensure the page page doesn't contain too many "you", "your", "we", "us", etc.
- [ ] **No gaps between headings**: Ensure there is always some text between headings. That is, there is always text beetween, let's say, "## Heading" and "### Heading"
- [ ] **Image processing**: Remind the user to set borders for images of 1px width and #DCDCDC color

### Validation Phase

- [ ] **Completeness check**: Verify the new `<original-file>-reviewed.md` file is created and contains your refinements
- [ ] **Style compliance**: Ascertain that the page stick to the 'Microsoft Writing Style Guide'
- [ ] **Cleanup**: The Python virtual environment is cleaned up to keep the project clean and neat
