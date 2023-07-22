package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathas/gin-rest-api/database"
	"github.com/jonathas/gin-rest-api/models"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
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

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted!"})
}

func GetStudentByName(c *gin.Context) {
	var student models.Student
	name := c.Param("name")

	database.DB.Where(&models.Student{Name: name}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	c.JSON(http.StatusOK, student)
}
