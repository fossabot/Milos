package controllers

import (
	"Milos/models/config"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

// 设置cookie加密的密钥
var hashKey = []byte(config.AppConfig.HashKey)
var blockKey = []byte(config.AppConfig.BlockKey)
var s = securecookie.New(hashKey, blockKey)

// Login 处理登录请求，设置cookie
func Login(c *gin.Context) {
	username := config.AppConfig.UserName
	password := config.AppConfig.Password

	c.Request.ParseForm() //解析表单

	log.Println("POST Form:", c.Request.PostFormValue("username"), c.Request.PostFormValue("password"))
	if c.Request.PostFormValue("username") == username && c.Request.PostFormValue("password") == password {
		value := map[string]string{
			"password": password,
			"date":     fmt.Sprintf("%d", time.Now().Unix()),
		}

		encoded, _ := s.Encode("session_id", value)

		//设置一个1月过期的Cookie
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    encoded,
			HttpOnly: true,
			MaxAge:   2629743,
		}
		http.SetCookie(c.Writer, cookie)
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta http-equiv="refresh" content="10">
	<meta http-equiv="refresh" content="1;url=/"> 
	</head>
</html>`) //跳转页面
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
	}

}

// AuthMiddleWare 路由中间件，验证cookie是否正确
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {

			value := make(map[string]string)
			err = s.Decode("session_id", cookie.Value, &value)
			if err != nil {
				log.Println("cookie decode error:", err)
			}
			log.Println("date is:", value["date"])
			cookietime, _ := strconv.ParseInt(value["date"], 10, 64)

			if value["password"] == config.AppConfig.Password && time.Now().Unix()-cookietime <= 2629743 { //密码相同且未过期
				log.Println("Auth passed!")
				c.Next()
				return
			}
		}
		c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
		c.Abort()
		return
	}
}
