package db

import (
	"github.com/ISSuh/monolith-sakila/internal/common"
	"gorm.io/gorm"
)

func Pagnation(e *gorm.DB, page *common.Pagnation) func(db *gorm.DB) *gorm.DB {
	return func(e *gorm.DB) *gorm.DB {
		if len(page.Sort) == 0 || len(page.SortBy) == 0 {
			return e.Offset(page.CalculateOffset()).Limit(page.Limit)
		}
		return e.Offset(page.CalculateOffset()).Limit(page.Limit).Order(page.SortSQL())
	}
}
