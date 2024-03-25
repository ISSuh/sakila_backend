package db

import "github.com/ISSuh/sakila_backend/pkg/db/database"

type ConnectionPoolConfig struct {
	IdleConnection int `yaml:"idle_connection"`
	OpenConnection int `yaml:"open_connection"`
}

type ORMConfig struct {
	LogLevel       string               `yaml:"loglevel"`
	ConnectionPool ConnectionPoolConfig `yaml:"connection_pool"`
}

type Config struct {
	Database database.Config `yaml:"engine"`
	ORM      ORMConfig       `yaml:"orm"`
}
