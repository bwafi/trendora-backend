package model

import "time"

type CustomerSessionRequest struct {
	ID           string    `json:"id,omitempty"`
	CustomerID   string    `json:"customer_id"`
	RefreshToken string    `json:"refresh_token"`
	IsRevoked    string    `json:"is_revoked"`
	ExpiresAt    time.Time `json:"expires_at"`
}
