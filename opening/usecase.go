package opening

import (
	"github.com/pedropassos06/gopportunities/schemas"
)

type OpeningUsecase interface {
	CreateOpening(opening *schemas.Opening) error
	GetOpeningByID(id uint) (*schemas.Opening, error)
	GetAllOpenings() ([]*schemas.Opening, error)
	UpdateOpening(request schemas.Opening, opening *schemas.Opening) error
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

func (u *OpeningUsecaseImpl) UpdateOpening(request schemas.Opening, opening *schemas.Opening) error {
	checkFieldsToUpdate(request, opening)
	return u.repo.UpdateOpening(opening)
}

func (u *OpeningUsecaseImpl) DeleteOpening(id uint) error {
	return u.repo.DeleteOpening(id)
}

func (u *OpeningUsecaseImpl) GetFilteredOpenings(filters map[string]interface{}) ([]schemas.Opening, error) {
	return u.repo.GetFilteredOpenings(filters)
}

func checkFieldsToUpdate(request schemas.Opening, dest *schemas.Opening) {
	// Update only the fields that are provided in the request
	if request.Role != "" {
		dest.Role = request.Role
	}
	if request.Company != "" {
		dest.Company = request.Company
	}
	if request.Location != "" {
		dest.Location = request.Location
	}
	if request.TypeOfEmployment != "" {
		dest.TypeOfEmployment = request.TypeOfEmployment
	}
	if request.Salary != 0 { // Assuming Salary is a numeric value and 0 is considered as "not provided"
		dest.Salary = request.Salary
	}
	if request.CompanyLogoUrl != "" {
		dest.CompanyLogoUrl = request.CompanyLogoUrl
	}
	if request.Description != "" {
		dest.Description = request.Description
	}
	if request.Link != "" {
		dest.Link = request.Link
	}
}
