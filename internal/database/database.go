package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Matnazar-Matnazarov/notes-api-go/internal/config"
	"github.com/Matnazar-Matnazarov/notes-api-go/internal/models"
)

var (
	// DB is the global database connection instance
	// This will be initialized once and reused throughout the application
	DB *gorm.DB
)

// Connect establishes a connection to the PostgreSQL database
// It reads configuration and sets up GORM with proper settings
func Connect(cfg *config.Config) error {
	dsn := cfg.GetDSN()
	
	// Configure GORM logger (only log errors and slow queries)
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}
	
	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	
	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	
	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)  // Maximum idle connections
	sqlDB.SetMaxOpenConns(100) // Maximum open connections
	
	DB = db
	log.Println("✅ Database connected successfully")
	
	return nil
}

// Migrate runs database migrations
// This creates/updates tables based on the models
func Migrate() error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	
	// Auto-migrate the Note model
	// This will create the table if it doesn't exist
	// and update it if the schema has changed
	if err := DB.AutoMigrate(&models.Note{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	
	log.Println("✅ Database migrations completed")
	return nil
}

// Close closes the database connection
// Should be called when the application shuts down
func Close() error {
	if DB == nil {
		return nil
	}
	
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	
	return sqlDB.Close()
}
