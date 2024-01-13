package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/golang/mock/mockgen/model"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

// SQLStore defines all functions to execute SQL queries and transactions
// Embedding the Queries struct inside a SQLStore struct is called composition and is prefered over inheritance
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
// func NewStore(db *sql.DB) *Store {
func NewStore(db *sql.DB) Store {
	// return &Store{
	return &SQLStore{
		db: db,
		// The New function creates and returns a Queries object
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// tx, err := store.db.BeginTx(ctx, &sql.TxOptions{} )
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}



