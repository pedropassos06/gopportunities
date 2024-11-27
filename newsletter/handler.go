package newsletter

import (
	"github.com/pedropassos06/gopportunities/config"
	"gorm.io/gorm"
)

type NewsletterHandler struct {
	DB     *gorm.DB
	Logger config.Logger
}

// NewNewsletterHandler initializes and returns a NewsletterHandler instance
func NewNewsletterHandler(db *gorm.DB, logger config.Logger) *NewsletterHandler {
	return &NewsletterHandler{
		DB:     db,
		Logger: logger,
	}
}
