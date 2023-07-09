package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathas/gin-rest-api/database"
	"github.com/jonathas/gin-rest-api/models"
)

func GetStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Param("name")

	c.JSON(200, gin.H{
		"message": "Hello " + name + "!",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}
