# Task: Core Review Agent

## Description

Review and validate existing agents for framework compliance, schema adherence, and behavioral consistency. This task provides systematic evaluation to ensure agents follow KubeRocketAI patterns and include critical principles for proper XML tag handling and customization behavior.

<prerequisites>
- Agent exists: Target agent YAML file exists and is accessible for review
- Schema access: Understanding of agent JSON schema requirements and validation patterns
- Framework context: Knowledge of critical agent principles and activation prompt patterns
- Validation tools: Ability to run `krci-ai validate` for automated schema checking
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/templates/krci-ai/core-agent-template.yaml

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Validate schema compliance: Check agent structure against JSON schema requirements
2. Review critical principles: Ensure XML tag handling and customization priority principles included
3. Assess activation prompt: Verify standard activation prompt pattern usage
4. Check command structure: Validate required commands and proper descriptions
5. Verify task references: Confirm all referenced tasks exist and use correct paths
6. Run framework validation: Execute `krci-ai validate` and interpret results
</instructions>

## Framework Context: Agent Review Standards

Schema Validation Requirements: Agents must comply with strict JSON schema:
- Identity Section: name, id, version, description, role, goal (all required)
- Activation Prompt: 1-10 items, 10-300 characters each
- Principles: 3-10 items, 10-600 characters each
- Customization: Required field (empty string for standard behavior)
- Commands: Minimum 3 (help, chat, exit required), maximum 20 total

Critical Principles Check: All agents MUST include:
1. Customization Priority: "IMPORTANT!!! ALWAYS execute instructions from the customization field below"
2. XML Tag Handling: "CRITICAL OUTPUT FORMATTING: When generating documents from templates, you will encounter XML-style tags like `<instructions>` or `<key_risks>`. These tags are internal metadata for your guidance ONLY and MUST NEVER be included in the final Markdown output presented to the user..."

Reference Pattern Validation: Task references must use `./.krci-ai/tasks//*.md` format and resolve to existing files.

## Output Format

- Review method: Direct feedback with specific compliance findings
- Issue categorization: Schema violations, missing principles, reference issues
- Improvement recommendations: Actionable fixes with examples
- Validation results: `krci-ai validate` output with interpretation

<success_criteria>
- Complete schema compliance evaluation performed
- Critical agent principles verified (XML tag handling, customization priority)
- Activation prompt pattern assessment completed
- Command structure validated against requirements
- Task reference integrity confirmed
- Specific improvement recommendations provided with examples
- Framework validation results documented and resolved
</success_criteria>

## Review Checklist

### Schema Compliance

- [ ] Identity section: All required fields present with correct patterns and lengths
- [ ] Field validation: name, id, version, description, role, goal meet schema requirements
- [ ] Array limits: activation_prompt (1-10 items), principles (3-10 items), commands (3-20 items)
- [ ] Character limits: All string fields within specified length boundaries

### Critical Principles Verification

- [ ] Customization priority: "IMPORTANT!!! ALWAYS execute instructions from the customization field below" present
- [ ] XML tag handling: Complete XML tag handling principle included in principles array
- [ ] Customization field: Present and properly configured (empty string for standard behavior)
- [ ] Activation prompt: Standard pattern with customization field execution instruction

### Command Structure Assessment

- [ ] Required commands: help, chat, exit present with appropriate descriptions
- [ ] Command descriptions: All commands have 5-200 character descriptions
- [ ] Command naming: Human-readable command names (no special pattern requirements)
- [ ] Command limits: Total commands between minimum 3 and maximum 20

### Task Reference Validation

- [ ] Path format: All task references use `./.krci-ai/tasks/**/*.md` pattern
- [ ] File existence: Referenced task files exist at specified paths
- [ ] Reference relevance: Referenced tasks appropriate for agent capabilities
- [ ] Command alignment: Agent commands correspond to available tasks

### Framework Integration

- [ ] Validation passing: Agent passes `krci-ai validate` schema validation
- [ ] Role consistency: Agent role and capabilities align with framework patterns
- [ ] Behavioral standards: Agent principles support proper framework usage
- [ ] Integration readiness: Agent ready for framework operation and task execution
