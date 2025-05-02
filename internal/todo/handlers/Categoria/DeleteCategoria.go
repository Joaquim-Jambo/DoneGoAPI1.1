package handlers

import (
	"net/http"
	"strconv"

	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Categoria"
	"github.com/gin-gonic/gin"
)

// DeleteCategoriaHandler godoc
// @Summary      Remove uma categoria
// @Description  Exclui uma categoria com base no ID
// @Tags         Categorias
// @Produce      json
// @Param        id path int true "ID da categoria"
// @Success      200 {object} models.ResponseExample "Categoria deletada com sucesso"
// @Failure      400 {object} models.ErrorExample "ID inválido ou mal formado"
// @Failure      404 {object} models.ErrorExample "Categoria não encontrada"
// @Router       /category/{id} [delete]
func DeleteCategoriaHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, err1 := strconv.Atoi(id)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "ID inválido ou mal formado",
		})
		return
	}
	data, err := repository.DeleteCategoria(id2)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Categoria não encontrada",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"sucess": "Categoria deletada com sucesso",
		"data":   data,
		"error":  nil,
	})
}
