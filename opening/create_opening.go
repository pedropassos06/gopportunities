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
// @Param request body schemas.Opening true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opening [post]
func (h *OpeningHandler) CreateOpeningHandler(ctx *gin.Context) {
	var opening schemas.Opening
	if err := ctx.BindJSON(&opening); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Create opening
	err := h.Usecase.CreateOpening(&opening)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "could not create opening")
		return
	}

	utils.SendSuccess(ctx, "create-opening", opening)
}
