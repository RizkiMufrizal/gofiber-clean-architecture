package repository

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"gorm.io/gorm"
)

func NewTransactionRepositoryImpl(DB *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{DB: DB}
}

type transactionRepositoryImpl struct {
	*gorm.DB
}

func (transactionRepository *transactionRepositoryImpl) Insert(ctx context.Context, transaction entity.Transaction) entity.Transaction {
	err := transactionRepository.DB.WithContext(ctx).Create(&transaction).Error
	exception.PanicLogging(err)
	return transaction
}

func (transactionRepository *transactionRepositoryImpl) Delete(ctx context.Context, transaction entity.Transaction) {
	transactionRepository.DB.WithContext(ctx).Delete(&transaction)
}

func (transactionRepository *transactionRepositoryImpl) FindById(ctx context.Context, id string) (entity.Transaction, error) {
	var transaction entity.Transaction
	result := transactionRepository.DB.WithContext(ctx).Where("transaction_id = ?", id).First(&transaction)
	if result.RowsAffected == 0 {
		return entity.Transaction{}, errors.New("transaction Not Found")
	}
	return transaction, nil
}

func (transactionRepository *transactionRepositoryImpl) FindAll(ctx context.Context) []entity.Transaction {
	var transactions []entity.Transaction
	transactionRepository.DB.WithContext(ctx).Find(&transactions)
	return transactions
}
