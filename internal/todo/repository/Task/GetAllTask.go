package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func GetAllTask() ([]models.Todo, error) {
	var tasks []models.Todo
	err := config.Db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
