package controllers

import (
	"fmt"
	"image"
	_ "image/gif" //必须导入解码包
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"path"

	"Milos/models/aliyun"

	"github.com/gin-gonic/gin"
)

//PostFile 上传单个文件
func PostFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("smfile") //在表单的 input name="uploadpic" 指定key
	if err != nil {
		log.Fatal("getfile err:", err)
	}
	defer file.Close()

	// 获取图片属性
	imgconf, _, err := image.DecodeConfig(file)
	fmt.Println("width = ", imgconf.Width)
	fmt.Println("height = ", imgconf.Height)

	file.Seek(0, os.SEEK_SET) //重置文件指针
	fileurl, objectname := aliyun.Upload(file, path.Ext(header.Filename))
	delurl := "delete/" + objectname

	// TODO:实现SM的API

	c.JSON(http.StatusOK, gin.H{
		"code": "success",
		"data": gin.H{
			"width":    imgconf.Width,
			"height":   imgconf.Height,
			"filename": header.Filename,
			"size":     header.Size,
			"ip":       c.ClientIP(),
			"url":      fileurl,
			"delete":   delurl,
		},
	})
}
