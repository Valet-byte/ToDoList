package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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
		Database struct {
			Port     string `yaml:"port"`
			Host     string `yaml:"host"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			DbName   string `yaml:"dbName"`
			SslMode  string `yaml:"sslMode"`
		} `yaml:"database"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`

	Jwt struct {
		ExpiresAt int
		JwtKye    string
	} `yaml:"jwt"`
	PasswordSalt string
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

	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	config.Server.Database.Password = os.Getenv("DB_PASSWORD")
	config.Jwt.JwtKye = os.Getenv("JWT_KEY")
	config.PasswordSalt = os.Getenv("PASSWORD_SALT")

	logrusInit(config)

	return config, nil
}

func logrusInit(config *Config) {
	level, err := logrus.ParseLevel(config.Log.Level)

	if err != nil {
		logrus.Fatalln("Fatal! Failed set log level! err : %s", err)
	}

	logrus.SetLevel(level)
}
