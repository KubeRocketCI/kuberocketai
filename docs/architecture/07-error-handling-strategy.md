# 7. Error Handling Strategy

KubeRocketAI uses a two-tier error handling approach: CLI-detected errors (caught by krci-ai tool) and LLM-detected errors (caught during AI agent execution).

## 7.1 CLI Error Detection

The CLI tool catches errors before LLM execution starts. These are structural and file system errors that prevent the system from running properly.

CLI catches:

- Missing framework files (agents, tasks, templates, data)
- File permission issues and invalid file paths
- Invalid YAML frontmatter in components
- Missing required fields in component definitions
- Framework installation failures
- Directory creation permission issues

CLI error response:

- Exit with non-zero code
- Display clear error message with file path
- Provide actionable resolution steps

Examples:

```bash
[ERROR] Agent file not found: ./krci-ai/agents/developer.md
[ERROR] Invalid YAML frontmatter in: ./krci-ai/tasks/deploy.md
[ERROR] Permission denied creating directory: ./krci-ai/
[ERROR] Missing required field 'commands' in agent definition
```

## 7.2 LLM Error Detection

The LLM catches errors during runtime execution. These are logical and contextual errors that can only be detected when the AI agent is processing user requests.

LLM catches:

- Referenced components that don't exist during execution
- Invalid component relationships discovered at runtime
- Missing template variables or data references
- Circular dependencies in agent workflows
- Ambiguous instructions or unclear requirements
- Conflicting agent principles or task instructions
- Failed task execution due to missing dependencies
- Template rendering failures

LLM error response:

- Continue with degraded functionality when possible
- Provide user-friendly error explanation
- Suggest alternative approaches or fixes
- Request clarification when needed

Examples:

```bash
[ERROR] Referenced template not found during execution: analyze-template.md
[ERROR] Circular dependency detected: `agent-a` → `task-b` → `agent-a`
[ERROR] Unable to process ambiguous instruction: "make it better"
[ERROR] Missing required data for task execution: project-specs.md
```

## 7.3 Implementation Guidelines

For CLI tool:

- Validate all file references before starting LLM
- Check YAML structure and required fields
- Verify permissions and file accessibility
- Provide helpful error messages with context

For LLM processing:

- Load all components and validate references
- Check for circular dependencies
- Handle missing components gracefully
- Provide fallback behavior for common errors
- Give clear explanations to users about what went wrong
