package model

type ProductResponse struct {
	ID             string                    `json:"id,omitempty"`
	StyleCode      string                    `json:"style_code,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Description    string                    `json:"description,omitempty"`
	Gender         string                    `json:"gender,omitempty"`
	CategoryId     string                    `json:"category_id,omitempty"`
	SubCategoryId  string                    `json:"sub_category_id,omitempty"`
	BasePrice      float32                   `json:"base_price,omitempty"`
	IsVisible      bool                      `json:"is_visible,omitempty"`
	ReleaseDate    int64                     `json:"release_date,omitempty"`
	Category       CategoryResponse          `json:"category,omitempty"`
	ProductVariant []*ProductVariantResponse `json:"variant,omitempty"`
	ProductImages  []*ImageResponse          `json:"product_images,omitempty"`
	CreatedAt      int64                     `json:"created_at,omitempty"`
	UpdatedAt      int64                     `json:"updated_at,omitempty"`
}

type CreateProductRequest struct {
	ID              string                  `json:"id,omitempty" form:"id"`
	StyleCode       string                  `json:"style_code" form:"style_code"`
	Name            string                  `json:"name" form:"name"`
	Description     string                  `json:"description" form:"description"`
	Gender          string                  `json:"gender" form:"gender"`
	CategoryId      string                  `json:"category_id" form:"category_id"`
	SubCategoryId   string                  `json:"sub_category_id" form:"sub_category_id"`
	BasePrice       float32                 `json:"base_price" form:"base_price"`
	IsVisible       bool                    `json:"is_visible" form:"is_visible"`
	ReleaseDate     int64                   `json:"release_date" form:"release_date"`
	ProductImages   []ImageRequest          `json:"product_images" form:"product_images"`
	ProductVariants []ProductVariantRequest `json:"product_variant" form:"product_variant"`
}

type ProductGetRequest struct {
	ID string `json:"id,omitempty"`
}

type ProductGetListRequest struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Gender        string `json:"gender,omitempty"`
	CategoryId    string `json:"category_id,omitempty"`
	SubCategoryId string `json:"sub_category_id,omitempty"`
}
