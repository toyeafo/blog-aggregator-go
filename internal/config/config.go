package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url    string `json:"db_url"`
	User_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	read_file, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(read_file)
	if err != nil {
		return Config{}, err
	}

	fileConfig := Config{}
	data_err := json.Unmarshal(data, &fileConfig)
	if data_err != nil {
		return Config{}, data_err
	}

	return fileConfig, nil
}

func (c *Config) SetUser(username string) error {
	c.User_name = username
	return write(*c)
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := homePath + "/" + configFileName
	return filePath, nil
}

func write(cfg Config) error {
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
