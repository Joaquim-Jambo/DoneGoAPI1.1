package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

// GetByOneTaskHandler godoc
// @Summary      Busca uma tarefa por ID
// @Description  Retorna os detalhes de uma tarefa específica com base no seu ID
// @Tags         Tarefas
// @Produce      json
// @Param        id path int true "ID da tarefa a ser buscada"
// @Success      200 {object} models.ResponseExample
// @Failure      400 {object} models.ErrorExample "ID inválido"
// @Failure      404 {object} models.ErrorExample "Tarefa não encontrada!"
// @Router       /todo/{id} [get]
func GetByOneTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, errId := strconv.Atoi(id)
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "ID inválido",
		})
		return
	}
	data, err := repository.GetByOneTask(id2)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Tarefa não encontrada!",
		})
		return
	}
	ctx.JSON(http.StatusFound, gin.H{
		"data":  data,
		"error": nil,
	})
}
