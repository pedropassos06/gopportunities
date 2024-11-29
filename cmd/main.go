package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/auth"
	"github.com/pedropassos06/gopportunities/config"
	"github.com/pedropassos06/gopportunities/newsletter"
	"github.com/pedropassos06/gopportunities/opening"
	"github.com/pedropassos06/gopportunities/resume"
	"github.com/pedropassos06/gopportunities/router"
	"github.com/pedropassos06/gopportunities/utils"
)

var (
	dbPath = "./db/main.db"
)

func init() {
	// Load .env file
	utils.LoadEnv()
}

func main() {
	// Create Logger
	logger := config.GetLogger("main")

	// Create SQLite instance
	sqlite := &config.SQLite{
		Path:   dbPath,
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

	ginRouter := gin.Default()
	// init specific handlers
	authHandler := auth.NewAuthHandler(config.GetSQLite(), logger)
	newsletterHandler := newsletter.NewNewsletterHandler(config.GetSQLite(), logger)
	resumeHandler := resume.NewResumeHandler(config.GetSQLite(), logger)

	openingRepo := opening.NewOpeningRepository(config.GetSQLite())
	openingUsecase := opening.NewOpeningUsecase(openingRepo)
	openingHandler := opening.NewOpeningHandler(openingUsecase)

	// Initialize Router
	router.InitializeRoutes(ginRouter, authHandler, resumeHandler, openingHandler, newsletterHandler)
}
