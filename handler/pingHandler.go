package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
