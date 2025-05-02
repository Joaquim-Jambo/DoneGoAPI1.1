package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func AddTaskCategory(categoria models.CategoriaDTO) (models.Categoria, error) {
	newCategoria := models.Categoria{
		Nome: categoria.Nome,
	}
	err := config.Db.Create(&newCategoria).Error
	if err != nil {
		return models.Categoria{}, err
	}
	return newCategoria, nil
}
