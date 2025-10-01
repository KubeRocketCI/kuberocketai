# Task: {{task_name}}

## Description

{{task_description}}

<instructions>
{{task_instructions_natural_flow}}

Example instruction pattern:
Identify target component for operation. Ask user for exact file path or component specification. Confirm existence and accessibility before proceeding.

Review framework standards for operation type. Read relevant sections from core-framework-standards.yaml. Understand requirements and validation criteria.

Execute primary operation steps. Perform main task actions with clear validation at each step. Verify outputs meet framework requirements. Confirm integration with framework ecosystem.

Validate operation results. Run krci-ai validate to verify compliance. Resolve any issues identified. Confirm successful completion against success criteria.
</instructions>

## Framework Context: {{context_area}}

<instructions>
Provide educational context about framework patterns, validation requirements, or architectural concepts relevant to this task. This section ensures self-contained usage and context-free understanding.
</instructions>

{{framework_context_explanation}}

## Output Format

<instructions>
Define expected output structure, file locations, naming conventions, and formatting requirements. Include specific paths and validation criteria.
</instructions>

{{output_format_specifications}}

<success_criteria>
{{success_validation_criteria}}
</success_criteria>

## Execution Checklist

<instructions>
Provide detailed checklist organized by execution phases. Use checkbox format with specific, actionable items that can be verified during task execution.
</instructions>

### {{execution_phase_1}}

{{execution_checklist_items_1}}

### {{execution_phase_2}}

{{execution_checklist_items_2}}

### {{execution_phase_3}}

{{execution_checklist_items_3}}

<instructions>
TEMPLATE USAGE GUIDANCE:

Variable Substitution: Replace all {{variable_name}} placeholders with appropriate content for specific task being created.

XML Tag Handling: Maintain XML guidance tags (instructions, success_criteria) for LLM processing. These provide section identification and must never appear in final user output.

Instruction Format: Use natural paragraph flow without numbered lists. Start paragraphs with imperative verbs. Include inline conditionals like "before proceeding" or "if missing" for HALT conditions.

Target Identification: For review or create tasks, first instruction paragraph must ask user to specify exact component (file path, name, purpose). Confirm existence or define specification before proceeding with operation.

Reference Integration: Ensure all framework references in frontmatter use correct dependency.data and dependencies.templates sections. Verify paths resolve to existing files.

Self-Contained Context: Include Framework Context section explaining relevant patterns, standards, and architectural concepts. Enable context-free task usage without external knowledge requirements.

Validation Ready: Structure task to pass krci-ai validate framework validation. Test all dependency paths. Confirm XML tag closure. Verify section ordering matches framework standards.

CRITICAL REMINDERS:
- XML tags (instructions, success_criteria) are for LLM guidance ONLY
- Final user output must be clean Markdown without XML tags
- All framework references must resolve to existing files
- Instructions use natural paragraphs, not numbered lists
- Target specification required as first action for component-specific tasks
- Task structure must support autonomous execution
</instructions>
