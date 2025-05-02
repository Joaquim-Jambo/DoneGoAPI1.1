package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func DeleteCategoria(id int) (models.Categoria, error) {
	var categoria models.Categoria
	if err := config.Db.First(&categoria, id).Error; err != nil {
		return models.Categoria{}, err //
	}

	if err := config.Db.Delete(&categoria).Error; err != nil {
		return models.Categoria{}, err
	}

	return categoria, nil
}
