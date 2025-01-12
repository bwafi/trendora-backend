package model

type CartItemResponse struct {
	ID         string `json:"id"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	VariantId  string `json:"variant_id"`
	Quantity   int    `json:"quantity"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`

	Customers      *CustomerResponse       `json:"customer,omitempty"`
	Product        *ProductResponse        `json:"products,omitempty"`
	ProductVariant *ProductVariantResponse `json:"variant,omitempty"`
}

type CartItemRequest struct {
	CustomerId string `json:"customer_id,omitempty" validate:"required"`
	ProductId  string `json:"product_id,omitempty" validate:"required"`
	VariantId  string `json:"variant_id,omitempty" validate:"required"`
	Quantity   int    `json:"quantity,omitempty" validate:"required,gte=1"`
}

type CartItemUpdateRequest struct {
	ID        string `json:"id,omitempty" validate:"required"`
	Quantity  int    `json:"quantity,omitempty" validate:"required,gte=1"`
	Operation string `json:"operation,omitempty" validate:"required,oneof=INCREASE DECREASE"`
}

type CartItemDeleteRequest struct {
	ID string `json:"id,omitempty" validate:"required"`
}

type CartItemGetRequest struct {
	ID         string `json:"id,omitempty" validate:"required"`
	CustomerId string `json:"customer_id,omitempty" validate:"required"`
}
