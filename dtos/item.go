package dtos

type ItemList struct {
	Items []*ItemDto `json:"items"`
}

type ItemDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
	Link        string `json:"link"`
}

type CreateItemDto struct {
	BagId       string `json:"bagId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

type UpdateItemDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
	Link        string `json:"link"`
}

type DeleteItemDto struct {
}
