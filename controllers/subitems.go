package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"web-stash-api/database"
	"web-stash-api/dtos"
	"web-stash-api/ent/subitem"
	"web-stash-api/mapping"
)

type SubItemController struct{}

func (s SubItemController) GetSubItem(c *gin.Context) {
	subItemId := c.Param("subItemId")
	if subItemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing subItemId"})
		c.Abort()
		return
	}

	subItemUUID, err := uuid.Parse(subItemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sub item id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.SubItem.
		Query().
		Where(subitem.IDEQ(subItemUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "sub item not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapSubItem(all[0]))
}

func (s SubItemController) CreateSubItem(c *gin.Context) {
	var createSubItemDto dtos.CreateSubItemDto
	if err := c.ShouldBind(&createSubItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	parentUUID, err := uuid.Parse(createSubItemDto.ItemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	subItem, err := db.SubItem.Create().
		SetName(createSubItemDto.Name).
		SetIcon(createSubItemDto.Icon).
		SetDescription(createSubItemDto.Description).
		SetParentID(parentUUID).
		SetLink(createSubItemDto.Link).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapSubItem(subItem))
}

func (s SubItemController) DeleteSubItem(c *gin.Context) {
	subItemId := c.Param("subItemId")
	if subItemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing subItemId"})
		c.Abort()
		return
	}

	subItemUUID, err := uuid.Parse(subItemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sub item id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	var deleteSubItemDto dtos.DeleteSubItemDto
	if err := c.ShouldBind(&deleteSubItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.SubItem.
		Query().
		Where(subitem.IDEQ(subItemUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "sub item not found"})
		c.Abort()
		return
	}

	subItem := all[0]

	err = db.SubItem.
		DeleteOne(subItem).
		Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapSubItem(subItem))
}

func (s SubItemController) UpdateSubItem(c *gin.Context) {
	subItemId := c.Param("subItemId")
	if subItemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing subItemId"})
		c.Abort()
		return
	}

	subItemUUID, err := uuid.Parse(subItemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sub item id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	var updateSubItemDto dtos.UpdateSubItemDto
	if err := c.ShouldBind(&updateSubItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.SubItem.
		Query().
		Where(subitem.IDEQ(subItemUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "sub item not found"})
		c.Abort()
		return
	}

	subItem := all[0]
	update := subItem.Update()

	if updateSubItemDto.Description != "" {
		update.SetDescription(updateSubItemDto.Description)
	}
	if updateSubItemDto.Icon != "" {
		update.SetIcon(updateSubItemDto.Icon)
	}
	if updateSubItemDto.Name != "" {
		update.SetName(updateSubItemDto.Name)
	}
	if updateSubItemDto.Link != "" {
		update.SetLink(updateSubItemDto.Link)
	}

	result, err := update.Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapSubItem(result))
}
