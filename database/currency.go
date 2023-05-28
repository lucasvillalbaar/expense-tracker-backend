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

func (repo *PostgresRepository) DeleteCurrency(ctx context.Context, currencyID string) error {
	query := `DELETE FROM currencies WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query, currencyID)
	if err != nil {
		return fmt.Errorf("failed to delete currency (DeleteCurrencyByID): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get number of rows affected (DeleteCurrencyByID): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("currency not found with ID: %s", currencyID)
	}

	return nil
}

func (repo *PostgresRepository) UpdateCurrency(ctx context.Context, currency *transaction.Currency) error {
	query := `UPDATE currencies SET name = $2, symbol = $3 WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query, currency.ID, currency.Name, currency.Symbol)
	if err != nil {
		return fmt.Errorf("failed to update currency (UpdateCurrencyByID): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get number of rows affected (UpdateCurrencyByID): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("currency not found with ID: %s", currency.ID)
	}

	return nil
}
