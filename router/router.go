package router

import (
	"github.com/gin-gonic/gin"
	"github.com/justoxh/go-server/config"
	"github.com/justoxh/go-server/handler"
	"github.com/justoxh/go-server/router/engine"
)

func InitRouter(cfg *config.Configuration) *gin.Engine {
	r := engine.New(cfg)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handler.Ping())
	}
	return r
}
