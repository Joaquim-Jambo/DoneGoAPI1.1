package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func DeleteTask(id int) (models.Todo, error) {
	var task models.Todo
	if err := config.Db.First(&task, id).Error; err != nil {
		return models.Todo{}, err //
	}

	if err := config.Db.Delete(&task).Error; err != nil {
		return models.Todo{}, err
	}

	return task, nil
}
