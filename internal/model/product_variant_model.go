package model

type ProductVariantResponse struct {
	ID            string               `json:"id"`
	ProductId     string               `json:"product_id"`
	SKU           string               `json:"sku"`
	ColorName     string               `json:"color_name"`
	Size          string               `json:"size"`
	Discount      float32              `json:"discount"`
	Price         float32              `json:"price"`
	StockQuantity int                  `json:"stock_quantity"`
	Weight        float32              `json:"weight"`
	IsAvailable   bool                 `json:"is_available"`
	Product       ProductResponse      `json:"product"`
	Image         ProductImageResponse `json:"image"`
}
