// models/news.go
package models

// News represents a detailed structure for a news article
type News struct {
	Title       string   `json:"title"`        // The title of the news article
	Description string   `json:"description"`  // A brief summary of the news article
	URL         string   `json:"url"`          // The URL to the full news article
	ImageURL    string   `json:"image_url"`    // The URL of the featured image
	Source      string   `json:"source"`       // The news source (e.g., BBC, CNN)
	PublishedAt string   `json:"published_at"` // The publication date of the article
	Author      string   `json:"author"`       // The author of the news article
	Category    string   `json:"category"`     // The category of the news (e.g., Politics, Sports)
	Tags        []string `json:"tags"`         // A list of tags associated with the article
	Content     string   `json:"content"`      // The full content of the article
}
