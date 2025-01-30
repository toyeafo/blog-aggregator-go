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

func (c *Config) SetUser(username string) error {
	c.User_name = username
	return write(*c)
}

func Read() (Config, error) {
	read_file, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.Open(read_file)
	if err != nil {
		return Config{}, err
	}
	defer data.Close()

	decoder := json.NewDecoder(data)
	fileConfig := Config{}
	err = decoder.Decode(&fileConfig)
	if err != nil {
		return Config{}, err
	}

	return fileConfig, nil
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
	readFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(readFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
