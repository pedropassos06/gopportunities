package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role             string
	Company          string
	Location         string // remote or site
	TypeOfEmployment string // full-time, part-time, internship
	Salary           int64
	CompanyLogoUrl   string
	Description      string
	Link             string
}

type OpeningResponse struct {
	ID               uint      `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
	Role             string    `json:"role"`
	Company          string    `json:"company"`
	Location         string    `json:"location"`
	TypeOfEmployment string    `json:"type_of_employment"`
	Salary           int64     `json:"salary"`
	CompanyLogoUrl   string    `json:"company_logo_url"`
	Description      string    `json:"description"`
	Link             string    `json:"link"`
}
