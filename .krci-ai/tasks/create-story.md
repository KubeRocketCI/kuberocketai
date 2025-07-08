# Task: create-story

**Description:**
Create detailed user story with acceptance criteria, implementation plan, and QA checklist

## Instructions

1. Gather story requirements including feature description and business value
2. Assign story number and create clear, descriptive story title
3. Identify the target persona and user role for the story
4. Define clear user goal and business value proposition using "As a..., I want..., so that..." format
5. Break down the requirement into specific, testable acceptance criteria
6. Estimate story points (1, 2, 3, 5, 8, 13) based on complexity and effort required
7. Set initial status (e.g., "Draft", "Ready for Development", "In Progress")
8. Identify epic association and priority level (Low, Medium, High, Critical)
9. Identify dependencies on other stories, features, or technical components
10. Create detailed description explaining the context and rationale
11. Create detailed implementation plan with technical approach and development steps
12. Define placeholder for implementation results (to be filled during development)
13. Define comprehensive QA checklist covering functional and non-functional aspects
14. Include Jira ticket reference if applicable (or mark as "TBD")
15. Format the story using [template](./.krci-ai/templates/story.md) ensuring **ALL** template variables are populated:
    - `{{story_number}}`: Sequential story identifier
    - `{{story_title}}`: Clear, descriptive title
    - `{{status}}`: Current development status
    - `{{epic}}`: Associated epic name/ID
    - `{{priority}}`: Business priority level
    - `{{story_points}}`: Effort estimation
    - `{{jira_ticket}}`: Jira reference or "TBD"
    - `{{dependencies}}`: List of dependencies or "None"
    - `{{persona}}`: Target user persona
    - `{{goal}}`: What the user wants to achieve
    - `{{business_value}}`: Why this is valuable
    - `{{acceptance_criteria}}`: Specific, testable criteria
    - `{{description}}`: Detailed context and rationale
    - `{{implementation_plan}}`: Technical approach and steps
    - `{{implementation_results}}`: Placeholder for actual results
    - `{{qa_checklist}}`: Comprehensive testing checklist
18. Ensure no template variable is left unfilled or with placeholder text
