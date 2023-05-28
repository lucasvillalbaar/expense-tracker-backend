package repository

import (
	"context"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

type Repository interface {
	Close()
	// Transactions
	InsertTransaction(ctx context.Context, transaction *transaction.Transaction) error
	DeleteTransaction(ctx context.Context, transactionID string) error
	// Currencies
	InsertCurrency(ctx context.Context, currency *transaction.Currency) error
	// Accounts
	InsertAccount(ctx context.Context, account *transaction.Account) error
	// Categories
	InsertCategory(ctx context.Context, category *transaction.Category) error
}

var repository Repository

func SetRepository(r Repository) {
	repository = r
}

func Close() {
	repository.Close()
}
