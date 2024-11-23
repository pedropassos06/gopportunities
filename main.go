package main

import (
	"github.com/pedropassos06/gopportunities/config"
	"github.com/pedropassos06/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	// Create Logger
	logger := config.GetLogger("main")

	// Create SQLite instance
	sqlite := &config.SQLite{
		Path:   "./db/main.db",
		Logger: logger,
	}

	// Initialize configs
	err := config.Init(config.Config{
		DB:     sqlite,
		Logger: logger,
	})
	if err != nil {
		logger.Err(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
