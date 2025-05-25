package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"luterana-biguacu-api/handlers"
)

func main() {
	// Create a default gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// API routes
	api := router.Group("/api")
	{
		// Content endpoints
		api.GET("/content", handlers.GetAllContent)
		api.GET("/content/:id", handlers.GetContentByID)
		api.GET("/content/type/:type", handlers.GetContentByType)

		// Contact form endpoint
		api.POST("/contact", handlers.SubmitContactForm)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start the server
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
