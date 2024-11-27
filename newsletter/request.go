package newsletter

import "github.com/pedropassos06/gopportunities/helper"

type NewsletterSubscriptionRequest struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

func (r *NewsletterSubscriptionRequest) Validate() error {
	if r.Email == "" {
		return helper.ErrParamIsRequired("email", "string")
	}
	return nil
}
