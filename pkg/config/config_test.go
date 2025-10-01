package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Error("GetInstance should return the same singleton instance")
	}
}

func TestAddConfig(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	testConfig := ConfigItem{
		EnvName:      "TEST_VAR",
		IsRequired:   false,
		DefaultValue: "test_default",
	}

	os.Setenv("TEST_VAR", "test_value")
	defer os.Unsetenv("TEST_VAR")

	err := config.AddConfig(testConfig)
	if err != nil {
		t.Fatalf("AddConfig failed: %v", err)
	}

	if len(config.configs) != 1 {
		t.Errorf("Expected 1 config, got %d", len(config.configs))
	}

	if config.Get("TEST_VAR") != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", config.Get("TEST_VAR"))
	}
}

func TestAddConfigs(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	testConfigs := []ConfigItem{
		{EnvName: "TEST_VAR1", IsRequired: false, DefaultValue: "default1"},
		{EnvName: "TEST_VAR2", IsRequired: false, DefaultValue: "default2"},
	}

	os.Setenv("TEST_VAR1", "value1")
	os.Setenv("TEST_VAR2", "value2")
	defer func() {
		os.Unsetenv("TEST_VAR1")
		os.Unsetenv("TEST_VAR2")
	}()

	err := config.AddConfigs(testConfigs)
	if err != nil {
		t.Fatalf("AddConfigs failed: %v", err)
	}

	if len(config.configs) != 2 {
		t.Errorf("Expected 2 configs, got %d", len(config.configs))
	}

	if config.Get("TEST_VAR1") != "value1" {
		t.Errorf("Expected 'value1', got '%s'", config.Get("TEST_VAR1"))
	}
}

func TestLoadFromEnvironmentWithDefaults(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	testConfig := ConfigItem{
		EnvName:      "UNSET_VAR",
		IsRequired:   false,
		DefaultValue: "default_value",
	}

	os.Unsetenv("UNSET_VAR")

	err := config.AddConfig(testConfig)
	if err != nil {
		t.Fatalf("AddConfig failed: %v", err)
	}

	if config.Get("UNSET_VAR") != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", config.Get("UNSET_VAR"))
	}
}

func TestLoadFromEnvironmentWithJSONStore(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	testConfig := ConfigItem{
		EnvName:      "JSON_TEST_VAR",
		IsRequired:   false,
		DefaultValue: "",
	}

	envStore := map[string]string{
		"JSON_TEST_VAR": "json_value",
	}

	jsonData, _ := json.Marshal(envStore)
	os.Setenv("TA_ENV_STORE", string(jsonData))
	defer func() {
		os.Unsetenv("TA_ENV_STORE")
		os.Unsetenv("JSON_TEST_VAR")
	}()

	err := config.AddConfig(testConfig)
	if err != nil {
		t.Fatalf("AddConfig failed: %v", err)
	}

	if config.Get("JSON_TEST_VAR") != "json_value" {
		t.Errorf("Expected 'json_value', got '%s'", config.Get("JSON_TEST_VAR"))
	}
}

func TestLoadFromEnvironmentWithGlobalJSONStore(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	testConfig := ConfigItem{
		EnvName:      "GLOBAL_JSON_TEST_VAR",
		IsRequired:   false,
		DefaultValue: "",
	}

	globalEnvStore := map[string]string{
		"GLOBAL_JSON_TEST_VAR": "global_json_value",
	}

	jsonData, _ := json.Marshal(globalEnvStore)
	os.Setenv("TA_ENV_GLOBAL_STORE", string(jsonData))
	defer func() {
		os.Unsetenv("TA_ENV_GLOBAL_STORE")
		os.Unsetenv("GLOBAL_JSON_TEST_VAR")
	}()

	err := config.AddConfig(testConfig)
	if err != nil {
		t.Fatalf("AddConfig failed: %v", err)
	}

	if config.Get("GLOBAL_JSON_TEST_VAR") != "global_json_value" {
		t.Errorf("Expected 'global_json_value', got '%s'", config.Get("GLOBAL_JSON_TEST_VAR"))
	}
}

func TestValidateRequiredKeys(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	testConfig := ConfigItem{
		EnvName:      "REQUIRED_VAR",
		IsRequired:   true,
		DefaultValue: "",
	}

	os.Setenv("REQUIRED_VAR", "required_value")
	defer os.Unsetenv("REQUIRED_VAR")

	err := config.AddConfig(testConfig)
	if err != nil {
		t.Fatalf("AddConfig should succeed with required value set: %v", err)
	}
}

func TestValidateRequiredKeysFails(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	testConfig := ConfigItem{
		EnvName:      "MISSING_REQUIRED_VAR",
		IsRequired:   true,
		DefaultValue: "",
	}

	os.Unsetenv("MISSING_REQUIRED_VAR")

	err := config.AddConfig(testConfig)
	if err == nil {
		t.Error("AddConfig should fail when required variable is missing")
	}

	expectedError := "missing required configuration key: MISSING_REQUIRED_VAR"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestGet(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	config.Set("TEST_KEY", "test_value")

	value := config.Get("TEST_KEY")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	missingValue := config.Get("MISSING_KEY")
	if missingValue != "" {
		t.Errorf("Expected empty string for missing key, got '%s'", missingValue)
	}
}

func TestSet(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	config.Set("SET_TEST_KEY", "set_value")

	value := config.Get("SET_TEST_KEY")
	if value != "set_value" {
		t.Errorf("Expected 'set_value', got '%s'", value)
	}
}

func TestHas(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	config.Set("EXISTS_KEY", "some_value")

	if !config.Has("EXISTS_KEY") {
		t.Error("Has should return true for existing key")
	}

	if config.Has("MISSING_KEY") {
		t.Error("Has should return false for missing key")
	}
}

func TestConcurrentAccess(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	var wg sync.WaitGroup
	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("CONCURRENT_KEY_%d", index)
			value := fmt.Sprintf("value_%d", index)

			config.Set(key, value)
			retrieved := config.Get(key)

			if retrieved != value {
				t.Errorf("Concurrent access failed: expected %s, got %s", value, retrieved)
			}
		}(i)
	}

	wg.Wait()
}

func TestLoadFromYAML(t *testing.T) {
	config := &AppConfig{
		values:  make(map[string]string),
		configs: make([]ConfigItem, 0),
	}

	tempFile, err := os.CreateTemp("", "test_config_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	yamlContent := `
test_key: test_value
nested:
  key: nested_value
`

	if _, err := tempFile.WriteString(yamlContent); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}
	tempFile.Close()

	err = config.LoadFromYAML(tempFile.Name())
	if err != nil {
		t.Fatalf("LoadFromYAML failed: %v", err)
	}

	if config.Get("test_key") != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", config.Get("test_key"))
	}
}
