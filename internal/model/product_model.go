package model

type ProductResponse struct {
	ID             string                   `json:"id,omitempty"`
	StyleCode      string                   `json:"style_code"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	Gender         string                   `json:"gender"`
	CategoryId     string                   `json:"category_id"`
	SubCategoryId  string                   `json:"sub_category_id"`
	BasePrice      float32                  `json:"base_price"`
	IsVisible      bool                     `json:"is_visible"`
	ReleaseDate    int64                    `json:"release_date"`
	Category       CategoryResponse         `json:"category"`
	ProductVariant []ProductVariantResponse `json:"variant"`
	Image          ProductImageResponse     `json:"image"`
	CreatedAt      int64                    `json:"created_at"`
	UpdatedAt      int64                    `json:"updated_at"`
}

type CreateProductRequest struct {
	ID             string                   `json:"id,omitempty"`
	StyleCode      string                   `json:"style_code" validate:"required"`
	Name           string                   `json:"name" validate:"required"`
	Description    string                   `json:"description"`
	Gender         string                   `json:"gender"`
	CategoryId     string                   `json:"category_id" validate:"required"`
	SubCategoryId  string                   `json:"sub_category_id"`
	BasePrice      float32                  `json:"base_price" validate:"required"`
	IsVisible      bool                     `json:"is_visible" validate:"required"`
	ReleaseDate    int64                    `json:"release_date" validate:"required"`
	Category       CategoryResponse         `json:"category"`
	ProductVariant []ProductVariantResponse `json:"variant"`
	Image          ProductImageResponse     `json:"image"`
}
