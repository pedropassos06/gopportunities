package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Init Router with default configs from Gin
	router := gin.Default()

	// Defining ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	// Run route
	router.Run(":8080")
}
