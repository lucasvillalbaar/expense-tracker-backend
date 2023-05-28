package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertCategory(ctx context.Context, category *transaction.Category) error {
	return repository.InsertCategory(ctx, category)
}
