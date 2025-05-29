package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello_world", helloWorldHandler)
	return router
}

func helloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func main() {
	// Create Gin router with default middleware
	router := getRouter()

	// Start the server
	log.Println("Server starting on :8080")
	router.Run(":8080")
}
