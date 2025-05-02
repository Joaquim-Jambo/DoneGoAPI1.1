package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Categoria"
	"github.com/gin-gonic/gin"
)

// UpdateCategoryHandler godoc
// @Summary      Atualiza uma categoria existente
// @Description  Atualiza os detalhes de uma categoria com base no ID fornecido
// @Tags         Categorias
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da categoria a ser atualizada"
// @Param        categoria body models.CategoriaDTO true "Dados da categoria para atualização"
// @Success      200 {object} models.ResponseExample
// @Failure      400 {object} models.ErrorExample "JSON inválido"
// @Failure      404 {object} models.ErrorExample "Categoria não encontrada ou falha ao atualizar"
// @Router       /category/{id} [put]
func UpdateCategoryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, errId := strconv.Atoi(id)
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "ID inválido",
		})
		return
	}
	var categoria models.CategoriaDTO
	if err1 := ctx.BindJSON(&categoria); err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "JSON INVALIDO",
		})
		return
	}

	data, err := repository.UpdateCategory(id2, categoria)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Falha ao atualizar categoria", // Corrigi a mensagem de erro para ser mais específica
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  data,
		"error": nil,
	})

}
