package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"luterana-biguacu-api/models"
)

// YouTubeAPIBaseURL is the base URL for YouTube API
const YouTubeAPIBaseURL = "https://www.googleapis.com/youtube/v3"

// GetYouTubeLives returns live streams from the configured YouTube channel
func GetYouTubeLives(c *gin.Context) {
	// In a production environment, you would get API key and channel ID from environment variables
	// and make actual API calls to YouTube
	// For now, we'll just use sample data

	// In a real application, you would make an API call to YouTube
	// For now, we'll return some sample data
	// This would be replaced with actual API calls in production

	// Sample data for demonstration
	lives := []models.YouTubeLive{
		{
			ID:          "video123",
			Title:       "Sunday Service - Live Stream",
			Description: "Join us for our weekly Sunday service.",
			ThumbnailURL: "https://example.com/thumbnail1.jpg",
			VideoURL:    "https://www.youtube.com/watch?v=video123",
			StartTime:   time.Now().Format(time.RFC3339),
			IsLive:      true,
		},
		{
			ID:          "video456",
			Title:       "Bible Study - Live Stream",
			Description: "Weekly Bible study session.",
			ThumbnailURL: "https://example.com/thumbnail2.jpg",
			VideoURL:    "https://www.youtube.com/watch?v=video456",
			StartTime:   time.Now().AddDate(0, 0, -7).Format(time.RFC3339),
			EndTime:     time.Now().AddDate(0, 0, -7).Add(time.Hour * 2).Format(time.RFC3339),
			IsLive:      false,
		},
	}

	c.JSON(http.StatusOK, models.YouTubeResponse{
		Success: true,
		Data:    lives,
	})
}

// Note: In a production environment, you would implement a function to fetch live streams
// from the YouTube API using the channel ID and API key from environment variables.
// This would involve making HTTP requests to the YouTube API, parsing the JSON response,
// and converting it to our YouTubeLive model.
