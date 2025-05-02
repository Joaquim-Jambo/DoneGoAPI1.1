package handlers

import (
	"net/http"

	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Task"
	"github.com/gin-gonic/gin"
)

// GetAllTask godoc
// @Summary      Lista todas as tarefas
// @Description  Retorna todas as tarefas da base de dados
// @Tags         Tarefas
// @Produce      json
// @Success      200 {array} models.ResponseExample
// @Failure      500 {object} models.ErrorExample "Falha ao listar as tarefas"
// @Router       /todo [get]
func GetAllTask(ctx *gin.Context) {
	response, err := repository.GetAllTask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": "Falha ao listar os todos",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  response,
		"error": nil,
	})
}
