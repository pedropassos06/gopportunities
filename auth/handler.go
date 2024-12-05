package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/user"
)

type AuthHandler interface {
	GoogleAuthHandler(c *gin.Context)
	GoogleCallbackHandler(c *gin.Context)
}

type AuthHandlerImpl struct {
	UserUseCase user.UserUsecase
}

func NewAuthHandler(usecase user.UserUsecase) AuthHandler {
	return &AuthHandlerImpl{
		UserUseCase: usecase,
	}
}
