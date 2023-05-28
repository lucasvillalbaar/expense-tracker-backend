package database

import (
	"context"
	"fmt"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func (repo *PostgresRepository) InsertAccount(ctx context.Context, account *transaction.Account) error {
	query := `INSERT INTO accounts (id, type, name, balance, currency) VALUES ($1, $2, $3, $4, $5)`

	_, err := repo.db.ExecContext(ctx, query,
		account.ID,
		account.Type,
		account.Name,
		account.Balance,
		account.Currency,
	)
	if err != nil {
		return fmt.Errorf("failed to insert account (InsertAccount): %v", err)
	}

	return nil
}
