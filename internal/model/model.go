package model

type WebResponse[T any] struct {
	Status  string         `json:"data,omitempty"`
	Message string         `json:"data,omitempty"`
	Data    T              `json:"data"`
	Paging  *PageMetadata  `json:"paging,omitempty"`
	Errors  *ErrorResponse `json:"errors,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
