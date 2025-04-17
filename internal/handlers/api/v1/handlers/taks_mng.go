package handlers

import (
	"exec_flow/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateTask(ctx *fiber.Ctx) error {
	var taskReq models.TaskRequest
	if err := ctx.BodyParser(&taskReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	id, err := h.service.CreateTask(&taskReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.Response{
		Status:  "success",
		Message: id,
	})
}
