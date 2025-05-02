package repository

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func GetAllCategoria() ([]models.Categoria, error) {
	var categorias []models.Categoria
	err := config.Db.Model(&models.Categoria{}).Preload("Todos").Find(&categorias).Error
	if err != nil {
		return nil, err
	}
	return categorias, nil
}
