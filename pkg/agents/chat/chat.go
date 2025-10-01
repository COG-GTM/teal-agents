package chat

import (
	"github.com/COG-GTM/teal-agents/pkg/types"
)

type ChatAgentConfig struct {
	Agent AgentConfig `yaml:"agent"`
}

type AgentConfig struct {
	Name          string   `yaml:"name"`
	Model         string   `yaml:"model"`
	SystemPrompt  string   `yaml:"system_prompt"`
	Temperature   *float64 `yaml:"temperature,omitempty"`
	Plugins       []string `yaml:"plugins,omitempty"`
	RemotePlugins []string `yaml:"remote_plugins,omitempty"`
}

type ChatAgent interface {
	types.BaseHandler

	GetAgentName() string
}

type Builder interface {
	Build(config ChatAgentConfig) (ChatAgent, error)
}
