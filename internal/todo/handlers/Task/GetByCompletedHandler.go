package handlers

import (
	"net/http"

	_ "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Task"
	"github.com/gin-gonic/gin"
)

// GetByCompletedHandler godoc
// @Summary      Lista tarefas completadas
// @Description  Retorna todas as tarefas que foram marcadas como completas
// @Tags         Tarefas
// @Produce      json
// @Success      200 {array} models.ResponseExample
// @Failure      500 {object} models.ErrorExample "Erro ao buscar tarefas completadas"
// @Router       /todo/completed [get]
func GetByCompletedHandler(ctx *gin.Context) {
	data, err := repository.GetByCompleted()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": "Erro ao completar a tarefa", // Mensagem um pouco genérica, sugiro "Erro ao buscar tarefas completadas"
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  data,
	})
}
