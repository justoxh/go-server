package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justoxh/go-server/log"
	"github.com/sirupsen/logrus"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		var body []byte
		if c.Request.Body != nil {
			body, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()
		stop := time.Now()
		latency := stop.Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		if latency > time.Minute {
			latency = latency - latency%time.Second
		}
		logrusFields := logrus.Fields{
			"timestamp":   int(start.UnixNano() / 1e6),
			"status":      statusCode,
			"latency":     fmt.Sprintf("%s", latency),
			"client_ip":   clientIP,
			"method":      c.Request.Method,
			"path":        path,
			"referer":     referer,
			"data_length": dataLength,
			"userAgent":   clientUserAgent,
			"body":        string(body),
		}
		if len(c.Errors) > 0 {
			log.Log.WithFields(logrusFields).Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			log.Log.WithFields(logrusFields).Info("request_log")
		}
	}
}
