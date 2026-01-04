package services

import (
	"errors"
	"fmt"

	"github.com/Matnazar-Matnazarov/notes-api-go/internal/database"
	"github.com/Matnazar-Matnazarov/notes-api-go/internal/models"
)

// NoteService handles all business logic related to notes
// This layer separates business logic from HTTP handlers
type NoteService struct{}

// NewNoteService creates a new instance of NoteService
func NewNoteService() *NoteService {
	return &NoteService{}
}

// CreateNote creates a new note in the database
// Returns the created note or an error
func (s *NoteService) CreateNote(req *models.CreateNoteRequest) (*models.Note, error) {
	// Validate input
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	
	if len(req.Title) > 255 {
		return nil, errors.New("title must be 255 characters or less")
	}
	
	if len(req.Content) > 10000 {
		return nil, errors.New("content must be 10000 characters or less")
	}
	
	// Create note model
	note := &models.Note{
		Title:   req.Title,
		Content: req.Content,
	}
	
	// Save to database
	if err := database.DB.Create(note).Error; err != nil {
		return nil, fmt.Errorf("failed to create note: %w", err)
	}
	
	return note, nil
}

// GetAllNotes retrieves all notes from the database
// Returns a slice of notes or an error
func (s *NoteService) GetAllNotes() ([]models.Note, error) {
	var notes []models.Note
	
	// Find all notes (excluding soft-deleted ones)
	if err := database.DB.Find(&notes).Error; err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}
	
	return notes, nil
}

// GetNoteByID retrieves a single note by its ID
// Returns the note or an error if not found
func (s *NoteService) GetNoteByID(id uint) (*models.Note, error) {
	var note models.Note
	
	// Find note by ID (excluding soft-deleted ones)
	if err := database.DB.First(&note, id).Error; err != nil {
		return nil, fmt.Errorf("note not found: %w", err)
	}
	
	return &note, nil
}

// UpdateNote updates an existing note
// Only provided fields will be updated
// Returns the updated note or an error
func (s *NoteService) UpdateNote(id uint, req *models.UpdateNoteRequest) (*models.Note, error) {
	// First, check if note exists
	note, err := s.GetNoteByID(id)
	if err != nil {
		return nil, err
	}
	
	// Update fields if provided
	if req.Title != nil {
		if *req.Title == "" {
			return nil, errors.New("title cannot be empty")
		}
		if len(*req.Title) > 255 {
			return nil, errors.New("title must be 255 characters or less")
		}
		note.Title = *req.Title
	}
	
	if req.Content != nil {
		if len(*req.Content) > 10000 {
			return nil, errors.New("content must be 10000 characters or less")
		}
		note.Content = *req.Content
	}
	
	// Save changes to database
	if err := database.DB.Save(note).Error; err != nil {
		return nil, fmt.Errorf("failed to update note: %w", err)
	}
	
	return note, nil
}

// DeleteNote soft-deletes a note (marks it as deleted but doesn't remove from DB)
// Returns an error if note not found
func (s *NoteService) DeleteNote(id uint) error {
	// Check if note exists
	_, err := s.GetNoteByID(id)
	if err != nil {
		return err
	}
	
	// Soft delete (GORM will set DeletedAt timestamp)
	if err := database.DB.Delete(&models.Note{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete note: %w", err)
	}
	
	return nil
}
