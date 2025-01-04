package model

type ProductVariantResponse struct {
	ID            string          `json:"id"`
	ProductId     string          `json:"product_id"`
	SKU           string          `json:"sku"`
	ColorName     string          `json:"color_name"`
	Size          string          `json:"size"`
	Discount      float32         `json:"discount"`
	Price         float32         `json:"price"`
	StockQuantity int             `json:"stock_quantity"`
	Weight        float32         `json:"weight"`
	IsAvailable   bool            `json:"is_available"`
	VariantImages []ImageResponse `json:"variant_images"`
}

type ProductVariantRequest struct {
	ProductId     string         `json:"product_id" form:"product_id"`
	SKU           string         `json:"sku" form:"sku"`
	ColorName     string         `json:"color_name" form:"color_name"`
	Size          string         `json:"size" form:"size"`
	Discount      float32        `json:"discount" form:"discount"`
	Price         float32        `json:"price" form:"price"`
	StockQuantity int            `json:"stock_quantity" form:"stock_quantity"`
	Weight        float32        `json:"weight" form:"weight"`
	IsAvailable   bool           `json:"is_available" form:"is_available"`
	VariantImages []ImageRequest `json:"variant_images" form:"variant_images"`
}
