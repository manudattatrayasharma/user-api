package main

import (
	"user-api/config"
	db "user-api/db/sqlc"
	"user-api/internal/handler"
	"user-api/internal/logger"
	"user-api/internal/middleware"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	zapLogger := logger.NewLogger()
	defer zapLogger.Sync()

	
	sqlDB, err := config.NewMySQLDB()
	if err != nil {
		zapLogger.Fatal("failed to connect to database", zap.Error(err))
	}

	queries := db.New(sqlDB)

	app := fiber.New()

	app.Use(middleware.RequestLogger(zapLogger))

	userHandler := handler.NewUserHandler(
		service.NewUserService(queries),
	)

	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Get("/users", userHandler.ListUsers)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)

	zapLogger.Info("server started", zap.String("port", "3000"))

	if err := app.Listen(":3000"); err != nil {
		zapLogger.Fatal("server failed to start", zap.Error(err))
	}
}
