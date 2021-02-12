package dtos

type CreateBagDto struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

type UpdateBagDto struct {
	Name string `json:"name"`
	Icon string `json:"id"`
}
