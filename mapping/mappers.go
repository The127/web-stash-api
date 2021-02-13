package mapping

import (
	"web-stash-api/dtos"
	"web-stash-api/ent"
)

func MapBag(bag *ent.Bag) *dtos.BagDto {
	return &dtos.BagDto{
		Id:   bag.ID.String(),
		Name: bag.Name,
		Icon: bag.Icon,
	}
}

func MapItem(item *ent.BagItem) *dtos.ItemDto {
	return &dtos.ItemDto{
		Id:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		Icon:        item.Icon,
		Image:       item.Image,
	}
}
