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

// Generate a new config file
func Gen(path string) error {
	// Create file
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// Prettify Json
	j, err := json.MarshalIndent(Config{
		Host:     "localhost",
		Port:     25575,
		Password: "password",
	}, "", "  ")
	if err != nil {
		return err
	}

	// Write pretty Json to file
	_, err = file.Write(j)
	if err != nil {
		return err
	}

	return nil
}

// Load a config file into a struct
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
