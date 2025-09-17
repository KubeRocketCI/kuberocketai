package processor

// AgentIdentityYamlRepresentation represents the identity section of an agent YAML.
type AgentIdentityYamlRepresentation struct {
	Name        string `yaml:"name"`
	ID          string `yaml:"id"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Role        string `yaml:"role"`
	Goal        string `yaml:"goal"`
	Icon        string `yaml:"icon"`
}

// AgentYamlRepresentation represents the structure of an agent YAML file.
type AgentYamlRepresentation struct {
	Agent struct {
		Identity AgentIdentityYamlRepresentation `yaml:"identity"`
		Tasks    []string                        `yaml:"tasks"`
	} `yaml:"agent"`
}

type TaskDependenciesYamlRepresentation struct {
	Dependencies struct {
		Templates  []string `yaml:"templates"`
		DataFiles  []string `yaml:"data"`
		Tasks      []string `yaml:"tasks"`
		McpServers []string `yaml:"mcp_servers"`
	} `yaml:"dependencies"`
}
