package dtos

type SubItemList struct {
	SubItems []*SubItemDto `json:"subItems"`
}

type SubItemDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Link        string `json:"link"`
}

type CreateSubItemDto struct {
	ItemId      string `json:"itemId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

type UpdateSubItemDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Link        string `json:"link"`
}

type DeleteSubItemDto struct {
}
