package opening

import (
	"github.com/pedropassos06/gopportunities/config"
	"gorm.io/gorm"
)

type OpeningHandler struct {
	DB     *gorm.DB
	Logger config.Logger
}

// NewOpeningHandler initializes and returns a OpeningHandler instance
func NewOpeningHandler(db *gorm.DB, logger config.Logger) *OpeningHandler {
	return &OpeningHandler{
		DB:     db,
		Logger: logger,
	}
}
