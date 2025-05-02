package main

import (
	config "github.com/Joaquim-Jambo/DoneGoAPI/config"
	"github.com/Joaquim-Jambo/DoneGoAPI/internal/todo/models"
)

func main() {
	config.Connect()
	var tarefas = []models.Todo{
		{Title: "Estudar Go", Description: "Ler documentação oficial do Go e praticar com exercícios.", Completed: false},
		{Title: "Revisar projeto DoneGoAPI", Description: "Verificar se todos os handlers estão documentados e funcionando.", Completed: false},
		{Title: "Testar Swagger", Description: "Acessar a interface Swagger e testar todos os endpoints.", Completed: true},
		{Title: "Preparar vídeo para o LinkedIn", Description: "Gravar apresentação do projeto e publicar no perfil profissional.", Completed: false},
		{Title: "Subir projeto no GitHub", Description: "Organizar repositório, criar README e .gitignore corretamente.", Completed: true},
	}

	for _, tarefa := range tarefas {
		config.Db.Create(&tarefa)
	}

}
