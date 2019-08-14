package controllers

import (
	"html/template"
	"io/ioutil"
	"log"

	"time"

	"github.com/gin-gonic/gin"
)

// Index 主页面
func Index(c *gin.Context) {
	b, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		log.Println("err:,no such file:views/index.html")
	}
	s := string(b)

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, s)
}

// LoginPage 登陆页面
func LoginPage(c *gin.Context) {
	html := template.Must(template.ParseFiles("views/login.html"))

	data := struct {
		Year int
	}{
		Year: time.Now().Year(),
	}

	html.Execute(c.Writer, data)
}

// About 关于页面
func About(c *gin.Context) {
	html := template.Must(template.ParseFiles("views/about.html"))

	data := struct {
		Year int
	}{
		Year: time.Now().Year(),
	}

	html.Execute(c.Writer, data)
}
