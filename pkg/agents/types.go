package agents

import (
	"context"

	"github.com/COG-GTM/teal-agents/pkg/kernel"
	"github.com/COG-GTM/teal-agents/pkg/plugins"
)

type SKAgent struct {
	ModelName      string
	ModelType      kernel.ModelType
	Agent          ChatCompletionAgent
	SOSupported    bool
	ExtraCollector plugins.ExtraDataCollector
}

type ChatCompletionAgent interface {
	Invoke(ctx context.Context, history *kernel.ChatHistory) (<-chan kernel.ChatMessageContent, error)
	InvokeStream(ctx context.Context, history *kernel.ChatHistory) (<-chan StreamingContent, error)
}

type StreamingContent struct {
	Content      string
	Role         kernel.AuthorRole
	IsComplete   bool
	FunctionCall *kernel.FunctionCallContent
}

type AgentConfig struct {
	Name          string   `yaml:"name"`
	Model         string   `yaml:"model"`
	SystemPrompt  string   `yaml:"system_prompt"`
	Temperature   *float64 `yaml:"temperature,omitempty"`
	Plugins       []string `yaml:"plugins,omitempty"`
	RemotePlugins []string `yaml:"remote_plugins,omitempty"`
}

type PromptExecutionSettings struct {
	Temperature            *float64
	MaxTokens              *int
	FunctionChoiceBehavior kernel.FunctionChoiceBehavior
	ResponseFormat         interface{}
}
