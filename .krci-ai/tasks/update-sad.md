# Task: Update Architecture Documentation

## Description

Update existing multi-file architecture documentation to reflect new requirements, Epic changes, or technical decisions while maintaining system consistency and development guidance. This update ensures architecture remains aligned with current PRD requirements and Epic implementations across all relevant architecture sections.

## Prerequisites

- [ ] **Existing architecture**: Current architecture files exist in `/docs/architecture/` directory following SAD appendix structure
- [ ] **Change trigger**: Clear reason for update (PRD changes, Epic updates, technical constraints)
- [ ] **Updated requirements**: New or modified BR/NFR requirements from PRD updates
- [ ] **Impact scope**: Understanding of which architectural sections and components are affected

### Reference Assets

Dependencies:

- ./.krci-ai/templates/sad-template.md
- ./.krci-ai/data/architecture-principles.md
- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/templates/architecture-review.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

1. **Follow SDLC workflow**: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for change management process and impact assessment
2. **Apply design principles**: Maintain consistency with [architecture-principles.md](./.krci-ai/data/architecture-principles.md) and [design-patterns.md](./.krci-ai/data/design-patterns.md)
3. **Update relevant sections**: Modify appropriate architecture files based on change scope using [sad-template.md](./.krci-ai/templates/sad-template.md) structure
4. **Document decisions**: Update 08-architectural-decisions.md with new ADR entries for significant changes

## Output Format

**Updated Architecture Files** - Modify existing numbered section files in `/docs/architecture/`:

### Common Update Targets

- [ ] **02-introduction.md** - Update PRD requirements mapping, scope changes, stakeholder updates
- [ ] **06-target-architecture.md** - Modify target architecture, solution strategy changes
- [ ] **07-transition-migration.md** - Update Epic breakdown guidance and migration approach
- [ ] **08-architectural-decisions.md** - Add new Architecture Decision Records for significant changes

### Conditional Updates (Based on Change Type)

- [ ] **01-executive-summary.md** - Business context or strategic changes
- [ ] **03-context.md** - Technology strategy or infrastructure changes
- [ ] **04-requirements.md** - Functional/non-functional requirement updates
- [ ] **05-baseline-architecture.md** - Current state changes
- [ ] **09-cross-cutting-concerns.md** - Security, scalability, or observability updates
- [ ] **10-quality-assurance.md** - Testing strategy changes
- [ ] **11-appendices.md** - Reference material updates

## Success Criteria

- [ ] **Files updated** - All affected architecture sections reflect changes accurately
- [ ] **Change documented** - Clear record of what changed and architectural rationale in 08-architectural-decisions.md
- [ ] **Requirements aligned** - Updated BR/NFR requirements properly addressed in 02-introduction.md and other relevant sections
- [ ] **Epic impact assessed** - Identified which Epics need updates due to architectural changes in 07-transition-migration.md
- [ ] **Consistency maintained** - Architecture decisions remain coherent across all sections
- [ ] **Quality preserved** - Documentation maintains professional architecture standards per [sad-template.md](./.krci-ai/templates/sad-template.md)
