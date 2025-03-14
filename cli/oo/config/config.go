package config

import (
	"fmt"

	"os"
	"path/filepath"

	types "github.com/openorch/openorch/cli/oo/types"

	"gopkg.in/yaml.v2"
)

func LoadConfig() (types.Config, error) {
	var config types.Config
	configDir := filepath.Join(os.Getenv("HOME"), ".openorch")
	configPath := filepath.Join(configDir, "cliConfig.yaml")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			return config, fmt.Errorf(
				"failed to create config directory: %v",
				err,
			)
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, err := os.Create(configPath)
		if err != nil {
			return config, fmt.Errorf("failed to create config file: %v", err)
		}
		file.Close()
	}

	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return config, fmt.Errorf("failed to stat config file: %v", err)
	}
	if fileInfo.Size() == 0 {
		return config, nil
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("failed to decode config file: %v", err)
	}

	return config, nil
}

func SaveConfig(config types.Config) error {
	configPath := filepath.Join(
		os.Getenv("HOME"),
		".openorch",
		"cliConfig.yaml",
	)

	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %v", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config file: %v", err)
	}

	return nil
}
