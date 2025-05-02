package handlers

import (
	"net/http"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

func GetByCompletedHandler(ctx *gin.Context) {
	data, err := repository.GetByCompleted()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Erro ao completar a tarefa",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  data,
	})
}
