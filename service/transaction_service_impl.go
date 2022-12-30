package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/google/uuid"
)

func NewTransactionServiceImpl(transactionRepository *repository.TransactionRepository) TransactionService {
	return &transactionServiceImpl{TransactionRepository: *transactionRepository}
}

type transactionServiceImpl struct {
	repository.TransactionRepository
}

func (transactionService *transactionServiceImpl) Create(ctx context.Context, transactionModel model.TransactionModel) model.TransactionModel {
	uuidGenerate := uuid.New()
	var transactionDetails []entity.TransactionDetail
	var totalPrice int64 = 0

	for _, detail := range transactionModel.TransactionDetails {
		totalPrice = totalPrice + detail.SubTotalPrice
		transactionDetails = append(transactionDetails, entity.TransactionDetail{
			TransactionId: uuidGenerate,
			ProductId:     detail.ProductId,
			Id:            uuid.New(),
		})
	}

	transaction := entity.Transaction{
		Id:                 uuid.New(),
		TotalPrice:         totalPrice,
		TransactionDetails: transactionDetails,
	}

	transactionService.TransactionRepository.Insert(ctx, transaction)
	return transactionModel
}

func (transactionService *transactionServiceImpl) Delete(ctx context.Context, id string) {
	transaction, err := transactionService.TransactionRepository.FindById(ctx, id)
	exception.PanicLogging(err)
	transactionService.TransactionRepository.Delete(ctx, transaction)
}

func (transactionService *transactionServiceImpl) FindById(ctx context.Context, id string) model.TransactionModel {
	transaction, err := transactionService.TransactionRepository.FindById(ctx, id)
	exception.PanicLogging(err)
	var transactionDetails []model.TransactionDetailModel
	for _, detail := range transaction.TransactionDetails {
		transactionDetails = append(transactionDetails, model.TransactionDetailModel{
			Id:            detail.Id.String(),
			SubTotalPrice: detail.SubTotalPrice,
			Price:         detail.Price,
			Quantity:      detail.Quantity,
			Product: model.ProductModel{
				Id:       detail.Product.Id.String(),
				Name:     detail.Product.Name,
				Price:    detail.Product.Price,
				Quantity: detail.Product.Quantity,
			},
		})
	}

	return model.TransactionModel{
		Id:                 transaction.Id.String(),
		TotalPrice:         transaction.TotalPrice,
		TransactionDetails: transactionDetails,
	}
}

func (transactionService *transactionServiceImpl) FindAll(ctx context.Context) (responses []model.TransactionModel) {
	transactions := transactionService.TransactionRepository.FindAll(ctx)
	for _, transaction := range transactions {
		var transactionDetails []model.TransactionDetailModel
		for _, detail := range transaction.TransactionDetails {
			transactionDetails = append(transactionDetails, model.TransactionDetailModel{
				Id:            detail.Id.String(),
				SubTotalPrice: detail.SubTotalPrice,
				Price:         detail.Price,
				Quantity:      detail.Quantity,
				Product: model.ProductModel{
					Id:       detail.Product.Id.String(),
					Name:     detail.Product.Name,
					Price:    detail.Product.Price,
					Quantity: detail.Product.Quantity,
				},
			})
		}

		responses = append(responses, model.TransactionModel{
			Id:                 transaction.Id.String(),
			TotalPrice:         transaction.TotalPrice,
			TransactionDetails: transactionDetails,
		})
	}

	return responses
}
