package opening

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id query string true "Opening id"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /opening [delete]
func (h *OpeningHandlerImpl) DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "queryParameter").Error())
	}

	// Convert the ID to a uint
	uintID, err := utils.StringToUint(id)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "id must be a number")
		return
	}

	// Find opening
	opening, err := h.Usecase.GetOpeningByID(uintID)
	if err != nil {
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Delete
	if err = h.Usecase.DeleteOpening(uintID); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	utils.SendSuccess(ctx, "delete-opening", opening)
}
