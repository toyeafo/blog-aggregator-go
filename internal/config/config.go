package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db_url    string `json:"db_url"`
	User_name string `json:"current_user_name"`
}

func Read() Config {
	read_file, err := getConfigFilePath()
	if err != nil {
		return Config{}
	}

	data, err := os.ReadFile(read_file)
	if err != nil {
		return Config{}
	}

	fileConfig := Config{}
	data_err := json.Unmarshal(data, &fileConfig)
	if data_err != nil {
		return Config{}
	}

	return fileConfig
}

func (c *Config) SetUser(username string) {
	c.User_name = username
	write(c)
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := homePath + "/" + configFileName
	return filePath, nil
}

func write(cfg *Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	readFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(readFile, dat, 0666)
	if err != nil {
		return err
	}

	return nil
}
