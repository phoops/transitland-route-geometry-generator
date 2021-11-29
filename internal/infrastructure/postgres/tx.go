package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TransactionError struct {
	RollbackNeeded bool
	Error          error
}

type TxFn func(ctx context.Context, tx *sqlx.Tx) *TransactionError

func WithTransaction(logger *zap.SugaredLogger, db *sqlx.DB, ctx context.Context, fn TxFn) (err error) {
	var txResult *TransactionError
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = tx.Rollback()
			logger.Error("transaction panicked", zap.Error(err))
			panic(p)
		} else if err != nil && txResult.RollbackNeeded {
			// something went wrong, rollback
			_ = tx.Rollback()
			logger.Error("transaction errored, rollbacked", zap.Error(err))
		} else {
			// all good, commit
			commitErr := tx.Commit()
			if commitErr == nil {
				logger.Debug("transaction commited")
			} else {
				logger.Error("error during transaction commit", zap.Error(err))
			}
		}
	}()

	txResult = fn(ctx, tx)
	if txResult != nil {
		err = txResult.Error
	}
	return err
}
