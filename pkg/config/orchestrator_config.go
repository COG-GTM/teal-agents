package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type OrchestratorConfig struct {
	APIVersion  string      `yaml:"apiVersion"`
	Kind        string      `yaml:"kind"`
	Description string      `yaml:"description"`
	ServiceName string      `yaml:"service_name"`
	Version     string      `yaml:"version"`
	Spec        interface{} `yaml:"spec"`
}

type TeamOrchestratorSpec struct {
	MaxRounds    int      `yaml:"max_rounds"`
	ManagerAgent string   `yaml:"manager_agent"`
	Agents       []string `yaml:"agents"`
}

type PlanningOrchestratorSpec struct {
	PlanningAgent  string   `yaml:"planning_agent"`
	Agents         []string `yaml:"agents"`
	HumanInTheLoop bool     `yaml:"human_in_the_loop"`
	HitlTimeout    int      `yaml:"hitl_timeout"`
}

type AssistantOrchestratorSpec struct {
	FallbackAgent string   `yaml:"fallback_agent"`
	AgentChooser  string   `yaml:"agent_chooser"`
	Agents        []string `yaml:"agents"`
}

func LoadOrchestratorConfig(filename string) (*OrchestratorConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read orchestrator config file: %w", err)
	}

	var config OrchestratorConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse orchestrator YAML: %w", err)
	}

	switch config.Kind {
	case "TeamOrchestrator":
		var spec TeamOrchestratorSpec
		specData, err := yaml.Marshal(config.Spec)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal spec for TeamOrchestrator: %w", err)
		}
		if err := yaml.Unmarshal(specData, &spec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal spec for TeamOrchestrator: %w", err)
		}
		config.Spec = spec

	case "PlanningOrchestrator":
		var spec PlanningOrchestratorSpec
		specData, err := yaml.Marshal(config.Spec)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal spec for PlanningOrchestrator: %w", err)
		}
		if err := yaml.Unmarshal(specData, &spec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal spec for PlanningOrchestrator: %w", err)
		}
		config.Spec = spec

	case "AssistantOrchestrator":
		var spec AssistantOrchestratorSpec
		specData, err := yaml.Marshal(config.Spec)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal spec for AssistantOrchestrator: %w", err)
		}
		if err := yaml.Unmarshal(specData, &spec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal spec for AssistantOrchestrator: %w", err)
		}
		config.Spec = spec

	default:
		return nil, fmt.Errorf("unsupported orchestrator kind: %s", config.Kind)
	}

	return &config, nil
}
