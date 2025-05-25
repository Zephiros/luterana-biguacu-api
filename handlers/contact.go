package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"luterana-biguacu-api/models"
)

// SubmitContactForm handles contact form submissions
func SubmitContactForm(c *gin.Context) {
	var form models.ContactForm

	// Bind JSON to the form struct and validate
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, models.ContactResponse{
			Success: false,
			Error:   "Invalid form data: " + err.Error(),
		})
		return
	}

	// In a real application, you would:
	// 1. Save the form data to a database
	// 2. Send an email notification
	// 3. Possibly trigger other actions

	// For now, we'll just return a success response
	c.JSON(http.StatusOK, models.ContactResponse{
		Success: true,
		Message: "Thank you for your message. We will get back to you soon!",
	})
}