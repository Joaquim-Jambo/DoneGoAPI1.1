package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

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
