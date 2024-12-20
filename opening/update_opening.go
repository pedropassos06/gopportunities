package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	utils "github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1/

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id query string true "Opening ID"
// @Param opening body schemas.Opening true "Opening data to update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opening [put]
func (h *OpeningHandlerImpl) UpdateOpeningHandler(ctx *gin.Context) {
	// Bind the request JSON to the UpdateOpeningRequest struct
	var request schemas.Opening
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "Invalid input data.")
		return
	}

	// Retrieve the opening ID from query params
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	// Convert the ID to a uint
	uintID, err := utils.StringToUint(id)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "id must be a number")
		return
	}

	// find opening first
	opening, err := h.Usecase.GetOpeningByID(uintID)
	if err != nil {
		utils.SendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	// Now update the opening with the new data
	if err := h.Usecase.UpdateOpening(request, opening); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Error updating opening.")
		return
	}

	// Return the updated opening
	utils.SendSuccess(ctx, "update-opening", opening)
}
