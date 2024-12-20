package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/auth"
	"github.com/pedropassos06/gopportunities/config"
	"github.com/pedropassos06/gopportunities/newsletter"
	"github.com/pedropassos06/gopportunities/opening"
	"github.com/pedropassos06/gopportunities/resume"
	"github.com/pedropassos06/gopportunities/router"
	"github.com/pedropassos06/gopportunities/user"
	"github.com/pedropassos06/gopportunities/utils"
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
		Path:   os.Getenv("DB_PATH"),
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
	// user handler
	userRepo := user.NewUserRepository(config.GetSQLite())
	userUsecase := user.NewUserUsecase(userRepo)
	userHandler := user.NewUserHandler(userUsecase)

	// auth handler
	authHandler := auth.NewAuthHandler(userUsecase)

	// init opening handler
	openingRepo := opening.NewOpeningRepository(config.GetSQLite())
	openingUsecase := opening.NewOpeningUsecase(openingRepo)
	openingHandler := opening.NewOpeningHandler(openingUsecase)

	// init newsletter handler
	newsletterRepo := newsletter.NewNewsletterRepository(config.GetSQLite())
	newsletterUsecase := newsletter.NewNewsletterUsecase(newsletterRepo)
	newsletterHandler := newsletter.NewNewsletterHandler(newsletterUsecase)

	// init resume handler
	resumeRepo := resume.NewResumeRepository(config.GetSQLite())
	resumeUsecase := resume.NewResumeUsecase(resumeRepo)
	resumeHandler := resume.NewResumeHandler(resumeUsecase)

	// Initialize Router
	router.InitializeRoutes(ginRouter, authHandler, resumeHandler, openingHandler, newsletterHandler, userHandler)
}
