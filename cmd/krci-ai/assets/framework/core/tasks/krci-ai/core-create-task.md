---
dependencies:
  data:
    - krci-ai/core-framework-standards.yaml
    - krci-ai/core-validation-checklist.md
  templates:
    - krci-ai/core-task-template.md
---

# Task: Core Create Task

## Description

Guide user through creating framework-compliant tasks following KubeRocketAI patterns. This task provides step-by-step instructions for creating tasks that use XML guidance tags, frontmatter dependencies, and proper structure for LLM processing within the AI-as-Code framework.

<instructions>
Define task specification. Ask user for exact task name, target agent, and task purpose. Clarify what the task will accomplish and which agent will reference it (example: task core-analyze-metrics.md for advisor agent).

Review framework standards for task creation. Read task structure requirements from core-framework-standards.yaml. Understand required sections, XML guidance system, dependency frontmatter format, and validation requirements.

Apply XML guidance system. Use instructions and success_criteria XML tags for LLM section identification and processing guidance. Structure instructions as natural paragraph flow with imperative verbs and inline HALT conditions. Define success_criteria with specific, measurable validation criteria.

Follow naming conventions. Use descriptive task names with appropriate prefixes. Apply core- prefix for framework tasks. Use kebab-case for file naming. Ensure name clearly indicates task purpose and scope.

Manage dependencies in frontmatter. Create YAML frontmatter at file beginning. List all referenced data files under dependencies.data section. List all referenced templates under dependencies.templates section. Verify all dependency paths resolve to existing files.

Structure task following template pattern. Use core-task-template.md as structural guide. Include required sections: Title, Description, Instructions (XML), Framework Context, Output Format, Success Criteria (XML), Execution Checklist. Follow logical section progression for clear task flow.

Ensure self-contained guidance. Include explanations of framework patterns and validation requirements. Provide context for XML tag system purpose and usage. Add educational content about architectural concepts relevant to task. Enable context-free usage without requiring external knowledge.

Format output for framework compliance. Save to ./.krci-ai/tasks/{task-name}.md following path conventions. Populate all template variables with task-specific content. Ensure instructions use natural paragraph flow without numbered lists. Verify all framework references use correct inline format.

Run framework validation. Execute krci-ai validate and resolve identified issues. Confirm all frontmatter dependencies resolve correctly. Verify task passes automated compliance checks. Test task provides sufficient guidance for autonomous execution.
</instructions>

## Framework Context: Task Architecture and XML Tag System

XML Tags Purpose: XML-style tags like `<instructions>` are internal metadata for LLM guidance only. They help LLMs identify section boundaries and processing requirements. These tags MUST NEVER appear in final user output - only clean Markdown should be presented to users.

Task Structure Requirements: All framework tasks must follow exact section ordering from template:
1. Title → 2. Description → 3. Instructions (XML) → 4. Framework Context → 5. Output Format → 6. Success Criteria (XML) → 7. Execution Checklist

Common XML Tags and Usage:
- `<instructions>` - Natural paragraph flow guidance for LLM processing with imperative verbs and inline conditionals
- `<success_criteria>` - Validation criteria for completion with measurable outcomes
- `<content_guidelines>` - Content quality and formatting requirements for consistency

Framework Integration Patterns: Tasks must be self-contained with context-free usage capability, comprehensive frontmatter dependencies, and autonomous execution guidance.

## Output Format

- Location: `./.krci-ai/tasks/{task-name}.md` following naming conventions
- Structure: Follow task template with XML guidance tags and exact section ordering
- Dependencies: All framework dependencies declared in YAML frontmatter with validation
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

- Framework validation: Run `krci-ai validate` to ensure clean starting state
- Dependency verification: Confirm all reference assets exist at specified paths
- Context gathering: Review user requirements and framework constraints
- Task naming: Define descriptive name following framework conventions
- Agent mapping: Identify target agent that will reference this task
- Purpose definition: Clear statement of what task accomplishes

### Execution Phase

- Template application: Use [core-task-template.md](./.krci-ai/templates/krci-ai/core-task-template.md) structure completely
- Content generation: Populate all sections with framework-compliant content
- Dependency integration: Add all required dependencies in YAML frontmatter using correct format
- XML guidance: Include `<instructions>`, `<success_criteria>` sections with natural paragraph flow
- Framework context: Add educational explanations and architectural guidance
- Self-contained context: Include XML tag explanations and framework pattern guidance

### Validation Phase

- Structure validation: Verify task follows template requirements exactly
- Dependency resolution: Confirm all frontmatter dependencies resolve to existing files
- Framework validation: Run `krci-ai validate` and resolve any identified issues
- Content completeness: Confirm task provides sufficient guidance for autonomous execution
- Quality assurance: Review against framework standards and validation checklist
- Integration testing: Verify task integrates properly with framework ecosystem
