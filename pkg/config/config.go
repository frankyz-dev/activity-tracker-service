package config

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configFile embed.FS

type RootConfig struct {
	DB Config `yaml:"db"`
}

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func LoadConfig() (*Config, error) {
	var rootConfig RootConfig

	// Load configuration from file
	file, err := configFile.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &rootConfig)
	if err != nil {
		return nil, err
	}

	// Override with environment variables if available
	if host, exists := os.LookupEnv("DB_HOST"); exists {
		rootConfig.DB.Host = host
	}
	if port, exists := os.LookupEnv("DB_PORT"); exists {
		rootConfig.DB.Port, err = strconv.Atoi(port)
		if err != nil {
			return nil, fmt.Errorf("invalid DB_PORT: %v", err)
		}
	}
	if user, exists := os.LookupEnv("DB_USER"); exists {
		rootConfig.DB.User = user
	}
	if password, exists := os.LookupEnv("DB_PASSWORD"); exists {
		rootConfig.DB.Password = password
	}
	if dbname, exists := os.LookupEnv("DB_NAME"); exists {
		rootConfig.DB.DBName = dbname
	}

	// Check if the user is unconfigured
	if rootConfig.DB.User == "TBD" {
		return nil, errors.New("Please configure the database credentials in the config.yaml file")
	}

	return &rootConfig.DB, nil
}
