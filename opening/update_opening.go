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
	uint64ID, err := utils.StringToUint(id)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "id must be a number")
		return
	}

	// find opening first
	opening, err := h.Usecase.GetOpeningByID(uint(uint64ID))
	if err != nil {
		utils.SendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	// Update only the fields that are provided in the request
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.TypeOfEmployment != "" {
		opening.TypeOfEmployment = request.TypeOfEmployment
	}
	if request.Salary != 0 { // Assuming Salary is a numeric value and 0 is considered as "not provided"
		opening.Salary = request.Salary
	}
	if request.CompanyLogoUrl != "" {
		opening.CompanyLogoUrl = request.CompanyLogoUrl
	}
	if request.Description != "" {
		opening.Description = request.Description
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	// Now update the opening with the new data
	if err := h.Usecase.UpdateOpening(opening); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Error updating opening.")
		return
	}

	// Return the updated opening
	utils.SendSuccess(ctx, "update-opening", opening)
}
