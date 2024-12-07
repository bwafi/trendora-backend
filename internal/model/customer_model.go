package model

type CustomerResponse struct {
	ID           string  `json:"id,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Name         string  `json:"name,omitempty"`
	Gender       bool    `json:"gender,omitempty"`
	DateOfBirth  int64   `json:"date_of_birth,omitempty"`
	Token        string  `json:"token,omitempty"`
	CreatedAt    int64   `json:"created_at,omitempty"`
	UpdatedAt    int64   `json:"updated_at,omitempty"`
}

type CustomerRegisterRequest struct {
	EmailAddress *string `json:"email_address,omitempty" validate:"email"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Name         string  `json:"name,omitempty" validate:"required"`
	Password     string  `json:"password,omitempty" validate:"required"`
	Gender       bool    `json:"gender,omitempty"`
	DateOfBirth  int64   `json:"date_of_birth,omitempty"`
}

type CustomerUpdateRequest struct {
	ID           string  `json:"id,omitempty" validate:"required"`
	EmailAddress *string `json:"email_address,omitempty" validate:"email"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Name         string  `json:"name,omitempty"`
	Password     string  `json:"password,omitempty"`
	Gender       bool    `json:"gender,omitempty"`
	DateOfBirth  int64   `json:"date_of_birth,omitempty"`
}
