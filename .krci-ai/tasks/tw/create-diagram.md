---
dependencies:
  data:
  mcp_servers:
---

# Task: Documentation Diagram Design

## Description

This task helps users create clear and consistent diagrams for Markdown documentation pages using Mermaid syntax.

The goal of this task is to transform a user-provided algorithm, workflow, or interaction description into a Mermaid diagram that can be embedded directly into Markdown documentation.

The task focuses on visualizing:

- component interactions,
- system logic and control flow,
- step-by-step algorithms and workflows.

## Prerequisites

<prerequisites>
- User has an algorithm, workflow, or logic description they want to visualize **or**.
- User knows where this logic is already described (documentation page, section, or code reference).
- Mermaid configuration: The documentation system (Docusaurus, MkDocs, etc.) has Mermaid support enabled with required modules and configurations for the diagram type being created.
</prerequisites>

## Instructions

<instructions>
1. When the **"doc-diagram"** task is requested, ensure the user has provided:
   - an algorithm, workflow, or logic description **or**
   - a reference to where this logic is already documented. If not provided, ask the user to describe the logic in a step-by-step or structured form.

2. Check Mermaid configuration: Verify that the project's Mermaid configuration includes the required modules and settings for the diagram type you plan to create. Check configuration files (e.g., `docusaurus.config.js`, `mkdocs.yml`, or project-specific Mermaid config files). If modules or configurations are missing, update them accordingly. Common requirements include:
   - Enabling specific Mermaid diagram types (flowchart, sequenceDiagram, stateDiagram, etc.),
   - Configuring theme and styling options,
   - Setting up security policies if needed,
   - Ensuring proper initialization in the documentation framework.

3. Show the user an example of the expected input format, such as:

   CI/CD pipeline logic:
   1. Pipeline is triggered by an external event,
   2. Application is built,
   3. Code quality checks are executed (linters, scanners, tests),
   4. Application is deployed,
   5. Application is promoted to the next environment.

4. Analyze the provided logic and determine:
   - the most suitable Mermaid diagram type, for example:
     - flowchart — for algorithms and step-by-step processes,
     - sequenceDiagram — for component interactions,
     - stateDiagram — for lifecycle or state transitions,
     - graph TD/LR — for high-level architecture,
   - whether a single diagram is sufficient or the logic should be split into multiple diagrams,
   - required Mermaid modules and configurations for the selected diagram type.

5. Build a Mermaid diagram that accurately reflects the provided logic, uses clear and meaningful node names, follows a readable layout direction (TD, LR, etc.), and avoids unnecessary visual complexity.
6. Generate the diagram in pure Mermaid syntax, ready to be embedded into Markdown, for example:

   ```mermaid
   flowchart TD
     A[External Trigger] --> B[Build Application]
     B --> C[Quality Checks]
     C --> D[Deploy Application]
     D --> E[Promote to Next Environment]
   ```

7. If the logic is ambiguous, incomplete, or overloaded, ask clarifying questions before generating the diagram, or propose assumptions explicitly and ask the user to confirm them.
8. When relevant, recommend splitting the logic into sub-diagrams, adding diagram titles or comments, and placing the diagram in a specific documentation section or page.
9. Verify configuration completeness: After generating the diagram, confirm that all required Mermaid modules and configurations are present in the project configuration files. If any are missing, update the configuration files to include them, ensuring the diagram will render correctly in the documentation system.
</instructions>

## Output Format

As an output, you are expected to provide:

- A Mermaid diagram wrapped in a Markdown-compatible code block.
- A short explanation of:
  - what the diagram represents,
  - why this Mermaid diagram type was chosen.
- Optional improvement suggestions (if applicable), such as:
  - alternative diagram types,
  - additional diagrams for clarity,
  - alignment with existing documentation pages.

## Mermaid Diagram Guidelines

- Prefer simplicity over visual density.
- Use descriptive node labels instead of internal abbreviations when possible.
- Keep diagrams readable in documentation layouts (Docusaurus, MkDocs).
- Do not include platform-specific styling unless explicitly requested.
- Ensure the diagram can be rendered by standard Mermaid implementations.
- Verify required modules: Different diagram types may require specific Mermaid modules to be enabled in the project configuration (e.g., `mermaid.flowchart`, `mermaid.sequence`, `mermaid.state`).
- Check configuration files: Ensure the documentation framework's Mermaid configuration includes all necessary modules and settings for the diagram type being used.
- If proper configuration is missing, ask user for approval for changing project documentation.

## Framework Context: Project Documentation

Diagrams are primarily used to:

- explain CI/CD workflows and deployment flows,
- visualize interactions between components and third-party tools,
- support Use Cases and Guides with visual explanations.

Diagrams should complement the text and not replace step-by-step instructions entirely.

## Success Criteria

<success_criteria>
- User-provided logic analyzed and understood.
- Appropriate Mermaid diagram type selected.
- Diagram accurately reflects the algorithm or interaction.
- Output is valid Mermaid syntax and Markdown-ready.
- Diagram improves understanding of the documented process.
- User receives clear next steps or improvement suggestions (if needed).
</success_criteria>

## Execution Checklist

### Preparation Phase

<execution_phase>
- User intent and diagram goal identified.
- Algorithm or logic description collected or referenced.
</execution_phase>

### Analysis Phase

<execution_phase>
- Diagram type selected with justification.
- Diagram scope defined (single vs multiple diagrams).
- Ambiguities identified and resolved or documented.
</execution_phase>

### Generation Phase

<execution_phase>
- Mermaid diagram generated using best practices.
- Node naming and flow validated for clarity.
- Diagram wrapped in Markdown-compatible format.
- Required Mermaid modules and configurations verified and updated if missing.
</execution_phase>

### Validation Phase

<validation_phase>
- Diagram ready for embedding into documentation.
- Mermaid configuration verified and complete for the diagram type.
- User informed about placement and possible extensions.
- Task completed with minimal iteration friction.
</validation_phase>