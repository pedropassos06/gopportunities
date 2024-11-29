package schemas

import (
	"time"
)

type Resume struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `gorm:"not null"`
	Filetype  string    `gorm:"not null"` // File type (e.g., application/pdf)
	Filepath  string    `gorm:"not null"` // Full path to the file
}
