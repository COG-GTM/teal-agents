package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	mu     sync.RWMutex
	values map[string]string
}

var (
	instance *AppConfig
	once     sync.Once
)

func GetInstance() *AppConfig {
	once.Do(func() {
		instance = &AppConfig{
			values: make(map[string]string),
		}
		instance.loadFromEnvironment()
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

func (c *AppConfig) loadFromEnvironment() {
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

	return nil
}
