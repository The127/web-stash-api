package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"web-stash-api/database"
	"web-stash-api/dtos"
	"web-stash-api/ent/bag"
	"web-stash-api/mapping"
)

type BagController struct{}

func (b BagController) GetBagItems(c *gin.Context) {
	bagId := c.Param("bagId")

	if bagId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing bagId"})
		c.Abort()
		return
	}

	bagUUID, err := uuid.Parse(bagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bag id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.Bag.
		Query().
		Where(bag.IDEQ(bagUUID)).
		WithItems().
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	bag := all[0]

	itemList := dtos.ItemList{
		Items: make([]*dtos.ItemDto, len(bag.Edges.Items)),
	}
	for i := range bag.Edges.Items {
		itemList.Items[i] = mapping.MapItem(bag.Edges.Items[i])
	}

	c.JSON(http.StatusOK, itemList)
}

func (b BagController) GetBags(c *gin.Context) {
	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.Bag.
		Query().
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	bagList := dtos.BagList{
		Bags: make([]*dtos.BagDto, len(all)),
	}
	for i := range all {
		bagList.Bags[i] = mapping.MapBag(all[i])
	}

	c.JSON(http.StatusOK, bagList)
}

func (b BagController) GetBag(c *gin.Context) {
	bagId := c.Param("bagId")

	if bagId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing bagId"})
		c.Abort()
		return
	}

	bagUUID, err := uuid.Parse(bagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bag id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.Bag.
		Query().
		Where(bag.IDEQ(bagUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "bag not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapBag(all[0]))
}

func (b BagController) CreateBag(c *gin.Context) {
	var createBagDto dtos.CreateBagDto
	if err := c.ShouldBind(&createBagDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	//TODO: real user id for multi tenancy
	userId, _ := uuid.Parse("52940bde-6e33-11eb-9439-0242ac130002")

	result, err := db.Bag.Create().
		SetIcon(createBagDto.Icon).
		SetName(createBagDto.Name).
		SetUserID(userId).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, mapping.MapBag(result))
}

func (b BagController) DeleteBag(c *gin.Context) {
	bagId := c.Param("bagId")

	if bagId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing bagId"})
		c.Abort()
		return
	}

	var deleteBagDto dtos.DeleteBagDto
	if err := c.ShouldBind(&deleteBagDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	bagUUID, err := uuid.Parse(bagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bag id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.Bag.
		Query().
		Where(bag.IDEQ(bagUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "bag not found"})
		c.Abort()
		return
	}

	bag := all[0]

	err = db.Bag.
		DeleteOne(bag).
		Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapBag(bag))
}

func (b BagController) UpdateBag(c *gin.Context) {
	bagId := c.Param("bagId")

	if bagId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing bagId"})
		c.Abort()
		return
	}

	var updateBagDto dtos.UpdateBagDto
	if err := c.ShouldBind(&updateBagDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	bagUUID, err := uuid.Parse(bagId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bag id cannot be parsed as valid uuid"})
		c.Abort()
		return
	}

	db := database.OpenDb()
	defer db.Close()

	ctx := context.Background()
	all, err := db.Bag.
		Query().
		Where(bag.IDEQ(bagUUID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "bag not found"})
		c.Abort()
		return
	}

	bag := all[0]
	update := bag.Update()

	if updateBagDto.Icon != "" {
		update.SetIcon(updateBagDto.Icon)
	}
	if updateBagDto.Name != "" {
		update.SetName(updateBagDto.Name)
	}

	result, err := update.Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, mapping.MapBag(result))
}
