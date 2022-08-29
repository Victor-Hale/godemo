package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 读取 cookie
	engine.GET("/read_cookie", func(context *gin.Context) {
		val, _ := context.Cookie("name")
		context.String(200, "Cookie:%s", val)
	})

	// 写入 cookie
	engine.GET("/write_cookie", func(context *gin.Context) {
		context.SetCookie("name", "Shimin Li", 24*60*60, "/", "localhost", false, true)
	})

	// 清理 cookie
	engine.GET("/clear_cookie", func(context *gin.Context) {
		context.SetCookie("name", "Shimin Li", -1, "/", "localhost", false, true)
	})

	engine.Run()
}