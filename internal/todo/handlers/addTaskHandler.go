package handlers

import (
	"net/http"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

func AddTaskHandler(ctx *gin.Context) {
	var newTask models.TodoDTO

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "JSON INVALIDO",
		})
		return
	}
	if newTask.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "O campo 'title' é obrigatório",
		})
		return
	}
	response, err := repository.AddTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "Falha ao adicionar tarefa",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  response,
		"error": nil,
	})
}
