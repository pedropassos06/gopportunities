package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1/

// @Summary List all openings
// @Description List all job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param role query string false "Role"
// @Param location query string false "Location"
// @Param company query string false "Company"
// @Param minSalary query string false "Minimum Salary"
// @Success 200 {object} ListOpeningsResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /openings [get]
func (h *OpeningHandlerImpl) ListOpeningsHandler(ctx *gin.Context) {
	// Prepare filters from query parameters
	filters := h.prepareFilters(ctx)

	// Get filtered openings from usecase
	openings, err := h.Usecase.GetFilteredOpenings(filters)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "could not retrieve openings")
		return
	}

	// Send success response
	utils.SendSuccess(ctx, "list-openings", openings)
}

// prepareFilters extracts filter parameters from the query
func (h *OpeningHandlerImpl) prepareFilters(ctx *gin.Context) map[string]interface{} {
	filters := make(map[string]interface{})

	// Extract query parameters and add to filters if present
	if role := ctx.Query("role"); role != "" {
		filters["role"] = role
	}
	if location := ctx.Query("location"); location != "" {
		filters["location"] = location
	}
	if company := ctx.Query("company"); company != "" {
		filters["company"] = company
	}
	if minSalary := ctx.Query("minSalary"); minSalary != "" {
		filters["salary >"] = minSalary
	}

	return filters
}
