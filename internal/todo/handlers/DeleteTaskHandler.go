package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

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
