package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/config"
	"github.com/pedropassos06/gopportunities/handler"
	"github.com/pedropassos06/gopportunities/router"
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

	handler := handler.NewHandler(config.GetSQLite(), logger)
	ginRouter := gin.Default()

	// Initialize Router
	router.InitializeRoutes(ginRouter, handler)
}
