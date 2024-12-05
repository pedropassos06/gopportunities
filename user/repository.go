package user

import (
	"github.com/pedropassos06/gopportunities/schemas"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user schemas.User) error
	DeleteUser(email string) error
	GetUser(email string) (*schemas.User, error)
	UpdateUser(schemas.User) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) CreateUser(user schemas.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepositoryImpl) DeleteUser(email string) error {
	return r.DB.Where("email = ?", email).Delete(&schemas.User{}).Error
}

func (r *UserRepositoryImpl) GetUser(email string) (*schemas.User, error) {
	user := &schemas.User{}
	err := r.DB.Where("email = ?", email).First(user).Error
	return user, err
}

func (r *UserRepositoryImpl) UpdateUser(user schemas.User) error {
	return r.DB.Save(&user).Error
}
