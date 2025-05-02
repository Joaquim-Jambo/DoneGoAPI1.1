package handlers

import (
	"net/http"
	"strconv"

	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	repository "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/repository/Task"
	"github.com/gin-gonic/gin"
)

// UpdateTaskHandler godoc
// @Summary      Atualiza uma tarefa existente
// @Description  Atualiza os detalhes de uma tarefa com base no ID fornecido
// @Tags         Tarefas
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da tarefa a ser atualizada"
// @Param        tarefa body models.TodoUpdateDTO true "Dados da tarefa para atualização"
// @Success      200 {object} models.ResponseExample
// @Failure      400 {object} models.ErrorExample "JSON inválido"
// @Failure      404 {object} models.ErrorExample "Tarefa não encontrada ou falha ao atualizar"
// @Router       /todo/{id} [put]
func UpdateTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, errId := strconv.Atoi(id)
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "ID inválido",
		})
		return
	}
	var task models.TodoUpdateDTO
	if err1 := ctx.BindJSON(&task); err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": "JSON INVALIDO",
		})
		return
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
