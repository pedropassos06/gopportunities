package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Resume struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Filetype string `gorm:"not null"` // File type (e.g., application/pdf)
	Filepath string `gorm:"not null"` // Full path to the file
}

type ResumeResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	UserID    uint      `json:"user_id"`
	Filename  string    `json:"file_name"`
	Filetype  string    `json:"file_type"`
	Filepath  string    `json:"file_path"`
}
