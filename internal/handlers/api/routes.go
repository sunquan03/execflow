package v1

import (
	"github.com/gofiber/fiber/v2"
	"task_scheduler/internal/handlers/api/v1/handlers"
)

func Routes(app *fiber.App, handler *handlers.Handler) {
	v1 := app.Group("/api/v1")
	tasks := v1.Group("/tasks")

	tasks.Post("/", handler.CreateTask)
}
