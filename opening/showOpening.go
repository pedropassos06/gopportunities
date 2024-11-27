package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/pedropassos06/gopportunities/helper"
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
// @Failure 400 {object} helper.ErrorResponse "ID query parameter is missing"
// @Failure 404 {object} helper.ErrorResponse "Opening not found"
// @Router /opening [get]
func (h *OpeningHandler) ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, helper.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := h.DB.First(&opening, id).Error; err != nil {
		helper.SendError(ctx, http.StatusNotFound, "opening not found.")
		return
	}

	helper.SendSuccess(ctx, "show-opening", opening)
}
