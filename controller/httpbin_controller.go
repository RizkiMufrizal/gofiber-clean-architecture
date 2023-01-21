package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type HttpBinController struct {
	service.HttpBinService
}

func NewHttpBinController(httpBinService *service.HttpBinService) *HttpBinController {
	return &HttpBinController{HttpBinService: *httpBinService}
}

func (controller HttpBinController) Route(app *fiber.App) {
	app.Get("/v1/api/httpbin", controller.PostHttpBin)
}

func (controller HttpBinController) PostHttpBin(c *fiber.Ctx) error {

	controller.HttpBinService.PostMethod(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    nil,
	})
}
