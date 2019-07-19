package controllers

import (
	"html/template"

	"time"

	"github.com/gin-gonic/gin"
)

// Index 主页面
func Index(c *gin.Context) {
	html := template.Must(template.ParseFiles("views/index.html"))

	data := struct {
		Year int
	}{
		Year: time.Now().Year(),
	}

	html.Execute(c.Writer, data)
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
