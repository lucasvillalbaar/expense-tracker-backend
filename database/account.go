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

func (repo *PostgresRepository) DeleteAccount(ctx context.Context, accountID string) error {
	query := `DELETE FROM accounts WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query, accountID)
	if err != nil {
		return fmt.Errorf("failed to delete account (DeleteAccount): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected (DeleteAccount): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("account not found")
	}

	return nil
}

func (repo *PostgresRepository) UpdateAccount(ctx context.Context, account *transaction.Account) error {
	query := `UPDATE accounts SET type = $2, name = $3, balance = $4, currency = $5 WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query,
		account.ID,
		account.Type,
		account.Name,
		account.Balance,
		account.Currency,
	)
	if err != nil {
		return fmt.Errorf("failed to update account (UpdateAccount): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected (UpdateAccount): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("account not found")
	}

	return nil
}
