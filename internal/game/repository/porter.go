package repository

import (
	"context"
	"porters/internal/game/model"
)

// GetPorterInfo �������� ���������� � ��������
func (s Storage) GetPorterInfo(ctx context.Context, porterID int) (model.GetAndCreatePorterInfo, error) {
	porter := model.GetAndCreatePorterInfo{}
	if err := s.pgPool.QueryRow(ctx, "SELECT max_weight, drunk, fatigue, salary FROM porters WHERE user_id = $1", porterID).
		Scan(&porter.MaxWeight, &porter.Drunk, &porter.Fatigue, &porter.Salary); err != nil {
		return model.GetAndCreatePorterInfo{}, err
	}

	return porter, nil
}

// GetCompletedPorterTasks �������� ����������� ������ ��������
func (s Storage) GetCompletedPorterTasks(ctx context.Context, porterID int) ([]model.GetCompletedPorterTasksResponse, error) {
	rows, err := s.pgPool.Query(ctx, "SELECT id, name, weight, porter_id FROM tasks WHERE porter_id = $1", porterID)
	if err != nil {
		return nil, err
	}

	var tasks []model.GetCompletedPorterTasksResponse
	for rows.Next() {
		task := model.GetCompletedPorterTasksResponse{}
		if err := rows.Scan(&task.ID, &task.Name, &task.Weight, &task.PorterID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// UpdatePorter ��������� ������� ��������� ��������
func (s Storage) UpdatePorter(ctx context.Context, porterID int, fatigue float64) error {
	_, err := s.pgPool.Exec(ctx, "UPDATE porters SET fatigue = $1 WHERE user_id = $2", fatigue, porterID)
	return err
}