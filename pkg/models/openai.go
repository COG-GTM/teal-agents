package models

import (
	"github.com/COG-GTM/teal-agents/pkg/kernel"
)

type OpenAIProvider struct {
	apiKey string
	models map[string]bool
}

func NewOpenAIProvider(apiKey string) *OpenAIProvider {
	return &OpenAIProvider{
		apiKey: apiKey,
		models: map[string]bool{
			"gpt-4o":        true,
			"gpt-4o-mini":   true,
			"gpt-4":         true,
			"gpt-3.5-turbo": true,
		},
	}
}

func (p *OpenAIProvider) GetModelType() kernel.ModelType {
	return kernel.ModelTypeOpenAI
}

func (p *OpenAIProvider) CreateClient(modelName, serviceID string) (kernel.ChatCompletionClient, error) {
	return nil, &NotImplementedError{Feature: "OpenAI client creation"}
}

func (p *OpenAIProvider) SupportsModel(modelName string) bool {
	return p.models[modelName]
}

func (p *OpenAIProvider) SupportsStructuredOutput(modelName string) bool {
	return modelName == "gpt-4o" || modelName == "gpt-4o-mini"
}
