package models

import (
	"time"

	"gorm.io/gorm"
)

// Note represents a note/reminder in the system
// This is the main entity that users can create, read, update, and delete
type Note struct {
	ID        uint           `json:"id" gorm:"primaryKey"`                    // Unique identifier
	Title     string         `json:"title" gorm:"not null;size:255"`          // Note title (required, max 255 chars)
	Content   string         `json:"content" gorm:"type:text"`                // Note content (can be long text)
	CreatedAt time.Time      `json:"created_at"`                             // When the note was created
	UpdatedAt time.Time      `json:"updated_at"`                             // When the note was last updated
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                         // Soft delete timestamp (hidden from JSON)
}

// TableName specifies the table name for GORM
func (Note) TableName() string {
	return "notes"
}

// CreateNoteRequest represents the data needed to create a new note
// Used for request validation
type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`   // Title is required, 1-255 characters
	Content string `json:"content" binding:"max=10000"`              // Content is optional, max 10000 characters
}

// UpdateNoteRequest represents the data needed to update an existing note
// All fields are optional - only provided fields will be updated
type UpdateNoteRequest struct {
	Title   *string `json:"title" binding:"omitempty,min=1,max=255"`   // Optional title update
	Content *string `json:"content" binding:"omitempty,max=10000"`      // Optional content update
}

// NoteResponse represents the note data sent to the client
// This is what the API returns in responses
type NoteResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse converts a Note model to NoteResponse
// This is useful for hiding internal fields (like DeletedAt) from the API
func (n *Note) ToResponse() *NoteResponse {
	return &NoteResponse{
		ID:        n.ID,
		Title:     n.Title,
		Content:   n.Content,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}
