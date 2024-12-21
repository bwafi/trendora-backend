package model

type AddressResponse struct {
	ID            string `json:"id"`
	CustomerID    string `json:"customer_id,omitempty"`
	RecipientName string `json:"recipient_name,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	AddressType   string `json:"address_type,omitempty"`
	City          string `json:"city,omitempty"`
	Province      string `json:"province,omitempty"`
	SubDistrict   string `json:"sub_district,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	CourierNotes  string `json:"courier_notes,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
	UpdatedAt     int64  `json:"updated_at,omitempty"`
}

type CreateAddressRequest struct {
	CustomerID    string `json:"customer_id,omitempty"`
	RecipientName string `json:"recipient_name,omitempty"`
	PhoneNumber   string `json:"phone_number,"`
	AddressType   string `json:"address_type"`
	City          string `json:"city"`
	Province      string `json:"province"`
	SubDistrict   string `json:"sub_district"`
	PostalCode    string `json:"postal_code"`
	CourierNotes  string `json:"courier_notes,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
	UpdatedAt     int64  `json:"updated_at,omitempty"`
}

type UpdateAddressRequest struct {
	ID            string `json:"id,omitempty"`
	CustomerID    string `json:"customer_id,omitempty"`
	RecipientName string `json:"recipient_name,omitempty"`
	PhoneNumber   string `json:"phone_number,"`
	AddressType   string `json:"address_type"`
	City          string `json:"city"`
	Province      string `json:"province"`
	SubDistrict   string `json:"sub_district"`
	PostalCode    string `json:"postal_code"`
	CourierNotes  string `json:"courier_notes,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
	UpdatedAt     int64  `json:"updated_at,omitempty"`
}

type DeleteAddressRequest struct {
	ID         string `json:"id,omitempty" validate:"required"`
	CustomerID string `json:"customer_id,omitempty" validate:"red"`
}

type GetAddressRequest struct {
	ID         string `json:"id,omitempty"`
	CustomerID string `json:"customer_id"`
}
