package repository

import (
	"context"
	"math/rand"
	"porters/internal/game/model"
	"time"
)

// GetTask получает задачу по идентификатору
func (s Storage) GetTask(ctx context.Context, taskID int) (model.GetTaskResponse, error) {
	var task model.GetTaskResponse
	if err := s.pgPool.QueryRow(ctx, "SELECT name, weight, customer_id FROM tasks WHERE id = $1", taskID).
		Scan(&task.Name, &task.Weight, &task.CustomerID); err != nil {
		return model.GetTaskResponse{}, err
	}
	return task, nil
}

// UpdateTaskAsDone обновляет задачу, помечая ее как сделанную
func (s Storage) UpdateTaskAsDone(ctx context.Context, taskID, porterID int) error {
	_, err := s.pgPool.Exec(ctx, "UPDATE tasks SET porter_id = $1 WHERE id = $2", porterID, taskID)
	return err
}

// CreateRandomTasks создает случайные задачи
func (s Storage) CreateRandomTasks(ctx context.Context, customerUsername string) error {
	// открываем транзакцию
	tx, err := s.pgPool.Begin(ctx)
	if err != nil {
		return err
	}

	// сначала получаем ID заказчика
	var customerID int
	if err = tx.QueryRow(ctx, "SELECT id FROM users WHERE username = $1", customerUsername).
		Scan(&customerID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// генерируем случайные задачи (от 2 до 7)
	numOfOrders := r.Intn(7-2+1) + 2
	for i := 0; i < numOfOrders; i++ {
		if _, err = tx.Exec(ctx, "INSERT INTO tasks (name, weight, customer_id) VALUES ($1, $2, $3)",
			model.GenerateRandomTaskName(), r.Intn(80-10+1)+10, customerID); err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	return tx.Commit(ctx)
}
