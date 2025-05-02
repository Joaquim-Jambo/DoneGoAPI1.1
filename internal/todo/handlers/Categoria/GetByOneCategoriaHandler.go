package handlers

import (
	"net/http"
	"strconv"

	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Categoria"
	"github.com/gin-gonic/gin"
)

// GetByOneCategoriaHandler godoc
// @Summary      Busca uma categoria por ID
// @Description  Retorna os detalhes de uma categoria específica com base no seu ID
// @Tags         Categorias
// @Produce      json
// @Param        id path int true "ID da categoria a ser buscada"
// @Success      200 {object} models.ResponseExample
// @Failure      400 {object} models.ErrorExample "ID inválido"
// @Failure      404 {object} models.ErrorExample "Categoria não encontrada!"
// @Router       /category/{id} [get]
func GetByOneCategoriaHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, errId := strconv.Atoi(id)
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "ID inválido",
		})
		return
	}
	data, err := repository.GetByOneCategoria(id2)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Categoria não encontrada!",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  data,
		"error": nil,
	})
}
