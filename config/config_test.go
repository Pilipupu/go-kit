package config

import (
	"github.com/Pilipupu/go-kit/log"
	"testing"
)

var logger = log.New()

func TestYamlConfig(t *testing.T) {
	c, err := LoadYamlConfigByFsPath[*TestConfig]("test.yaml", &TestConfig{})
	if err != nil {
		logger.Error(err, "load yaml fail")
	}

	logger.Info("Yaml unmarshal", "name", c.Container[0].Name)
}

type Container struct {
	Command string `yaml:"command"`
	Name    string `yaml:"name"`
}

type TestConfig struct {
	ApiVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Container  []Container `yaml:"containers"`
}
