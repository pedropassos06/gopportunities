package newsletter

import (
	"github.com/pedropassos06/gopportunities/schemas"
	"gorm.io/gorm"
)

type NewsletterRepository interface {
	Subscribe(subscription schemas.NewsletterSubscription) error
	Unsubscribe(subscription schemas.NewsletterSubscription) error
	Find(email string) (schemas.NewsletterSubscription, error)
}

type NewsletterRepositoryImpl struct {
	DB *gorm.DB
}

// NewNewsletterRepository initializes and returns a NewsletterRepository instance
func NewNewsletterRepository(db *gorm.DB) NewsletterRepository {
	return &NewsletterRepositoryImpl{
		DB: db,
	}
}

func (r *NewsletterRepositoryImpl) Subscribe(subscription schemas.NewsletterSubscription) error {
	return r.DB.Create(&subscription).Error
}

func (r *NewsletterRepositoryImpl) Unsubscribe(subscription schemas.NewsletterSubscription) error {
	return r.DB.Save(subscription).Error
}

func (r *NewsletterRepositoryImpl) Find(email string) (schemas.NewsletterSubscription, error) {
	var subscription schemas.NewsletterSubscription
	if err := r.DB.First(&subscription, "email = ?", email).Error; err != nil {
		return schemas.NewsletterSubscription{}, err
	}
	return subscription, nil
}
