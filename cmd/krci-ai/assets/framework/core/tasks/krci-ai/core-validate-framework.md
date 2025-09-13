# Task: Core Validate Framework

## Description

Execute comprehensive framework validation using CLI tools and manual inspection to ensure component compliance, reference integrity, and operational readiness. This task provides systematic validation procedures for complete framework quality assurance.

<prerequisites>
- Framework installed: .krci-ai directory exists with complete structure (agents/, tasks/, templates/, data/)
- CLI access: `krci-ai` command available with validate functionality
- Component knowledge: Understanding of framework components and validation requirements
- Issue resolution: Ability to address validation findings and implement fixes
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/krci-ai/core-framework-standards.yaml
- ./.krci-ai/data/krci-ai/core-validation-checklist.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

<instructions>
1. Execute CLI validation: Run `krci-ai validate` with appropriate verbosity for comprehensive checking
2. Analyze validation results: Interpret automated validation findings and categorize issues
3. Perform manual inspection: Review components for framework pattern compliance beyond automated checks
4. Validate reference integrity: Ensure all `[filename](path)` references resolve to existing files
5. Check component relationships: Verify agents reference existing tasks and tasks reference existing templates/data
6. Validate against [core-validation-checklist](./.krci-ai/data/krci-ai/core-validation-checklist.md)
7. Document findings: Create comprehensive validation report with specific remediation actions
</instructions>

## Framework Context: Validation Architecture

CLI Validation Capabilities: `krci-ai validate` provides automated validation including:
- Agent Schema Compliance: JSON schema validation for agent YAML structure
- Task Path Validation: Verification that agent task references resolve to existing files
- Template Accessibility: Confirmation that templates are accessible from framework structure
- Markdown Link Validation: Checking `[filename](path)` references for existing targets
- Cross-platform Compatibility: File accessibility across different operating systems

Validation Scope Areas:
- Static Validation: Design-time checking via CLI tool (token-free comprehensive analysis)
- Reference Integrity: All framework component references resolve correctly
- Schema Compliance: Agents meet JSON schema requirements precisely
- Pattern Consistency: Components follow established framework conventions

## Output Format

- Validation execution: CLI command results with appropriate verbosity
- Issue categorization: Findings organized by component type and severity
- Remediation guidance: Specific actions required to resolve validation failures
- Compliance confirmation: Verification of successful validation after fixes

<success_criteria>
- CLI validation executed successfully with comprehensive results
- All validation issues identified and categorized by component and severity
- Reference integrity confirmed for all framework component dependencies
- Schema compliance verified for all agents against JSON schema requirements
- Specific remediation actions provided for any validation failures
- Framework validation passes completely after issue resolution
</success_criteria>

## Execution Checklist

### CLI Validation Execution

- [ ] Basic validation: Run `krci-ai validate` for standard framework checking
- [ ] Verbose validation: Execute `krci-ai validate --verbose` for detailed analysis
- [ ] Quiet validation: Use `krci-ai validate --quiet` for summary-only results if needed
- [ ] Result interpretation: Analyze CLI output for issues and compliance status

### Validation Result Analysis

- [ ] Issue categorization: Organize findings by component type (agents, tasks, templates, data)
- [ ] Severity assessment: Classify issues as critical, warning, or informational
- [ ] Priority ranking: Determine resolution order based on impact and dependencies
- [ ] Root cause analysis: Understand underlying causes of validation failures

### Component Compliance Review

- [ ] Agent schema validation: Verify all agents meet JSON schema requirements exactly
- [ ] Task structure validation: Confirm tasks follow framework patterns and XML guidance
- [ ] Template validation: Check templates for proper variable usage and structure
- [ ] Data validation: Verify data files are accessible and properly organized

### Reference Integrity Verification

- [ ] Agent-task references: Confirm all agent.tasks entries resolve to existing task files
- [ ] Task-template references: Verify task `[filename](path)` links resolve to existing templates
- [ ] Task-data references: Confirm task references to data files resolve correctly
- [ ] Agent-data references: Verify agent references to behavioral data resolve properly

### Manual Inspection and Quality Assurance

- [ ] Framework pattern compliance: Review components for adherence to established patterns
- [ ] Critical principle inclusion: Verify agents include XML tag handling and customization principles
- [ ] XML guidance usage: Confirm tasks use XML tags for LLM processing guidance
- [ ] Content quality: Assess component content for completeness and effectiveness

### Issue Resolution and Remediation

- [ ] Fix implementation: Address validation failures with specific corrections
- [ ] Revalidation: Re-run CLI validation after implementing fixes
- [ ] Compliance confirmation: Verify complete validation success after remediation
- [ ] Documentation update: Record validation process and resolution actions for reference
