package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertCurrency(ctx context.Context, currency *transaction.Currency) error {
	return repository.InsertCurrency(ctx, currency)
}
