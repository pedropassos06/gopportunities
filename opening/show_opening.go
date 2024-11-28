package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	utils "github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1/

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} utils.ErrorResponse "ID query parameter is missing"
// @Failure 404 {object} utils.ErrorResponse "Opening not found"
// @Router /opening [get]
func (h *OpeningHandler) ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := h.DB.First(&opening, id).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, "opening not found.")
		return
	}

	utils.SendSuccess(ctx, "show-opening", opening)
}
