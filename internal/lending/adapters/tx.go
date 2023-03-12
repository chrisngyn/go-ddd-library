package adapters

import (
	"context"
	"database/sql"
	"fmt"
)

func WithTx(ctx context.Context, db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rErr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
