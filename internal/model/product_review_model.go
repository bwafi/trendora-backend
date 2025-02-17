package model

type ProductReviewRequest struct {
	ID         string  `json:"id,omitempty"`
	ProductId  string  `json:"product_id" validate:"required"`
	CustomerID string  `json:"customer_id" validate:"required"`
	Rating     float64 `json:"rating" validate:"required,gte=1,lte=5"`
	Comment    string  `json:"comment" validate:"required"`
}

type ProductReviewResponse struct {
	ID         string  `json:"id,omitempty"`
	ProductId  string  `json:"product_id" validate:"required"`
	CustomerID string  `json:"customer_id" validate:"required"`
	Rating     float64 `json:"rating" validate:"required,gte=1,lte=5"`
	Comment    string  `json:"comment" validate:"required"`
	CreatedAt  int64   `json:"created_at,omitempty"`
	UpdatedAt  int64   `json:"updated_at,omitempty"`
}
