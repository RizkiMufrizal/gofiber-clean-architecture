package impl

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewProductRepositoryImpl(DB *gorm.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: DB}
}

type productRepositoryImpl struct {
	*gorm.DB
}

func (repository *productRepositoryImpl) Insert(ctx context.Context, product entity.Product) entity.Product {
	product.Id = uuid.New()
	err := repository.DB.WithContext(ctx).Create(&product).Error
	exception.PanicLogging(err)
	return product
}

func (repository *productRepositoryImpl) Update(ctx context.Context, product entity.Product) entity.Product {
	err := repository.DB.WithContext(ctx).Where("product_id = ?", product.Id).Updates(&product).Error
	exception.PanicLogging(err)
	return product
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, product entity.Product) {
	err := repository.DB.WithContext(ctx).Delete(&product).Error
	exception.PanicLogging(err)
}

func (repository *productRepositoryImpl) FindById(ctx context.Context, id string) (entity.Product, error) {
	var product entity.Product
	result := repository.DB.WithContext(ctx).Unscoped().Where("product_id = ?", id).First(&product)
	if result.RowsAffected == 0 {
		return entity.Product{}, errors.New("product Not Found")
	}
	return product, nil
}

func (repository *productRepositoryImpl) FindAl(ctx context.Context) []entity.Product {
	var products []entity.Product
	repository.DB.WithContext(ctx).Find(&products)
	return products
}
