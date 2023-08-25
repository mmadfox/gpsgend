package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Service string `yaml:"service"`

	Logger struct {
		Format string `yaml:"format"`
		Level  string `yaml:"level"`
	} `yaml:"logger"`

	API struct {
		HTTP struct {
			Addr string `yml:"addr"`
		} `yaml:"http"`
	} `yaml:"api"`

	Storage struct {
		Mongodb struct {
			URI            string `yml:"uri"`
			CollectionName string `yaml:"collectionName"`
			DatabaseName   string `yaml:"databaseName"`
		} `yaml:"mongodb"`
	} `yaml:"storage"`
}

func FromFile(filename string) (*Config, error) {
	conf := new(Config)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read gpsgend config: %w", err)
	}
	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, err
	}
	return conf, err
}
