package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// CreateRandomTasks создание случайных задач
func (h *Handler) CreateRandomTasks(ctx *fiber.Ctx) error {
	reqBody := ctx.Body()
	m := make(map[string]string)
	if err := json.Unmarshal(reqBody, &m); err != nil {
		log.Errorf("CreateRandomTasks: %v", err)
		return err
	}

	customerUsername, ok := m["customerUsername"]
	if !ok {
		err := errors.New("empty input")
		log.Errorf("CreateRandomTasks: %v", err)
		return err
	}

	if err := h.gameService.CreateRandomTasks(context.Background(), customerUsername); err != nil {
		log.Errorf("CreateRandomTasks: %v", err)
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}
