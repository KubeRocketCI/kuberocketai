# Task: Core Review Task

## Description

Review and validate existing tasks for framework compliance, XML tag usage, and structural integrity. This task provides systematic evaluation criteria to ensure tasks follow KubeRocketAI patterns and provide proper LLM guidance.

<prerequisites>
- Task exists: Target task file exists and is accessible for review
- Framework context: Understanding of XML tag system, inline references, and validation requirements
- Validation access: Ability to run `krci-ai validate` for automated checking
- Review authority: Knowledge of framework patterns to provide actionable feedback
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/data/krci-ai/core-validation-checklist.md
- ./.krci-ai/templates/krci-ai/core-task-template.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Validate XML guidance system: Ensure task uses proper XML tags for LLM section identification
2. Check inline references: Verify all `[filename](path)` references resolve to existing files
3. Review structure compliance: Confirm task follows template pattern with required sections
4. Assess self-contained guidance: Evaluate if task provides context-free usage explanations
5. Test framework validation: Run `krci-ai validate` to identify automated compliance issues
6. Provide improvement recommendations: Offer specific, actionable feedback for enhancement
</instructions>

## Framework Context: Review Standards

XML Tag Evaluation: Tasks must use XML-style tags (`<prerequisites>`, `<instructions>`, `<success_criteria>`) for LLM guidance. These tags are internal metadata only and must never appear in user-facing output. Review should verify:
- Proper XML tag placement and closure
- Content appropriateness within tags
- LLM guidance effectiveness

Reference Pattern Validation: All framework references must use inline `[filename](./.krci-ai/path/to/file)` format. Check for:
- Correct path resolution
- Existing target files
- Appropriate reference context

## Output Format

- Review method: Direct feedback on existing task file or separate review document
- Feedback structure: Organized by compliance areas (XML tags, references, structure, content)
- Actionable recommendations: Specific improvements with examples
- Validation results: Include `krci-ai validate` output and interpretation

<success_criteria>
- Complete evaluation of task against framework standards
- XML tag usage properly assessed for LLM guidance effectiveness
- All inline references validated for existence and correctness
- Structural compliance confirmed against template requirements
- Specific, actionable improvement recommendations provided
- Framework validation results documented and interpreted
</success_criteria>

## Review Checklist

### XML Guidance Assessment

- [ ] XML tag presence: Task includes `<prerequisites>`, `<instructions>`, `<success_criteria>` sections
- [ ] Tag closure: All XML tags properly opened and closed
- [ ] Content appropriateness: XML tag content provides effective LLM guidance
- [ ] User output safety: Verify XML tags won't appear in final user-facing content

### Reference Validation

- [ ] Inline format: All references use `[filename](./.krci-ai/path/to/file)` pattern
- [ ] Path accuracy: Referenced files exist at specified locations
- [ ] Reference relevance: Referenced files are appropriate for task context
- [ ] Link resolution: All `[filename](path)` links resolve correctly

### Structure Compliance

- [ ] Template adherence: Task follows krci-ai/core-task-template.md structure
- [ ] Section completeness: Required sections (Description, Prerequisites, Instructions, etc.) present
- [ ] Content organization: Information logically organized and accessible
- [ ] Self-contained guidance: Task provides context-free usage explanations

### Framework Integration

- [ ] Validation passing: Task passes `krci-ai validate` without errors
- [ ] Agent integration: Task can be properly referenced by target agents
- [ ] Framework patterns: Task follows established framework conventions
- [ ] Quality standards: Content meets framework quality and completeness requirements
