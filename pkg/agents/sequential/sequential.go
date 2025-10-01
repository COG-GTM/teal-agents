package sequential

import (
	"context"

	"github.com/COG-GTM/teal-agents/pkg/types"
)

type TaskConfig struct {
	Name         string `yaml:"name"`
	TaskNo       int    `yaml:"task_no"`
	Description  string `yaml:"description"`
	Instructions string `yaml:"instructions"`
	Agent        string `yaml:"agent"`
}

type AgentConfig struct {
	Name          string   `yaml:"name"`
	Model         string   `yaml:"model"`
	SystemPrompt  string   `yaml:"system_prompt"`
	Temperature   *float64 `yaml:"temperature,omitempty"`
	Plugins       []string `yaml:"plugins,omitempty"`
	RemotePlugins []string `yaml:"remote_plugins,omitempty"`
}

type SequentialAgentConfig struct {
	Agents []AgentConfig `yaml:"agents"`
	Tasks  []TaskConfig  `yaml:"tasks"`
}

type SequentialAgent interface {
	types.BaseHandler

	GetTasks() []Task
}

type Task interface {
	GetName() string

	GetDescription() string

	Invoke(ctx context.Context, history []types.HistoryMessage, inputs map[string]interface{}) (*types.InvokeResponse, error)

	InvokeStream(ctx context.Context, history []types.HistoryMessage, inputs map[string]interface{}) (<-chan types.StreamResponse, error)
}

type Builder interface {
	Build(config SequentialAgentConfig) (SequentialAgent, error)
}
