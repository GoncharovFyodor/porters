package model

const (
	// CustomerRole представляет собой роль "заказчик"
	CustomerRole = "customer"
)

// Customer заказчик
type Customer struct {
	UserID       int `db:"user_id" json:"user_id"`
	StartCapital int `db:"start_capital" json:"start_capital"`
}

// StartGameRequest представляет собой запрос начала игры
type StartGameRequest struct {
	PorterIDs []int `json:"porterIDs"`
	TaskID    int   `json:"taskID"`
}

// GetCustomerInfoResponse представляет собой ответ на запрос информации о заказчике: стартовый капитал и грузчики
type GetCustomerInfoResponse struct {
	CustomerStartCapital int      `json:"start_capital"`
	Porters              []Porter `json:"porters"`
}

// GetAvailableTasksForCustomerResponse представляет собой ответ на запрос доступных задач у заказчика
type GetAvailableTasksForCustomerResponse struct {
	ID     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Weight int    `db:"weight" json:"weight"`
}
