package model

import "github.com/go-playground/validator/v10"

var (
	// CustomerRole роль "заказчик"
	CustomerRole = "customer"

	// PorterRole роль "грузчик"
	PorterRole = "porter"

	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

// User учетная запись пользователя
type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

// SignupUserRequest запрос регистрации пользователя
type SignupUserRequest struct {
	Username string `db:"username" json:"username" validate:"required,gte=2"`
	Password string `db:"password" json:"password" validate:"required,gte=6"`
	Role     string `db:"role" json:"role" validate:"required"`
}

// Validate валидация учетной записи пользователя при регистрации
func (i *SignupUserRequest) Validate() error {
	return validate.Struct(i)
}

// LoginUserRequest запрос логина пользователя
type LoginUserRequest struct {
	Username string `db:"username" json:"username" validate:"required,gte=2"`
	Password string `db:"password" json:"password" validate:"required,gte=6"`
}

// Validate валидация учетной записи пользователя при логине
func (i *LoginUserRequest) Validate() error {
	return validate.Struct(i)
}
