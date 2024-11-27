package newsletter

import "github.com/pedropassos06/gopportunities/utils"

type NewsletterSubscriptionRequest struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

func (r *NewsletterSubscriptionRequest) Validate() error {
	if r.Email == "" {
		return utils.ErrParamIsRequired("email", "string")
	}
	return nil
}
