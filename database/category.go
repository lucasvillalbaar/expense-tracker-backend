package database

import (
	"context"
	"fmt"

	transaction "github.com/lucasvillalbaar/expense-tracker-backend/pkg/transactions"
)

func (repo *PostgresRepository) InsertCategory(ctx context.Context, category *transaction.Category) error {
	query := `INSERT INTO categories (id, category_name, subcategory_name) VALUES ($1, $2, $3)`

	_, err := repo.db.ExecContext(ctx, query,
		category.ID,
		category.CategoryName,
		category.SubcategoryName,
	)
	if err != nil {
		return fmt.Errorf("failed to insert category (InsertCategory): %v", err)
	}

	return nil
}
