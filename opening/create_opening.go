package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	utils "github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opening [post]
func (h *OpeningHandler) CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		h.Logger.Errf("validation error: %v", err.Error())
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:             request.Role,
		Company:          request.Company,
		Location:         request.Location,
		TypeOfEmployment: request.TypeOfEmployment,
		Salary:           request.Salary,
		CompanyLogoUrl:   request.CompanyLogoUrl,
		Description:      request.Description,
		Link:             request.Link,
	}

	if err := h.DB.Create(&opening).Error; err != nil {
		h.Logger.Errf("error creating opening: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	utils.SendSuccess(ctx, "create-opening", opening)
}
