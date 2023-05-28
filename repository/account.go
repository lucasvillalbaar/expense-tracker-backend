package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func InsertAccount(ctx context.Context, account *transaction.Account) error {
	return repository.InsertAccount(ctx, account)
}
