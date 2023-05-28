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

func (repo *PostgresRepository) DeleteCategory(ctx context.Context, categoryID string) error {
	query := `DELETE FROM categories WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query, categoryID)
	if err != nil {
		return fmt.Errorf("failed to delete category (DeleteCategory): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected (DeleteCategory): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found with ID: %s", categoryID)
	}

	return nil
}

func (repo *PostgresRepository) UpdateCategory(ctx context.Context, category *transaction.Category) error {
	query := `UPDATE categories SET category_name = $2, subcategory_name = $3 WHERE id = $1`

	result, err := repo.db.ExecContext(ctx, query,
		category.ID,
		category.CategoryName,
		category.SubcategoryName,
	)
	if err != nil {
		return fmt.Errorf("failed to update category (UpdateCategory): %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected (UpdateCategory): %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found with ID: %s", category.ID)
	}

	return nil
}
