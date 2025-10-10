package models

import (
	"github.com/COG-GTM/teal-agents/pkg/kernel"
)

type AnthropicProvider struct {
	apiKey string
	models map[string]bool
}

func NewAnthropicProvider(apiKey string) *AnthropicProvider {
	return &AnthropicProvider{
		apiKey: apiKey,
		models: map[string]bool{
			"claude-3-5-sonnet-20241022": true,
			"claude-3-opus-20240229":     true,
			"claude-3-sonnet-20240229":   true,
		},
	}
}

func (p *AnthropicProvider) GetModelType() kernel.ModelType {
	return kernel.ModelTypeAnthropic
}

func (p *AnthropicProvider) CreateClient(modelName, serviceID string) (kernel.ChatCompletionClient, error) {
	return nil, &NotImplementedError{Feature: "Anthropic client creation"}
}

func (p *AnthropicProvider) SupportsModel(modelName string) bool {
	return p.models[modelName]
}

func (p *AnthropicProvider) SupportsStructuredOutput(modelName string) bool {
	return true
}
