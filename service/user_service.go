package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) entity.User
}
