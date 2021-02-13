package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"web-stash-api/config"
)

func CorsMiddleware(cfg config.Config) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Server.Cors.AllowOrigins
	return cors.New(corsConfig)
}
