package main

import (
	"fmt"
	"net/http"

	config "github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/routes"
)

func main() {
	// Connect to the database
	config.Connect()
	config.Db.AutoMigrate(&models.Todo{})
	routes.Initialize()

	data, err := http.Get("http://localhost:8080/api/v1/todo")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
