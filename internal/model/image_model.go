package model

import "mime/multipart"

type ImageResponse struct {
	ID           string `json:"id"`
	ProductId    string `json:"product_id,omitempty"`
	VarianId     string `json:"variant_id,omitempty"`
	ImageUrl     string `json:"image_url"`
	DisplayOrder int    `json:"display_order"`
}

type ImageRequest struct {
	ProductId    string                `json:"product_id,omitempty" form:"product_id"`
	VarianId     string                `json:"variant_id,omitempty" form:"variant_id"`
	Image        *multipart.FileHeader `form:"image"`
	DisplayOrder int                   `json:"display_order" form:"display_order"`
}
