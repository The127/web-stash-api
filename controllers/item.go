package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"web-stash-api/database"
	"web-stash-api/dtos"
	"web-stash-api/ent/bagitem"
	"web-stash-api/mapping"
)

type ItemController struct{}

func (i ItemController) GetItem(c *gin.Context) {
	itemId := c.Param("itemId")

	if itemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing itemId"})
		c.Abort()
		return
	}

	itemUUID, err := uuid.Parse(itemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.BagItem.
		Query().
		Where(bagitem.IDEQ(itemUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapItem(all[0]))

}

func (i ItemController) CreateItem(c *gin.Context) {
	var createItemDto dtos.CreateItemDto
	if err := c.ShouldBind(&createItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	bagId, err := uuid.Parse(createItemDto.BagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bag id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()

	result, err := db.BagItem.Create().
		SetName(createItemDto.Name).
		SetIcon(createItemDto.Icon).
		SetBagID(bagId).
		SetDescription(createItemDto.Description).
		SetImage(createItemDto.Image).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, mapping.MapItem(result))
}

func (i ItemController) DeleteItem(c *gin.Context) {
	itemId := c.Param("itemId")

	if itemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing itemId"})
		c.Abort()
		return
	}

	var deleteItemDto dtos.DeleteItemDto
	if err := c.ShouldBind(&deleteItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	itemUUID, err := uuid.Parse(itemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.BagItem.
		Query().
		Where(bagitem.IDEQ(itemUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		c.Abort()
		return
	}

	item := all[0]

	err = db.BagItem.
		DeleteOne(item).
		Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapItem(item))
}

func (i ItemController) UpdateItem(c *gin.Context) {
	itemId := c.Param("itemId")

	if itemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing itemId"})
		c.Abort()
		return
	}

	var updateItemDto dtos.UpdateItemDto
	if err := c.ShouldBind(&updateItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	itemUUID, err := uuid.Parse(itemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.BagItem.
		Query().
		Where(bagitem.IDEQ(itemUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		c.Abort()
		return
	}

	item := all[0]
	update := item.Update()

	if updateItemDto.Image != "" {
		update.SetImage(updateItemDto.Image)
	}
	if updateItemDto.Description != "" {
		update.SetDescription(updateItemDto.Description)
	}
	if updateItemDto.Icon != "" {
		update.SetIcon(updateItemDto.Icon)
	}
	if updateItemDto.Name != "" {
		update.SetName(updateItemDto.Name)
	}

	result, err := update.Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapItem(result))
}
