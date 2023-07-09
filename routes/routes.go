package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathas/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/students", controllers.GetStudents)
	r.POST("/api/students", controllers.CreateStudent)
	r.GET("/api/:name", controllers.Greeting)
	r.Run()
}