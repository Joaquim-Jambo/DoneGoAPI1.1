package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

// CompletedTaskHandler godoc
// @Summary      Marca uma tarefa como concluída
// @Description  Atualiza o status de uma tarefa para concluída com base no ID
// @Tags         Tarefas
// @Produce      json
// @Param        id path int true "ID da tarefa"
// @Success      200 {object} models.ResponseExample
// @Failure      400 {object} models.ErrorExample "ID inválido"
// @Failure      404 {object} models.ErrorExample "Tarefa não encontrada"
// @Router       /todo/{id}/completed [patch]
func CompletedTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, errId := strconv.Atoi(id)
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"error":   "ID inválido",
		})
		return
	}
	data, err := repository.CompletedTask(id2)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"error":   "Tarefa não encontrada",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"error":   nil,
	})
}
