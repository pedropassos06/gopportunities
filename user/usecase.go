package user

import "github.com/pedropassos06/gopportunities/schemas"

type UserUsecase interface {
	CreateUser(user schemas.User) error
	DeleteUser(email string) error
	GetUser(email string) (*schemas.User, error)
	UpdateUser(schemas.User) error
}

type UserUsecaseImpl struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repo: repo,
	}
}

func (u *UserUsecaseImpl) CreateUser(user schemas.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserUsecaseImpl) DeleteUser(email string) error {
	return u.repo.DeleteUser(email)
}

func (u *UserUsecaseImpl) GetUser(email string) (*schemas.User, error) {
	return u.repo.GetUser(email)
}

func (u *UserUsecaseImpl) UpdateUser(user schemas.User) error {
	return u.repo.UpdateUser(user)
}
