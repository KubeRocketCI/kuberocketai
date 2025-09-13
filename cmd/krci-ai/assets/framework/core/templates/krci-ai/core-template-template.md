# {{template_title}}

<instructions>
TEMPLATE STRUCTURE GUIDANCE: This template provides a framework for creating new templates. Templates use `{{variable}}` placeholders for dynamic content and include `<instructions>` XML tags for LLM processing guidance.

XML GUIDANCE SYSTEM: XML tags like `<instructions>` provide processing guidance for LLMs and help identify section boundaries. These tags are internal metadata ONLY and must never appear in final user output.

VARIABLE SYSTEM: Use `{{variable_name}}` placeholders for dynamic content that will be substituted during template usage. Choose descriptive variable names that clearly indicate expected content.
</instructions>

## {{main_section_header}}

<instructions>
Primary content section with variable placeholders for main template content. This section should contain the core structure and information that the template is designed to format.
</instructions>

{{main_content_area}}

## {{secondary_section_header}}

<instructions>
Secondary content section for additional template structure. Include variable placeholders and organizational elements that support the template's intended output format.
</instructions>

{{secondary_content_area}}

### {{subsection_header}}

<instructions>
Subsection for detailed template elements. Use this pattern for creating hierarchical template structure that supports comprehensive content organization.
</instructions>

{{subsection_content}}

## {{optional_section_header}}

<instructions>
Optional content section that may be populated based on context and requirements. Design template sections to be flexible for different usage scenarios.
</instructions>

{{optional_content_area}}

<instructions>
TEMPLATE CREATION GUIDANCE:

1. **Variable Design Principles**:
   - Use descriptive names: `{{project_name}}` not `{{name}}`
   - Follow consistent patterns: snake_case recommended
   - Group related variables logically within sections
   - Indicate variable types through naming (lists, sections, values)

2. **Variable Types and Usage**:
   - **Simple Values**: `{{title}}`, `{{description}}`, `{{author_name}}`
   - **Content Lists**: `{{requirements_list}}`, `{{feature_items}}`
   - **Large Sections**: `{{technical_details}}`, `{{implementation_approach}}`
   - **Optional Content**: `{{additional_notes}}`, `{{optional_appendix}}`

3. **XML Guidance Integration**:
   - Use `<instructions>` tags to guide LLM content generation
   - Provide examples and hints within instruction tags
   - Explain variable purpose and expected content format
   - Include formatting guidance for consistent output

4. **Template Structure Requirements**:
   - Follow markdown format with clear heading hierarchy
   - Organize content logically for intended output structure
   - Include sufficient variable placeholders for flexibility
   - Design for reusability across similar contexts

5. **LLM Processing Optimization**:
   - Structure templates for natural language processing
   - Include clear section boundaries and organization
   - Provide context clues for content generation
   - Optimize for consistent, high-quality output

TEMPLATE VALIDATION CHECKLIST:
- [ ] All variables use `{{variable_name}}` format consistently
- [ ] Variable names are descriptive and follow naming conventions
- [ ] `<instructions>` tags provide clear guidance for LLM processing
- [ ] Template structure supports intended output format
- [ ] Content organization is logical and accessible
- [ ] Template design enables reusability and flexibility

CRITICAL REMINDERS:
- XML tags (`<instructions>`) are for LLM guidance ONLY - never in final output
- Variables must be populated during template usage
- Template structure should facilitate high-quality content generation
- Design templates for use across multiple similar contexts
</instructions>
