package handler

import (
	"github.com/pedropassos06/gopportunities/config"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Logger config.Logger
}

// NewHandler initializes and returns a Handler instance
func NewHandler(db *gorm.DB, logger config.Logger) *Handler {
	return &Handler{
		DB:     db,
		Logger: logger,
	}
}
