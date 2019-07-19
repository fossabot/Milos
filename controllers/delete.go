package controllers

import (
	"Milos/models/aliyun"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	objectName := c.Param("objectName")[1:] //必须删掉路径开头的“/”才能删除文件
	aliyun.Delete(objectName)
	log.Println(objectName)
	html := template.Must(template.ParseFiles("views/deleted.html"))

	html.Execute(c.Writer, nil)
}
