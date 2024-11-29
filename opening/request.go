package opening

// type CreateOpeningRequest struct {
// 	Role             string `json:"role"`
// 	Company          string `json:"company"`
// 	Location         string `json:"location"`
// 	TypeOfEmployment string `json:"type_of_employment"`
// 	Salary           int64  `json:"salary"`
// 	CompanyLogoUrl   string `json:"company_logo_url"`
// 	Description      string `json:"description"`
// 	Link             string `json:"link"`
// }

// func (r *CreateOpeningRequest) Validate() error {
// 	if r.Role == "" {
// 		return utils.ErrParamIsRequired("role", "string")
// 	}
// 	if r.Company == "" {
// 		return utils.ErrParamIsRequired("company", "string")
// 	}
// 	if r.Location == "" {
// 		return utils.ErrParamIsRequired("location", "string")
// 	}
// 	if r.TypeOfEmployment == "" {
// 		return utils.ErrParamIsRequired("type_of_employment", "string")
// 	}
// 	if r.Salary <= 0 {
// 		return utils.ErrParamIsRequired("salary", "int64")
// 	}
// 	if r.CompanyLogoUrl == "" {
// 		return utils.ErrParamIsRequired("company_logo_url", "string")
// 	}
// 	if r.Description == "" {
// 		return utils.ErrParamIsRequired("description", "string")
// 	}
// 	if r.Link == "" {
// 		return utils.ErrParamIsRequired("link", "string")
// 	}
// 	return nil
// }

// type UpdateOpeningRequest struct {
// 	Role             string `json:"role"`
// 	Company          string `json:"company"`
// 	Location         string `json:"location"`
// 	TypeOfEmployment string `json:"type_of_employment"`
// 	Salary           int64  `json:"salary"`
// 	CompanyLogoUrl   string `json:"company_logo_url"`
// 	Description      string `json:"description"`
// 	Link             string `json:"link"`
// }

// func (r *UpdateOpeningRequest) Validate() error {
// 	// if any field is provided, validation is true
// 	if r.Role != "" || r.Company != "" || r.Location != "" || r.TypeOfEmployment != "" || r.Salary > 0 || r.CompanyLogoUrl != "" || r.Description != "" || r.Link != "" {
// 		return nil
// 	}

// 	return fmt.Errorf("at least one valid field must be provided")
// }
