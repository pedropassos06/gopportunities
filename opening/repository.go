package opening

import (
	"github.com/pedropassos06/gopportunities/schemas"
	"gorm.io/gorm"
)

type OpeningRepository interface {
	CreateOpening(opening *schemas.Opening) error
	GetOpeningByID(id uint) (*schemas.Opening, error)
	GetAllOpenings() ([]*schemas.Opening, error)
	UpdateOpening(opening *schemas.Opening) error
	DeleteOpening(id uint) error
	GetFilteredOpenings(filters map[string]interface{}) ([]schemas.Opening, error)
}

type OpeningRepositoryImpl struct {
	DB *gorm.DB
}

func NewOpeningRepository(db *gorm.DB) OpeningRepository {
	return &OpeningRepositoryImpl{DB: db}
}

func (r *OpeningRepositoryImpl) CreateOpening(opening *schemas.Opening) error {
	return r.DB.Create(opening).Error
}

func (r *OpeningRepositoryImpl) GetOpeningByID(id uint) (*schemas.Opening, error) {
	var opening schemas.Opening
	err := r.DB.First(&opening, id).Error
	return &opening, err
}

func (r *OpeningRepositoryImpl) GetAllOpenings() ([]*schemas.Opening, error) {
	var openings []*schemas.Opening
	err := r.DB.Find(&openings).Error
	return openings, err
}

func (r *OpeningRepositoryImpl) UpdateOpening(opening *schemas.Opening) error {
	return r.DB.Save(opening).Error
}

func (r *OpeningRepositoryImpl) DeleteOpening(id uint) error {
	return r.DB.Delete(&schemas.Opening{}, id).Error
}

func (r *OpeningRepositoryImpl) GetFilteredOpenings(filters map[string]interface{}) ([]schemas.Opening, error) {
	var openings []schemas.Opening
	query := r.DB.Model(&schemas.Opening{})

	// Dynamically apply filters
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	// Execute query
	if err := query.Find(&openings).Error; err != nil {
		return nil, err
	}

	return openings, nil
}
