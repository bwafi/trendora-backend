package model

type CartItemResponse struct {
	ID         string `json:"id"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	VariantId  string `json:"variant_id"`
	Quantity   int    `json:"quantity"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`

	Customers      CustomerResponse       `json:"customer,omitempty"`
	Product        ProductResponse        `json:"products,omitempty"`
	ProductVariant ProductVariantResponse `json:"variant,omitempty"`
}

type CartItemRequest struct {
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	VariantId  string `json:"variant_id"`
	Quantity   int    `json:"quantity"`
}
