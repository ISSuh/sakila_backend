package config

import (
    "errors"
    "gopkg.in/yaml.v3"
    "os"
)

type Config struct {
}

func NewConfig(path string) (Config, error) {
    if len(path) == 0 {
        return Config{}, errors.New("can not found config file")
    }

    var buffer []byte
    var err error
    buffer, err = loadFile(path)
    if err != nil {
        return Config{}, err
    }

    var config Config
    if err = yaml.Unmarshal(buffer, config); err != nil {
        return Config{}, nil
    }
    return config, nil
}

func loadFile(path string) (buffer []byte, err error) {
    if buffer, err = os.ReadFile(path); err != nil {
        return nil, err
    }
    return buffer, nil
}
