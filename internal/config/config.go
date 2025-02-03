package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl    string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func Read() (Config, error) {
	var result Config
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		return Config{}, err
	}
	return result, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.Username = username
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/" + configFileName, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err = encoder.Encode(cfg); err != nil {
		return err
	}
	return nil

}
