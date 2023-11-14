package service

import (
	"context"
	"porters/internal/game/model"
)

// CustomerRepository репозиторий заказчика
type CustomerRepository interface {
	GetCustomerInfo(ctx context.Context, customerID int) (model.GetCustomerInfoResponse, error)
	GetAvailableTasksForCustomer(ctx context.Context, customerID int) ([]model.GetAvailableTasksForCustomerResponse, error)
	GetCustomerAndPortersStats(ctx context.Context, customerID int, porterIDs []int, taskID int) (int, map[int]model.GetAndCreatePorterInfo, error)
	UpdateCustomer(ctx context.Context, customerID int, customerStartCapital int) error
}

// GameRepository репозиторий игры
type GameRepository interface {
	CustomerRepository
	PorterRepository
	TaskRepository
}

// GameService сервис игры
type GameService struct {
	gameRepository GameRepository
}

// NewGame создает новую игру
func NewGame(gameRepository GameRepository) GameService {
	return GameService{gameRepository: gameRepository}
}

// GetCustomerInfo получает информацию о заказчиках
func (gs GameService) GetCustomerInfo(ctx context.Context, customerID int) (model.GetCustomerInfoResponse, error) {
	return gs.gameRepository.GetCustomerInfo(ctx, customerID)
}

// GetAvailableTasksForCustomer получает список доступных задач для заказчика
func (gs GameService) GetAvailableTasksForCustomer(ctx context.Context, customerID int) ([]model.GetAvailableTasksForCustomerResponse, error) {
	return gs.gameRepository.GetAvailableTasksForCustomer(ctx, customerID)
}

// StartGame запускает игру
func (gs GameService) StartGame(ctx context.Context, customerID int, porterIDs []int, taskID int) (bool, error) {
	customerStartCapital, porters, err := gs.gameRepository.GetCustomerAndPortersStats(ctx, customerID, porterIDs, taskID)
	if err != nil {
		return false, err
	}

	task, err := gs.gameRepository.GetTask(ctx, taskID)
	if err != nil {
		return false, err
	}

	var totalPortersSalary int
	var totalMaxWeight float64
	for _, porter := range porters {
		totalPortersSalary += porter.Salary

		// расчет поднятого веса для пьяных и трезвых грузчиков
		if porter.Drunk {
			totalMaxWeight += float64(porter.MaxWeight) * ((100 - porter.Fatigue) / 100) * (porter.Fatigue + 50/100)
		} else {
			totalMaxWeight += float64(porter.MaxWeight) * ((100 - porter.Fatigue) / 100)
		}
	}

	success := true
	if totalPortersSalary > customerStartCapital || totalMaxWeight < float64(task.Weight) {
		success = false
	}

	if err = gs.gameRepository.UpdateCustomer(ctx, customerID, customerStartCapital-totalPortersSalary); err != nil {
		return false, err
	}

	for porterID, porter := range porters {
		if porter.Fatigue+20 >= 100 {
			porter.Fatigue = 100
			if err = gs.gameRepository.UpdatePorter(ctx, porterID, porter.Fatigue); err != nil {
				return false, err
			}
		} else {
			porter.Fatigue += 20
			if err = gs.gameRepository.UpdatePorter(ctx, porterID, porter.Fatigue); err != nil {
				return false, err
			}
		}

		if success {
			if err = gs.gameRepository.UpdateTaskAsDone(ctx, taskID, porterID); err != nil {
				return false, err
			}
		}
	}
	if err = gs.gameRepository.UpdateCustomer(ctx, customerID, customerStartCapital-totalPortersSalary); err != nil {
		return false, err
	}
	return success, nil
}
