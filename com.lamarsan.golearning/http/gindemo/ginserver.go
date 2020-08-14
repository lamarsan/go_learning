package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)
import "go.uber.org/zap"

/**
 * description gin demo
 *
 * @author lamar
 * @date 2020/8/8 5:24 下午
 */
const keyRequestId = "requestId"

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// 做一个中间层，在进入具体的方法之前，先做一个拦截，打个日志
	// context可以传递参数
	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run()
}
