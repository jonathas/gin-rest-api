package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathas/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.POST("/students", controllers.CreateStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/:id", controllers.GetStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/students/name/:name", controllers.GetStudentsByName)
	r.GET("/:name", controllers.Greeting)
	r.NoRoute((controllers.EndpointNotFound))
	r.Run()
}