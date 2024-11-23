package config

import (
	"os"

	"github.com/pedropassos06/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	Connect() (*gorm.DB, error)
	Migrate(db *gorm.DB) error
}

type SQLite struct {
	Path   string
	Logger Logger
}

func (s *SQLite) Connect() (*gorm.DB, error) {
	// Check if the db exists
	_, err := os.Stat(s.Path)
	if os.IsNotExist(err) {
		s.Logger.Info("Database file not found, creating...")

		// Create the db file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(s.Path)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(s.Path), &gorm.Config{})
	if err != nil {
		s.Logger.Errf("SQLite opening error: %v", err)
		return nil, err
	}

	return db, nil
}

func (s *SQLite) Migrate(db *gorm.DB) error {
	// Migrate schema
	err := db.AutoMigrate(&schemas.Opening{}, &schemas.Resume{}, &schemas.NewsletterSubscription{})
	if err != nil {
		s.Logger.Errf("SQLite auto migrate error: %v", err)
		return err
	}
	return nil
}
