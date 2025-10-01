package agents

import (
	"context"
	"fmt"

	"github.com/COG-GTM/teal-agents/pkg/kernel"
	"github.com/COG-GTM/teal-agents/pkg/plugins"
)

type AgentBuilder struct {
	kernelBuilder KernelBuilder
	authorization string
}

type KernelBuilder interface {
	BuildKernel(ctx context.Context, modelName, serviceID string, pluginNames, remotePluginNames []string, auth string, extraCollector plugins.ExtraDataCollector) (kernel.Kernel, error)
	GetModelTypeForName(modelName string) kernel.ModelType
	ModelSupportsStructuredOutput(modelName string) bool
}

func NewAgentBuilder(kernelBuilder KernelBuilder, authorization string) *AgentBuilder {
	return &AgentBuilder{
		kernelBuilder: kernelBuilder,
		authorization: authorization,
	}
}

func (b *AgentBuilder) BuildAgent(ctx context.Context, config AgentConfig, extraCollector plugins.ExtraDataCollector, outputType string) (*SKAgent, error) {
	kern, err := b.kernelBuilder.BuildKernel(
		ctx,
		config.Model,
		config.Name,
		config.Plugins,
		config.RemotePlugins,
		b.authorization,
		extraCollector,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build kernel: %w", err)
	}

	soSupported := b.kernelBuilder.ModelSupportsStructuredOutput(config.Model)

	modelType := b.kernelBuilder.GetModelTypeForName(config.Model)

	service, err := kern.GetService(config.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	settings := &PromptExecutionSettings{
		Temperature:            config.Temperature,
		FunctionChoiceBehavior: kernel.FunctionChoiceAuto,
	}

	if soSupported && outputType != "" {
		settings.ResponseFormat = outputType
	}

	agent := &basicChatAgent{
		kernel:       kern,
		service:      service,
		instructions: config.SystemPrompt,
		name:         config.Name,
		settings:     settings,
	}

	return &SKAgent{
		ModelName:      config.Model,
		ModelType:      modelType,
		Agent:          agent,
		SOSupported:    soSupported,
		ExtraCollector: extraCollector,
	}, nil
}

type basicChatAgent struct {
	kernel       kernel.Kernel
	service      kernel.ChatCompletionClient
	instructions string
	name         string
	settings     *PromptExecutionSettings
}

func (a *basicChatAgent) Invoke(ctx context.Context, history *kernel.ChatHistory) (<-chan kernel.ChatMessageContent, error) {
	resultChan := make(chan kernel.ChatMessageContent, 1)
	close(resultChan)
	return resultChan, fmt.Errorf("not implemented yet - to be completed in Session 4")
}

func (a *basicChatAgent) InvokeStream(ctx context.Context, history *kernel.ChatHistory) (<-chan StreamingContent, error) {
	resultChan := make(chan StreamingContent, 1)
	close(resultChan)
	return resultChan, fmt.Errorf("not implemented yet - to be completed in Session 4")
}
