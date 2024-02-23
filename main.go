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
	route.GET("/alltask", routes.GetAllTask)
	route.POST("/createTask", routes.CreateTask)
	route.DELETE("/deleteTask/:id", routes.DeleteTask)
	route.PUT("/updateTask/:id", routes.UpdateTask)
	route.GET("/getTask/:id", routes.GetTask)

	route.Run(":5000")
}
