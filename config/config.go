package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func LoadYamlConfigByFsPath[T any](name string, config T) (T, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(b, config)
	if err != nil {
		return config, err
	}

	return config, nil
}
