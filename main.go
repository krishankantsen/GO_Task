package main

import (
	"GOTASK/models"
	"GOTASK/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()
	// Continue with your Gin routes and other setup
	route.GET("/tasks", routes.GetAllTask)
	route.POST("/tasks", routes.CreateTask)
	route.DELETE("/tasks/:id", routes.DeleteTask)
	route.PUT("/tasks/:id", routes.UpdateTask)
	route.GET("/tasks/:id", routes.GetTask)

	route.Run(":5000")
}
