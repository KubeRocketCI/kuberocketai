# Task: Documentation Review

Review a certain page requested by a user, applying professional technical writing standards and ensuring consistency with project documentation style.

## Prerequisites

<prerequisites>
- Page exists: Page for review is presented as a file or text in the chat
- Page is allowed to read: The file has read permissions for everyone
- Technical Writing standards available: Microsoft Writing Style Guide accessible for reference
- Project context: Understanding of project documentation style and standards
</prerequisites>

### Reference Assets

Dependencies:

- ./.krci-ai/data/common/sdlc-framework.md
- ./.krci-ai/templates/documentation-review.md

Validation: Verify all dependencies exist at specified paths before proceeding. HALT if any missing.

## Instructions

<instructions>
1. Get data to review: Read a page user provided you with
2. Review data: Based on the rules and best practices outlined in the [Microsoft Writing Style Guide](https://docs.microsoft.com/en-us/style-guide/), review the page
3. Refer to source: Examine the project structure you work in. If this is a documentation project, read a number of pages to understand the project style and report to the user that you refer to these pages as examples
4. Be professional: Respond as a Senior Technical Writer. Notify the user about what you have changed and why
5. Follow SDLC workflow: Reference [sdlc-framework.md](./.krci-ai/data/common/sdlc-framework.md) for documentation review standards
</instructions>

## Output Format

**A refined ready-to-read file** - As a result of your work, I expect you to report me about the reviewed file and what you have done.

## Success Criteria

<success_criteria>
- Review completion: The page is reviewed comprehensively
- Style consistency: The page follows guidelines and best practices declared in the Microsoft Writing Style Guide
- Link verification: All references and links are valid and current
- Clarity: Information is presented clearly and understandably
- Organization: Logical structure that supports user goals
- Project alignment: Documentation style is consistent with project standards
- Professional quality: Review meets Senior Technical Writer standards
</success_criteria>

## Execution Checklist

<execution_checklist>

### Discovery Phase

- Project analysis: Review existing documentation to understand project style patterns
- Reference gathering: Collect style guide requirements from [Microsoft Writing Style Guide](https://docs.microsoft.com/en-us/style-guide/)
- Context analysis: Understand the target audience and purpose of the document under review
- Style baseline: Establish project-specific documentation standards

### Review Phase

- Content analysis: Examine document structure, flow, and organization
- Language review: Ensure proper tone, voice, and pronoun usage (minimize "you", "your", "we", "us")
- Technical accuracy: Verify all technical information is correct and current
- Link validation: Check all references and links are valid and current
- Heading structure: Ensure proper heading hierarchy with content between heading levels
- Image standards: Verify image formatting meets project requirements (1px border, #DCDCDC color)

### Validation Phase

- Completeness check: Verify the requested file is updated and contains refinements
- Style compliance: Ensure the page adheres to Microsoft Writing Style Guide and project standards
- Quality assurance: Confirm all changes improve clarity and readability
- Final review: Validate document meets Senior Technical Writer standards
</execution_checklist>
