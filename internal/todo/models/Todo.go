package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
type TodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoUpdateDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
