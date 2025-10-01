package models

import (
	"context"

	"github.com/COG-GTM/teal-agents/pkg/kernel"
)

type Provider interface {
	GetModelType() kernel.ModelType

	CreateClient(modelName, serviceID string) (kernel.ChatCompletionClient, error)

	SupportsModel(modelName string) bool

	SupportsStructuredOutput(modelName string) bool
}

type ProviderFactory interface {
	GetProvider(modelName string) (Provider, error)
	RegisterProvider(provider Provider) error
	GetProviderForModel(modelName string) (Provider, error)
}

type ChatCompletionFactory struct {
	providers []Provider
	appConfig interface{}
}

func NewChatCompletionFactory(appConfig interface{}) *ChatCompletionFactory {
	factory := &ChatCompletionFactory{
		providers: make([]Provider, 0),
		appConfig: appConfig,
	}
	return factory
}

func (f *ChatCompletionFactory) RegisterProvider(provider Provider) error {
	f.providers = append(f.providers, provider)
	return nil
}

func (f *ChatCompletionFactory) GetChatCompletionForModel(ctx context.Context, modelName, serviceID string) (kernel.ChatCompletionClient, error) {
	for _, provider := range f.providers {
		if provider.SupportsModel(modelName) {
			return provider.CreateClient(modelName, serviceID)
		}
	}
	return nil, &ModelNotFoundError{ModelName: modelName}
}

func (f *ChatCompletionFactory) GetModelType(modelName string) (kernel.ModelType, error) {
	for _, provider := range f.providers {
		if provider.SupportsModel(modelName) {
			return provider.GetModelType(), nil
		}
	}
	return "", &ModelNotFoundError{ModelName: modelName}
}

func (f *ChatCompletionFactory) ModelSupportsStructuredOutput(modelName string) bool {
	for _, provider := range f.providers {
		if provider.SupportsModel(modelName) {
			return provider.SupportsStructuredOutput(modelName)
		}
	}
	return false
}
