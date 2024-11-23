package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/pedropassos06/gopportunities/docs"
	"github.com/pedropassos06/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/resumes/upload/:user_id", handler.UploadResumeHandler)
		v1.GET("/ping", handler.PingHandler)
	}

	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
