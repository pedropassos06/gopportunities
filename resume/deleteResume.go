package resume

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// DeleteResumeHandler godoc
// @Summary Deletes Resume
// @Description deletes a resume given its id
// @Tags Resume
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param resume_id path string true "ID of the resume"
// @Success 200 {object} DeleteResumeResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /resume/{resume_id} [delete]
func (h *ResumeHandler) DeleteResumeHandler(ctx *gin.Context) {
	// grab resume id
	id := ctx.Param("resume_id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, "resume id is required")
		return
	}

	if err := h.DB.First(&schemas.Resume{}, id).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, "resume not found")
		return
	}

	// delete resume
	if err := h.DB.Delete(&schemas.Resume{}, id).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "server is unable to delete resume at the moment")
		return
	}

	utils.SendSuccess(ctx, "delete-resume", nil)
}
