package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespJsonData struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func RespJson(c *gin.Context, code int, msg string, data interface{}) {
	result := &RespJsonData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}

func Json(c *gin.Context, code int, data interface{}) {
	result := &RespJsonData{
		Code: code,
		Msg:  getMessage(code),
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}

func AbortJson(c *gin.Context, code int, data interface{}) {
	result := &RespJsonData{
		Code: code,
		Msg:  getMessage(code),
		Data: data,
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}
