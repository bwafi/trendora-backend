package model

type CustomerResponse struct {
	ID           string  `json:"id,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Name         string  `json:"name,omitempty"`
	Gender       bool    `json:"gender,omitempty"`
	DateOfBirth  int64   `json:"date_of_birth,omitempty"`
	RefreshToken string  `json:"refresh_token,omitempty"`
	AccessToken  string  `json:"access_token,omitempty"`
	CreatedAt    int64   `json:"created_at,omitempty"`
	UpdatedAt    int64   `json:"updated_at,omitempty"`
}

type CustomerRegisterRequest struct {
	EmailAddress *string `json:"email_address,omitempty" validate:"omitempty,email"`
	PhoneNumber  *string `json:"phone_number,omitempty" validate:"omitempty,min=10,max=15"`
	Name         string  `json:"name,omitempty" validate:"required,min=3,max=50"`
	Password     string  `json:"password,omitempty" validate:"required,min=6"`
	Gender       bool    `json:"gender,omitempty"`
	DateOfBirth  int64   `json:"date_of_birth,omitempty"`
}

type CustomerUpdateRequest struct {
	ID           string  `json:"id,omitempty" validate:"required"`
	EmailAddress *string `json:"email_address,omitempty" validate:"omitempty,email"`
	PhoneNumber  *string `json:"phone_number,omitempty" validate:"omitempty,min=10,max=15"`
	Name         string  `json:"name,omitempty" validate:"required;min=3,max=50"`
	Password     string  `json:"password,omitempty" validate:"required,min=6"`
	Gender       bool    `json:"gender,omitempty"`
	DateOfBirth  int64   `json:"date_of_birth,omitempty"`
}

type CustomerLoginRequest struct {
	EmailAddress *string `json:"email_address,omitempty" validate:"omitempty,email"`
	PhoneNumber  *string `json:"phone_number,omitempty" validate:"omitempty,min=10,max=15"`
	Password     string  `json:"password,omitempty" validate:"required,min=6"`
}

type CustomerDeleteRequest struct {
	ID string `json:"id,omitempty" validate:"required"`
}
