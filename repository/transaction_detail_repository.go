package repository

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
)

type TransactionDetailRepository interface {
	FindById(ctx context.Context, id string) (entity.TransactionDetail, error)
}
