package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/pedropassos06/gopportunities/helper"
	"github.com/pedropassos06/gopportunities/schemas"
)

// @BasePath /api/v1/

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Param opening body UpdateOpeningRequest true "Opening data to update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /opening [put]
func (h *OpeningHandler) UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		h.Logger.Errf("validation error: %v", err.Error())
		helper.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest, helper.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := h.DB.First(&opening, id).Error; err != nil {
		helper.SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	// update opening
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
	if request.Salary <= 0 {
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
	// Save opening
	if err := h.DB.Save(&opening).Error; err != nil {
		h.Logger.Errf("error updating opening: %v", err.Error())
		helper.SendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	helper.SendSuccess(ctx, "update-opening", opening)
}
