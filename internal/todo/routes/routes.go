package routes

import (
	handlers "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/handlers/Categoria"
	handler "github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/handlers/Task"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Grupo de rotas para a API
	router := r.Group("/api/v1")
	{
		router.POST("/todo", handler.AddTaskHandler)
		router.GET("/todo", handler.GetAllTask)
		router.GET("/todo/:id", handler.GetByOneTaskHandler)
		router.PUT("/todo/:id", handler.UpdateTaskHandler)
		router.DELETE("/todo/:id", handler.DeleteTaskHandler)
		router.PATCH("/todo/:id/completed", handler.CompletedTaskHandler)
		router.GET("/todo/completed", handler.GetByCompletedHandler)

		router.POST("/category", handlers.AddTaskCategory)
		router.GET("/category", handlers.GetCategoriaTask)
		router.GET("/category/:id", handlers.GetByOneCategoriaHandler)
		router.PUT("/category/:id", handlers.UpdateCategoryHandler)
		router.DELETE("/category/:id", handlers.DeleteCategoriaHandler)
	}
}
