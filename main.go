package main

import (
	"Milos/routers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 设置日志文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.DefaultWriter = io.Writer(os.Stdout)
	// 使用日志中间件
	r.Use(gin.Logger())
	// 设置静态文件夹
	r.Static("/static", "./static")
	r.Static("/favicon.ico","./static/favicon.ico")
	// 加载路由
	routers.LoadRouters(r)
	r.Run(":8080")
}
