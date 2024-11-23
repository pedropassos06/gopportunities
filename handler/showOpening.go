package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
)

// @BasePath /api/v1/

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse "ID query parameter is missing"
// @Failure 404 {object} ErrorResponse "Opening not found"
// @Router /opening [get]
func (h *Handler) ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := h.DB.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found.")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
