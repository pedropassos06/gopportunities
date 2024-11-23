package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Upload Resume
// @Description Uploads a resume to the database
// @Tags Resumes
// @Accept multipart/form-data
// @Produce json
// @Param user_id path string true "User ID of the resume owner"
// @Param resume formData file true "Resume file to upload"
// @Success 200 {object} UploadResumeResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /resumes/upload [post]
func UploadResumeHandler(ctx *gin.Context) {
	// grab user id from Params
	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id is required to upload resume",
		})
		return
	}

	file, err := ctx.FormFile("resume")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save file",
		})
		return
	}

	filePath := generateFilePath(file.Filename)

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	resume := schemas.Resume{
		UserID:   123,
		Filetype: filepath.Ext(file.Filename),
		Filepath: filePath,
	}

	if err := db.Create(&resume).Error; err != nil {
		logger.Errf("error uploading resume: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error uploading resume on database")
		return
	}

	sendSuccess(ctx, "upload-resume", resume)
}

func generateFilePath(filename string) string {
	// timestamp-based filename
	timestamp := time.Now().Unix()
	uniqueFilename := fmt.Sprintf("%d_%s", timestamp, filename)
	return fmt.Sprintf("./uploads/resumes/%s", uniqueFilename)
}
