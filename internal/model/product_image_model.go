package model

type ProductImageResponse struct {
	ID           string `json:"id"`
	ProductId    string `json:"product_id"`
	VarianId     string `json:"variant_id"`
	ImageUrl     string `json:"image_url"`
	DisplayOrder int    `json:"display_order"`
}
