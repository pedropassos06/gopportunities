package newsletter

import "github.com/pedropassos06/gopportunities/schemas"

type NewsletterSubscriptionResponse struct {
	Message string                                 `json:"message"`
	Data    schemas.NewsletterSubscriptionResponse `json:"data"`
}

type NewsletterUnsubscribeResponse struct {
	Message string                                 `json:"message"`
	Data    schemas.NewsletterSubscriptionResponse `json:"data"`
}
