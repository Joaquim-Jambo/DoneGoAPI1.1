package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func AddTask(task models.TodoDTO) (models.Todo, error) {
	newTask := models.Todo{
		Title:       task.Title,
		Completed:   false,
		Description: task.Description,
		CategoriaID: task.CategoriaID,
	}
	err := config.Db.Create(&newTask).Error
	if err != nil {
		return models.Todo{}, err
	}
	return newTask, nil
}
