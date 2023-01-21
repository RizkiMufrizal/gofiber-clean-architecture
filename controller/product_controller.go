package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service.ProductService
	configuration.Config
}

func NewProductController(productService *service.ProductService, config configuration.Config) *ProductController {
	return &ProductController{ProductService: *productService, Config: config}
}

func (controller ProductController) Route(app *fiber.App) {
	app.Post("/v1/api/product", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Create)
	app.Put("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Update)
	app.Delete("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Delete)
	app.Get("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.FindById)
	app.Get("/v1/api/product", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.FindAll)
}

// Create func create product.
// @Description create product.
// @Summary create product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product [post]
func (controller ProductController) Create(c *fiber.Ctx) error {
	var request model.ProductCreateOrUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Create(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Update func update one exists product.
// @Description update one exists product.
// @Summary update one exists product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Param id path string true "Product Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product/{id} [put]
func (controller ProductController) Update(c *fiber.Ctx) error {
	var request model.ProductCreateOrUpdateModel
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Update(c.Context(), request, id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Delete func delete one exists product.
// @Description delete one exists product.
// @Summary delete one exists product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product/{id} [delete]
func (controller ProductController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.ProductService.Delete(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

// FindById func gets one exists product.
// @Description Get one exists product.
// @Summary get one exists product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product/{id} [get]
func (controller ProductController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.ProductService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// FindAll func gets all exists products.
// @Description Get all exists products.
// @Summary get all exists products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product [get]
func (controller ProductController) FindAll(c *fiber.Ctx) error {
	result := controller.ProductService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
