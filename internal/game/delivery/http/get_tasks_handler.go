package delivery

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"porters/internal/game/model"
)

// GetTasks получение задач
func (h *Handler) GetTasks(ctx *fiber.Ctx) error {
	id, role := ctx.Locals("id").(int), ctx.Locals("role").(string)
	if role == model.CustomerRole {
		customerTasks, err := h.gameService.GetAvailableTasksForCustomer(context.Background(), id)
		if err != nil {
			log.Errorf("GetTasks: %v", err)
			return err
		} else if len(customerTasks) == 0 {
			return ctx.Status(fiber.StatusOK).SendString("No tasks assigned")
		}
		return ctx.JSON(customerTasks)
	} else if role == model.PorterRole {
		porterTasks, err := h.gameService.GetCompletedPorterTasks(context.Background(), id)
		if err != nil {
			log.Errorf("GetTasks: %v", err)
			return err
		} else if len(porterTasks) == 0 {
			return ctx.Status(fiber.StatusOK).SendString("No tasks you already done")
		}
		return ctx.JSON(porterTasks)
	}
	return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("unknown user role %s", role))
}
