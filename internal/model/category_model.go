package model

type CategoryResponse struct {
	ID                 string             `json:"id"`
	ParentId           string             `json:"parent_id"`
	Name               string             `json:"name"`
	Slug               string             `json:"slug"`
	ParentCategory     *CategoryResponse  `json:"parent"`
	ChildrenCategories []CategoryResponse `json:"children"`
}
