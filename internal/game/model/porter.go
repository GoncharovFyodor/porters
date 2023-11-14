package model

const (
	// PorterRole представляет собой роль "грузчик"
	PorterRole = "porter"
)

// Porter грузчик
type Porter struct {
	UserID    int     `db:"user_id" json:"user_id"`
	MaxWeight int     `db:"max_weight" json:"max_weight"`
	Drunk     bool    `db:"drunk" json:"drunk"`
	Fatigue   float64 `db:"fatigue" json:"fatigue"`
	Salary    int     `db:"salary" json:"salary"`
}

// GetAndCreatePorterInfo представляет собой возвращаемую информацию о грузчиках
type GetAndCreatePorterInfo struct {
	MaxWeight int     `db:"max_weight" json:"max_weight"`
	Drunk     bool    `db:"drunk" json:"drunk"`
	Fatigue   float64 `db:"fatigue" json:"fatigue"`
	Salary    int     `db:"salary" json:"salary"`
}

// GetCompletedPorterTasksResponse представляет собой ответ на запрос о завершенных задачах
type GetCompletedPorterTasksResponse struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Weight   int    `db:"weight" json:"weight"`
	PorterID int    `db:"porter_id" json:"porter_id"`
}
