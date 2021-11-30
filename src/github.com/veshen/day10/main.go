package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// create router use gin
	r := gin.Default()
	// 绑定路由规则
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/veshen", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "veshen",
		})
	})
	// 获取api参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	// 获取 query 参数
	r.GET("/user/query", func(c *gin.Context) {
		name := c.Query("name")
		action := c.Query("action")
		message := name + " is " + action
		c.String(200, message)
	})
	// 启动服务 默认8080端口
	r.Run()
}
