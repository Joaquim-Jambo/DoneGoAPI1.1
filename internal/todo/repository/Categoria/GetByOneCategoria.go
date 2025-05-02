package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func GetByOneCategoria(id int) (models.Categoria, error) {
	var categoria models.Categoria
	err := config.Db.Preload("Todos").First(&categoria, id).Error
	if err != nil {
		return models.Categoria{}, err
	}
	return categoria, nil
}
