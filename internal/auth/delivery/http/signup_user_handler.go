package delivery

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"porters/internal/auth/model"
)

func (h *Handler) signupUser(ctx *fiber.Ctx) error {
	reqBody := ctx.Body()

	var request model.SignupUserRequest
	if err := json.Unmarshal(reqBody, &request); err != nil {
		log.Errorf("signupUser: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := request.Validate(); err != nil {
		log.Errorf("signupUser: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if request.Role != model.CustomerRole && request.Role != model.PorterRole {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid role, choose between porter and customer")
	}

	if err := h.userService.SignupUser(context.Background(), request); err != nil {
		log.Errorf("signupUser: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
