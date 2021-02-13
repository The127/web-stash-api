package server

import (
	"github.com/gin-gonic/gin"
	"web-stash-api/config"
	"web-stash-api/controllers"
	"web-stash-api/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		bagGroup := v1.Group("bags")
		{
			bag := new(controllers.BagController)
			bagGroup.GET("/", bag.GetBags)
			bagGroup.GET("/:bagId", bag.GetBag)
			bagGroup.POST("/", bag.CreateBag)
			bagGroup.PUT("/:bagId", bag.UpdateBag)
			bagGroup.DELETE("/:bagId", bag.DeleteBag)
		}

		itemGroup := bagGroup.Group("items")
		{
			item := new(controllers.ItemController)
			itemGroup.GET("/:itemId", item.GetItem)
			itemGroup.POST("/", item.CreateItem)
			itemGroup.PUT("/:itemId", item.UpdateItem)
			itemGroup.DELETE("/:itemId", item.DeleteItem)
		}
	}

	router.Use(middlewares.CorsMiddleware(config.Cfg))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
