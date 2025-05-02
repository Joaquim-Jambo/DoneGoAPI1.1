package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func UpdateCategory(id int, updateCategory models.CategoriaDTO) (models.Categoria, error) {
	var categoria models.Categoria
	err := config.Db.First(&categoria, &id).Error
	if err != nil {
		return models.Categoria{}, err
	}
	if updateCategory.Nome == "" {
		categoria.Nome = updateCategory.Nome
	}
	err = config.Db.Save(&categoria).Error
	if err != nil {
		return models.Categoria{}, err
	}
	return categoria, nil
}
