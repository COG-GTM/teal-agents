package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadTeamOrchestratorConfig(t *testing.T) {
	config, err := LoadOrchestratorConfig("testdata/team_orchestrator.yaml")
	if err != nil {
		t.Fatalf("LoadOrchestratorConfig failed: %v", err)
	}

	if config.APIVersion != "skagents/v1" {
		t.Errorf("Expected apiVersion 'skagents/v1', got '%s'", config.APIVersion)
	}

	if config.Kind != "TeamOrchestrator" {
		t.Errorf("Expected kind 'TeamOrchestrator', got '%s'", config.Kind)
	}

	if config.ServiceName != "CollaborationOrchestrator" {
		t.Errorf("Expected service_name 'CollaborationOrchestrator', got '%s'", config.ServiceName)
	}

	if config.Version != "0.1" {
		t.Errorf("Expected version '0.1', got '%s'", config.Version)
	}

	spec, ok := config.Spec.(TeamOrchestratorSpec)
	if !ok {
		t.Fatalf("Expected TeamOrchestratorSpec, got %T", config.Spec)
	}

	if spec.MaxRounds != 10 {
		t.Errorf("Expected max_rounds 10, got %d", spec.MaxRounds)
	}

	if spec.ManagerAgent != "TeamManagerAgent:0.1" {
		t.Errorf("Expected manager_agent 'TeamManagerAgent:0.1', got '%s'", spec.ManagerAgent)
	}

	expectedAgents := []string{"WikipediaAgent:0.1", "ArxivSearchAgent:0.1", "AssistantAgent:0.1"}
	if len(spec.Agents) != len(expectedAgents) {
		t.Errorf("Expected %d agents, got %d", len(expectedAgents), len(spec.Agents))
	}

	for i, expected := range expectedAgents {
		if i >= len(spec.Agents) || spec.Agents[i] != expected {
			t.Errorf("Expected agent %d to be '%s', got '%s'", i, expected, spec.Agents[i])
		}
	}
}

func TestLoadPlanningOrchestratorConfig(t *testing.T) {
	config, err := LoadOrchestratorConfig("testdata/planning_orchestrator.yaml")
	if err != nil {
		t.Fatalf("LoadOrchestratorConfig failed: %v", err)
	}

	if config.Kind != "PlanningOrchestrator" {
		t.Errorf("Expected kind 'PlanningOrchestrator', got '%s'", config.Kind)
	}

	spec, ok := config.Spec.(PlanningOrchestratorSpec)
	if !ok {
		t.Fatalf("Expected PlanningOrchestratorSpec, got %T", config.Spec)
	}

	if spec.PlanningAgent != "PlanningAgent:0.1" {
		t.Errorf("Expected planning_agent 'PlanningAgent:0.1', got '%s'", spec.PlanningAgent)
	}

	if spec.HumanInTheLoop != false {
		t.Errorf("Expected human_in_the_loop false, got %t", spec.HumanInTheLoop)
	}

	if spec.HitlTimeout != 0 {
		t.Errorf("Expected hitl_timeout 0, got %d", spec.HitlTimeout)
	}

	expectedAgents := []string{"WikipediaAgent:0.1", "ArxivSearchAgent:0.1", "AssistantAgent:0.1"}
	if len(spec.Agents) != len(expectedAgents) {
		t.Errorf("Expected %d agents, got %d", len(expectedAgents), len(spec.Agents))
	}
}

func TestLoadAssistantOrchestratorConfig(t *testing.T) {
	config, err := LoadOrchestratorConfig("testdata/assistant_orchestrator.yaml")
	if err != nil {
		t.Fatalf("LoadOrchestratorConfig failed: %v", err)
	}

	if config.Kind != "AssistantOrchestrator" {
		t.Errorf("Expected kind 'AssistantOrchestrator', got '%s'", config.Kind)
	}

	if config.ServiceName != "DemoAgentOrchestrator" {
		t.Errorf("Expected service_name 'DemoAgentOrchestrator', got '%s'", config.ServiceName)
	}

	spec, ok := config.Spec.(AssistantOrchestratorSpec)
	if !ok {
		t.Fatalf("Expected AssistantOrchestratorSpec, got %T", config.Spec)
	}

	if spec.FallbackAgent != "GeneralAgent:0.1" {
		t.Errorf("Expected fallback_agent 'GeneralAgent:0.1', got '%s'", spec.FallbackAgent)
	}

	if spec.AgentChooser != "AgentSelectorAgent:0.1" {
		t.Errorf("Expected agent_chooser 'AgentSelectorAgent:0.1', got '%s'", spec.AgentChooser)
	}

	expectedAgents := []string{"MathAgent:0.1", "WeatherAgent:0.1"}
	if len(spec.Agents) != len(expectedAgents) {
		t.Errorf("Expected %d agents, got %d", len(expectedAgents), len(spec.Agents))
	}

	for i, expected := range expectedAgents {
		if i >= len(spec.Agents) || spec.Agents[i] != expected {
			t.Errorf("Expected agent %d to be '%s', got '%s'", i, expected, spec.Agents[i])
		}
	}
}

func TestInvalidYAML(t *testing.T) {
	_, err := LoadOrchestratorConfig("testdata/invalid.yaml")
	if err == nil {
		t.Error("LoadOrchestratorConfig should fail for invalid YAML")
	}
}

func TestMissingFile(t *testing.T) {
	_, err := LoadOrchestratorConfig("testdata/nonexistent.yaml")
	if err == nil {
		t.Error("LoadOrchestratorConfig should fail for missing file")
	}
}

func TestInvalidOrchestratorKind(t *testing.T) {
	tempFile, err := os.CreateTemp("", "invalid_kind_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	yamlContent := `
apiVersion: skagents/v1
kind: UnknownOrchestrator
description: Test config with unknown kind
service_name: TestOrchestrator
version: 0.1
spec:
  test_field: test_value
`

	if _, err := tempFile.WriteString(yamlContent); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}
	tempFile.Close()

	_, err = LoadOrchestratorConfig(tempFile.Name())
	if err == nil {
		t.Error("LoadOrchestratorConfig should fail for unknown orchestrator kind")
	}

	expectedError := "unsupported orchestrator kind: UnknownOrchestrator"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestConfigurationFilesParsing(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		kind     string
	}{
		{"Team", "testdata/team_orchestrator.yaml", "TeamOrchestrator"},
		{"Planning", "testdata/planning_orchestrator.yaml", "PlanningOrchestrator"},
		{"Assistant", "testdata/assistant_orchestrator.yaml", "AssistantOrchestrator"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := LoadOrchestratorConfig(tc.filename)
			if err != nil {
				t.Fatalf("LoadOrchestratorConfig failed for %s: %v", tc.name, err)
			}

			if config.Kind != tc.kind {
				t.Errorf("Expected kind '%s', got '%s'", tc.kind, config.Kind)
			}

			if config.APIVersion != "skagents/v1" {
				t.Errorf("Expected apiVersion 'skagents/v1', got '%s'", config.APIVersion)
			}

			if config.Spec == nil {
				t.Error("Spec should not be nil")
			}
		})
	}
}

func TestWorkingDirectory(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	testdataPath := filepath.Join(wd, "testdata")
	if _, err := os.Stat(testdataPath); os.IsNotExist(err) {
		t.Fatalf("testdata directory does not exist at %s", testdataPath)
	}

	files := []string{
		"team_orchestrator.yaml",
		"planning_orchestrator.yaml",
		"assistant_orchestrator.yaml",
		"invalid.yaml",
	}

	for _, file := range files {
		filePath := filepath.Join(testdataPath, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Test file does not exist: %s", filePath)
		}
	}
}
