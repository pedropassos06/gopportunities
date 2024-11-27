package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	helper "github.com/pedropassos06/gopportunities/helper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// grab auth header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header not found"})
			return
		}

		// check if auth header is bearer
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		// validate token
		token := parts[1]
		if !helper.ValidateToken(token) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// continue
		ctx.Next()
	}
}
