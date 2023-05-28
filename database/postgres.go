package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

func (repo *PostgresRepository) InsertTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	query := `INSERT INTO transactions (id, created_at, type, category, description, account, original_amount, currency, base_amount)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := repo.db.ExecContext(ctx, query,
		transaction.ID,
		transaction.CreatedAt,
		transaction.Type,
		transaction.Category,
		transaction.Description,
		transaction.Account,
		transaction.OriginalAmount,
		transaction.Currency,
		transaction.BaseAmount,
	)
	if err != nil {
		return fmt.Errorf("failed to insert expense transaction (InsertTransaction): %v", err)
	}

	return nil
}

func (repo *PostgresRepository) DeleteTransaction(ctx context.Context, transactionID string) error {
	query := `DELETE FROM transactions WHERE id = $1`

	_, err := repo.db.ExecContext(ctx, query, transactionID)
	if err != nil {
		return fmt.Errorf("failed to delete transaction (DeleteTransaction): %v", err)
	}

	return nil
}

func (repo *PostgresRepository) UpdateTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	query := `UPDATE transactions SET created_at = $2, type = $3, category = $4, description = $5, account = $6, original_amount = $7, currency = $8, base_amount = $9 WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query,
		transaction.ID,
		transaction.CreatedAt,
		transaction.Type,
		transaction.Category,
		transaction.Description,
		transaction.Account,
		transaction.OriginalAmount,
		transaction.Currency,
		transaction.BaseAmount,
	)
	if err != nil {
		return fmt.Errorf("failed to update transaction (UpdateTransaction): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected (UpdateTransaction): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("transaction not found")
	}

	return nil
}
