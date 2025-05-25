package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"luterana-biguacu-api/models"
)

// GetAllContent returns all content
func GetAllContent(c *gin.Context) {
	// In a real application, this would fetch from a database
	// For now, we'll return some sample data
	contents := []models.Content{
		{
			ID:      "1",
			Title:   "Welcome to our Church",
			Body:    "This is the main page of our church website.",
			Type:    "page",
			Created: time.Now().AddDate(0, 0, -30).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
		{
			ID:      "2",
			Title:   "Upcoming Events",
			Body:    "Check out our upcoming events and activities.",
			Type:    "page",
			Created: time.Now().AddDate(0, 0, -15).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
		{
			ID:      "3",
			Title:   "Sunday Service",
			Body:    "Join us for our Sunday service at 10:00 AM.",
			Type:    "event",
			Created: time.Now().AddDate(0, 0, -7).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
	}

	c.JSON(http.StatusOK, models.ContentResponse{
		Success: true,
		Data:    contents,
	})
}

// GetContentByID returns content by ID
func GetContentByID(c *gin.Context) {
	id := c.Param("id")

	// In a real application, this would fetch from a database
	// For now, we'll return sample data based on the ID
	var content models.Content
	found := false

	// Sample data
	contents := []models.Content{
		{
			ID:      "1",
			Title:   "Welcome to our Church",
			Body:    "This is the main page of our church website.",
			Type:    "page",
			Created: time.Now().AddDate(0, 0, -30).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
		{
			ID:      "2",
			Title:   "Upcoming Events",
			Body:    "Check out our upcoming events and activities.",
			Type:    "page",
			Created: time.Now().AddDate(0, 0, -15).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
		{
			ID:      "3",
			Title:   "Sunday Service",
			Body:    "Join us for our Sunday service at 10:00 AM.",
			Type:    "event",
			Created: time.Now().AddDate(0, 0, -7).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
	}

	for _, c := range contents {
		if c.ID == id {
			content = c
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, models.ContentResponse{
			Success: false,
			Error:   "Content not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.ContentResponse{
		Success: true,
		Data:    []models.Content{content},
	})
}

// GetContentByType returns content by type
func GetContentByType(c *gin.Context) {
	contentType := c.Param("type")

	// In a real application, this would fetch from a database
	// For now, we'll return sample data based on the type
	var filteredContents []models.Content

	// Sample data
	contents := []models.Content{
		{
			ID:      "1",
			Title:   "Welcome to our Church",
			Body:    "This is the main page of our church website.",
			Type:    "page",
			Created: time.Now().AddDate(0, 0, -30).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
		{
			ID:      "2",
			Title:   "Upcoming Events",
			Body:    "Check out our upcoming events and activities.",
			Type:    "page",
			Created: time.Now().AddDate(0, 0, -15).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
		{
			ID:      "3",
			Title:   "Sunday Service",
			Body:    "Join us for our Sunday service at 10:00 AM.",
			Type:    "event",
			Created: time.Now().AddDate(0, 0, -7).Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
	}

	for _, c := range contents {
		if c.Type == contentType {
			filteredContents = append(filteredContents, c)
		}
	}

	if len(filteredContents) == 0 {
		c.JSON(http.StatusOK, models.ContentResponse{
			Success: true,
			Data:    []models.Content{},
		})
		return
	}

	c.JSON(http.StatusOK, models.ContentResponse{
		Success: true,
		Data:    filteredContents,
	})
}