package user

import "github.com/gin-gonic/gin"

type UserHandler interface {
	DeleteUserHandler(ctx *gin.Context)
	GetUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
}

type UserHandlerImpl struct {
	Usecase UserUsecase
}

func NewUserHandler(usecase UserUsecase) UserHandler {
	return &UserHandlerImpl{
		Usecase: usecase,
	}
}
