package model

type ProductSizeResponse struct {
	ID            string  `json:"id"`
	VariantId     string  `json:"variant_id"`
	SKU           string  `json:"sku"`
	Size          string  `json:"size"`
	Discount      float32 `json:"discount"`
	Price         float32 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	IsAvailable   bool    `json:"column:is_available"`
}

type ProductSizeRequest struct {
	VariantId     string  `json:"variant_id"`
	SKU           string  `json:"sku"`
	Size          string  `json:"size"`
	Discount      float32 `json:"discount"`
	Price         float32 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}
