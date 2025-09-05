# MD2XML Refactoring Plan

!!!IMPORTANT: ALL CHANGES MUST BE UNDER: "./cmd/krci-ai/assets/framework/core" path

## Overview

This document provides a systematic approach to refactor framework tasks from pure Markdown structure to hybrid Markdown + XML tags structure for optimal LLM parsing while maintaining human readability.

## Background Context

**Original Problem**: Pure Markdown structure made it difficult for LLM agents to:
- Quickly identify section boundaries
- Parse structured content consistently
- Hand off work between different agent personas
- Maintain scope boundaries during task execution

**Solution**: Hybrid structure combining XML tags for LLM parsing with Markdown headings for human navigation.

## Pre-Refactoring Analysis

### 1. Commit History Review

```bash
# Identify the baseline commit that established XML patterns
git show b67cffedc899aa98a6d395c6d92b2b0bbb7325d0

# Examine existing XML implementations
grep -r "<\|>" cmd/krci-ai/assets/framework/core/tasks/create-story.md
```

### 2. Framework Consistency Check

- Review `create-story.md` for established XML patterns
- Identify sections that should be XML-tagged vs. remain as headers

### 3. Target Task Analysis

Before refactoring any task file:

```bash
# Check current structure
grep -n "^##" path/to/target-task.md

# Count sections and identify patterns
wc -l path/to/target-task.md
```

## Template Dependencies Check

### Before Task Refactoring: Verify Template Alignment

**Critical**: Check if the refactoring requires template updates. In commit `b67cffe`, both tasks AND templates were modified with XML tags.

```bash
# Check if templates reference the target task
grep -r "target-task-name" cmd/krci-ai/assets/framework/core/templates/

# Example: story.md template was extensively modified with XML tags
git show b67cffedc899aa98a6d395c6d92b2b0bbb7325d0 -- cmd/krci-ai/assets/framework/core/templates/story.md
```

**Template XML Changes Made in Reference Commit:**
- Added `<instructions>` sections with detailed formatting guidance
- Added `<status>`, `<dependencies>`, `<user_story>` etc. XML wrappers
- Converted HTML comments to XML instruction blocks
- Simplified formatting requirements (removed complex sub-bullets)
- Added comprehensive section ordering requirements

### Template Refactoring Requirements

When tasks reference templates (e.g., story creation tasks reference `story.md`):

1. **Template Structure Alignment**: Templates must match the XML structure expected by refactored tasks
2. **Instruction Consistency**: Template guidance should use `<instructions>` blocks, not HTML comments
3. **Content Simplification**: Templates should guide toward simplified content structure for LLM parsing
4. **Section Ordering**: Templates must enforce proper section sequence through XML instructions

### Template Conversion Example

**BEFORE (HTML comments):**

```markdown
## Status

| Field                  | Value                       |
|------------------------|-----------------------------|
| Status                 | {{status}}                  |
| Epic Reference         | {{epic_reference}}          |

<!-- Status tracking and Epic/PRD traceability -->
<!-- Enables progress monitoring and dependency validation -->

<!-- Template Guidance:
Status Options: Draft -> Approved -> In Progress -> Done -> Completed
Epic Reference Example: "Epic 1 - Unified Agent Activation"
Priority Example: Critical, High, Medium, Low
 -->
```

**AFTER (XML structure with instructions INSIDE):**

```markdown
## Status

<status>
| Field                  | Value                       |
|------------------------|-----------------------------|
| Status                 | {{status}}                  |
| Epic Reference         | {{epic_reference}}          |

<instructions>
Status tracking and Epic/PRD traceability. Enables progress monitoring and dependency validation.

Status Options: Draft -> Approved -> In Progress -> Done -> Completed
Epic Reference Example: Epic 1 - Unified Agent Activation
Priority Example: Critical, High, Medium, Low

CRITICAL: Status section contains ONLY these fields. Do NOT add Dependencies or other fields here.
</instructions>
</status>
```

**Key Template Changes:**
1. **HTML comments → XML instructions**: `<!-- comment -->` becomes `<instructions>content</instructions>`
2. **Add XML section wrappers**: Content gets wrapped in semantic tags like `<status>`
3. **CRITICAL: Instructions INSIDE sections**: `<instructions>` blocks must be positioned INSIDE XML section wrappers, not after them
4. **Simplify formatting guidance**: Remove complex bullet structures, use simple text
5. **Add critical constraints**: Explicit requirements like "CRITICAL: Status section contains ONLY..."
6. **Remove example formatting**: Convert formatted examples to plain text for LLM clarity

## Refactoring Process

### Phase 1: XML Tag Implementation

**Step 1: Identify Core Sections for XML Tagging**
Tag these standard sections across all task files:
- `<prerequisites>` - Dependencies and requirements
- `<instructions>` - Step-by-step guidance
- `<success_criteria>` - Measurable completion indicators

**Step 2: Identify Role-Specific Sections**
Based on task type, tag specialized sections:
- **PO tasks**: `<business_validation>`, `<story_format_review>`, `<acceptance_criteria_validation>`
- **Dev tasks**: `<technical_completeness>`, `<task_implementation_review>`, `<architecture_validation>`
- **Architect tasks**: `<system_architecture_alignment>`, `<component_design_validation>`, `<integration_pattern_review>`

**Step 3: Apply XML Tags to Content**

```markdown
<!-- BEFORE -->
## Prerequisites

- [ ] **Story exists**: Target story file exists...
- [ ] **Epic context**: Understanding of parent Epic...

<!-- AFTER -->
## Prerequisites

<prerequisites>
- Story exists: Target story file exists...
- Epic context: Understanding of parent Epic...
</prerequisites>
```

**XML Tagging Rules:**
1. Remove `[ ]` checkbox formatting from list items within XML tags
2. Remove `**bold**` formatting from key terms within XML tags
3. Keep content structure but simplify formatting for LLM parsing
4. Ensure every opening tag has a corresponding closing tag

### Phase 2: Header Structure Optimization

**Step 4: Move Headers Outside XML Tags**

```markdown
<!-- CORRECT STRUCTURE -->
### Business Requirements Validation

<business_validation>
- User value assessment: Confirm story delivers clear...
- Business justification: Validate business need...
</business_validation>
```

**Step 5: Fix Heading Hierarchy**
- Ensure proper H1 → H2 → H3 progression
- Convert H4 to H3 to avoid MD001 lint errors
- Maintain logical document structure

### Phase 3: Structure Validation

**Step 6: Verify XML Tag Integrity**

```bash
# Check all tags are properly closed
grep -n "<\|>" path/to/task.md

# Count opening vs closing tags (should be equal)
grep -c "<[^/]" path/to/task.md  # Opening tags
grep -c "</" path/to/task.md     # Closing tags
```

**Step 7: Restore Missing H2 Headers**
Ensure these critical H2 sections exist:

```markdown
## Prerequisites
## Instructions
## Success Criteria
## Execution Checklist  # (if applicable)
```

**Step 8: Content Quality Check**
- XML tags contain pure content without redundant headers
- Headers provide document navigation structure
- Content flows logically between tagged and untagged sections

## Quality Assurance Checklist

### Structure Validation

- [ ] All XML tags properly opened and closed
- [ ] No orphaned tags detected via grep
- [ ] Headers follow H1 → H2 → H3 hierarchy
- [ ] Major sections (Prerequisites, Instructions, Success Criteria) have H2 headers

### Content Validation

- [ ] XML tagged content simplified (no `**bold**`, no `[ ]` checkboxes)
- [ ] Headers positioned outside XML tags
- [ ] Template instructions positioned INSIDE XML sections (not after them)
- [ ] Original content meaning preserved
- [ ] Framework consistency maintained with `create-story.md` patterns

### Agent Experience Validation

- [ ] Section boundaries clearly identified for LLM parsing
- [ ] Cross-agent handoff information easily extractable
- [ ] Scope boundaries maintained through tag structure
- [ ] Human readability preserved through header hierarchy

## Implementation Commands

### Single Task + Template Refactoring

```bash
# 1. Backup originals
cp path/to/task.md path/to/task.md.backup
cp path/to/template.md path/to/template.md.backup  # if applicable

# 2. Check template dependencies
grep -r "template-name" path/to/task.md

# 3. Refactor template FIRST (if referenced by task)
#    - Convert HTML comments to <instructions> blocks
#    - Add XML section wrappers
#    - Position <instructions> blocks INSIDE XML sections
#    - Simplify formatting guidance
#    - Add section ordering requirements

# 4. Refactor task file
#    - Apply XML tags (use MultiEdit tool)
#    - Move headers outside tags
#    - Fix heading hierarchy
#    - Restore H2 headers
#    - Validate structure

# 5. Test template-task integration
#    - Verify task can generate content using updated template
#    - Check XML structure consistency
```

### Batch Refactoring

```bash
# Find all task-template pairs needing refactoring
find cmd/krci-ai/assets/framework/core/tasks/ -name "*.md" -not -name "create-story.md"
find cmd/krci-ai/assets/framework/core/templates/ -name "*.md"

# Process template files first, then their dependent tasks
# Follow the 5-step process above for each pair
```

## Framework Benefits Post-Refactoring

### For LLM Agents

1. **Rapid Section Navigation**: XML tags enable instant content location
2. **Clean Parsing**: Structured content without formatting interference
3. **Cross-Agent Integration**: Standardized data extraction between agent types
4. **Scope Clarity**: XML boundaries prevent scope creep during task execution

### For Human Users

1. **Visual Structure**: H2/H3 headers maintain document outline
2. **Easy Navigation**: Traditional Markdown TOC functionality
3. **Content Flow**: Logical reading experience preserved
4. **Maintenance**: Clear separation between navigation and content

## Post-Implementation Validation

### Lint Compliance

```bash
# Should pass without MD001, MD033 errors
markdownlint-cli2 path/to/refactored-task.md
```

### Agent Testing

Test refactored task with agent personas to ensure:
- Quick section identification
- Proper scope adherence
- Effective cross-agent handoffs
- Template compliance

### Framework Consistency

```bash
# Compare XML patterns across tasks
grep -A5 -B5 "<prerequisites>" cmd/krci-ai/assets/framework/core/tasks/*.md
grep -A5 -B5 "<instructions>" cmd/krci-ai/assets/framework/core/tasks/*.md
```

## Success Metrics

- **Structure**: All XML tags properly nested and closed
- **Hierarchy**: No MD001 heading increment violations
- **Consistency**: XML patterns match established framework standards
- **Functionality**: LLM agents can parse and execute tasks effectively
- **Readability**: Human users can navigate and understand content easily

## Common Pitfalls to Avoid

### Task Refactoring Pitfalls

1. **Don't nest headers inside XML tags** - breaks parsing
2. **Don't leave orphaned XML tags** - causes structure errors
3. **Don't skip H2 restoration** - breaks document navigation
4. **Don't change content meaning** - preserve original intent
5. **Don't ignore heading hierarchy** - creates lint violations

### Template Refactoring Pitfalls

1. **Don't refactor tasks before templates** - creates structure mismatches
2. **Don't leave HTML comments in templates** - inconsistent with XML instruction pattern
3. **Don't position instructions outside XML sections** - breaks reference commit pattern, instructions must be INSIDE sections
4. **Don't create complex formatting in templates** - LLMs prefer simplified structure
5. **Don't ignore template-task dependencies** - breaks content generation workflow
6. **Don't skip template testing** - templates must generate valid XML-structured output

## Future Framework Evolution

This hybrid Markdown+XML approach provides:
- **Scalability**: Easy to extend with new XML tag types
- **Maintainability**: Clear separation of structure vs. content
- **Flexibility**: Works across different agent personas and task types
- **Standards Compliance**: Maintains Markdown compatibility while adding LLM optimization

Use this plan as a template for systematically upgrading framework tasks to the hybrid structure that optimizes for both human readability and LLM processing efficiency.

## Post-Implementation Notes

### Missing Dependencies Handling

When implementing XML structure updates, agents should:
1. **Check for missing Reference Assets section** - If task lacks dependencies section, add it following pattern from create-story.md, create-prd.md
2. **Validate dependency paths** - Ensure referenced .krci-ai paths exist or ask user for feedback on correct dependencies
3. **Add markdown links** - Convert plain references to proper markdown links where applicable
4. **Request clarification** - If unsure about task-specific dependencies, ask user for guidance rather than assume
