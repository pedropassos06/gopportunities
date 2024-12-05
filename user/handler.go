package user

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUserHandler(ctx *gin.Context)
	DeleteUserHandler(ctx *gin.Context)
	GetUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
}

type UserHandlerImpl struct {
	usecase UserUsecase
}

func NewUserHandler(usecase UserUsecase) *UserHandlerImpl {
	return &UserHandlerImpl{
		usecase: usecase,
	}
}
