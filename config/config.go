package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host     string
	Port     int
	Password string
}

func Load(path string) (error, *Config) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return err, &Config{}
	}

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return err, &Config{}
	}

	return nil, config
}
