package resume

import "github.com/pedropassos06/gopportunities/schemas"

type UploadResumeResponse struct {
	Message string                 `json:"message"`
	Data    schemas.ResumeResponse `json:"data"`
}

type GetResumeResponse struct {
	Message string                 `json:"message"`
	Data    schemas.ResumeResponse `json:"data"`
}

type DeleteResumeResponse struct {
	Message string                 `json:"message"`
	Data    schemas.ResumeResponse `json:"data"`
}
