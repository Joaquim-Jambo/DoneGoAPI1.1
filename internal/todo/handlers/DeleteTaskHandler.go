package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

// DeleteTaskHandler godoc
// @Summary      Remove uma tarefa
// @Description  Exclui uma tarefa com base no ID
// @Tags         Tarefas
// @Produce      json
// @Param        id path int true "ID da tarefa"
// @Success      200 {object} models.ResponseExample "Tarefa deletada com sucesso"
// @Failure      400 {object} models.ErrorExample "ID inválido ou mal formado"
// @Failure      404 {object} models.ErrorExample "Tarefa não encontrada"
// @Router       /todo/{id} [delete]
func DeleteTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, _ := strconv.Atoi(id)
	data, err := repository.DeleteTask(id2)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Falha ao deletar tarefa",
		})
		return
	}
	ctx.JSON(http.StatusFound, gin.H{
		"data":  data,
		"error": nil,
	})
}
