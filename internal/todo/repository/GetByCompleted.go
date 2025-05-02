package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func GetByCompleted() ([]models.Todo, error) {
	var todos []models.Todo
	err := config.Db.Where("Completed = ?", true).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
