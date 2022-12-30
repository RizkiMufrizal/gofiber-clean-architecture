package repository

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewTransactionRepositoryImpl(DB *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{DB: DB}
}

type transactionRepositoryImpl struct {
	*gorm.DB
}

func (transactionRepository *transactionRepositoryImpl) Insert(ctx context.Context, transaction entity.Transaction) entity.Transaction {
	transaction.Id = uuid.New()
	err := transactionRepository.DB.WithContext(ctx).Create(&transaction)
	exception.PanicLogging(err)
	return transaction
}

func (transactionRepository *transactionRepositoryImpl) Delete(ctx context.Context, transaction entity.Transaction) {
	err := transactionRepository.DB.WithContext(ctx).Where("id = ?", transaction.Id).Delete(&transaction)
	exception.PanicLogging(err)
}

func (transactionRepository *transactionRepositoryImpl) FindById(ctx context.Context, id string) (entity.Transaction, error) {
	var transaction entity.Transaction
	result := transactionRepository.DB.WithContext(ctx).Where("id = ?", id).First(&transaction)
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
