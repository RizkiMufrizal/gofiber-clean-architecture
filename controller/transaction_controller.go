package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	service.TransactionService
	configuration.Config
}

func NewTransactionController(transactionService *service.TransactionService, config configuration.Config) *TransactionController {
	return &TransactionController{TransactionService: *transactionService, Config: config}
}

func (controller TransactionController) Route(app *fiber.App) {
	app.Post("/v1/api/transaction", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.Create)
	app.Delete("/v1/api/transaction/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.Delete)
	app.Get("/v1/api/transaction/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindById)
	app.Get("/v1/api/transaction", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindAll)
}

// Create func create transaction.
// @Description create transaction.
// @Summary create transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body model.TransactionCreateUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction [post]
func (controller TransactionController) Create(c *fiber.Ctx) error {
	var request model.TransactionCreateUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.TransactionService.Create(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Delete func delete one exists transaction.
// @Description delete one exists transaction.
// @Summary delete one exists transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction/{id} [delete]
func (controller TransactionController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.TransactionService.Delete(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

// FindById func gets one exists transaction.
// @Description Get one exists transaction.
// @Summary get one exists transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction/{id} [get]
func (controller TransactionController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.TransactionService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// FindAll func gets all exists transaction.
// @Description Get all exists transaction.
// @Summary get all exists transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction [get]
func (controller TransactionController) FindAll(c *fiber.Ctx) error {
	result := controller.TransactionService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
