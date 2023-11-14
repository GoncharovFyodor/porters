package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"porters/hashing"
	authHandler "porters/internal/auth/delivery/http"
	authRepo "porters/internal/auth/repository"
	authService "porters/internal/auth/service"
	"porters/internal/config"
	gameHandler "porters/internal/game/delivery/http"
	gameRepo "porters/internal/game/repository"
	gameService "porters/internal/game/service"
)

var (
	secret = os.Getenv("JWT_SECRET")
	salt   = os.Getenv("JWT_SALT")
)

const (
	configPath = "./configs/config.yaml"
)

func main() {
	runApp()
}

func runApp() {
	app := fiber.New()
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	pgPool, err := pgxpool.New(context.Background(), cfg.DB.Conn)
	if err != nil {
		log.Fatal(err)
	}
	err = pgPool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	hasher := hashing.NewSHA1Hasher(salt)

	authStorage := authRepo.NewStorage(pgPool)
	authServ := authService.NewUser(authStorage, hasher, []byte(secret))
	app = authHandler.NewHandler(app, authServ)

	gameStorage := gameRepo.NewStorage(pgPool)
	gameServ := gameService.NewGame(gameStorage)
	app = gameHandler.NewHandler(app, gameServ)

	log.Info("Server started")
	log.Fatal(app.Listen(cfg.Srv.Addr))
}
