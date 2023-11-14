package delivery

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"porters/internal/auth/model"
)

func (h *Handler) loginUser(ctx *fiber.Ctx) error {
	reqBody := ctx.Body()

	var request model.LoginUserRequest
	if err := json.Unmarshal(reqBody, &request); err != nil {
		log.Errorf("loginUser: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := request.Validate(); err != nil {
		log.Errorf("loginUser: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token, err := h.userService.LoginUser(context.Background(), request)
	if err != nil {
		log.Errorf("loginUser: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(map[string]string{
		"token": token,
	})
}
