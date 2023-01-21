package impl

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"gorm.io/gorm"
)

func NewTransactionRepositoryImpl(DB *gorm.DB) repository.TransactionRepository {
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
	result := transactionRepository.DB.WithContext(ctx).
		Table("tb_transaction").
		Select("tb_transaction.transaction_id, tb_transaction.total_price, tb_transaction_detail.transaction_detail_id, tb_transaction_detail.sub_total_price, tb_transaction_detail.price, tb_transaction_detail.quantity, tb_product.product_id, tb_product.name, tb_product.price, tb_product.quantity").
		Joins("join tb_transaction_detail on tb_transaction_detail.transaction_id = tb_transaction.transaction_id").
		Joins("join tb_product on tb_product.product_id = tb_transaction_detail.product_id").
		Preload("TransactionDetails").
		Preload("TransactionDetails.Product").
		Where("tb_transaction.transaction_id = ?", id).
		First(&transaction)
	if result.RowsAffected == 0 {
		return entity.Transaction{}, errors.New("transaction Not Found")
	}
	return transaction, nil
}

func (transactionRepository *transactionRepositoryImpl) FindAll(ctx context.Context) []entity.Transaction {
	var transactions []entity.Transaction
	transactionRepository.DB.WithContext(ctx).
		Table("tb_transaction").
		Select("tb_transaction.transaction_id, tb_transaction.total_price, tb_transaction_detail.transaction_detail_id, tb_transaction_detail.sub_total_price, tb_transaction_detail.price, tb_transaction_detail.quantity, tb_product.product_id, tb_product.name, tb_product.price, tb_product.quantity").
		Joins("join tb_transaction_detail on tb_transaction_detail.transaction_id = tb_transaction.transaction_id").
		Joins("join tb_product on tb_product.product_id = tb_transaction_detail.product_id").
		Preload("TransactionDetails").
		Preload("TransactionDetails.Product").
		Find(&transactions)
	return transactions
}
