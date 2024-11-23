package schemas

import (
	"time"

	"gorm.io/gorm"
)

type NewsletterSubscription struct {
	gorm.Model
	UserID     uint
	Email      string
	Subscribed bool
}

type NewsletterSubscriptionResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserID     uint      `json:"user_id"`
	Email      string    `json:"email"`
	Subscribed bool      `json:"subscribed"`
}
