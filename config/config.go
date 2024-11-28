package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger Logger
)

type Config struct {
	DB     Database
	Logger Logger
}

func Init(config Config) error {
	var err error

	// Initialize the logger
	logger = config.Logger

	// Initialize the database
	db, err = config.DB.Connect()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// Migrate the schema
	err = config.DB.Migrate(db)
	if err != nil {
		return fmt.Errorf("error migrating database: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) Logger {
	logger = NewLogger(p)
	return logger
}
