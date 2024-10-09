// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Set up routes from another file
	setupRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")
}
