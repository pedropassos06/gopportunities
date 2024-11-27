package resume

import (
	"github.com/pedropassos06/gopportunities/config"
	"gorm.io/gorm"
)

type ResumeHandler struct {
	DB     *gorm.DB
	Logger config.Logger
}

// NewResumeHandler initializes and returns a ResumeHandler instance
func NewResumeHandler(db *gorm.DB, logger config.Logger) *ResumeHandler {
	return &ResumeHandler{
		DB:     db,
		Logger: logger,
	}
}
