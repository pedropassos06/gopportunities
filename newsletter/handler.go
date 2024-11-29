package newsletter

import (
	"github.com/gin-gonic/gin"
)

type NewsletterHandler interface {
	SubscribeHandler(ctx *gin.Context)
	UnsubscribeHandler(ctx *gin.Context)
}

type NewsletterHandlerImpl struct {
	Usecase NewsletterUsecase
}

// NewNewsletterHandler initializes and returns a NewsletterHandler instance
func NewNewsletterHandler(usecase NewsletterUsecase) NewsletterHandler {
	return &NewsletterHandlerImpl{
		Usecase: usecase,
	}
}
