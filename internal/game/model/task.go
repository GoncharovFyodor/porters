package model

import (
	"fmt"
	"math/rand"
	"time"
)

// Task задача
type Task struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Weight     int    `db:"weight" json:"weight"`
	CustomerID int    `db:"customer_id" json:"customer_id"`
	PorterID   int    `db:"porter_id" json:"porter_id"`
}

// GetTaskResponse представляет собой ответ на получение от
type GetTaskResponse struct {
	Name       string `db:"name" json:"name"`
	Weight     int    `db:"weight" json:"weight"`
	CustomerID int    `db:"customer_id" json:"customer_id"`
}

// GenerateRandomTaskName генерирует случайные названия задач
func GenerateRandomTaskName() string {
	words := []string{"apple", "banana", "cherry", "pear", "orange", "phone", "table", "book", "sofa", "TV", "bed"}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomIndex := r.Intn(len(words))
	randomIndex2 := r.Intn(len(words))
	randomIndex3 := r.Intn(len(words))
	return fmt.Sprintf("%s, %s, %s", words[randomIndex], words[randomIndex2], words[randomIndex3])
}
