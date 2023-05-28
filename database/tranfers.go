package database

import (
	"context"
	"fmt"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func (repo *PostgresRepository) InsertTransfer(ctx context.Context, transfer *transaction.Transfer) error {
	query := `INSERT INTO transfers (id, created_at, source_account, source_amount, source_fee, destination_account, destination_amount, destination_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repo.db.ExecContext(ctx, query,
		transfer.ID,
		transfer.CreatedAt,
		transfer.SourceAccount,
		transfer.SourceAmount,
		transfer.SourceFee,
		transfer.DestinationAccount,
		transfer.DestinationAmount,
		transfer.DestinationFee,
	)
	if err != nil {
		return fmt.Errorf("failed to insert transfer (InsertTransfer): %v", err)
	}

	return nil
}

func (repo *PostgresRepository) DeleteTransfer(ctx context.Context, transferID string) error {
	query := `DELETE FROM transfers WHERE id = $1`

	_, err := repo.db.ExecContext(ctx, query, transferID)
	if err != nil {
		return fmt.Errorf("failed to delete transfer (DeleteTransfer): %v", err)
	}

	return nil
}

func (repo *PostgresRepository) UpdateTransfer(ctx context.Context, transfer *transaction.Transfer) error {
	query := `UPDATE transfers SET created_at = $2, source_account = $3, source_amount = $4, source_fee = $5, destination_account = $6, destination_amount = $7, destination_fee = $8 WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query,
		transfer.ID,
		transfer.CreatedAt,
		transfer.SourceAccount,
		transfer.SourceAmount,
		transfer.SourceFee,
		transfer.DestinationAccount,
		transfer.DestinationAmount,
		transfer.DestinationFee,
	)
	if err != nil {
		return fmt.Errorf("failed to update transfer (UpdateTransfer): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected (UpdateTransfer): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("transfer not found")
	}

	return nil
}
