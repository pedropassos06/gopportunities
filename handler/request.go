package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role             string `json:"role"`
	Company          string `json:"company"`
	Location         string `json:"location"`
	TypeOfEmployment string `json:"type_of_employment"`
	Salary           int64  `json:"salary"`
	CompanyLogoUrl   string `json:"company_logo_url"`
	Description      string `json:"description"`
	Link             string `json:"link"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.TypeOfEmployment == "" {
		return errParamIsRequired("type_of_employment", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	if r.CompanyLogoUrl == "" {
		return errParamIsRequired("company_logo_url", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role             string `json:"role"`
	Company          string `json:"company"`
	Location         string `json:"location"`
	TypeOfEmployment string `json:"type_of_employment"`
	Salary           int64  `json:"salary"`
	CompanyLogoUrl   string `json:"company_logo_url"`
	Description      string `json:"description"`
	Link             string `json:"link"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// if any field is provided, validation is true
	if r.Role != "" || r.Company != "" || r.Location != "" || r.TypeOfEmployment != "" || r.Salary > 0 || r.CompanyLogoUrl != "" || r.Description != "" || r.Link != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}

type NewsletterSubscriptionRequest struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

func (r *NewsletterSubscriptionRequest) Validate() error {
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	return nil
}
