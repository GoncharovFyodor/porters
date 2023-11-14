package repository

import (
	"context"
	"porters/internal/game/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Storafe хранилище
type Storage struct {
	pgPool *pgxpool.Pool
}

// NewStorage возвращает новое хранилище
func NewStorage(pool *pgxpool.Pool) Storage {
	return Storage{pool}
}

// Close закрывает соединение с хранилищем
func (s Storage) Close() {
	s.pgPool.Close()
}

// GetCustomerInfo получает информацию о заказчике
func (s Storage) GetCustomerInfo(ctx context.Context, customerID int) (model.GetCustomerInfoResponse, error) {
	tx, err := s.pgPool.Begin(ctx)
	if err != nil {
		return model.GetCustomerInfoResponse{}, err
	}

	availableAttributes := model.GetCustomerInfoResponse{}
	if err = tx.QueryRow(ctx, "SELECT start_capital FROM customers WHERE user_id = $1", customerID).
		Scan(&availableAttributes.CustomerStartCapital); err != nil {
		tx.Rollback(ctx)
		return model.GetCustomerInfoResponse{}, err
	}

	rows, err := tx.Query(ctx, "SELECT user_id, max_weight, drunk, fatigue, salary FROM porters")
	for rows.Next() {
		porter := model.Porter{}
		if err = rows.Scan(&porter.UserID, &porter.MaxWeight, &porter.Drunk, &porter.Fatigue, &porter.Salary); err != nil {
			tx.Rollback(ctx)
			return model.GetCustomerInfoResponse{}, err
		}
		availableAttributes.Porters = append(availableAttributes.Porters, porter)
	}
	return availableAttributes, nil
}

// GetAvailableTasksForCustomer получает доступные задачи для заказчика
func (s Storage) GetAvailableTasksForCustomer(ctx context.Context, customerID int) ([]model.GetAvailableTasksForCustomerResponse, error) {
	rows, err := s.pgPool.Query(ctx, "SELECT id, name, weight FROM tasks WHERE customer_id = $1 AND porter_id IS NULL;", customerID)
	if err != nil {
		return nil, err
	}

	var tasks []model.GetAvailableTasksForCustomerResponse
	for rows.Next() {
		task := model.GetAvailableTasksForCustomerResponse{}
		if err = rows.Scan(&task.ID, &task.Name, &task.Weight); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetCustomerAndPortersStats получает характеристики заказчика и грузчиков
func (s Storage) GetCustomerAndPortersStats(ctx context.Context, customerID int, porterIDs []int, taskID int) (int, map[int]model.GetAndCreatePorterInfo, error) {
	tx, err := s.pgPool.Begin(ctx)
	if err != nil {
		return 0, nil, err
	}

	// получение стартового капитала
	var customerStartCapital int
	if err = tx.QueryRow(ctx, "SELECT start_capital FROM customers WHERE user_id = $1", customerID).
		Scan(&customerStartCapital); err != nil {
		tx.Rollback(ctx)
		return 0, nil, err
	}

	// получение информации о грузчиках
	porters := make(map[int]model.GetAndCreatePorterInfo)
	for _, porterID := range porterIDs {
		var porter model.GetAndCreatePorterInfo
		if err = tx.QueryRow(ctx, "SELECT max_weight, drunk, fatigue, salary FROM porters WHERE user_id = $1", porterID).
			Scan(&porter.MaxWeight, &porter.Drunk, &porter.Fatigue, &porter.Salary); err != nil {
			tx.Rollback(ctx)
			return 0, nil, err
		}
		porters[porterID] = porter
	}

	if err = tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return 0, nil, err
	}

	return customerStartCapital, porters, nil
}

// UpdateCustomer обновляет стартовый капитал заказчика
func (s Storage) UpdateCustomer(ctx context.Context, customerID int, customerStartCapital int) error {
	_, err := s.pgPool.Exec(ctx, "UPDATE customers SET start_capital = $1 WHERE user_id = $2", customerStartCapital, customerID)
	return err
}
