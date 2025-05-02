package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func UpdateTask(id int, updateTask models.TodoUpdateDTO) (models.Todo, error) {
	var task models.Todo

	err := config.Db.First(&task, id).Error
	if err != nil {
		return models.Todo{}, err
	}
	if updateTask.Title != "" {
		task.Title = updateTask.Title
	}
	if updateTask.Description != "" {
		task.Description = updateTask.Description
	}
	err = config.Db.Save(&task).Error
	if err != nil {
		return models.Todo{}, err
	}
	return task, nil
}
