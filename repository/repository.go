package repository

import (
	"context"

	"github.com/lucasvillalbaar/expense-tracker-backend/models"
)

type Repository interface {
	Close()
	InsertTransaction(ctx context.Context, transaction *models.Transaction) error
}

var repository Repository

func SetRepository(r Repository) {
	repository = r
}

func Close() {
	repository.Close()
}

func InsertTransaction(ctx context.Context, transaction *models.Transaction) error {
	return repository.InsertTransaction(ctx, transaction)
}
