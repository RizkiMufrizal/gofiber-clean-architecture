package controller

import (
	"bytes"
	"encoding/json"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository/impl"
	impl2 "github.com/RizkiMufrizal/gofiber-clean-architecture/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http/httptest"
)

func createTestApp() *fiber.App {
	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	productController.Route(app)
	transactionController.Route(app)
	transactionDetailController.Route(app)
	userController.Route(app)

	return app
}

// setup configuration
var config = configuration.New("../.env.test")
var database = configuration.NewDatabase(config)
var redis = configuration.NewRedis(config)

// repository
var productRepository = impl.NewProductRepositoryImpl(database)
var transactionRepository = impl.NewTransactionRepositoryImpl(database)
var transactionDetailRepository = impl.NewTransactionDetailRepositoryImpl(database)
var userRepository = impl.NewUserRepositoryImpl(database)

// service
var productService = impl2.NewProductServiceImpl(&productRepository, redis)
var transactionService = impl2.NewTransactionServiceImpl(&transactionRepository)
var transactionDetailService = impl2.NewTransactionDetailServiceImpl(&transactionDetailRepository)
var userService = impl2.NewUserServiceImpl(&userRepository)

// controller
var productController = NewProductController(&productService, config)
var transactionController = NewTransactionController(&transactionService, config)
var transactionDetailController = NewTransactionDetailController(&transactionDetailService, config)
var userController = NewUserController(&userService, config)

var appTest = createTestApp()

func authenticationCreate() map[string]interface{} {
	userRepository.DeleteAll()

	password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	exception.PanicLogging(err)
	roles := []string{"ROLE_ADMIN", "ROLE_USER"}
	userRepository.Create("admin", string(password), roles)

	userModel := model.UserModel{
		Username: "admin",
		Password: "admin",
	}

	userRequestBody, _ := json.Marshal(userModel)

	userRequest := httptest.NewRequest("POST", "/v1/api/authentication", bytes.NewBuffer(userRequestBody))
	userRequest.Header.Set("Content-Type", "application/json")
	userRequest.Header.Set("Accept", "application/json")

	userResponse, _ := appTest.Test(userRequest)

	userResponseBody, _ := io.ReadAll(userResponse.Body)
	userWebResponse := model.GeneralResponse{}
	_ = json.Unmarshal(userResponseBody, &userWebResponse)

	userJsonData, _ := json.Marshal(userWebResponse.Data)

	tokenResponse := map[string]interface{}{}
	_ = json.Unmarshal(userJsonData, &tokenResponse)

	return tokenResponse
}
