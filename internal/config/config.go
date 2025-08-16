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

func (cfg Config) SetUser(user string) {
	cfg.Current_user_name = user
}

func Read() (Config, error) {
	address, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	address = address + configFileName

	cfg := Config{}
	err = json.Unmarshal(address, &cfg)
	if err != nil {
		return Config{}, err
	}
	return
}
