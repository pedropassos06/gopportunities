package auth

import (
	"github.com/pedropassos06/gopportunities/config"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB     *gorm.DB
	Logger config.Logger
}

func NewAuthHandler(db *gorm.DB, logger config.Logger) *AuthHandler {
	return &AuthHandler{
		DB:     db,
		Logger: logger,
	}
}
