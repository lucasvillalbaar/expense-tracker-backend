package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertCurrency(ctx context.Context, currency *transaction.Currency) error {
	return repository.InsertCurrency(ctx, currency)
}

func DeleteCurrency(ctx context.Context, currencyID string) error {
	return repository.DeleteCurrency(ctx, currencyID)
}

func UpdateCurrency(ctx context.Context, currency *transaction.Currency) error {
	return repository.UpdateCurrency(ctx, currency)
}
