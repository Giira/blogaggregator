package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func (cfg Config) SetUser(user string) error {
	cfg.Current_user_name = user
	err := cfg.Write()
	if err != nil {
		return err
	}
	return nil
}

func (cfg Config) Write() error {
	address, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	address = address + "/" + configFileName
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(address, jsonData, 0666)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	address, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	address = address + "/" + configFileName

	data, err := os.ReadFile(address)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
