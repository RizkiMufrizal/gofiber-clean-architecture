package repository

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
)

type TransactionRepository interface {
	Insert(ctx context.Context, transaction entity.Transaction) entity.Transaction
	Delete(ctx context.Context, transaction entity.Transaction)
	FindById(ctx context.Context, id string) (entity.Transaction, error)
	FindAll(ctx context.Context) []entity.Transaction
}
