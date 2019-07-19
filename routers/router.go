package routers

import (
	ctrs "Milos/controllers"

	"github.com/gin-gonic/gin"
)

// LoadRouters 初始化router
func LoadRouters(router *gin.Engine) {
	loadRouters(router)
}

func loadRouters(router *gin.Engine) {

	// 路由控制函数，我们全部放在controllers目录下
	router.GET("/", ctrs.AuthMiddleWare(), ctrs.Index)
	router.POST("/", ctrs.AuthMiddleWare(), ctrs.PostFile)
	router.GET("/about", ctrs.About)
	router.GET("/delete/*objectName", ctrs.DeleteFile)
	router.GET("/loginpage", ctrs.LoginPage)
	router.POST("/loginpage", ctrs.LoginPage)
	router.POST("/login", ctrs.Login)

	// ......  很多很多路由。。。
}
