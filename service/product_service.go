package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
)

type ProductService interface {
	Create(ctx context.Context, model model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel
	Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id int) model.ProductCreateOrUpdateModel
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) model.ProductModel
	FindAll(ctx context.Context) []model.ProductModel
}
