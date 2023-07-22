package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jonathas/gin-rest-api/controllers"
	"github.com/jonathas/gin-rest-api/database"
	"github.com/jonathas/gin-rest-api/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func createMockStudent() {
	student := models.Student{
		Name: "Jonathas",
		Age:  23,
		Document: "12345678901",
	}

	database.DB.Create(&student)
	ID = int(student.ID)
}

func deleteMockStudent() {
	database.DB.Delete(&models.Student{}, ID)
}

func TestVerifyGreetingStatusCode(t *testing.T) {
	r := setup()

	r.GET("/:name", controllers.Greeting)

	req, _ := http.NewRequest("GET", "/jonathas", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "OK response is expected")

	resMock := `{"message":"Hello jonathas!"}`

	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, resMock, string(resBody), "Response body should contain the mock")
}

func TestListAllStudents(t *testing.T) {
	database.Connect()
	createMockStudent()
	defer deleteMockStudent() // This will run after the test is done

	r := setup()

	r.GET("/students", controllers.GetStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "OK response is expected")
}

func TestGetStudentsByName(t *testing.T) {
	database.Connect()
	createMockStudent()
	defer deleteMockStudent()

	r := setup()

	r.GET("/students/name/:name", controllers.GetStudentsByName)

	req, _ := http.NewRequest("GET", "/students/name/Jonathas", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Contains(t, res.Body.String(), "Jonathas")
	assert.Contains(t, res.Body.String(), "23")
	assert.Contains(t, res.Body.String(), "12345678901")
}

func TestGetStudent(t *testing.T) {
	database.Connect()
	createMockStudent()
	defer deleteMockStudent()

	r := setup()

	r.GET("/students/:id", controllers.GetStudent)

	req, _ := http.NewRequest("GET", "/students/" + strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var mockedStudent models.Student
	json.Unmarshal(res.Body.Bytes(), &mockedStudent)

	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Equal(t, "Jonathas", mockedStudent.Name)
	assert.Equal(t, 23, mockedStudent.Age)
	assert.Equal(t, "12345678901", mockedStudent.Document)
}

func TestCreateStudent(t *testing.T) {
	database.Connect()
	defer deleteMockStudent()

	r := setup()

	r.POST("/students", controllers.CreateStudent)

	student := models.Student{
		Name: "Jonathas",
		Age:  23,
		Document: "12345678901",
	}

	studentJSON, _ := json.Marshal(student)

	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(studentJSON))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var mockedStudent models.Student
	json.Unmarshal(res.Body.Bytes(), &mockedStudent)

	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Equal(t, "Jonathas", mockedStudent.Name)
	assert.Equal(t, 23, mockedStudent.Age)
	assert.Equal(t, "12345678901", mockedStudent.Document)
}

func TestUpdateStudent(t *testing.T) {
	database.Connect()
	createMockStudent()
	defer deleteMockStudent()

	r := setup()

	r.PATCH("/students/:id", controllers.UpdateStudent)

	student := models.Student{
		Name: "Jonathas",
		Age:  25,
		Document: "12345678901",
	}

	studentJSON, _ := json.Marshal(student)

	req, _ := http.NewRequest("PATCH", "/students/" + strconv.Itoa(ID), bytes.NewBuffer(studentJSON))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var mockedStudent models.Student
	json.Unmarshal(res.Body.Bytes(), &mockedStudent)

	assert.Equal(t, 200, res.Code, "OK response is expected")
	assert.Equal(t, "Jonathas", mockedStudent.Name)
	assert.Equal(t, 25, mockedStudent.Age)
	assert.Equal(t, "12345678901", mockedStudent.Document)
}

func TestDeleteStudent(t *testing.T) {
	database.Connect()
	createMockStudent()

	r := setup()

	r.DELETE("/students/:id", controllers.DeleteStudent)

	req, _ := http.NewRequest("DELETE", "/students/" + strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "OK response is expected")
}
