package routes

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Grupo de rotas para a API
	router := r.Group("/api/v1")
	{
		router.POST("/todo", handlers.AddTaskHandler)
		router.GET("/todo", handlers.GetAllTask)
		router.GET("/todo/:id", handlers.GetByOneTaskHandler)
		router.PUT("/todo/:id", handlers.UpdateTaskHandler)
		router.DELETE("/todo/:id", handlers.DeleteTaskHandler)
		router.PATCH("/todo/:id/completed", handlers.CompletedTaskHandler)
		router.GET("/todo/completed", handlers.GetByCompletedHandler)
	}
}
