package dtos

type BagList struct {
	Bags []*BagDto `json:"bags"`
}

type BagDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CreateBagDto struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

type UpdateBagDto struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type DeleteBagDto struct {
}
