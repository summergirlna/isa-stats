package postgresql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"isa-stats/pkg/repository"
)

type TransactionManager struct {
	db *sqlx.DB
}

var txKey = struct{}{}

func (t TransactionManager) Execute(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error) {
	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, &txKey, tx)

	v, err := f(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return v, nil
}

func NewTransactionManager(db *sqlx.DB) repository.Transaction {
	return &TransactionManager{db: db}
}

func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}
