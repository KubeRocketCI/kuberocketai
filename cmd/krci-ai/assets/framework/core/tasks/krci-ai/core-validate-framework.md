---
dependencies:
  data:
    - krci-ai/core-framework-standards.yaml
    - krci-ai/core-validation-checklist.md
---

# Task: Core Validate Framework

## Description

Execute comprehensive framework validation using CLI tools and manual inspection to ensure component compliance, reference integrity, and operational readiness. This task provides systematic validation procedures for complete framework quality assurance.

<instructions>
Execute CLI validation with appropriate verbosity. Run krci-ai validate for standard framework checking. Use krci-ai validate --verbose for detailed analysis when investigating specific issues. Apply krci-ai validate --quiet for summary-only results if needed. Capture and analyze CLI output for issues and compliance status.

Analyze validation results systematically. Interpret automated validation findings from CLI output. Categorize issues by component type (agents, tasks, templates, data files). Classify severity as critical, warning, or informational. Prioritize resolution order based on impact and dependencies. Understand root causes of validation failures for effective remediation.

Perform manual inspection beyond automated checks. Review agent schema compliance against JSON schema requirements exactly. Confirm tasks follow framework patterns with proper XML guidance structure. Check templates for proper variable usage, naming consistency, and structure. Verify data files are accessible, properly organized, and appropriately formatted.

Validate reference integrity across all components. Confirm agent tasks entries resolve to existing task files at specified paths. Verify task frontmatter dependencies resolve to existing templates under dependencies.templates section. Test task data references resolve correctly under dependencies.data section. Check agent data references to behavioral guidance files resolve properly.

Check component relationship alignment. Verify agents reference tasks appropriate for their capabilities and role. Confirm tasks reference templates that match their output requirements. Test that tasks reference data files providing relevant standards or specifications. Validate command-to-task mapping follows framework patterns across agents.

Assess framework pattern compliance manually. Review components for adherence to established conventions beyond automated validation. Verify agents include critical XML tag handling and customization principles exactly. Confirm tasks use natural paragraph instruction flow without numbered lists. Check self-contained context and framework pattern explanations present.

Validate against framework standards and checklist. Reference core-validation-checklist.md for comprehensive quality assurance criteria. Check components against all critical validation requirements. Verify warning-level best practices are followed. Assess informational excellence criteria where applicable.

Document findings comprehensively. Create validation report organized by component type and severity. List specific issues with exact file paths and line references where applicable. Provide clear remediation guidance for each finding category. Include priority ranking for resolution efforts. Document validation process and timeline for reference.

Implement fixes and revalidate. Address validation failures with specific corrections per remediation guidance. Re-run krci-ai validate after implementing fixes to verify resolution. Confirm complete validation success with zero critical errors. Update validation documentation to record resolution actions and final compliance status.
</instructions>

## Framework Context: Validation Architecture

CLI Validation Capabilities: `krci-ai validate` provides automated validation including:
- Agent Schema Compliance: JSON schema validation for agent YAML structure
- Task Path Validation: Verification that agent task references resolve to existing files
- Template Accessibility: Confirmation that templates are accessible from framework structure
- File Path Validation: Checking file path references for existing targets
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

- Basic validation: Run `krci-ai validate` for standard framework checking
- Verbose validation: Execute `krci-ai validate --verbose` for detailed analysis
- Quiet validation: Use `krci-ai validate --quiet` for summary-only results if needed
- Result interpretation: Analyze CLI output for issues and compliance status

### Validation Result Analysis

- Issue categorization: Organize findings by component type (agents, tasks, templates, data)
- Severity assessment: Classify issues as critical, warning, or informational
- Priority ranking: Determine resolution order based on impact and dependencies
- Root cause analysis: Understand underlying causes of validation failures

### Component Compliance Review

- Agent schema validation: Verify all agents meet JSON schema requirements exactly
- Task structure validation: Confirm tasks follow framework patterns and XML guidance
- Template validation: Check templates for proper variable usage and structure
- Data validation: Verify data files are accessible and properly organized

### Reference Integrity Verification

- Agent-task references: Confirm all agent.tasks entries resolve to existing task files
- Task-template references: Verify task frontmatter dependencies resolve to existing templates
- Task-data references: Confirm task references to data files resolve correctly
- Agent-data references: Verify agent references to behavioral data resolve properly

### Manual Inspection and Quality Assurance

- Framework pattern compliance: Review components for adherence to established patterns
- Critical principle inclusion: Verify agents include XML tag handling and customization principles
- XML guidance usage: Confirm tasks use XML tags for LLM processing guidance
- Content quality: Assess component content for completeness and effectiveness

### Issue Resolution and Remediation

- Fix implementation: Address validation failures with specific corrections
- Revalidation: Re-run CLI validation after implementing fixes
- Compliance confirmation: Verify complete validation success after remediation
- Documentation update: Record validation process and resolution actions for reference
