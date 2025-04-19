package v1

import (
	"exec_flow/internal/handlers/api/v1/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler *handlers.Handler) {
	v1 := app.Group("/api/v1")
	tasks := v1.Group("/tasks")

	tasks.Post("/", handler.CreateTask)
}
