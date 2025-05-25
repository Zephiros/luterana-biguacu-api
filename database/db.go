package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"luterana-biguacu-api/models"
)

// DB is the database connection
var DB *gorm.DB

// Connect establishes a connection to the database
func Connect() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection
	var err error
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Enable Logger
	DB.LogMode(true)

	// Auto Migrate
	DB.AutoMigrate(&models.Testimonial{})

	// Seed data if the testimonials table is empty
	var count int
	DB.Model(&models.Testimonial{}).Count(&count)
	if count == 0 {
		log.Println("Seeding testimonials data...")
		seedTestimonials()
	}

	log.Println("Connected to database")
	return nil
}

// Close closes the database connection
func Close() error {
	return DB.Close()
}

// seedTestimonials seeds the database with some testimonials
func seedTestimonials() {
	testimonials := []models.Testimonial{
		{
			ID:          "1",
			Title:       "Community Outreach",
			Description: "Helping those in need in our community.",
			Date:        time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
			Location:    "Downtown Community Center",
			ImageURL:    "https://example.com/images/outreach.jpg",
			Created:     time.Now().AddDate(0, 0, -5).Format(time.RFC3339),
			Updated:     time.Now().Format(time.RFC3339),
		},
		{
			ID:          "2",
			Title:       "Youth Group Meeting",
			Description: "Weekly gathering for our church youth.",
			Date:        time.Now().AddDate(0, 0, 3).Format("2006-01-02"),
			Location:    "Church Hall",
			ImageURL:    "https://example.com/images/youth.jpg",
			Created:     time.Now().AddDate(0, 0, -2).Format(time.RFC3339),
			Updated:     time.Now().Format(time.RFC3339),
		},
	}

	for _, testimonial := range testimonials {
		DB.Create(&testimonial)
	}
}
