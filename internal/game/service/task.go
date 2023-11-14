package service

import (
	"context"
	"porters/internal/game/model"
)

// TaskRepository ����������� ������
type TaskRepository interface {
	CreateRandomTasks(ctx context.Context, customerUsername string) error
	GetTask(ctx context.Context, taskID int) (model.GetTaskResponse, error)
	UpdateTaskAsDone(ctx context.Context, taskId, porterID int) error
}

// CreateRandomTasks ������� ��������� ������
func (gs GameService) CreateRandomTasks(ctx context.Context, customerUsername string) error {
	return gs.gameRepository.CreateRandomTasks(ctx, customerUsername)
}