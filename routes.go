// routes.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRoutes(router *gin.Engine) {
	// Define a simple GET request handler
	router.GET("/api/news", func(c *gin.Context) {
		// This could return mock data for now
		c.JSON(http.StatusOK, gin.H{
			"message": "News API working!",
		})
	})
}
