package delivery

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"strings"
	"time"
)

var secret = os.Getenv("JWT_SECRET")

// middleware аутентификации
func (h *Handler) authMiddleware(ctx *fiber.Ctx) error {
	token, err := getTokenFromRequest(ctx)
	if err != nil {
		return err
	}

	parser := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Alg()}}
	t, err := parser.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		err = errors.New("invalid claims")
		log.Errorf("authMiddleware: %v", err)
		return err
	}

	if err = claims.Valid(); err != nil {
		log.Errorf("authMiddleware: %v", err)
		return errors.New("invalid token")
	}

	if !t.Valid {
		log.Errorf("authMiddleware: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	expires := claims.ExpiresAt
	if expires == 0 {
		log.Errorf("authMiddleware: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString("token expiration is not set")
	} else if expires < time.Now().Unix() {
		log.Errorf("authMiddleware: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString("login token expired, please get new one")
	}

	subj := claims.Subject
	if subj == "" {
		log.Errorf("authMiddleware: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString("user id is not set")
	}

	id, err := strconv.Atoi(subj)
	if err != nil {
		log.Errorf("authMiddleware: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("can't convert id to integer")
	}
	ctx.Locals("id", id)

	role := claims.Issuer
	if role == "" {
		log.Errorf("authMiddleware: %v", err)
		return ctx.Status(fiber.StatusBadRequest).SendString("user role is not set")
	}
	ctx.Locals("role", role)
	return ctx.Next()
}

func getTokenFromRequest(ctx *fiber.Ctx) (string, error) {
	header := ctx.Get("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}
	return headerParts[1], nil
}
