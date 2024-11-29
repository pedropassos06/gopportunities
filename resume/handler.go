package resume

import (
	"github.com/gin-gonic/gin"
)

type ResumeHandler interface {
	UploadResumeHandler(ctx *gin.Context)
	GetResumeHandler(ctx *gin.Context)
	DeleteResumeHandler(ctx *gin.Context)
}

type ResumeHandlerImpl struct {
	Usecase ResumeUsecase
}

// NewResumeHandler initializes and returns a ResumeHandler instance
func NewResumeHandler(usecase ResumeUsecase) ResumeHandler {
	return &ResumeHandlerImpl{
		Usecase: usecase,
	}
}
