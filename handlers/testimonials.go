package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"luterana-biguacu-api/database"
	"luterana-biguacu-api/models"
	"net/http"
	"time"
)

// GetAllTestimonials returns all testimonials
func GetAllTestimonials(c *gin.Context) {
	var testimonials []models.Testimonial
	if err := database.DB.Find(&testimonials).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.TestimonialResponse{
			Success: false,
			Error:   "Failed to fetch testimonials: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.TestimonialResponse{
		Success: true,
		Data:    testimonials,
	})
}

// GetTestimonialByID returns a testimonial by ID
func GetTestimonialByID(c *gin.Context) {
	id := c.Param("id")
	var testimonial models.Testimonial

	if err := database.DB.Where("id = ?", id).First(&testimonial).Error; err != nil {
		c.JSON(http.StatusNotFound, models.TestimonialResponse{
			Success: false,
			Error:   "Testimonial not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.TestimonialResponse{
		Success: true,
		Data:    []models.Testimonial{testimonial},
	})
}

// CreateTestimonial creates a new testimonial
func CreateTestimonial(c *gin.Context) {
	var testimonial models.Testimonial

	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, models.TestimonialResponse{
			Success: false,
			Error:   "Invalid testimonial data: " + err.Error(),
		})
		return
	}

	// Generate a new ID (simple timestamp-based ID)
	testimonial.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	// Set created and updated timestamps
	now := time.Now().Format(time.RFC3339)
	testimonial.Created = now
	testimonial.Updated = now

	// Save to database
	if err := database.DB.Create(&testimonial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.TestimonialResponse{
			Success: false,
			Error:   "Failed to create testimonial: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.TestimonialResponse{
		Success: true,
		Data:    []models.Testimonial{testimonial},
	})
}

// UpdateTestimonial updates an existing testimonial
func UpdateTestimonial(c *gin.Context) {
	id := c.Param("id")
	var testimonial models.Testimonial
	var updatedTestimonial models.Testimonial

	// Check if testimonial exists
	if err := database.DB.Where("id = ?", id).First(&testimonial).Error; err != nil {
		c.JSON(http.StatusNotFound, models.TestimonialResponse{
			Success: false,
			Error:   "Testimonial not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&updatedTestimonial); err != nil {
		c.JSON(http.StatusBadRequest, models.TestimonialResponse{
			Success: false,
			Error:   "Invalid testimonial data: " + err.Error(),
		})
		return
	}

	// Preserve the original ID and created timestamp
	updatedTestimonial.ID = testimonial.ID
	updatedTestimonial.Created = testimonial.Created
	updatedTestimonial.Updated = time.Now().Format(time.RFC3339)

	// Update in database
	if err := database.DB.Model(&testimonial).Updates(updatedTestimonial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.TestimonialResponse{
			Success: false,
			Error:   "Failed to update testimonial: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.TestimonialResponse{
		Success: true,
		Data:    []models.Testimonial{updatedTestimonial},
	})
}

// DeleteTestimonial deletes a testimonial
func DeleteTestimonial(c *gin.Context) {
	id := c.Param("id")
	var testimonial models.Testimonial

	// Check if testimonial exists
	if err := database.DB.Where("id = ?", id).First(&testimonial).Error; err != nil {
		c.JSON(http.StatusNotFound, models.TestimonialResponse{
			Success: false,
			Error:   "Testimonial not found",
		})
		return
	}

	// Delete from database
	if err := database.DB.Delete(&testimonial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.TestimonialResponse{
			Success: false,
			Error:   "Failed to delete testimonial: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.TestimonialResponse{
		Success: true,
		Message: "Testimonial deleted successfully",
	})
}