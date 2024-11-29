package schemas

import (
	"time"
)

type NewsletterSubscription struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserID     uint      `json:"user_id"`
	Email      string    `json:"email"`
	Subscribed bool      `json:"subscribed" gorm:"default:true"`
}
