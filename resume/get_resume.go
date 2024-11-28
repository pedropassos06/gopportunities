package resume

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// GetResumeHandler godoc
// @Summary Retrieve a resume
// @Description retrieves a resume by user id
// @Tags Resume
// @Accept json
// @Produce multipart/form-data
// @Param Authorization header string true "Bearer Token"
// @Param user_id path string true "User ID of the resume owner"
// @Success 200 {object} GetResumeResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /resume/{user_id} [get]
func (rh *ResumeHandler) GetResumeHandler(ctx *gin.Context) {
	// get user
	userId := ctx.Param("user_id")
	if userId == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("id", "path parameter").Error())
	}

	var resume schemas.Resume
	// make sure it exists
	if err := rh.DB.First(&resume, "user_id = ?", userId).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, "resume not found")
		return
	}

	// split path to get filename
	parts := strings.Split(resume.Filepath, "/")
	filename := parts[len(parts)-1]

	// return file as attachment
	ctx.FileAttachment(resume.Filepath, filename)
}