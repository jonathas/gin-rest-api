package main

import (
	"github.com/jonathas/gin-rest-api/database"
	"github.com/jonathas/gin-rest-api/models"
	"github.com/jonathas/gin-rest-api/routes"
)

func main() {
	database.Connect()

	models.Students = []models.Student{
		{Name: "Jonathas", Document: "123456789"},
		{Name: "João", Document: "987654321"},
	}
	routes.HandleRequests()
}
