package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

type Repository interface {
	Close()
	InsertTransaction(ctx context.Context, transaction *transaction.Transaction) error
	InsertCurrency(ctx context.Context, currency *transaction.Currency) error
}

var repository Repository

func SetRepository(r Repository) {
	repository = r
}

func Close() {
	repository.Close()
}
