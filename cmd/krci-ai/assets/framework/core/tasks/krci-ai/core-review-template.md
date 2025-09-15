# Task: Core Review Template

## Description

Review and validate existing templates for variable consistency, LLM guidance effectiveness, and structural integrity. This task provides evaluation criteria to ensure templates support proper content generation and framework compliance with comprehensive quality assessment.

<prerequisites>
- Template exists: Target template file exists and is accessible for review
- Template context: Understanding of template's intended usage and output format
- Variable knowledge: Familiarity with `{{variable}}` placeholder system and naming conventions
- LLM processing: Knowledge of how templates guide content generation and framework integration
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/templates/krci-ai/core-template-template.md
- ./.krci-ai/data/krci-ai/core-validation-checklist.md

CRITICAL: Load all dependencies by reading their complete content before task execution. HALT if any missing.

<instructions>
1. Validate variable system: Check `{{variable}}` placeholder consistency and naming patterns with type appropriateness
2. Review XML guidance: Assess `<instructions>` tags for effective LLM processing guidance and content generation
3. Evaluate structure: Confirm template supports intended output format and organization with logical flow
4. Check reusability: Assess template design for use across multiple contexts with flexible adaptation
5. Test LLM compatibility: Review template for natural language processing optimization and clarity
6. Provide improvements: Offer specific recommendations for enhanced effectiveness with implementation examples
</instructions>

## Framework Context: Template Review Standards and Quality Assessment

**Variable System Evaluation**: Templates should use consistent `{{variable_name}}` patterns with comprehensive validation:

- Naming consistency: Variables follow descriptive naming conventions with clear purpose indication
- Type appropriateness: Simple values, content lists, large sections, and optional content used correctly
- Documentation clarity: Variable purpose clear through naming, context, and instruction comments
- Optional handling: Optional vs required variables clearly indicated and properly documented

**XML Guidance Assessment**: Review instruction tags for LLM processing effectiveness:

- `<instructions>` tags provide clear, actionable content generation guidance
- Comments explain variable usage and formatting expectations with examples
- Processing hints facilitate proper LLM content generation and template usage
- Section organization guidance maintains consistent structure across outputs

**Structure and Reusability Standards**: Templates should support flexible usage while maintaining consistency, logical content flow, and framework integration patterns.

**LLM Optimization Requirements**: Templates must be optimized for natural language processing with clear guidance, appropriate variable placement, and autonomous content generation capability.

## Output Format

- Review method: Comprehensive feedback on template effectiveness, compliance, and improvement opportunities
- Evaluation structure: Organized by variable system, XML guidance, structure, reusability, LLM compatibility
- Improvement recommendations: Specific enhancements with implementation examples and best practices
- Quality assessment: Framework compliance status and integration readiness evaluation

<success_criteria>
- Framework compliance verified: Template passes all automated validation checks without errors
- Pattern adherence confirmed: Template follows established framework conventions exactly
- Reference integrity validated: All references resolve correctly and appropriately
- Quality standards met: Template meets completeness, clarity, and maintainability requirements
- Integration readiness achieved: Template ready for framework operation and usage
- Documentation completeness confirmed: All required sections populated with actionable content
</success_criteria>

## Execution Checklist

### Preparation Phase

- [ ] Framework validation: Run `krci-ai validate` to establish baseline template status
- [ ] Dependency verification: Confirm all reference assets exist at specified paths
- [ ] Context gathering: Review template specifications and intended usage scenarios
- [ ] Review scope definition: Identify specific areas requiring evaluation and feedback

### Execution Phase

- [ ] Variable system validation: Check `{{variable}}` placeholder consistency and naming patterns
- [ ] Variable type assessment: Verify appropriate mix of simple values, lists, and content sections
- [ ] XML guidance review: Assess `<instructions>` tags for effective LLM processing guidance
- [ ] Structure evaluation: Confirm template supports intended output format with logical organization
- [ ] Reusability testing: Assess template design for use across multiple similar contexts
- [ ] LLM compatibility verification: Review template for natural language processing optimization
- [ ] Framework integration assessment: Verify template follows established patterns and conventions

### Validation Phase

- [ ] Template effectiveness evaluation: Assess overall template quality and usability
- [ ] Issue categorization: Organize findings by variable system, XML guidance, structure, reusability
- [ ] Improvement recommendations: Provide specific, actionable feedback with implementation examples
- [ ] Quality standards verification: Confirm template meets framework completeness requirements
- [ ] Integration readiness assessment: Verify template capability for effective framework usage
- [ ] Documentation completeness review: Ensure template provides comprehensive guidance and examples
