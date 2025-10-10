package plugins

import (
	"context"

	"github.com/COG-GTM/teal-agents/pkg/kernel"
)

type Plugin interface {
	GetName() string
	GetDescription() string
	GetFunctions() []kernel.KernelFunction
}

type BasePlugin struct {
	Name          string
	Description   string
	Authorization string
	ExtraData     ExtraDataCollector
}

type ExtraDataElement struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ExtraDataCollector interface {
	AddExtraData(key, value string)
	GetExtraData() []ExtraDataElement
	IsEmpty() bool
}

type PluginFunction struct {
	Name        string
	Description string
	PluginName  string
	Parameters  map[string]ParameterSpec
	Handler     func(context.Context, map[string]interface{}) (interface{}, error)
}

type ParameterSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
}

func (f *PluginFunction) Invoke(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	return f.Handler(ctx, args)
}

func (f *PluginFunction) GetName() string        { return f.Name }
func (f *PluginFunction) GetDescription() string { return f.Description }
func (f *PluginFunction) GetPluginName() string  { return f.PluginName }

type PluginLoader interface {
	RegisterPlugin(name string, plugin Plugin) error
	GetPlugin(name string) (Plugin, error)
	ListPlugins() []string
}
