package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"web-stash-api/config"
)

func CorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.Cfg.Server.Cors.AllowOrigins
	return cors.New(corsConfig)
}
