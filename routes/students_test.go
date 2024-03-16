package routes

// write test case for updateStudent
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update student",
	})
}
