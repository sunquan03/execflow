package main

import (
	v1 "exec_flow/internal/handlers/api"
	"exec_flow/internal/handlers/api/v1/handlers"
	"exec_flow/internal/repositories"
	"exec_flow/internal/services"
	"exec_flow/pkg/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

const idleTimeout = 1 * time.Second

func main() {
	fmt.Println("Check")

	redisClient := repositories.NewRedisClient()

	repository := repositories.NewRepository(redisClient)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)

	_app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			return c.Status(code).JSON(models.Response{
				Status:  "error",
				Message: err.Error(),
			})
		},
		DisableStartupMessage: true,
	})

	v1.Routes(_app, handler)

	err := _app.Listen(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}
}
