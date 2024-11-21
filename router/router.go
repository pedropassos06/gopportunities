package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Init Router with default configs from Gin
	router := gin.Default()

	// Init Routes
	InitializeRoutes(router)

	// Run route
	router.Run(":8081")
}
