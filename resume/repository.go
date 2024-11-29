package resume

import (
	"github.com/pedropassos06/gopportunities/schemas"
	"gorm.io/gorm"
)

type ResumeRepository interface {
	UploadResume(schemas.Resume) error
	GetResumeByID(resumeID uint) (schemas.Resume, error)
	GetResumeByUserID(userID uint) (schemas.Resume, error)
	DeleteResume(resumeID uint) error
}

type ResumeRepositoryImpl struct {
	DB *gorm.DB
}

func NewResumeRepository(db *gorm.DB) *ResumeRepositoryImpl {
	return &ResumeRepositoryImpl{
		DB: db,
	}
}

func (r *ResumeRepositoryImpl) UploadResume(resume schemas.Resume) error {
	return r.DB.Create(&resume).Error
}

func (r *ResumeRepositoryImpl) GetResumeByID(resumeID uint) (schemas.Resume, error) {
	var resume schemas.Resume
	if err := r.DB.First(&resume, resumeID).Error; err != nil {
		return schemas.Resume{}, err
	}
	return resume, nil
}

func (r *ResumeRepositoryImpl) GetResumeByUserID(userID uint) (schemas.Resume, error) {
	var resume schemas.Resume
	if err := r.DB.First(&resume, "user_id = ?", userID).Error; err != nil {
		return schemas.Resume{}, err
	}
	return resume, nil
}

func (r *ResumeRepositoryImpl) DeleteResume(resumeID uint) error {
	return r.DB.Delete(&schemas.Resume{}, resumeID).Error
}
