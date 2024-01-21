package db

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

// SQLStore defines all functions to execute SQL queries and transactions
// Embedding the Queries struct inside a SQLStore struct is called composition and is prefered over inheritance
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
// func NewStore(db *sql.DB) *Store {
func NewStore(connPool *pgxpool.Pool) Store {
	// return &Store{
	return &SQLStore{
		connPool: connPool,
		// The New function creates and returns a Queries object
		Queries: New(connPool),
	}
}
