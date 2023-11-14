package delivery

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"porters/internal/auth/model"
)

type UserService interface {
	SignupUser(ctx context.Context, request model.SignupUserRequest) error
	LoginUser(ctx context.Context, request model.LoginUserRequest) (string, error)
}

type Handler struct {
	userService UserService
}

func NewHandler(app *fiber.App, userService UserService) *fiber.App {
	h := Handler{
		userService: userService,
	}

	app.Get("/login", h.loginUser)
	app.Post("/register", h.signupUser)

	return app
}