package handlers

import (
	"net/http"

	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Categoria"
	"github.com/gin-gonic/gin"
)

// GetCategoriaTask godoc
// @Summary      Lista todas as categorias
// @Description  Retorna todas as categorias da base de dados
// @Tags         Categorias
// @Produce      json
// @Success      200 {array} models.ResponseExample
// @Failure      500 {object} models.ErrorExample "Falha ao listar as categorias"
// @Router       /category [get]
func GetCategoriaTask(ctx *gin.Context) {
	response, err := repository.GetAllCategoria()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": "Falha ao listar as categorias",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  response,
		"error": nil,
	})
}
