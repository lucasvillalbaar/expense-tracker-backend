package database

import (
	"context"
	"fmt"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func (repo *PostgresRepository) InsertCurrency(ctx context.Context, currency *transaction.Currency) error {
	query := `INSERT INTO currencies (id, name, symbol) VALUES ($1, $2, $3)`

	_, err := repo.db.ExecContext(ctx, query,
		currency.ID,
		currency.Name,
		currency.Symbol,
	)
	if err != nil {
		return fmt.Errorf("failed to insert currency (InsertCurrency): %v", err)
	}

	return nil
}
