package repository

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) Authentication(ctx context.Context, username string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).
		Joins("inner join tb_user_role on tb_user_role.username = tb_user.username").
		Preload("UserRoles").
		Where("tb_user.username = ? and tb_user.is_active = ?", username, true).
		Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return userResult, nil
}
