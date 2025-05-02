package main

import (
	config "github.com/Joaquim-Jambo/DoneGoAPI/config"
	_ "github.com/Joaquim-Jambo/DoneGoAPI/docs"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/routes"
)

//	@title			DoneGoAPI
//	@version		1.0
//	@description	API para gerenciar tarefas e listas de tarefas.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Joaquim Jambo
//	@contact.url	https://www.linkedin.com/in/joaquimjambo
//	@contact.email	joaquimjambo12@gmail.com

//	@license.name	Joaquim Jambo
//	@license.url	https://www.linkedin.com/in/joaquimjambo

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	// Connect to the database
	config.Connect()
	config.Db.AutoMigrate(&models.Todo{})
	routes.Initialize()

}
