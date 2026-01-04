package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Matnazar-Matnazarov/notes-api-go/internal/models"
	"github.com/Matnazar-Matnazarov/notes-api-go/internal/services"
	"github.com/gin-gonic/gin"
)

// getResponseTime calculates and returns response time in milliseconds
// It gets the start time from context and calculates the duration
func getResponseTime(c *gin.Context) int64 {
	startTime, exists := c.Get("start_time")
	if !exists {
		return 0
	}

	start, ok := startTime.(time.Time)
	if !ok {
		return 0
	}

	duration := time.Since(start)
	return duration.Milliseconds()
}

// NoteHandler handles HTTP requests related to notes
// This is the HTTP layer that receives requests and returns responses
type NoteHandler struct {
	service *services.NoteService
}

// NewNoteHandler creates a new instance of NoteHandler
func NewNoteHandler() *NoteHandler {
	return &NoteHandler{
		service: services.NewNoteService(),
	}
}

// CreateNote creates a new note from the request body
// @Summary      Create a new note
// @Description  Create a new note with title and content. Returns the created note with response time in milliseconds.
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        note  body      models.CreateNoteRequest  true  "Note data"
// @Success      201   {object}  models.SuccessResponse  "Note created successfully"
// @Failure      400   {object}  models.ErrorResponse     "Invalid request data"
// @Failure      500   {object}  models.ErrorResponse     "Internal server error"
// @Router       /notes [post]
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req models.CreateNoteRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Invalid request data",
			"details":          err.Error(),
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Create note using service
	note, err := h.service.CreateNote(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":            "Failed to create note",
			"details":          err.Error(),
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Return success response with created note and response time
	c.JSON(http.StatusCreated, gin.H{
		"message":          "Note created successfully",
		"data":             note.ToResponse(),
		"response_time_ms": getResponseTime(c),
	})
}

// GetAllNotes returns all notes in the database
// @Summary      Get all notes
// @Description  Retrieve all notes from the database. Returns list of notes with response time in milliseconds.
// @Tags         notes
// @Produce      json
// @Success      200  {object}  models.SuccessResponse  "Notes retrieved successfully"
// @Failure      500  {object}  models.ErrorResponse    "Internal server error"
// @Router       /notes [get]
func (h *NoteHandler) GetAllNotes(c *gin.Context) {
	// Get all notes using service
	notes, err := h.service.GetAllNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":            "Failed to get notes",
			"details":          err.Error(),
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Convert to response format
	responses := make([]*models.NoteResponse, len(notes))
	for i := range notes {
		responses[i] = notes[i].ToResponse()
	}

	// Return success response with response time
	c.JSON(http.StatusOK, gin.H{
		"message":          "Notes retrieved successfully",
		"data":             responses,
		"count":            len(responses),
		"response_time_ms": getResponseTime(c),
	})
}

// GetNoteByID returns a single note by its ID
// @Summary      Get a note by ID
// @Description  Retrieve a single note by its unique identifier. Returns note with response time in milliseconds.
// @Tags         notes
// @Produce      json
// @Param        id   path      int  true  "Note ID"
// @Success      200  {object}  models.SuccessResponse  "Note retrieved successfully"
// @Failure      400  {object}  models.ErrorResponse    "Invalid note ID"
// @Failure      404  {object}  models.ErrorResponse    "Note not found"
// @Router       /notes/{id} [get]
func (h *NoteHandler) GetNoteByID(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Invalid note ID",
			"details":          "ID must be a valid number",
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Get note using service
	note, err := h.service.GetNoteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":            "Note not found",
			"details":          err.Error(),
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Return success response with response time
	c.JSON(http.StatusOK, gin.H{
		"message":          "Note retrieved successfully",
		"data":             note.ToResponse(),
		"response_time_ms": getResponseTime(c),
	})
}

// UpdateNote updates an existing note
// @Summary      Update a note
// @Description  Update an existing note by ID. Only provided fields will be updated. Returns updated note with response time in milliseconds.
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id    path      int                     true  "Note ID"
// @Param        note  body      models.UpdateNoteRequest  true  "Note data to update"
// @Success      200   {object}  models.SuccessResponse    "Note updated successfully"
// @Failure      400   {object}  models.ErrorResponse      "Invalid request data"
// @Failure      404   {object}  models.ErrorResponse      "Note not found"
// @Router       /notes/{id} [put]
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Invalid note ID",
			"details":          "ID must be a valid number",
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	var req models.UpdateNoteRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Invalid request data",
			"details":          err.Error(),
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Update note using service
	note, err := h.service.UpdateNote(uint(id), &req)
	if err != nil {
		// Check if it's a "not found" error
		if err.Error() == "note not found: record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error":            "Note not found",
				"details":          err.Error(),
				"response_time_ms": getResponseTime(c),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Failed to update note",
			"details":          err.Error(),
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Return success response with response time
	c.JSON(http.StatusOK, gin.H{
		"message":          "Note updated successfully",
		"data":             note.ToResponse(),
		"response_time_ms": getResponseTime(c),
	})
}

// DeleteNote soft-deletes a note
// @Summary      Delete a note
// @Description  Soft-delete a note by ID (marks as deleted but doesn't remove from database). Returns response time in milliseconds.
// @Tags         notes
// @Produce      json
// @Param        id   path      int  true  "Note ID"
// @Success      200  {object}  models.SuccessResponse  "Note deleted successfully"
// @Failure      400  {object}  models.ErrorResponse    "Invalid note ID"
// @Failure      404  {object}  models.ErrorResponse    "Note not found"
// @Failure      500  {object}  models.ErrorResponse    "Internal server error"
// @Router       /notes/{id} [delete]
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	// Parse ID from URL parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Invalid note ID",
			"details":          "ID must be a valid number",
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Delete note using service
	if err := h.service.DeleteNote(uint(id)); err != nil {
		// Check if it's a "not found" error
		if err.Error() == "note not found: record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error":            "Note not found",
				"details":          err.Error(),
				"response_time_ms": getResponseTime(c),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":            "Failed to delete note",
			"details":          err.Error(),
			"response_time_ms": getResponseTime(c),
		})
		return
	}

	// Return success response with response time
	c.JSON(http.StatusOK, gin.H{
		"message":          "Note deleted successfully",
		"response_time_ms": getResponseTime(c),
	})
}
