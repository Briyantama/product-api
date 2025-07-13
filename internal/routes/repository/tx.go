package repository

import (
	"context"
	"test-case-vhiweb/internal/constants"

	"gorm.io/gorm"
)

type txKey struct{}

type WithTx struct {
	conn *gorm.DB
}

func NewTxRepository(conn *gorm.DB) WithTx {
	return WithTx{conn: conn}
}

func ExtractTx(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if ok {
		return tx
	}

	return db
}

func AddForUpdate(ctx context.Context, query string) string {
	if _, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return query + " FOR UPDATE"
	}
	
	return query
}

func (wt *WithTx) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := wt.conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	ctxWithTx := context.WithValue(ctx, txKey{}, tx)

	if err := fn(ctxWithTx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return constants.New(
			constants.ERRCONFLICT,
			constants.ErrTransactionFail,
		)
	}

	return nil
}
