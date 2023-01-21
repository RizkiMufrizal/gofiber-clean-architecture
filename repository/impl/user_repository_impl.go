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

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) Create(username string, password string, roles []string) {
	var userRoles []entity.UserRole
	for _, role := range roles {
		userRoles = append(userRoles, entity.UserRole{
			Id:       uuid.New(),
			Username: username,
			Role:     role,
		})
	}
	user := entity.User{
		Username:  username,
		Password:  password,
		IsActive:  true,
		UserRoles: userRoles,
	}
	err := userRepository.DB.Create(&user).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) DeleteAll() {
	err := userRepository.DB.Where("1=1").Delete(&entity.User{}).Error
	exception.PanicLogging(err)
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
