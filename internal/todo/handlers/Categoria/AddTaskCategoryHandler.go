package handlers

import (
	"net/http"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Categoria"
	"github.com/gin-gonic/gin"
)

// AddTaskCategory godoc
// @Summary      Cria uma nova categoria
// @Description  Adiciona uma nova categoria à base de dados
// @Tags         Categorias
// @Accept       json
// @Produce      json
// @Param        categoria body models.CategoriaDTO true "Dados da categoria"
// @Success      200 {object} models.ResponseExample
// @Failure      400 {object} models.ErrorExample "JSON inválido"
// @Failure      404 {object} models.ErrorExample "Falha ao adicionar Categoria"
// @Router       /category [post]
func AddTaskCategory(ctx *gin.Context) {
	var newCategoria models.CategoriaDTO
	if err := ctx.BindJSON(&newCategoria); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Json inválido !",
		})
	}
	if newCategoria.Nome == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "O campo 'Nome' é obrigatório!",
		})
		return
	}
	response, err := repository.AddTaskCategory(newCategoria)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Falha ao adicionar a categoria",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  response,
		"error": nil,
	})
}
