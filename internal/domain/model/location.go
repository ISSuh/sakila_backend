package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (Location) GormDataType() string {
	return "geometry"
}

func (l Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL: "ST_PointFromText(?)",
		Vars: []interface{}{
			fmt.Sprintf("POINT(%f %f)", l.X, l.Y),
		},
	}
}

type JSON json.RawMessage

func (l *Location) Scan(value interface{}) error {
	location, ok := value.(Location)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	l.X = location.X
	l.Y = location.Y
	return nil
}
