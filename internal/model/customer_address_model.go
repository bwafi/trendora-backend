package model

type AddressResponse struct {
	ID            string `json:"id"`
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
	ID            string `json:"id,omitempty"`
	CustomerID    string `json:"customer_id"`
	RecipientName string `json:"recipient_name"`
	PhoneNumber   string `json:"phone_number"`
	AddressType   string `json:"address_type"`
	City          string `json:"city"`
	Province      string `json:"province"`
	SubDistrict   string `json:"sub_district"`
	PostalCode    string `json:"postal_code"`
	CourierNotes  string `json:"courier_notes,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
	UpdatedAt     int64  `json:"updated_at,omitempty"`
}

type GetAddressRequest struct {
	ID         string `json:"id,omitempty"`
	CustomerID string `json:"customer_id"`
}
