package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justoxh/go-server/common/render"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		render.Json(c, render.Ok, "pong")
	}
}
