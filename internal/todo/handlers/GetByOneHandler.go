package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

func GetByOneTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, _ := strconv.Atoi(id)
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
