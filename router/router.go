package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/auth"
	"github.com/pedropassos06/gopportunities/docs"
	"github.com/pedropassos06/gopportunities/middleware"
	"github.com/pedropassos06/gopportunities/newsletter"
	"github.com/pedropassos06/gopportunities/opening"
	"github.com/pedropassos06/gopportunities/resume"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine, ah auth.AuthHandler, rh *resume.ResumeHandler, oh *opening.OpeningHandler, nh *newsletter.NewsletterHandler) {
	// define base path for our api
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// public routes
	public := r.Group("/api/v1")
	{
		public.GET("/auth/google", ah.GoogleAuthHandler)
		public.GET("/auth/google/callback", ah.GoogleCallbackHandler)
	}

	// use auth middleware - protected routes
	protected := r.Group("/api/v1", middleware.AuthMiddleware())
	{
		// opening endpoints
		protected.GET("/opening", oh.ShowOpeningHandler)
		protected.POST("/opening", oh.CreateOpeningHandler)
		protected.DELETE("/opening", oh.DeleteOpeningHandler)
		protected.PUT("/opening", oh.UpdateOpeningHandler)
		protected.GET("/openings", oh.ListOpeningsHandler)
		// resume endpoints
		protected.POST("/resume/upload/:user_id", rh.UploadResumeHandler)
		protected.DELETE("/resume/:resume_id", rh.DeleteResumeHandler)
		protected.GET("/resume/:user_id", rh.GetResumeHandler)
		// newsletter endpoints
		protected.POST("/newsletter/subscribe", nh.SubscribeHandler)
		protected.PUT("/newsletter/unsubscribe/:user_email", nh.UnsubscribeHandler)
	}

	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // configure this if needed
	}
	r.Run(fmt.Sprintf(":%s", port))
}
