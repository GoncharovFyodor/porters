package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"porters/internal/game/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// StartGame запуск игры
func (h *Handler) StartGame(ctx *fiber.Ctx) error {
	id, role := ctx.Locals("id").(int), ctx.Locals("role").(string)

	// запустить игру может только заказчик
	if role == model.PorterRole {
		log.Errorf("StartGame: %v", errors.New("invalid role"))
		return ctx.Status(fiber.StatusForbidden).SendString("You have no permission to perform this request")
	}

	reqBody := ctx.Body()
	startGameRequest := model.StartGameRequest{}
	if err := json.Unmarshal(reqBody, &startGameRequest); err != nil {
		log.Errorf("StartGame: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid JSON input")
	}

	// Запуск игры
	success, err := h.gameService.StartGame(context.Background(), id, startGameRequest.PorterIDs, startGameRequest.TaskID)
	if err != nil {
		log.Errorf("StartGame: %v", err)
		return err
	}

	// Вывод сообщения в зависимости от результата игры
	if success {
		return ctx.Status(fiber.StatusOK).SendString("YOU WIN!")
	}
	return ctx.Status(fiber.StatusOK).SendString("YOU LOSE!")
}
