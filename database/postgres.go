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
