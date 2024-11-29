package opening

import "github.com/gin-gonic/gin"

type OpeningHandler interface {
	ShowOpeningHandler(c *gin.Context)
	CreateOpeningHandler(c *gin.Context)
	DeleteOpeningHandler(c *gin.Context)
	UpdateOpeningHandler(c *gin.Context)
	ListOpeningsHandler(c *gin.Context)
}

type OpeningHandlerImpl struct {
	Usecase OpeningUsecase
}

// NewOpeningHandler initializes and returns a OpeningHandler instance
func NewOpeningHandler(usecase OpeningUsecase) OpeningHandler {
	return &OpeningHandlerImpl{
		Usecase: usecase,
	}
}
