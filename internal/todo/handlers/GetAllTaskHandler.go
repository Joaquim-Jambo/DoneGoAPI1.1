package handlers

import (
	"net/http"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

func GetAllTask(ctx *gin.Context) {
	response, err := repository.GetAllTask()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
