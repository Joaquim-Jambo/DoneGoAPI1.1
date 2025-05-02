package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func CompletedTask(id int) (models.Todo, error) {
	var task models.Todo
	err := config.Db.First(&task, id).Error
	if err != nil {
		return models.Todo{}, err
	}
	task.Completed = !task.Completed
	config.Db.Save(&task)
	return task, nil
}
