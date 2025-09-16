---
dependencies:
  data:
    - krci-ai/core-validation-checklist.md
  templates:
    - krci-ai/core-task-template.md
---

# Task: Core Review Task

## Description

Review and validate existing tasks for framework compliance, XML tag usage, and structural integrity. This task provides systematic evaluation criteria to ensure tasks follow KubeRocketAI patterns and provide proper LLM guidance with comprehensive quality assessment.

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

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.

<instructions>
1. Validate XML guidance system: Ensure task uses proper XML tags for LLM section identification and guidance
2. Check inline references: Verify all `[filename](path)` references resolve to existing files with proper validation
3. Review structure compliance: Confirm task follows template pattern with required sections and ordering
4. Assess self-contained guidance: Evaluate if task provides context-free usage explanations and framework context
5. Test framework validation: Run `krci-ai validate` to identify automated compliance issues and resolution steps
6. Provide improvement recommendations: Offer specific, actionable feedback for enhancement with implementation examples
</instructions>

## Framework Context: Task Review Standards and Quality Assessment

**XML Tag Evaluation**: Tasks must use XML-style tags (`<prerequisites>`, `<instructions>`, `<success_criteria>`) for LLM guidance. These tags are internal metadata only and must never appear in user-facing output. Review should verify:

- Proper XML tag placement and closure with correct syntax
- Content appropriateness within tags providing clear LLM guidance
- LLM guidance effectiveness for autonomous task execution
- Section boundary identification and processing requirements clarity

**Reference Pattern Validation**: All framework references must use inline `[filename](./.krci-ai/path/to/file)` format. Comprehensive validation includes:

- Correct path resolution to existing framework components
- Target file existence and accessibility verification
- Appropriate reference context and integration patterns
- Dependency chain integrity and circular reference prevention

**Task Structure Requirements**: Tasks must follow exact template compliance with proper section ordering, comprehensive execution checklists, and autonomous execution capability.

## Output Format

- Review method: Direct feedback on existing task file or separate review document
- Feedback structure: Organized by compliance areas (XML tags, references, structure, content)
- Actionable recommendations: Specific improvements with examples
- Validation results: Include `krci-ai validate` output and interpretation

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

- [ ] Framework validation: Run `krci-ai validate` to establish baseline task status
- [ ] Dependency verification: Confirm all reference assets exist at specified paths
- [ ] Context gathering: Review task specifications and intended functionality
- [ ] Review scope definition: Identify specific areas requiring evaluation and feedback

### Execution Phase

- [ ] XML guidance assessment: Verify task includes proper `<prerequisites>`, `<instructions>`, `<success_criteria>` sections
- [ ] XML tag validation: Check all XML tags are properly opened, closed, and provide effective LLM guidance
- [ ] Reference pattern verification: Confirm all references use `[filename](./.krci-ai/path/to/file)` format correctly
- [ ] Reference resolution testing: Verify all referenced files exist at specified locations
- [ ] Structure compliance review: Confirm task follows template pattern with required sections and ordering
- [ ] Self-contained guidance assessment: Evaluate if task provides context-free usage explanations
- [ ] Framework context validation: Verify educational context and architectural guidance presence

### Validation Phase

- [ ] Framework validation execution: Run `krci-ai validate` and document results with interpretation
- [ ] Issue categorization: Organize findings by XML tags, references, structure, content quality
- [ ] Improvement recommendations: Provide specific, actionable feedback with implementation examples
- [ ] Quality assurance verification: Confirm task meets framework completeness and clarity requirements
- [ ] Integration readiness assessment: Verify task capability for framework operation and agent reference
- [ ] Documentation completeness review: Ensure all sections populated with actionable, comprehensive content
