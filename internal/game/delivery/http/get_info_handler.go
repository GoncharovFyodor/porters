package delivery

import (
	"context"
	"fmt"
	"porters/internal/game/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// GetInfo получение информации
func (h *Handler) GetInfo(ctx *fiber.Ctx) error {
	id, role := ctx.Locals("id").(int), ctx.Locals("role").(string)
	if role == model.CustomerRole {
		customerInfo, err := h.gameService.GetCustomerInfo(context.Background(), id)
		if err != nil {
			log.Errorf("GetInfo: %v", err)
		}
		return ctx.JSON(customerInfo)
	} else if role == model.PorterRole {
		porterInfo, err := h.gameService.GetPorterInfo(context.Background(), id)
		if err != nil {
			log.Errorf("GetInfo: %v", err)
		}
		return ctx.JSON(porterInfo)
	}
	return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("unknown user role %s", role))
}
