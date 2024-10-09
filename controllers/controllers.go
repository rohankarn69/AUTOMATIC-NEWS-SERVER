// controllers.go
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"myapp/models"
)

// NewsController handles the logic for the news API
type NewsController struct{}

// GetNews handles GET requests for the /api/news route
func (ctrl NewsController) GetNews(c *gin.Context) {
	// Mock data with enhanced News schema
	newsData := []models.News{
		{
			Title:       "Breaking News 1",
			Description: "Description for breaking news 1",
			URL:         "https://news1.com",
			ImageURL:    "https://news1.com/image.jpg",
			Source:      "News Source 1",
			PublishedAt: "2024-10-09",
			Author:      "Author 1",
			Category:    "World",
			Tags:        []string{"Breaking", "World", "Politics"},
			Content:     "This is the full content for the first news article.",
		},
		{
			Title:       "Breaking News 2",
			Description: "Description for breaking news 2",
			URL:         "https://news2.com",
			ImageURL:    "https://news2.com/image.jpg",
			Source:      "News Source 2",
			PublishedAt: "2024-10-08",
			Author:      "Author 2",
			Category:    "Technology",
			Tags:        []string{"Tech", "Innovation", "Startups"},
			Content:     "This is the full content for the second news article.",
		},
	}

	// Return the news data as JSON
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newsData,
	})
}
