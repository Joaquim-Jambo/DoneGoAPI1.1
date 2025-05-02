package repository

import (
	"fmt"

	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func DeleteTask(id int) (string, error) {
	var task models.Todo
	err := config.Db.Delete(&task, id).Error
	if err != nil {
		return "", err
	}
	config.Db.First(&task, id)
	msg := fmt.Sprintf("task %v deletada com sucesso !", task.Title)
	return msg, nil
}
