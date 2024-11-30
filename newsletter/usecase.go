package newsletter

import (
	"github.com/pedropassos06/gopportunities/schemas"
	"github.com/pedropassos06/gopportunities/utils"
)

type NewsletterUsecase interface {
	Subscribe(subscription schemas.NewsletterSubscription) error
	Unsubscribe(subscription schemas.NewsletterSubscription) error
	Find(email string) (schemas.NewsletterSubscription, error)
}

type NewsletterUsecaseImpl struct {
	Repository NewsletterRepository
}

// NewNewsletterUsecase initializes and returns a NewsletterUsecase instance
func NewNewsletterUsecase(repository NewsletterRepository) NewsletterUsecase {
	return &NewsletterUsecaseImpl{
		Repository: repository,
	}
}

func (u *NewsletterUsecaseImpl) Subscribe(subscription schemas.NewsletterSubscription) error {
	if subscription.Email == "" {
		return utils.ErrParamIsRequired("email", "string")
	}
	return u.Repository.Subscribe(subscription)
}

func (u *NewsletterUsecaseImpl) Unsubscribe(subscription schemas.NewsletterSubscription) error {
	// set subscribed att to false
	subscription.Subscribed = false
	return u.Repository.Unsubscribe(subscription)
}

func (u *NewsletterUsecaseImpl) Find(email string) (schemas.NewsletterSubscription, error) {
	return u.Repository.Find(email)
}
