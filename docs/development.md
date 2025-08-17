# KubeRocketAI Development Architecture Guide

## **Unified Discovery Method Architecture**

The following diagram shows the improved architecture after unifying the discovery methods, eliminating code duplication while maintaining all functionality:

```mermaid
graph TB
    %% CLI Entry Points
    A1["`**krci-ai install --agent X**
    cmd/install.go:57`"] --> B1{"`**Selective?**`"}
    A2["`**krci-ai validate**
    cmd/validate.go:52`"] --> V1["`**runValidate()**
    validate.go:70`"]
    A3["`**krci-ai bundle --agent X**
    cmd/bundle.go:64`"] --> BU1["`**runBundle()**
    bundle.go:217`"]
    A4["`**krci-ai tokens --agent X**
    cmd/tokens.go:60`"] --> T1["`**runTokensCommand()**
    tokens.go:79`"]

    %% INSTALL FLOW (Selective Path)
    B1 -->|Yes| C1["`**runSelectiveInstallation()**
    install.go:74`"]
    B1 -->|No| D1["`**runFullInstallation()**
    install.go:121`"]

    C1 --> E1["`**ParseAgentList()**
    bundle.go:104`"]
    E1 --> F1["`**NewInstaller()**
    installer.go:192`"]
    F1 --> G1["`**InstallSelective()**
    installer.go:485`"]

    %% VALIDATE FLOW
    V1 --> V2["`**NewFrameworkAnalyzer()**
    analyzer.go:93`"]
    V2 --> V3["`**OptimizedAnalyzeFramework()**
    analyzer.go:105`"]
    V3 --> V4["`**detectCriticalIssues()**
    analyzer.go:137`"]
    V3 --> V5["`**detectWarningIssues()**
    analyzer.go:121`"]
    V3 --> V6["`**generateFrameworkInsights()**
    analyzer.go:128`"]

    %% BUNDLE FLOW
    BU1 --> BU2["`**validateBundleFlags()**
    bundle.go:223`"]
    BU2 --> BU3["`**setupAndValidateFramework()**
    bundle.go:257`"]
    BU3 --> BU4["`**parseAndValidateAgents()**
    bundle.go:234`"]
    BU4 --> BU5["`**collectBundleContent()**
    bundle.go:238`"]

    %% TOKENS FLOW
    T1 --> T2["`**tokens.NewCalculator()**
    calculator.go:48`"]
    T2 --> T3["`**NewDiscovery()**
    discovery.go:81`"]
    T3 --> T4{"`**Agent or All?**`"}
    T4 -->|Agent| T5["`**DiscoverAgentWithDependencies()**
    discovery.go:269`"]
    T4 -->|All| T6["`**DiscoverAgentsWithDependencies()**
    discovery.go:217`"]

    %% UNIFIED DISCOVERY METHOD
    G1 --> H1["`**NewEmbeddedSource()**
    source.go:82`"]
    H1 --> I1["`**NewDiscoveryWithSource()**
    discovery.go:89`"]
    I1 --> UNIFIED["`**🔄 UNIFIED METHOD**
    DiscoverAgentsWithDependencies(agentNames ...string)
    discovery.go:227`"]

    V6 --> UNIFIED
    BU5 --> UNIFIED
    T5 --> UNIFIED
    T6 --> UNIFIED

    %% Unified Method Components
    UNIFIED --> UNI1["`**validateSourceRequirements()**
    discovery.go:252`"]
    UNIFIED --> UNI2["`**getValidationInsights()**
    discovery.go:260`"]
    UNIFIED --> UNI3["`**getTargetAgents()**
    discovery.go:309`"]
    UNIFIED --> UNI4["`**buildAgentDependencies()**
    discovery.go:338`"]

    %% Source-specific insights
    UNI2 --> UNI2A["`**getFilesystemInsights()**
    discovery.go:268`"]
    UNI2 --> UNI2B["`**getEmbeddedInsights()**
    discovery.go:278`"]

    %% Core Validation Engine (SHARED BY ALL)
    UNI2A --> CORE["`**CORE VALIDATION ENGINE**
    FrameworkAnalyzer.AnalyzeFramework()
    FrameworkAnalyzer.AnalyzeEmbeddedDependencies()
    analyzer.go (Enhanced)`"]
    UNI2B --> CORE

    %% Core Validation Components (REUSED BY ALL COMMANDS)
    CORE --> CORE1["`**buildAgentRelationship()**
    analyzer.go:700+`"]
    CORE --> CORE2["`**extractYAMLTasks()**
    analyzer.go:850+`"]
    CORE --> CORE3["`**processAgentTasks()**
    analyzer.go:900+`"]
    CORE --> CORE4["`**addTaskDependencies()**
    analyzer.go:950+`"]

    %% Dependency Resolution Output
    CORE1 --> OUTPUT["`**ComponentRelationship**
    validation/insights.go:44
    *UNIFIED OUTPUT FORMAT*`"]
    CORE2 --> OUTPUT
    CORE3 --> OUTPUT
    CORE4 --> OUTPUT

    %% Command-Specific Output Processing
    OUTPUT --> OUT1["`**install**: installFilteredAssets()
    installer.go:528`"]
    OUTPUT --> OUT2["`**validate**: ValidationReport
    validation/output.go:27`"]
    OUTPUT --> OUT3["`**bundle**: generateBundleMarkdown()
    bundle.go:800+`"]
    OUTPUT --> OUT4["`**tokens**: TokenCalculation
    tokens/calculator.go:100+`"]

    %% Error Handling & Validation Paths (SHARED)
    CORE --> ERR1["`**=ERROR HANDLING**
    detectBrokenInternalLinks()
    analyzer.go:179`"]
    ERR1 --> ERR2["`**detectMissingTaskFiles()**
    analyzer.go:148`"]
    ERR1 --> ERR3["`**detectFormatIssues()**
    analyzer.go:169`"]

    %% Unified Components Styling
    style E1 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style F1 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style I1 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style V2 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style V3 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style BU3 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style T2 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style T3 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style UNI1 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style UNI2 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style UNI3 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style UNI4 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style UNI2A fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style UNI2B fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style CORE1 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style CORE2 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style CORE3 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style CORE4 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style OUTPUT fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style ERR1 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style ERR2 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    style ERR3 fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px

    %% Unified Method Styling
    style UNIFIED fill:#fff3e0,stroke:#f57c00,stroke-width:4px
    %% Core Engine Styling
    style CORE fill:#ffeb3b,stroke:#f57f17,stroke-width:4px
```
