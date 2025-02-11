package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, nil
	}

	var config Config
	if err = json.Unmarshal(data, &config); err != nil {
		return Config{}, nil
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/" + configFileName, nil
}

func Write(cfg Config) error {
	filename, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
