package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Host           string `yaml:"host"`
		Port           string `yaml:"port"`
		MaxHeaderBytes int    `yaml:"maxHeaderBytes"`
		Timeout        struct {
			Read   time.Duration `yaml:"read"`
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}