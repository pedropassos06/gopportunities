package resume

import (
	"strconv"

	"github.com/pedropassos06/gopportunities/schemas"
)

type ResumeUsecase interface {
	UploadResume(resume schemas.Resume) error
	GetResumeByID(resumeID string) (schemas.Resume, error)
	GetResumeByUserID(userID string) (schemas.Resume, error)
	DeleteResume(resumeID uint) error
}

type ResumeUsecaseImpl struct {
	ResumeRepo ResumeRepository
}

func NewResumeUsecase(repo ResumeRepository) *ResumeUsecaseImpl {
	return &ResumeUsecaseImpl{
		ResumeRepo: repo,
	}
}

func (u *ResumeUsecaseImpl) UploadResume(resume schemas.Resume) error {
	return u.ResumeRepo.UploadResume(resume)
}

func (u *ResumeUsecaseImpl) GetResumeByID(resumeID string) (schemas.Resume, error) {
	// Convert userID to uint
	resumeIDuint, err := strconv.ParseUint(resumeID, 10, 32)
	if err != nil {
		return schemas.Resume{}, err
	}
	return u.ResumeRepo.GetResumeByID(uint(resumeIDuint))
}

func (u *ResumeUsecaseImpl) GetResumeByUserID(userID string) (schemas.Resume, error) {
	// Convert userID to uint
	userIDuint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return schemas.Resume{}, err
	}
	return u.ResumeRepo.GetResumeByUserID(uint(userIDuint))
}

func (u *ResumeUsecaseImpl) DeleteResume(resumeID uint) error {
	return u.ResumeRepo.DeleteResume(resumeID)
}
