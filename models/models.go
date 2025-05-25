package models

// Content represents website content
type Content struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Type    string `json:"type"` // e.g., "page", "post", "event"
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// ContactForm represents a contact form submission
type ContactForm struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// ContentResponse is the response for content endpoints
type ContentResponse struct {
	Success bool     `json:"success"`
	Data    []Content `json:"data,omitempty"`
	Error   string   `json:"error,omitempty"`
}

// ContactResponse is the response for contact form submission
type ContactResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}