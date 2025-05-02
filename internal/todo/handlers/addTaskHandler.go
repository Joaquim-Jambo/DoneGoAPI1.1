package handlers

import (
	"net/http"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository"
	"github.com/gin-gonic/gin"
)

// AddTaskHandler godoc
// @Summary      Cria uma nova tarefa
// @Description  Adiciona uma nova tarefa à base de dados
// @Tags         Tarefas
// @Accept       json
// @Produce      json
// @Param        tarefa body models.TodoDTO true "Dados da tarefa"
// @Success      200 {object} models.ResponseExample
// @Failure      400 {object} models.ErrorExample "JSON inválido"
// @Failure      404 {object} models.ErrorExample "Falha ao adicionar tarefa"
// @Router       /todo [post]
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
