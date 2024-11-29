package auth

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	GoogleAuthHandler(c *gin.Context)
	GoogleCallbackHandler(c *gin.Context)
}

type AuthHandlerImpl struct{}

func NewAuthHandler() AuthHandler {
	return &AuthHandlerImpl{}
}
