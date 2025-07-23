# Epic {{epic_number}}: {{epic_title}}

## Status

| Field                | Value                    |
|----------------------|--------------------------|
| Status               | {{status}}               |
| Priority             | {{priority}}            |
| Epic Owner           | {{owner}}               |
| Timeline             | {{timeline}}            |

<!-- Epic tracking and SDLC integration -->
<!-- Enables progress monitoring and PRD traceability validation -->

<!-- Template Guidance:
Status Options: Planning -> Approved -> In Progress -> Done -> Completed
Priority Example: Critical, High, Medium, Low (align with PRD requirement priorities)
Epic Owner Example: "Product Owner", "Tech Lead", "Development Team"
Timeline Example: "Sprint 1-3 (6 weeks)", "Q1 2024", "March-April 2024"
-->

## Overview

### Problem Statement

{{problem_statement}}

<!-- Clear, specific user problem this Epic solves from PRD analysis -->
<!-- Must trace back to PRD business requirements and user pain points -->

<!-- Template Guidance:
Problem Statement Example:
"Users struggle with agent discovery and IDE integration, leading to fragmented workflows and reduced productivity. Current manual agent selection requires deep technical knowledge, creating adoption barriers for 70% of target users."

Format Structure:
- Start with specific user pain point from PRD
- Include quantifiable impact or evidence
- Connect to PRD BR/NFR requirements
- Avoid solution-oriented language

✅ DO: Reference specific PRD requirements (BR1, BR2, NFR1...)
✅ DO: Include quantifiable user impact
✅ DO: Focus on user pain, not missing features
✅ DO: Use evidence from PRD user research
❌ DON'T: Describe what the system can't do
❌ DON'T: Include solution details in problem statement
❌ DON'T: Use vague terms like "users want better experience"
-->

### Goal

{{goal}}

<!-- Specific, measurable outcome that defines Epic completion -->
<!-- Must align with PRD success metrics and enable clear Story validation -->

<!-- Template Guidance:
Goal Example:
"Enable seamless agent discovery and IDE integration for 90% of target users within 2 weeks of first use, reducing agent selection time from 15 minutes to under 30 seconds while maintaining enterprise security standards."

Format Structure:
- Define specific, measurable outcome
- Include target user percentage or metrics
- Set clear timeline expectations
- Connect to PRD success criteria

✅ DO: Make goals specific and measurable
✅ DO: Include target metrics from PRD
✅ DO: Set realistic timeline boundaries
✅ DO: Connect to PRD success criteria
❌ DON'T: Use vague language like "improve user experience"
❌ DON'T: Set unmeasurable goals
❌ DON'T: Ignore PRD success metrics alignment
-->

### Target Users

{{target_users}}

<!-- Specific user personas from PRD who benefit from this Epic -->
<!-- Must align with PRD user segments and enable Story "As a user" scenarios -->

<!-- Template Guidance:
Target Users Example:
"Primary: Software Architects (60%) - designing system components and integration patterns
Secondary: Development Leads (30%) - implementing architectural decisions in code
Tertiary: Product Managers (10%) - validating technical feasibility and scope"

Format Structure:
- List primary, secondary, tertiary users with percentages
- Include specific roles and responsibilities
- Connect to PRD user segments
- Enable Story persona development

✅ DO: Use specific personas from PRD
✅ DO: Include user percentages or priority levels
✅ DO: Describe user context and responsibilities
✅ DO: Enable clear Story "As a [user]" scenarios
❌ DON'T: Use generic roles like "users" or "developers"
❌ DON'T: Create new personas not in PRD
❌ DON'T: Omit user context or responsibilities
-->

## Scope

### What's Included

{{in_scope}}

<!-- Specific features and functionality this Epic delivers -->
<!-- Must map to PRD requirements and enable Story breakdown -->

<!-- Template Guidance:
What's Included Example:
"1. Agent discovery UI with search and filter capabilities (BR2, BR3)
2. IDE integration framework for Cursor, VS Code, and JetBrains (NFR1)
3. Authentication and security layer for enterprise environments (NFR2)
4. Basic agent validation and health checking (BR4)"

Format Structure:
- Number items for clear tracking
- Reference specific PRD requirements in parentheses
- Focus on user-facing functionality
- Enable clear Story mapping

✅ DO: Reference specific PRD requirements (BR1, NFR2...)
✅ DO: Number items for clear tracking
✅ DO: Focus on user-facing functionality
✅ DO: Include enough detail for Story creation
❌ DON'T: Include technical implementation details
❌ DON'T: List features without PRD traceability
❌ DON'T: Use vague descriptions
-->

### What's Not Included

{{out_of_scope}}

<!-- Clear boundaries of what this Epic excludes -->
<!-- Prevents scope creep and guides Story prioritization -->

<!-- Template Guidance:
What's Not Included Example:
"1. Advanced analytics and usage monitoring (deferred to Epic 3)
2. Custom agent development frameworks (out of MVP scope)
3. Multi-language support beyond English (future roadmap item)
4. Legacy IDE support (VS 2019, Eclipse) - minimum viable platform focus"

Format Structure:
- Number excluded items for clarity
- Explain reasoning (deferred, out of scope, future)
- Reference other Epics for deferred items
- Set clear boundaries for Stories

✅ DO: Clearly state what's excluded and why
✅ DO: Reference future Epics for deferred items
✅ DO: Explain reasoning for exclusions
✅ DO: Set clear Story development boundaries
❌ DON'T: Leave scope boundaries unclear
❌ DON'T: Include items without rationale
❌ DON'T: Create artificial limitations
-->

### Dependencies

{{dependencies}}

<!-- Other Epics, systems, or external requirements this Epic needs -->
<!-- Critical for Story sequencing and implementation planning -->

<!-- Template Guidance:
Dependencies Example:
"Epic Dependencies:
- Epic 1: KubeRocketAI Baseline (foundation infrastructure)
- Epic 2: Core Engine (agent processing capabilities)

System Dependencies:
- Python 3.8+ runtime environment
- IDE extension APIs (Cursor, VS Code, JetBrains)
- Enterprise authentication systems (SSO integration)

External Dependencies:
- Design system components (UI framework)
- Security compliance approval (enterprise deployment)"

Format Structure:
- Group by type: Epic, System, External
- Include specific version requirements
- Note approval or integration needs
- Enable dependency validation

✅ DO: Group dependencies by type
✅ DO: Include specific version or requirement details
✅ DO: Note approval processes or integration needs
✅ DO: Enable clear dependency validation
❌ DON'T: List vague dependencies without specifics
❌ DON'T: Omit external approval processes
❌ DON'T: Ignore system requirement details
-->

## Solution Approach

{{solution_approach}}

<!-- High-level implementation strategy and architectural direction -->
<!-- Guides Story creation without prescribing technical details -->

<!-- Template Guidance:
Solution Approach Example:
"Implementation Strategy:
1. Modular UI framework with progressive enhancement for cross-IDE compatibility
2. Plugin architecture enabling independent IDE integration development
3. Centralized authentication service with configurable enterprise adapters
4. Agent validation pipeline with health monitoring and error recovery

Technical Approach:
- Frontend: React-based components for consistent IDE integration
- Backend: Python microservices with REST API interfaces
- Security: OAuth 2.0 with enterprise SSO federation
- Deployment: Containerized services with Kubernetes orchestration"

Format Structure:
- Separate implementation and technical approaches
- Focus on architectural decisions, not detailed implementation
- Enable Story technical context
- Include integration patterns and service boundaries

✅ DO: Provide architectural guidance for Stories
✅ DO: Include integration patterns and service boundaries
✅ DO: Focus on approach, not detailed implementation
✅ DO: Enable technical context for Story development
❌ DON'T: Include detailed code specifications
❌ DON'T: Prescribe exact implementation details
❌ DON'T: Ignore architectural considerations
-->

## Acceptance Criteria

{{acceptance_criteria}}

<!-- Specific, testable conditions that define Epic completion -->
<!-- Must include measurable outcomes and validation commands for Story verification -->

<!-- Template Guidance:
Acceptance Criteria Example:
"1. User can discover and select agents through IDE interface within 30 seconds
   - Validation: Time user workflows from agent search to selection
   - Command: `pytest tests/integration/agent_discovery_test.py`

2. IDE integration works across Cursor, VS Code, JetBrains with consistent UX
   - Validation: Cross-platform testing demonstrates feature parity
   - Command: `npm run test:integration -- --platform=all`

3. Enterprise authentication integrates with existing SSO systems
   - Validation: SSO login flow completes without manual intervention
   - Command: `python tests/auth/sso_integration_test.py`

4. Agent validation prevents malformed or insecure agent usage
   - Validation: Security scan passes with zero critical issues
   - Command: `security-scan --agents --threshold=critical`"

Format Structure:
- Number criteria for clear tracking and Story mapping
- Include validation methods and verification commands
- Define measurable success indicators
- Enable automated testing where possible

✅ DO: Include specific validation methods and commands
✅ DO: Make criteria measurable and testable
✅ DO: Number criteria for Story traceability
✅ DO: Enable automated verification where possible
❌ DON'T: Use subjective criteria like "works well"
❌ DON'T: Omit validation or verification steps
❌ DON'T: Create criteria that can't be tested
-->

## User Stories

{{planned_stories}}

<!-- Planned Stories that implement this Epic with clear breakdown -->
<!-- Enables immediate Story creation with Epic context and traceability -->

<!-- Template Guidance:
User Stories Example:
"Planned Stories for Implementation:

**Phase 1: Foundation (Sprint 1)**
- Story {{epic_number}}.01: Agent Discovery Interface
  - As a Software Architect, I want to search and filter available agents
  - Acceptance: Search returns relevant agents within 3 seconds
  - Dependencies: Epic 1 baseline infrastructure

- Story {{epic_number}}.02: Basic IDE Integration
  - As a Development Lead, I want to activate agents within my IDE
  - Acceptance: Agent activation works in Cursor and VS Code
  - Dependencies: Story {{epic_number}}.01 completion

**Phase 2: Enhancement (Sprint 2)**
- Story {{epic_number}}.03: Enterprise Authentication
  - As a Product Manager, I want secure agent access with SSO
  - Acceptance: SSO login integrates with existing enterprise systems
  - Dependencies: Security compliance approval

**Phase 3: Validation (Sprint 3)**
- Story {{epic_number}}.04: Agent Health Monitoring
  - As a Software Architect, I want to validate agent security and performance
  - Acceptance: Health dashboard shows agent status and metrics
  - Dependencies: All previous stories completed"

Format Structure:
- Group Stories by implementation phases/sprints
- Include user persona, goal, and basic acceptance criteria
- Note dependencies between Stories
- Use consistent naming: {{epic_number}}.01, {{epic_number}}.02, etc.

✅ DO: Group Stories by logical implementation phases
✅ DO: Include user persona and basic acceptance criteria
✅ DO: Note dependencies between Stories
✅ DO: Use consistent Story numbering within Epic
✅ DO: Enable immediate Story creation with this context
❌ DON'T: Include detailed Story implementation in Epic
❌ DON'T: Create Stories without clear user value
❌ DON'T: Ignore dependencies or sequencing requirements
-->
