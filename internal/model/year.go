package model

import (
	"database/sql/driver"
)

type Year struct {
	year int64
}

func (Year) GormDataType() string {
	return "year"
}

func (y *Year) Scan(value any) error {
	y.year = value.(int64)
	return nil
}

func (y Year) Value() (driver.Value, error) {
	return y.year, nil
}
