package repository

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"gorm.io/gorm"
)

func NewTransactionDetailRepositoryImpl(DB *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepositoryImpl{DB: DB}
}

type transactionDetailRepositoryImpl struct {
	*gorm.DB
}

func (transactionDetailRepository *transactionDetailRepositoryImpl) FindById(ctx context.Context, id string) (entity.TransactionDetail, error) {
	var transactionDetail entity.TransactionDetail
	result := transactionDetailRepository.DB.WithContext(ctx).Where("id = ?", id).First(&transactionDetail)
	if result.RowsAffected == 0 {
		return entity.TransactionDetail{}, errors.New("transaction Detail Not Found")
	}
	return transactionDetail, nil
}
