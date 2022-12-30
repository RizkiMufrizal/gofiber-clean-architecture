package configuration

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
