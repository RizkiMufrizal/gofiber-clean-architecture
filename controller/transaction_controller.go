package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	service.TransactionService
}

func NewTransactionController(transactionService *service.TransactionService) *TransactionController {
	return &TransactionController{TransactionService: *transactionService}
}

func (controller TransactionController) Route(app *fiber.App) {
	app.Post("/v1/api/transaction", controller.Create)
	app.Delete("/v1/api/transaction/:id", controller.Delete)
	app.Get("/v1/api/transaction/:id", controller.FindById)
	app.Get("/v1/api/transaction", controller.FindAll)
}

func (controller TransactionController) Create(c *fiber.Ctx) error {
	var request model.TransactionCreateUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.TransactionService.Create(c.Context(), request)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller TransactionController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.TransactionService.Delete(c.Context(), id)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

func (controller TransactionController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.TransactionService.FindById(c.Context(), id)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller TransactionController) FindAll(c *fiber.Ctx) error {
	result := controller.TransactionService.FindAll(c.Context())
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
