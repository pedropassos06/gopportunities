package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/docs"
	"github.com/pedropassos06/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine, h *handler.Handler) {
	// Initialize handler
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", h.ShowOpeningHandler)
		v1.POST("/opening", h.CreateOpeningHandler)
		v1.DELETE("/opening", h.DeleteOpeningHandler)
		v1.PUT("/opening", h.UpdateOpeningHandler)
		v1.GET("/openings", h.ListOpeningsHandler)
		v1.POST("/resumes/upload/:user_id", h.UploadResumeHandler)
		v1.POST("/newsletter/subscribe", h.SubscribeHandler)
		v1.GET("/ping", h.PingHandler)
	}

	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8081")
}
