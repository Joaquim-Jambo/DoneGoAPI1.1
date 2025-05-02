package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func GetByOneTask(id int) (models.Todo, error) {
	var task models.Todo
	err := config.Db.Preload("").First(&task, id).Error
	if err != nil {
		return models.Todo{}, err
	}
	return task, nil
}
