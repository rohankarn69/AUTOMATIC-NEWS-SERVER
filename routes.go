// routes.go
package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	// Create a new instance of the NewsController
	newsController := NewsController{}

	// Use the GetNews controller method to handle the /api/news GET route
	router.GET("/api/news", newsController.GetNews)
}
