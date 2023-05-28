package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

type Repository interface {
	Close()
	InsertTransaction(ctx context.Context, transaction *transaction.Transaction) error
}

var repository Repository

func SetRepository(r Repository) {
	repository = r
}

func Close() {
	repository.Close()
}

func InsertTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	return repository.InsertTransaction(ctx, transaction)
}
