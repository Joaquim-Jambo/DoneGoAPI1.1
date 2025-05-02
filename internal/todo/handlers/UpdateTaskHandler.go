package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, _ := strconv.Atoi(id)
	var task models.TodoUpdateDTO
	if err1 := ctx.BindJSON(&task); err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "JSON INVALIDO",
		})
	}
	data, err := repository.UpdateTask(id2, task)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Falha ao atualizar tarefa",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  data,
		"error": nil,
	})

}
