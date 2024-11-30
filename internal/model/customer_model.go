package model

type CustomerResponse struct {
	ID           string `json:"id,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	Token        string `json:"token,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
}
