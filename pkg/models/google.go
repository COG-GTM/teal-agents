package models

import (
	"github.com/COG-GTM/teal-agents/pkg/kernel"
)

type GoogleProvider struct {
	apiKey string
	models map[string]bool
}

func NewGoogleProvider(apiKey string) *GoogleProvider {
	return &GoogleProvider{
		apiKey: apiKey,
		models: map[string]bool{
			"gemini-pro":        true,
			"gemini-pro-vision": true,
			"gemini-1.5-pro":    true,
			"gemini-1.5-flash":  true,
		},
	}
}

func (p *GoogleProvider) GetModelType() kernel.ModelType {
	return kernel.ModelTypeGoogle
}

func (p *GoogleProvider) CreateClient(modelName, serviceID string) (kernel.ChatCompletionClient, error) {
	return nil, &NotImplementedError{Feature: "Google client creation"}
}

func (p *GoogleProvider) SupportsModel(modelName string) bool {
	return p.models[modelName]
}

func (p *GoogleProvider) SupportsStructuredOutput(modelName string) bool {
	return true
}
