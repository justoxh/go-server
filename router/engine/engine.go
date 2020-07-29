package engine

import (
	"net/http"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/justoxh/go-server/config"
	"github.com/justoxh/go-server/router/middleware"
)

func New(cfg *config.Configuration) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLog())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	corsConfig := getCorsConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	if cfg.Server.LimitConnection > 0 {
		r.Use(limit.MaxAllowed(cfg.Server.LimitConnection))
	}
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"result": false,
			"error":  "Method Not Allowed",
		})
		return
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"result": false,
			"error":  "Endpoint Not Found",
		})
		return
	})
	return r
}

func getCorsConfig() cors.Config {
	return cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Forwarded-For", "User-Agent", "Referer"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}
