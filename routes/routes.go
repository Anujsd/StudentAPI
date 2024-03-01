package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/v1/students", getStudents)
	server.POST("/v1/student", createStudent)
	server.GET("/v1/student/:id", getStudent)
	server.DELETE("/v1/student/:id", deleteStudent)
	server.PUT("/v1/student/:id", updateStudent)
}
