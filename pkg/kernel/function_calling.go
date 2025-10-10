package kernel

import (
	"context"
)

type FunctionCallContent struct {
	ID           string                 `json:"id"`
	PluginName   string                 `json:"plugin_name"`
	FunctionName string                 `json:"function_name"`
	Arguments    map[string]interface{} `json:"arguments"`
}

type FunctionResultContent struct {
	CallID       string      `json:"call_id"`
	PluginName   string      `json:"plugin_name"`
	FunctionName string      `json:"function_name"`
	Result       interface{} `json:"result"`
	Error        error       `json:"error,omitempty"`
}

type FunctionChoiceBehavior string

const (
	FunctionChoiceAuto     FunctionChoiceBehavior = "auto"
	FunctionChoiceNone     FunctionChoiceBehavior = "none"
	FunctionChoiceRequired FunctionChoiceBehavior = "required"
)

type KernelFunction interface {
	GetName() string
	GetDescription() string
	GetPluginName() string
	Invoke(ctx context.Context, args map[string]interface{}) (interface{}, error)
}

type Kernel interface {
	AddService(client ChatCompletionClient) error
	AddPlugin(name string, plugin Plugin) error
	GetFunction(pluginName, functionName string) (KernelFunction, error)
	GetService(serviceID string) (ChatCompletionClient, error)
}

type Plugin interface {
	GetName() string
	GetDescription() string
	GetFunctions() []KernelFunction
}
