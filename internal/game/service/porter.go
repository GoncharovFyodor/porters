package service

import (
	"context"
	"porters/internal/game/model"
)

// PorterRepository репозиторий грузчика
type PorterRepository interface {
	GetPorterInfo(ctx context.Context, porterID int) (model.GetAndCreatePorterInfo, error)
	GetCompletedPorterTasks(ctx context.Context, porterID int) ([]model.GetCompletedPorterTasksResponse, error)
	UpdatePorter(ctx context.Context, porterID int, fatigue float64) error
}

// GetPorterInfo получает информацию о грузчике
func (gs GameService) GetPorterInfo(ctx context.Context, porterID int) (model.GetAndCreatePorterInfo, error) {
	return gs.gameRepository.GetPorterInfo(ctx, porterID)
}

// GetCompletedPorterTasks получает список выполненных грузчиком задач
func (gs GameService) GetCompletedPorterTasks(ctx context.Context, porterID int) ([]model.GetCompletedPorterTasksResponse, error) {
	return gs.gameRepository.GetCompletedPorterTasks(ctx, porterID)
}
