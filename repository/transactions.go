package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	return repository.InsertTransaction(ctx, transaction)
}
