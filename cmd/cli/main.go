package main

import (
	"github.com/Joaquim-Jambo/DoneGoAPI/cmd/cli/config"
)

func main() {
	// Connect to the database
	config.Connect()

	// Migrate the schema
	// db.AutoMigrate(&models.Product{})
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Todo{})
}
