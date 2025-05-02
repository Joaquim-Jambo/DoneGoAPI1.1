package models

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type TodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// ResponseExample representa uma resposta de sucesso
type ResponseExample struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorExample representa uma resposta de erro
type ErrorExample struct {
	Error string `json:"error"`
}

type TodoUpdateDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
