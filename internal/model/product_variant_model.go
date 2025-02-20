package model

type ProductVariantResponse struct {
	ID            string                 `json:"id"`
	ProductId     string                 `json:"product_id"`
	SKU           string                 `json:"sku"`
	ColorName     string                 `json:"color_name"`
	Weight        float32                `json:"weight"`
	IsAvailable   bool                   `json:"is_available"`
	VariantImages []*ImageResponse       `json:"variant_images"`
	ProductSizes  []*ProductSizeResponse `json:"product_sizes"`
}

type ProductVariantRequest struct {
	ProductId     string                `json:"product_id" form:"product_id"`
	SKU           string                `json:"sku" form:"sku"`
	ColorName     string                `json:"color_name" form:"color_name"`
	Weight        float32               `json:"weight" form:"weight"`
	IsAvailable   bool                  `json:"is_available" form:"is_available"`
	VariantImages []ImageRequest        `json:"variant_images" form:"variant_images"`
	ProductSizes  []ProductSizeResponse `json:"product_sizes" form:"product_sizes"`
}
