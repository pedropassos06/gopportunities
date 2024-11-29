package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Param opening body UpdateOpeningRequest true "Opening data to update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /opening [put]
func (h *OpeningHandler) UpdateOpeningHandler(ctx *gin.Context) {
	// Bind the request JSON to the UpdateOpeningRequest struct
	var request UpdateOpeningRequest
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
	uint64ID, err := utils.StringToUint(id)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "id must be a number")
		return
	}

	// Usecase to update opening
	opening, err := h.Usecase.GetOpeningByID(uint(uint64ID))
	if err != nil {
		utils.SendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	// Apply updates from the request to the opening
	if err := h.Usecase.UpdateOpening(opening, &request); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Error updating opening.")
		return
	}

	// Return the updated opening
	utils.SendSuccess(ctx, "update-opening", opening)
}
