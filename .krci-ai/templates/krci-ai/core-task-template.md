# Task: {{task_name}}

<instructions>
TASK TEMPLATE STRUCTURE: Follow this exact section ordering for framework compliance:
1. Title (with task name) → 2. Description → 3. Prerequisites (XML) → 4. Reference Assets → 5. Instructions (XML) → 6. Framework Context → 7. Output Format → 8. Success Criteria (XML) → 9. Execution Checklist

XML GUIDANCE SYSTEM: Use `<prerequisites>`, `<instructions>`, `<success_criteria>` tags for LLM processing guidance. These tags help LLMs identify section boundaries and processing requirements. CRITICAL: These XML tags are internal metadata ONLY and must never appear in final user output.
</instructions>

## Description

{{task_description}}

<prerequisites>
{{task_prerequisites}}
</prerequisites>

### Reference Assets

<instructions>
List all framework dependencies using inline reference format. Each dependency must use `[filename](./.krci-ai/path/to/file)` pattern for proper framework integration.
</instructions>

Dependencies:

{{framework_dependencies}}

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
{{task_instructions}}
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

1. **Variable Substitution**: Replace all `{{variable_name}}` placeholders with appropriate content
2. **XML Tag Handling**: Maintain XML guidance tags for LLM processing - these provide section identification
3. **Reference Integration**: Ensure all framework references use `[filename](./.krci-ai/path)` format
4. **Self-Contained Context**: Include sufficient explanations for context-free usage
5. **Validation Ready**: Structure task to pass `krci-ai validate` framework validation

CRITICAL REMINDERS:
- XML tags (`<instructions>`, `<prerequisites>`, `<success_criteria>`) are for LLM guidance ONLY
- Final user output must be clean Markdown without XML tags
- All framework references must resolve to existing files
- Task structure must support autonomous execution
</instructions>
