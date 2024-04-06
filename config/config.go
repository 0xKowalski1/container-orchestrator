package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Namespace string `json:"namespace"` // Production, Development or Test
}

func LoadConfig(configFile string) (*Config, error) {
	// Read the file
	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	// Parse the file
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}