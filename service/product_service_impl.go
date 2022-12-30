package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/google/uuid"
)

func NewProductServiceImpl(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{ProductRepository: *productRepository}
}

type productServiceImpl struct {
	repository.ProductRepository
}

func (service *productServiceImpl) Create(ctx context.Context, productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	product := entity.Product{
		Name:     productModel.Name,
		Price:    productModel.Price,
		Quantity: productModel.Quantity,
	}
	service.ProductRepository.Insert(ctx, product)
	return productModel
}

func (service *productServiceImpl) Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel {
	product := entity.Product{
		Id:       uuid.MustParse(id),
		Name:     productModel.Name,
		Price:    productModel.Price,
		Quantity: productModel.Quantity,
	}
	service.ProductRepository.Update(ctx, product)
	return productModel
}

func (service *productServiceImpl) Delete(ctx context.Context, id string) {
	product, err := service.ProductRepository.FindById(ctx, id)
	exception.PanicLogging(err)
	service.ProductRepository.Delete(ctx, product)
}

func (service *productServiceImpl) FindById(ctx context.Context, id string) model.ProductModel {
	product, err := service.ProductRepository.FindById(ctx, id)
	exception.PanicLogging(err)
	return model.ProductModel{
		Id:       product.Id.String(),
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
}

func (service *productServiceImpl) FindAll(ctx context.Context) (responses []model.ProductModel) {
	products := service.ProductRepository.FindAl(ctx)
	for _, product := range products {
		responses = append(responses, model.ProductModel{
			Id:       product.Id.String(),
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	if len(products) == 0 {
		return []model.ProductModel{}
	}
	return responses
}
