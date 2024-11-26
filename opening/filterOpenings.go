package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	utils "github.com/pedropassos06/gopportunities/utils"
)

// handler for filtering listings
func (h *OpeningHandler) ListFilteredOpeningsHandler(ctx *gin.Context) {
	filters := make(map[string]interface{})

	// read query params and add to filters
	if role := ctx.Query("role"); role != "" {
		filters["role"] = role
	}
	if location := ctx.Query("location"); location != "" {
		filters["location"] = location
	}
	if remote := ctx.Query("remote"); remote != "" {
		filters["remote"] = remote
	}
	if minSalary := ctx.Query("minSalary"); minSalary != "" {
		filters["salary >"] = minSalary
	}

	// Call the generic filter function
	openings, err := h.FilterOpenings(filters)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "could not retrieve openings")
		return
	}

	utils.SendSuccess(ctx, "list-filtered-openings", openings)
}

// filters openings based on a filters array
func (h *OpeningHandler) FilterOpenings(filters map[string]interface{}) ([]schemas.Opening, error) {
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
