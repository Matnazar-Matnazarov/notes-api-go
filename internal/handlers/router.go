package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Matnazar-Matnazarov/notes-api-go/internal/middleware"
)

// SetupRoutes configures all API routes
// This is where we define all endpoints and connect them to handlers
func SetupRoutes() *gin.Engine {
	// Create Gin router
	// Use default middleware (logger, recovery)
	router := gin.Default()

	// Add response time middleware to all routes
	router.Use(middleware.ResponseTime())

	// Health check endpoint
	// Useful for monitoring and checking if server is running
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Swagger documentation endpoint
	// Access Swagger UI at http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes group
	// All note-related endpoints are under /api/v1/notes
	api := router.Group("/api/v1")
	{
		noteHandler := NewNoteHandler()

		notes := api.Group("/notes")
		{
			// CRUD operations
			notes.POST("", noteHandler.CreateNote)       // Create a new note
			notes.GET("", noteHandler.GetAllNotes)       // Get all notes
			notes.GET("/:id", noteHandler.GetNoteByID)   // Get a single note by ID
			notes.PUT("/:id", noteHandler.UpdateNote)    // Update a note
			notes.DELETE("/:id", noteHandler.DeleteNote) // Delete a note
		}
	}

	return router
}
