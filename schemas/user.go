package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
