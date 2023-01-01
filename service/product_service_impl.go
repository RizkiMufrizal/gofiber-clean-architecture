package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/validation"
	"github.com/go-redis/cache/v8"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func NewProductServiceImpl(productRepository *repository.ProductRepository, cache *cache.Cache) ProductService {
	return &productServiceImpl{ProductRepository: *productRepository, Cache: cache}
}

type productServiceImpl struct {
	repository.ProductRepository
	*cache.Cache
}

func (service *productServiceImpl) Create(ctx context.Context, productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	validation.Validate(productModel)
	product := entity.Product{
		Name:     productModel.Name,
		Price:    productModel.Price,
		Quantity: productModel.Quantity,
	}
	service.ProductRepository.Insert(ctx, product)
	return productModel
}

func (service *productServiceImpl) Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel {
	validation.Validate(productModel)
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
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.ProductRepository.Delete(ctx, product)
}

func (service *productServiceImpl) FindById(ctx context.Context, id string) model.ProductModel {
	var product entity.Product
	productCache := configuration.GetCache(service.Cache, ctx, "product_"+id)
	if productCache == nil {
		productFindById, err := service.ProductRepository.FindById(ctx, id)
		if err != nil {
			panic(exception.NotFoundError{
				Message: err.Error(),
			})
		}
		configuration.SetCache(service.Cache, ctx, "product_"+id, &productFindById)
		product = productFindById
	} else {
		err := mapstructure.Decode(productCache, &product)
		exception.PanicLogging(err)
	}

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
