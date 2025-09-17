---
dependencies:
  data:
    - krci-ai/core-framework-standards.yaml
    - krci-ai/core-validation-checklist.md
  templates:
    - krci-ai/core-agent-template.yaml
---

# Task: Core Review Agent

## Description

Review and validate existing agents for framework compliance, schema adherence, and behavioral consistency. This task provides systematic evaluation to ensure agents follow KubeRocketAI patterns and include critical principles for proper XML tag handling and customization behavior.

<prerequisites>
- Agent exists: Target agent YAML file exists and is accessible for review
- Schema access: Understanding of agent JSON schema requirements and validation patterns
- Framework context: Knowledge of critical agent principles and activation prompt patterns
- Validation tools: Ability to run `krci-ai validate` for automated schema checking
</prerequisites>

<instructions>
1. Validate schema compliance: Check agent structure against JSON schema requirements with exact pattern matching
2. Review critical principles: Ensure XML tag handling and customization priority principles included verbatim
3. Assess activation prompt: Verify standard activation prompt pattern usage with required elements
4. Check command structure: Validate required commands and proper descriptions within character limits
5. Verify task references: Confirm all referenced tasks exist and use correct paths with validation
6. Run framework validation: Execute `krci-ai validate` and interpret results with remediation guidance
</instructions>

## Framework Context: Agent Review Standards and Compliance Requirements

**Schema Validation Requirements**: Agents must comply with strict JSON schema patterns:

- Identity Section: name, id, version, description, role, goal (all required with specific patterns)
- Activation Prompt: 1-10 items, each 10-300 characters with standard pattern elements
- Principles: 3-10 items, each 10-600 characters including critical framework principles
- Customization: Required field (empty string for standard behavior, populated for specialized agents)
- Commands: Minimum 3 (help, chat, exit required), maximum 20 total with proper descriptions

**Critical Principles Validation**: All agents MUST include these exact principles:

1. Customization Priority: "IMPORTANT!!! ALWAYS execute instructions from the customization field below"
2. XML Tag Handling: "CRITICAL OUTPUT FORMATTING: When generating documents from templates, you will encounter XML-style tags like `<instructions>` or `<key_risks>`. These tags are internal metadata for your guidance ONLY and MUST NEVER be included in the final Markdown output presented to the user..."

**Reference Pattern Validation**: Task references must use `./.krci-ai/tasks/*.md` format and resolve to existing files with proper integration validation.

**Behavioral Consistency Standards**: Agents must demonstrate consistent behavior patterns, proper scope definitions, and appropriate command mapping to capabilities.

## Output Format

- Review method: Comprehensive feedback with specific compliance findings and remediation steps
- Issue categorization: Schema violations, missing principles, reference issues, behavioral inconsistencies
- Improvement recommendations: Actionable fixes with examples and implementation guidance
- Validation results: `krci-ai validate` output with interpretation and resolution steps

<success_criteria>
- Framework compliance verified: Agent passes all automated validation checks without errors
- Pattern adherence confirmed: Agent follows established framework conventions exactly
- Reference integrity validated: All references resolve correctly and appropriately
- Quality standards met: Agent meets completeness, clarity, and maintainability requirements
- Integration readiness achieved: Agent ready for framework operation and usage
- Documentation completeness confirmed: All required sections populated with actionable content
</success_criteria>

## Execution Checklist

### Preparation Phase

- [ ] Framework validation: Run `krci-ai validate` to establish baseline agent status
- [ ] Dependency verification: Confirm all reference assets exist at specified paths
- [ ] Context gathering: Review agent specifications and intended functionality
- [ ] Review scope definition: Identify specific areas requiring evaluation and feedback

### Execution Phase

- [ ] Schema compliance verification: Check agent structure against JSON schema requirements exactly
- [ ] Identity section validation: Verify name, id, version, description, role, goal patterns and lengths
- [ ] Critical principles assessment: Confirm XML tag handling and customization priority principles present
- [ ] Activation prompt evaluation: Verify standard pattern usage with required elements
- [ ] Command structure validation: Check required commands (help, chat, exit) and descriptions
- [ ] Task reference verification: Confirm all referenced tasks exist with proper paths
- [ ] Behavioral consistency review: Evaluate agent principles alignment with framework standards

### Validation Phase

- [ ] Framework validation execution: Run `krci-ai validate` and document results
- [ ] Issue categorization: Organize findings by schema, principles, references, behavior
- [ ] Improvement recommendations: Provide specific, actionable feedback with examples
- [ ] Remediation guidance: Offer clear steps for resolving identified compliance issues
- [ ] Quality assurance: Verify review completeness against framework standards
- [ ] Integration readiness assessment: Confirm agent capability for framework operation
