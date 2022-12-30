package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
)

type ProductService interface {
	Create(ctx context.Context, model model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel
	Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.ProductModel
	FindAll(ctx context.Context) []model.ProductModel
}
