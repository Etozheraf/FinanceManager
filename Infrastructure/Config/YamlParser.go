package Config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Configuration struct {
	Database DataBaseConfig `yaml:"database"`
	Telegram TelegramConfig `yaml:"telegram"`
}

func LoadConfigurations(filePath string) (*Configuration, error) {
	var config Configuration

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read yaml file: %s", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml data: %s", err)
	}

	return &config, nil
}

var Config = &Configuration{}

func init() {
	configPath := "source/application.yaml"
	var err error
	Config, err = LoadConfigurations(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Database Host: %s\n", Config.Database.Host)
	fmt.Printf("Telegram Secret: %s\n", Config.Telegram.Secret)
}
