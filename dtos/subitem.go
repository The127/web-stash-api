package dtos

type SubItemList struct {
	SubItems []*SubItemDto `json:"subItems"`
}

type SubItemDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CreateSubItemDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
}

type UpdateSubItemDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type DeleteSubItemDto struct {
}
