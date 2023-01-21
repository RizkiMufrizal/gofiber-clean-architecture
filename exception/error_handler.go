package exception

import (
	"encoding/json"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}
