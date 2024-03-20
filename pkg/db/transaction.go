package db

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type Transaction struct {
	db *Database
}

func NewTransaction(d *Database) *Transaction {
	return &Transaction{
		db: d,
	}
}

func (t *Transaction) DoTransaction(task func() error, opts ...*sql.TxOptions) error {
	if err := t.db.engine.Transaction(
		func(tx *gorm.DB) error {
			return task()
		}, opts...); err != nil {
		return errors.New("Transaction error")
	}
	return nil
}
