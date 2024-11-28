package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	utils "github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1/

// @Summary List all openings
// @Description List all job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /openings [get]
func (h *OpeningHandler) ListOpeningsHandler(ctx *gin.Context) {
	// Check if any query parameters are provided
	filters := make(map[string]interface{})
	if role := ctx.Query("role"); role != "" {
		filters["role"] = role
	}
	if location := ctx.Query("location"); location != "" {
		filters["location"] = location
	}
	if company := ctx.Query("company"); company != "" {
		filters["company"] = company
	}
	if remote := ctx.Query("remote"); remote != "" {
		filters["remote"] = remote
	}
	if minSalary := ctx.Query("minSalary"); minSalary != "" {
		filters["salary >"] = minSalary
	}

	// If no filters are provided, list all openings
	if len(filters) == 0 {
		var openings []schemas.Opening
		if err := h.DB.Find(&openings).Error; err != nil {
			utils.SendError(ctx, http.StatusInternalServerError, "could not retrieve openings")
			return
		}
		utils.SendSuccess(ctx, "list-openings", openings)
		return
	}

	// Otherwise, apply filters
	openings, err := h.filterOpenings(filters)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "could not retrieve openings")
		return
	}
	utils.SendSuccess(ctx, "list-filtered-openings", openings)
}

// filters openings based on a filters array
func (h *OpeningHandler) filterOpenings(filters map[string]interface{}) ([]schemas.Opening, error) {
	var openings []schemas.Opening

	// start the query
	query := h.DB.Model(&schemas.Opening{})

	// dynamically apply filters
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	// execute the query
	if err := query.Find(&openings).Error; err != nil {
		return nil, err
	}

	return openings, nil
}
