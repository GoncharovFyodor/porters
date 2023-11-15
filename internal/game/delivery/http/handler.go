package delivery

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"porters/internal/game/model"
)

// CustomerService сервис заказчика
type CustomerService interface {
	GetCustomerInfo(ctx context.Context, customerID int) (model.GetCustomerInfoResponse, error)
	GetAvailableTasksForCustomer(ctx context.Context, customerID int) ([]model.GetAvailableTasksForCustomerResponse, error)
	StartGame(ctx context.Context, customerID int, porterIDs []int, orderID int) (bool, error)
}

// PorterService сервис грузчика
type PorterService interface {
	GetPorterInfo(ctx context.Context, porterID int) (model.GetAndCreatePorterInfo, error)
	GetCompletedPorterTasks(ctx context.Context, porterID int) ([]model.GetCompletedPorterTasksResponse, error)
}

// TaskService сервис задачи
type TaskService interface {
	CreateRandomTasks(ctx context.Context, customerUsername string) error
}

// GameService сервис игры
type GameService interface {
	CustomerService
	PorterService
	TaskService
}

// Handler обработчик API
type Handler struct {
	salt        string
	gameService GameService
}

// NewHandler создает новый обработчик API
func NewHandler(app *fiber.App, gameService GameService) *fiber.App {
	h := Handler{
		gameService: gameService,
	}

	app.Get("/me", h.authMiddleware, h.GetInfo)
	app.Get("/tasks", h.authMiddleware, h.GetTasks)
	app.Post("/tasks", h.CreateRandomTasks)
	app.Post("/start", h.authMiddleware, h.StartGame)

	return app
}
