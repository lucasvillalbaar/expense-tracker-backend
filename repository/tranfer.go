package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertTransfer(ctx context.Context, transfer *transaction.Transfer) error {
	return repository.InsertTransfer(ctx, transfer)
}

func DeleteTransfer(ctx context.Context, transferID string) error {
	return repository.DeleteTransfer(ctx, transferID)
}

func UpdateTransfer(ctx context.Context, transfer *transaction.Transfer) error {
	return repository.UpdateTransfer(ctx, transfer)
}
