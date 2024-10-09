// services/news_service.go
package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"myapp/models"
)

// API_KEY is your NewsAPI key (replace with your actual key)
const API_KEY = "YOUR_NEWS_API_KEY"

// FetchNewsFromAPI fetches the latest news from an external news API and maps it to the News model
func FetchNewsFromAPI() ([]models.News, error) {
	// Define the API endpoint for fetching news
	apiURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", API_KEY)

	// Make an HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, errors.New("failed to fetch news from API")
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("news API responded with an error")
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read news API response body")
	}

	// Define a struct to map the raw API response (as per the external API's structure)
	type APINewsResponse struct {
		Status  string `json:"status"`
		TotalResults int `json:"totalResults"`
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			URLToImage  string `json:"urlToImage"`
			Source      struct {
				Name string `json:"name"`
			} `json:"source"`
			PublishedAt string `json:"publishedAt"`
			Author      string `json:"author"`
			Content     string `json:"content"`
		} `json:"articles"`
	}

	// Parse the API response
	var apiResponse APINewsResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to parse news API response")
	}

	// Map the API response to the News model
	var newsData []models.News
	for _, article := range apiResponse.Articles {
		newsData = append(newsData, models.News{
			Title:       article.Title,
			Description: article.Description,
			URL:         article.URL,
			ImageURL:    article.URLToImage,
			Source:      article.Source.Name,
			PublishedAt: article.PublishedAt,
			Author:      article.Author,
			Category:    "General", // You could assign this dynamically or add categories based on tags
			Tags:        []string{}, // This could be populated if the API supports tags
			Content:     article.Content,
		})
	}

	return newsData, nil
}
