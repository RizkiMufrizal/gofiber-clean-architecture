package main

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/controller"
	_ "github.com/RizkiMufrizal/gofiber-clean-architecture/docs"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Go Fiber Clean Architecture
// @version 1.0.0
// @description Baseline project using Go Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9999
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization For JWT
func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	redis := configuration.NewRedis(config)

	//repository
	productRepository := repository.NewProductRepositoryImpl(database)
	transactionRepository := repository.NewTransactionRepositoryImpl(database)
	transactionDetailRepository := repository.NewTransactionDetailRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)

	//service
	productService := service.NewProductServiceImpl(&productRepository, redis)
	transactionService := service.NewTransactionServiceImpl(&transactionRepository)
	transactionDetailService := service.NewTransactionDetailServiceImpl(&transactionDetailRepository)
	userService := service.NewUserServiceImpl(&userRepository)

	//controller
	productController := controller.NewProductController(&productService, config)
	transactionController := controller.NewTransactionController(&transactionService, config)
	transactionDetailController := controller.NewTransactionDetailController(&transactionDetailService, config)
	userController := controller.NewUserController(&userService, config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	productController.Route(app)
	transactionController.Route(app)
	transactionDetailController.Route(app)
	userController.Route(app)

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
