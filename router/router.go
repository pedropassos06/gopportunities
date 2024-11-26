package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/docs"
	"github.com/pedropassos06/gopportunities/handler"
	"github.com/pedropassos06/gopportunities/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine, h *handler.Handler) {
	// define base path for our api
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// public routes
	public := r.Group("/api/v1")
	{
		public.GET("/auth/google", h.GoogleAuthHandler)
		public.GET("/auth/google/callback", h.GoogleCallbackHandler)
		public.GET("/ping", h.PingHandler)
	}

	// use auth middleware - protected routes
	protected := r.Group("/api/v1", middleware.AuthMiddleware())
	{
		protected.GET("/opening", h.ShowOpeningHandler)
		protected.POST("/opening", h.CreateOpeningHandler)
		protected.DELETE("/opening", h.DeleteOpeningHandler)
		protected.PUT("/opening", h.UpdateOpeningHandler)
		protected.GET("/openings", h.ListOpeningsHandler)
		protected.POST("/resumes/upload/:user_id", h.UploadResumeHandler)
		protected.POST("/newsletter/subscribe", h.SubscribeHandler)
	}

	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081" // configure this if needed
	}
	r.Run(fmt.Sprintf(":%s", port))
}
