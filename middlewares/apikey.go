package middlewares

import (
	"github.com/gin-gonic/gin"
	"web-stash-api/config"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

func validateToken(c *gin.Context) {
	apiKey := c.Request.Header.Get("X-Auth-Token")

	if apiKey == "" {
		c.AbortWithStatus(401)
	} else if checkToken(apiKey) {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}

func checkToken(apiKey string) bool {
	for _, key := range config.Cfg.Authentication.AllowedApiKeys {
		if key == apiKey {
			return true
		}
	}
	return false
}
