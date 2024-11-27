package opening

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/helper"
	"github.com/pedropassos06/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening id"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Router /opening [delete]
func (h *OpeningHandler) DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, helper.ErrParamIsRequired("id", "queryParameter").Error())
	}

	opening := schemas.Opening{}

	// Find opening (need to find so then you can delete)
	if err := h.DB.First(&opening, id).Error; err != nil {
		helper.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Delete
	if err := h.DB.Delete(&opening, id).Error; err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	helper.SendSuccess(ctx, "delete-opening", opening)
}
