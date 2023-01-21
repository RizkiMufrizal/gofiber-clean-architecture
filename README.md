# Go Fiber Clean Architecture

This project was created to learn golang with go fiber framework

## How To Run

1. Run docker compose with command `docker compose up`
2. Run Migration DB `migrate -database "mysql://root:root@tcp(localhost:3306)/gofiber_clean_architecture" -path db/migrations up`
3. Run application with command `go run main.go`

## Feature

- [x] Database ORM
- [x] Database Relational
- [x] Json Validation
- [x] JWT Security
- [x] Database migration
- [x] Docker Support
- [x] Open API / Swagger
- [x] Integration Test
- [x] Http Client
- [x] Error Handling
- [x] Logging
- [x] Cache