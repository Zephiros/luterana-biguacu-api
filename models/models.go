package models

// ContactForm represents a contact form submission
type ContactForm struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// ContactResponse is the response for contact form submission
type ContactResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// YouTubeLive represents a YouTube live stream
type YouTubeLive struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	VideoURL    string `json:"videoUrl"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime,omitempty"`
	IsLive      bool   `json:"isLive"`
}

// YouTubeResponse is the response for YouTube endpoints
type YouTubeResponse struct {
	Success bool         `json:"success"`
	Data    []YouTubeLive `json:"data,omitempty"`
	Error   string       `json:"error,omitempty"`
}

// Testimonial represents a testimonial entry
type Testimonial struct {
	ID          string `json:"id" gorm:"primary_key"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Date        string `json:"date" gorm:"not null"`
	Location    string `json:"location" gorm:"not null"`
	ImageURL    string `json:"imageUrl,omitempty" gorm:"column:image_url"`
	Created     string `json:"created" gorm:"not null"`
	Updated     string `json:"updated" gorm:"not null"`
}

// TestimonialResponse is the response for testimonial endpoints
type TestimonialResponse struct {
	Success bool         `json:"success"`
	Data    []Testimonial `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
	Error   string       `json:"error,omitempty"`
}
