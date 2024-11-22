package config

import (
	"os"

	"github.com/pedropassos06/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// Get Logger
	logger := GetLogger("SQLite")
	dbPath := "./db/main.db"

	// Check if the db exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")

		// Create the db file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("SQLite opening error: %v", err)
		return nil, err
	}

	// Migrate schema
	err = db.AutoMigrate(&schemas.Opening{}, &schemas.Resume{})
	if err != nil {
		logger.Errf("SQLite auto migrate error: %v", err)
		return nil, err
	}

	// Return db
	return db, nil
}
