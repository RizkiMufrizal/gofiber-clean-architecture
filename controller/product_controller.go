package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: *productService}
}

func (controller ProductController) Route(app *fiber.App) {
	app.Post("/v1/api/product", controller.Create)
	app.Put("/v1/api/product/:id", controller.Update)
	app.Delete("/v1/api/product/:id", controller.Delete)
	app.Get("/v1/api/product/:id", controller.FindById)
	app.Get("/v1/api/product", controller.FindAll)
}

func (controller ProductController) Create(c *fiber.Ctx) error {
	var request model.ProductCreateOrUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Create(c.Context(), request)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller ProductController) Update(c *fiber.Ctx) error {
	var request model.ProductCreateOrUpdateModel
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Update(c.Context(), request, id)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller ProductController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.ProductService.Delete(c.Context(), id)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

func (controller ProductController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.ProductService.FindById(c.Context(), id)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller ProductController) FindAll(c *fiber.Ctx) error {
	result := controller.ProductService.FindAll(c.Context())
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
