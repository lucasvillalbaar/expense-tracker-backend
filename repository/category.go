package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertCategory(ctx context.Context, category *transaction.Category) error {
	return repository.InsertCategory(ctx, category)
}
func DeleteCategory(ctx context.Context, categoryID string) error {
	return repository.DeleteCategory(ctx, categoryID)
}

func UpdateCategory(ctx context.Context, category *transaction.Category) error {
	return repository.UpdateCategory(ctx, category)
}
