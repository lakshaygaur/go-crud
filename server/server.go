package server

import (
	"github.com/gin-gonic/gin"
)

const PORT = "8000"

func StartServer() {
	server := gin.Default()

	server.GET("/employee", GetEmployeeByID)
	server.POST("/employee", CreateEmployee)
	server.DELETE("/employee", DeleteEmployeeByID)

	server.Run("0.0.0.0:" + PORT) // listen and serve on 0.0.0.0:8080
}
