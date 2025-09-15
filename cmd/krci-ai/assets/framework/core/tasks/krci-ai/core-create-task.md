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

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.

<instructions>
1. Apply XML guidance system: Use XML tags (`<prerequisites>`, `<instructions>`, `<success_criteria>`) for LLM section identification and guidance
2. Follow naming convention: Use descriptive task names with appropriate prefixes (e.g., "core-" for framework tasks)
3. Create inline references: Use `[filename](./.krci-ai/path/to/file)` format for all framework references
4. Structure validation: Follow task template structure with Description, Prerequisites, Instructions, Success Criteria
5. Ensure self-contained guidance: Include explanations of XML tag system and framework patterns for context-free usage
6. Format output: Use [core-task-template.md](./.krci-ai/templates/krci-ai/core-task-template.md) for consistent structure
</instructions>

## Framework Context: Task Architecture and XML Tag System

**XML Tags Purpose**: XML-style tags like `<instructions>` are internal metadata for LLM guidance only. They help LLMs identify section boundaries and processing requirements. These tags MUST NEVER appear in final user output - only clean Markdown should be presented to users.

**Task Structure Requirements**: All framework tasks must follow exact section ordering from template:
1. Title → 2. Description → 3. Prerequisites (XML) → 4. Reference Assets → 5. Instructions (XML) → 6. Framework Context → 7. Output Format → 8. Success Criteria (XML) → 9. Execution Checklist

**Common XML Tags and Usage**:
- `<prerequisites>` - Requirements before task execution with specific validation criteria
- `<instructions>` - Step-by-step guidance for LLM processing with numbered, actionable items
- `<success_criteria>` - Validation criteria for completion with measurable outcomes
- `<content_guidelines>` - Content quality and formatting requirements for consistency

**Framework Integration Patterns**: Tasks must be self-contained with context-free usage capability, comprehensive inline references, and autonomous execution guidance.

## Output Format

- Location: `./.krci-ai/tasks/{task-name}.md` following naming conventions
- Structure: Follow task template with XML guidance tags and exact section ordering
- References: All framework references use inline `[filename](path)` format with validation
- Validation: Task passes framework validation via `krci-ai validate` with zero errors

<success_criteria>
- Framework compliance verified: Task passes all automated validation checks without errors
- Pattern adherence confirmed: Task follows established framework conventions exactly
- Reference integrity validated: All references resolve correctly and appropriately
- Quality standards met: Task meets completeness, clarity, and maintainability requirements
- Integration readiness achieved: Task ready for framework operation and usage
- Documentation completeness confirmed: All required sections populated with actionable content
</success_criteria>

## Execution Checklist

### Preparation Phase

- [ ] Framework validation: Run `krci-ai validate` to ensure clean starting state
- [ ] Dependency verification: Confirm all reference assets exist at specified paths
- [ ] Context gathering: Review user requirements and framework constraints
- [ ] Task naming: Define descriptive name following framework conventions
- [ ] Agent mapping: Identify target agent that will reference this task
- [ ] Purpose definition: Clear statement of what task accomplishes

### Execution Phase

- [ ] Template application: Use [core-task-template.md](./.krci-ai/templates/krci-ai/core-task-template.md) structure completely
- [ ] Content generation: Populate all sections with framework-compliant content
- [ ] Reference integration: Add all required inline references using correct format
- [ ] XML guidance: Include `<prerequisites>`, `<instructions>`, `<success_criteria>` sections
- [ ] Framework context: Add educational explanations and architectural guidance
- [ ] Self-contained context: Include XML tag explanations and framework pattern guidance

### Validation Phase

- [ ] Structure validation: Verify task follows template requirements exactly
- [ ] Reference resolution: Confirm all `[filename](path)` links resolve to existing files
- [ ] Framework validation: Run `krci-ai validate` and resolve any identified issues
- [ ] Content completeness: Confirm task provides sufficient guidance for autonomous execution
- [ ] Quality assurance: Review against framework standards and validation checklist
- [ ] Integration testing: Verify task integrates properly with framework ecosystem
