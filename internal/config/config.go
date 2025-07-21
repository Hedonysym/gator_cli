package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	var cfg Config
	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}
	file, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(jsonData, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(u string) error {
	cfg.Current_user_name = u
	return write(*cfg)
}
