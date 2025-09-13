# Task: Core Create Task

## Description

Guide user through creating framework-compliant tasks following KubeRocketAI patterns. This task provides step-by-step instructions for creating tasks that use XML guidance tags, inline references, and proper structure for LLM processing within the AI-as-Code framework.

<prerequisites>
- Framework installed: .krci-ai directory exists with agents/, tasks/, templates/, data/ structure
- Target agent identified: Specific agent that will use this new task
- Task purpose defined: Clear understanding of what the task will accomplish
- Framework patterns understood: Knowledge of XML tags, inline references, and validation requirements
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/templates/krci-ai/core-task-template.md
- ./.krci-ai/data/krci-ai/core-validation-checklist.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Apply XML guidance system: Use XML tags (`<prerequisites>`, `<instructions>`, `<success_criteria>`) for LLM section identification and guidance
2. Follow naming convention: Use descriptive task names with appropriate prefixes (e.g., "core-" for framework tasks)
3. Create inline references: Use `[filename](./.krci-ai/path/to/file)` format for all framework references
4. Structure validation: Follow task template structure with Description, Prerequisites, Instructions, Success Criteria
5. Ensure self-contained guidance: Include explanations of XML tag system and framework patterns for context-free usage
6. Format output: Use [core-task-template.md](./.krci-ai/templates/krci-ai/core-task-template.md) for consistent structure
</instructions>

## Framework Context: XML Tag System

**XML Tags Purpose**: XML-style tags like `<instructions>` are **internal metadata for LLM guidance only**. They help LLMs identify section boundaries and processing requirements. These tags **MUST NEVER appear in final user output** - only clean Markdown should be presented to users.

**Common XML Tags:**
- `<prerequisites>` - Requirements before task execution
- `<instructions>` - Step-by-step guidance for LLM processing
- `<success_criteria>` - Validation criteria for completion
- `<content_guidelines>` - Content quality and formatting requirements

## Output Format

- **Location**: `./.krci-ai/tasks/{task-name}.md` following naming conventions
- **Structure**: Follow task template with XML guidance tags
- **References**: All framework references use inline `[filename](path)` format
- **Validation**: Task passes framework validation via `krci-ai validate`

<success_criteria>
- Task file created with proper naming convention and location
- XML guidance tags included for LLM processing (prerequisites, instructions, success_criteria)
- All framework references use inline `[filename](path)` format correctly
- Task structure follows template pattern with required sections
- Self-contained explanations enable context-free usage
- Framework validation passes without errors
</success_criteria>

## Execution Checklist

### Task Planning

- [ ] **Task naming**: Define descriptive name following framework conventions
- [ ] **Agent mapping**: Identify target agent that will reference this task
- [ ] **Purpose definition**: Clear statement of what task accomplishes
- [ ] **Dependencies identification**: Determine required templates and data references

### Content Creation

- [ ] **Template application**: Use [core-task-template.md](./.krci-ai/templates/krci-ai/core-task-template.md) structure
- [ ] **XML guidance**: Include `<prerequisites>`, `<instructions>`, `<success_criteria>` sections
- [ ] **Framework references**: Add inline links to required templates/data using `[filename](path)` format
- [ ] **Self-contained context**: Include XML tag explanations and framework pattern guidance

### Validation

- [ ] **Structure validation**: Verify task follows template requirements exactly
- [ ] **Reference validation**: Ensure all `[filename](path)` links resolve to existing files
- [ ] **Framework compliance**: Run `krci-ai validate` to check for issues
- [ ] **Content completeness**: Confirm task provides sufficient guidance for autonomous execution
