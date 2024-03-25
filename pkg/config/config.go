package config

import (
	"errors"
	"os"

	"github.com/ISSuh/sakila_backend/pkg/db"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	LogLevel string       `yaml:"log"`
	Server   ServerConfig `yaml:"app"`
	Database db.Config    `yaml:"db"`
}

func NewConfig(path string) (*Config, error) {
	if len(path) == 0 {
		return nil, errors.New("can not found config file")
	}

	buffer, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	if err = yaml.Unmarshal(buffer, config); err != nil {
		return nil, nil
	}
	return config, nil
}

func loadFile(path string) (buffer []byte, err error) {
	if buffer, err = os.ReadFile(path); err != nil {
		return nil, err
	}
	return buffer, nil
}
