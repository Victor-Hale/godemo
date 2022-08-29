package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	engine := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	engine.GET("/okk", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin get method")
	})
	// Post 请求路由
	engine.POST("/1", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin post method")
	})

	// 3.监听端口，默认在8080
	err := engine.Run(":8080")
	if err != nil {
		return
	}




}
