package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/studentAPI/models"
	"github.com/gin-gonic/gin"
)

func updateStudent(context *gin.Context) {
	studentID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	var newStudent models.Student
	err = context.ShouldBindJSON(&newStudent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	_, err = models.GetStudentByID(studentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "student id not found. Try again later."})
		return
	}

	newStudent.ID = studentID
	err = newStudent.Update()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update requested data"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Data is updated successfully!"})
}

func deleteStudent(context *gin.Context) {
	studentID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	student, err := models.GetStudentByID(studentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "student id not found. Try again later."})
		return
	}

	err = student.Delete()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete requested data"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Data is deleted successfully!"})
}

func createStudent(context *gin.Context) {
	var student models.Student
	err := context.ShouldBindJSON(&student)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	err = student.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not insert data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Record Inserted Successfully!"})
}

func getStudents(context *gin.Context) {
	students, err := models.GetAllStudents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch students. Try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"length": len(students), "message": students})
}

func getStudent(context *gin.Context) {
	studentID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	student, err := models.GetStudentByID(studentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch student data. Try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": student})
}
