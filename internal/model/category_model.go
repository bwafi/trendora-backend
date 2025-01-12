package model

type CategoryResponse struct {
	ID                 string             `json:"id,omitempty"`
	ParentId           string             `json:"parent_id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Slug               string             `json:"slug,omitempty"`
	ParentCategory     *CategoryResponse  `json:"parent_category,omitempty"`
	ChildrenCategories []CategoryResponse `json:"children_category,omitempty"`
}
