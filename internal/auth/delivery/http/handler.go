package delivery

import (
	"context"
	"porters/internal/auth/model"

	"github.com/gofiber/fiber/v2"
)

// UserService сервис пользователя
type UserService interface {
	SignupUser(ctx context.Context, request model.SignupUserRequest) error
	LoginUser(ctx context.Context, request model.LoginUserRequest) (string, error)
}

// Handler обработчик API
type Handler struct {
	userService UserService
}

// NewHandler создает новый обработчик API
func NewHandler(app *fiber.App, userService UserService) *fiber.App {
	h := Handler{
		userService: userService,
	}

	app.Get("/login", h.loginUser)
	app.Post("/register", h.signupUser)

	return app
}
