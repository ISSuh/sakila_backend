package db

import (
	"errors"
	"time"

	"github.com/ISSuh/monolith-sakila/pkg/db/database"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	logLevelInfo  = "info"
	logLevelError = "error"
	logLevelWarn  = "warn"

	driverMysql = "mysql"
)

func SetLogLevel(level string) logger.LogLevel {
	switch level {
	case logLevelInfo:
		return logger.Info
	case logLevelWarn:
		return logger.Warn
	case logLevelError:
		return logger.Error
	default:
		return logger.Silent
	}
}

func DatabaseDriver(c Config) (gorm.Dialector, error) {
	switch c.Database.Driver {
	case driverMysql:
		return database.NewMysqlDialector(c.Database), nil
	default:
		return nil, errors.New("no support driver")
	}
}

type Database struct {
	config Config
	engine *gorm.DB
}

func NewDatabase(config Config) *Database {
	d := &Database{
		config: config,
	}
	return d
}

func (d *Database) Connect() error {
	dialector, err := DatabaseDriver(d.config)
	if err != nil {
		return err
	}

	config := gorm.Config{
		Logger: logger.Default.LogMode(SetLogLevel(d.config.ORM.LogLevel)),
	}

	db, err := gorm.Open(dialector, &config)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(d.config.ORM.ConnectionPool.IdleConnection)
	sqlDB.SetMaxOpenConns(d.config.ORM.ConnectionPool.OpenConnection)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	d.engine = db
	return nil
}

func (d *Database) Engine() *gorm.DB {
	return d.engine
}
