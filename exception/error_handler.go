package exception

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.JSON(model.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}
