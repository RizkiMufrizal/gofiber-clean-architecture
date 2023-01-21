package impl

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
)

func NewTransactionDetailServiceImpl(transactionDetailRepository *repository.TransactionDetailRepository) service.TransactionDetailService {
	return &transactionDetailServiceImpl{TransactionDetailRepository: *transactionDetailRepository}
}

type transactionDetailServiceImpl struct {
	repository.TransactionDetailRepository
}

func (transactionDetailService *transactionDetailServiceImpl) FindById(ctx context.Context, id string) model.TransactionDetailModel {
	transactionDetail, err := transactionDetailService.TransactionDetailRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return model.TransactionDetailModel{
		Id:            transactionDetail.Id.String(),
		SubTotalPrice: transactionDetail.SubTotalPrice,
		Price:         transactionDetail.Price,
		Quantity:      transactionDetail.Quantity,
		Product: model.ProductModel{
			Id:       transactionDetail.Product.Id.String(),
			Name:     transactionDetail.Product.Name,
			Price:    transactionDetail.Product.Price,
			Quantity: transactionDetail.Product.Quantity,
		},
	}
}
