package main

import (
	"github.com/jonathas/gin-rest-api/database"
	"github.com/jonathas/gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
