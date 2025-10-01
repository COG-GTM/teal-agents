package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type ConfigItem struct {
	EnvName      string
	IsRequired   bool
	DefaultValue string
}

type AppConfig struct {
	mu      sync.RWMutex
	values  map[string]string
	configs []ConfigItem
}

var (
	instance *AppConfig
	once     sync.Once
)

func GetInstance() *AppConfig {
	once.Do(func() {
		instance = &AppConfig{
			values:  make(map[string]string),
			configs: make([]ConfigItem, 0),
		}
	})
	return instance
}

func (c *AppConfig) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.values[key]
}

func (c *AppConfig) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

func (c *AppConfig) Has(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.values[key]
	return exists
}

func (c *AppConfig) AddConfig(config ConfigItem) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.configs = append(c.configs, config)
	return c.loadFromEnvironment()
}

func (c *AppConfig) AddConfigs(configs []ConfigItem) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.configs = append(c.configs, configs...)
	return c.loadFromEnvironment()
}

func (c *AppConfig) loadFromEnvironment() error {
	if err := c.parseJSONStores(); err != nil {
		return fmt.Errorf("failed to parse JSON stores: %w", err)
	}

	for _, config := range c.configs {
		value := os.Getenv(config.EnvName)
		if value == "" && config.DefaultValue != "" {
			value = config.DefaultValue
		}
		c.values[config.EnvName] = value
	}

	return c.validateRequiredKeys()
}

func (c *AppConfig) parseJSONStores() error {
	stores := []string{"TA_ENV_STORE", "TA_ENV_GLOBAL_STORE"}

	for _, storeName := range stores {
		storeData := os.Getenv(storeName)
		if storeData == "" {
			continue
		}

		var envVars map[string]string
		if err := json.Unmarshal([]byte(storeData), &envVars); err != nil {
			return fmt.Errorf("failed to parse %s JSON: %w", storeName, err)
		}

		for key, value := range envVars {
			if err := os.Setenv(key, value); err != nil {
				return fmt.Errorf("failed to set environment variable %s: %w", key, err)
			}
		}
	}

	return nil
}

func (c *AppConfig) validateRequiredKeys() error {
	for _, config := range c.configs {
		if config.IsRequired {
			value := c.values[config.EnvName]
			if value == "" {
				return fmt.Errorf("missing required configuration key: %s", config.EnvName)
			}
		}
	}
	return nil
}

func (c *AppConfig) LoadFromYAML(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var config map[string]interface{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range config {
		c.values[key] = fmt.Sprintf("%v", value)
	}

	return nil
}
