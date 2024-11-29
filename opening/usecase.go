package opening

import (
	"github.com/pedropassos06/gopportunities/schemas"
)

type OpeningUsecase interface {
	CreateOpening(opening *schemas.Opening) error
	GetOpeningByID(id uint) (*schemas.Opening, error)
	GetAllOpenings() ([]*schemas.Opening, error)
	UpdateOpening(opening *schemas.Opening) error
	DeleteOpening(id uint) error
	GetFilteredOpenings(filters map[string]interface{}) ([]schemas.Opening, error)
}

type OpeningUsecaseImpl struct {
	repo OpeningRepository
}

func NewOpeningUsecase(repo OpeningRepository) *OpeningUsecaseImpl {
	return &OpeningUsecaseImpl{repo}
}

func (u *OpeningUsecaseImpl) CreateOpening(opening *schemas.Opening) error {
	return u.repo.CreateOpening(opening)
}

func (u *OpeningUsecaseImpl) GetOpeningByID(id uint) (*schemas.Opening, error) {
	return u.repo.GetOpeningByID(id)
}

func (u *OpeningUsecaseImpl) GetAllOpenings() ([]*schemas.Opening, error) {
	return u.repo.GetAllOpenings()
}

func (u *OpeningUsecaseImpl) UpdateOpening(opening *schemas.Opening) error {
	return u.repo.UpdateOpening(opening)
}

func (u *OpeningUsecaseImpl) DeleteOpening(id uint) error {
	return u.repo.DeleteOpening(id)
}

func (u *OpeningUsecaseImpl) GetFilteredOpenings(filters map[string]interface{}) ([]schemas.Opening, error) {
	return u.repo.GetFilteredOpenings(filters)
}
