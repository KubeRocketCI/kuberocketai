# Task: Core Review Template

## Description

Review and validate existing templates for variable consistency, LLM guidance effectiveness, and structural integrity. This task provides evaluation criteria to ensure templates support proper content generation and framework compliance.

<prerequisites>
- Template exists: Target template file exists and is accessible for review
- Template context: Understanding of template's intended usage and output format
- Variable knowledge: Familiarity with `{{variable}}` placeholder system and naming conventions
- LLM processing: Knowledge of how templates guide content generation
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/templates/krci-ai/core-template-template.md
- ./.krci-ai/data/core-validation-checklist.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Validate variable system: Check `{{variable}}` placeholder consistency and naming patterns
2. Review XML guidance: Assess `<instructions>` tags for effective LLM processing guidance
3. Evaluate structure: Confirm template supports intended output format and organization
4. Check reusability: Assess template design for use across multiple contexts
5. Test LLM compatibility: Review template for natural language processing optimization
6. Provide improvements: Offer specific recommendations for enhanced effectiveness
</instructions>

## Framework Context: Template Review Standards

Variable System Evaluation: Templates should use consistent `{{variable_name}}` patterns:
- Naming Consistency: Variables follow descriptive naming conventions
- Type Appropriateness: Simple values, lists, and sections used correctly
- Documentation: Variable purpose clear through naming and context
- Optional Handling: Optional vs required variables clearly indicated

XML Guidance Assessment: Review instruction tags for LLM processing effectiveness:
- `<instructions>` tags provide clear content generation guidance
- Comments explain variable usage and formatting expectations
- Processing hints facilitate proper LLM content generation

Structure and Reusability: Templates should support flexible usage while maintaining consistency.

## Output Format

- Review method: Direct feedback on template effectiveness and compliance
- Issue identification: Variable problems, guidance gaps, structural issues
- Improvement recommendations: Specific enhancements with examples
- Usage assessment: Evaluation of template's effectiveness for intended purpose

<success_criteria>
- Complete variable system evaluation performed
- XML guidance tags assessed for LLM processing effectiveness
- Template structure validated for intended output support
- Reusability and flexibility assessment completed
- Specific improvement recommendations provided
- Framework compliance confirmed or issues identified
</success_criteria>

## Review Checklist

### Variable System Assessment

- [ ] Variable syntax: All variables use `{{variable_name}}` format correctly
- [ ] Naming consistency: Variable names follow descriptive, consistent patterns
- [ ] Variable types: Appropriate use of simple values, lists, and content sections
- [ ] Documentation clarity: Variable purpose clear from naming and context

### XML Guidance Evaluation

- [ ] Instruction presence: `<instructions>` tags provide clear LLM guidance where needed
- [ ] Content guidance: Instructions help LLMs generate appropriate content for variables
- [ ] Formatting hints: Guidance for consistent formatting and structure
- [ ] Processing support: Instructions facilitate effective template usage

### Structure and Organization

- [ ] Markdown compliance: Template uses proper markdown syntax and structure
- [ ] Content flow: Logical organization and content hierarchy
- [ ] Section clarity: Clear headings and content organization
- [ ] Output support: Template structure supports intended final output format

### Reusability and Flexibility

- [ ] Context flexibility: Template usable across multiple similar contexts
- [ ] Variable adaptability: Variable system accommodates diverse content needs
- [ ] Modification ease: Template structure allows for easy customization
- [ ] Framework integration: Template follows established framework patterns

### LLM Processing Optimization

- [ ] Natural language compatibility: Template structure facilitates LLM processing
- [ ] Content generation support: Template design aids in quality content creation
- [ ] Instruction effectiveness: XML guidance tags provide valuable processing hints
- [ ] User output quality: Template produces clean, professional final output

### Framework Compliance

- [ ] Standard adherence: Template follows framework template conventions
- [ ] Reference compatibility: Template can be properly referenced by tasks
- [ ] Validation readiness: Template meets framework validation requirements
- [ ] Integration capability: Template works effectively within framework ecosystem
