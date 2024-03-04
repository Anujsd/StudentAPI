package main

import (
	"fmt"

	"example.com/studentAPI/db"
	"example.com/studentAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		fmt.Print("Not able to start server")
	}
}
