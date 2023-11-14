package repository

import (
	"context"
	"math"
	"math/rand"
	authmodel "porters/internal/auth/model"
	"porters/internal/game/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Storage хранилище
type Storage struct {
	pgPool *pgxpool.Pool
}

// NewStorage создает новое хранилище
func NewStorage(pool *pgxpool.Pool) Storage {
	return Storage{pool}
}

// Close закрывает соединение с хранилищем
func (s Storage) Close() {
	s.pgPool.Close()
}

// CreateUser создает пользователя
func (s Storage) CreateUser(ctx context.Context, user authmodel.User) error {
	tx, err := s.pgPool.Begin(ctx)
	if err != nil {
		return err
	}

	// Добавление данных о пользователе в БД и возвращение идентификатора
	var id int
	if err = tx.QueryRow(ctx, "INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id",
		user.Username, user.Password, user.Role).Scan(&id); err != nil {
		tx.Rollback(ctx)
		return err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if user.Role == model.CustomerRole {
		// Добавление данных о заказчике в БД
		if _, err = tx.Exec(ctx, "INSERT INTO customers (user_id, start_capital) VALUES ($1, $2)",
			id, r.Intn(100000-10000+1)+10000); err != nil {
			tx.Rollback(ctx)
			return err
		}
	} else if user.Role == model.PorterRole {
		// Добавление данных о грузчике в БД
		if _, err = tx.Exec(ctx, "INSERT INTO porters (user_id, max_weight, drunk, fatigue, salary) VALUES ($1, $2, $3, $4, $5)",
			id, r.Intn(30-5+1)+5, r.Intn(2) == 1, math.Round((r.Float64())*100+1), r.Intn(30000-10000+1)+10000); err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	return tx.Commit(ctx)
}

// GetByCredentials получает данные пользователя по введенным имени пользователя и паролю
func (s Storage) GetByCredentials(ctx context.Context, username, password string) (authmodel.User, error) {
	var user authmodel.User
	if err := s.pgPool.QueryRow(ctx, "SELECT id, username, password, role FROM users WHERE username = $1 AND password = $2",
		username, password).Scan(&user.ID, &user.Username, &user.Password, &user.Role); err != nil {
		return authmodel.User{}, err
	}
	return user, nil
}
