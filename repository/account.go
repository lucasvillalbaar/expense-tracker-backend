package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertAccount(ctx context.Context, account *transaction.Account) error {
	return repository.InsertAccount(ctx, account)
}

func DeleteAccount(ctx context.Context, accountID string) error {
	return repository.DeleteAccount(ctx, accountID)
}

func UpdateAccount(ctx context.Context, account *transaction.Account) error {
	return repository.UpdateAccount(ctx, account)
}
