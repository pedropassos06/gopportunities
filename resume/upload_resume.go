package resume

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Upload Resume
// @Description Uploads a resume to the database
// @Tags Resume
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param user_id path string true "User ID of the resume owner"
// @Param resume formData file true "Resume file to upload"
// @Success 200 {object} UploadResumeResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /resume/upload/{user_id} [post]
func (h *ResumeHandler) UploadResumeHandler(ctx *gin.Context) {
	// grab user id from Params
	userIDStr := ctx.Param("user_id")
	if userIDStr == "" {
		utils.SendError(ctx, http.StatusBadRequest, "user_id is required to upload resume")
		return
	}

	file, err := ctx.FormFile("resume")
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "resume is required")
		return
	}

	filePath := generateFilePath(file.Filename)

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "failed to save file")
		return
	}

	// Convert userID to uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "user_id must be a valid integer")
		return
	}

	resume := schemas.Resume{
		UserID:   uint(userID),
		Filetype: filepath.Ext(file.Filename),
		Filepath: filePath,
	}

	if err := h.DB.Create(&resume).Error; err != nil {
		h.Logger.Errf("error uploading resume: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, "error uploading resume on database")
		return
	}

	utils.SendSuccess(ctx, "upload-resume", resume)
}

func generateFilePath(filename string) string {
	// timestamp-based filename
	timestamp := time.Now().Unix()
	uniqueFilename := fmt.Sprintf("%d_%s", timestamp, filename)
	return fmt.Sprintf("./uploads/resumes/%s", uniqueFilename)
}
