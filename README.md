# ✅ DoneGoAPI

**DoneGoAPI** é uma API RESTful escrita em Go com o framework Gin, utilizada para gerenciar tarefas (todos). O projeto utiliza MySQL como banco de dados e está dockerizado para facilitar o desenvolvimento e a implantação. A documentação da API é gerada automaticamente com Swagger.

---

## 🛠️ Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)
- [Docker](https://www.docker.com/)
- [Swagger (Swaggo)](https://github.com/swaggo/swag)

---

## 🚀 Como Executar o Projeto

### Pré-requisitos

- Go instalado
- Docker e Docker Compose
- Git
### 1. Clonar o repositório

```bash
git clone https://github.com/Joaquim-Jambo/DoneGoAPI.git
cd DoneGoAPI
````

### 2. Configurar variáveis de ambiente

Crie um arquivo `.env` na raiz com o seguinte conteúdo:

```env
MYSQL_ROOT_PASSWORD=root
MYSQL_DATABASE=todolist
```

### 3. Subir os containers

```bash
docker-compose up -d
```

Isso iniciará um container com MySQL rodando na porta 3306.

### 4. Rodar as migrações e seeds (opcional)

Você pode executar as migrações (caso existam) e inserir dados de exemplo:

```bash
go run main.go
```


## 📄 Endpoints Principais

A documentação completa pode ser acessada via Swagger:

> 📚 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

### Exemplos:

| Método | Rota                   | Descrição                   |
| ------ | ---------------------- | --------------------------- |
| GET    | `/todo`                | Lista todas as tarefas      |
| GET    | `/todo/completed`      | Lista tarefas concluídas    |
| GET    | `/todo/{id}`           | Busca tarefa por ID         |
| POST   | `/todo`                | Cria uma nova tarefa        |
| PUT    | `/todo/{id}`           | Atualiza tarefa existente   |
| PATCH  | `/todo/{id}/completed` | Marca tarefa como concluída |
| DELETE | `/todo/{id}`           | Deleta tarefa por ID        |

---

## 🧪 Exemplo de Requisição (JSON)

```json
{
  "title": "Estudar Golang",
  "description": "Ler documentação do Gin",
  "completed": false
}
```

---

## 🐳 Sobre o Docker

O projeto já inclui suporte a Docker com o arquivo `docker-compose.yml` para rodar o MySQL

## ✅ Funcionalidades

* [x] CRUD de tarefas
* [x] Marcar tarefa como concluída
* [x] Filtro de tarefas concluídas
* [x] Integração com banco de dados
* [x] Swagger para documentação
* [x] Docker para banco de dados

---

## 👤 Autor

**Joaquim Fariti Nacassi Jambo**
📍 Luanda - Angola
📧 [joaquimjambo12@gmail.com](mailto:joaquimjambo12@gmail.com)
📱 +244 940 612 772
🔗 [LinkedIn](https://www.linkedin.com/in/Joaquim-Jambo)

---

## 📌 Observações

* Lembre-se de não subir seu `.env` para o repositório.
* Para ambientes de produção, utilize variáveis mais seguras.

---

## 🧠 Licença

Este projeto é licenciado sob a **MIT License**. Sinta-se à vontade para usar, modificar e distribuir.

```

---
