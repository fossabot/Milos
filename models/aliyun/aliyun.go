package aliyun

import (
	"Milos/models/config"
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client

func init() {
	// 创建OSSClient实例。
	var err error
	client, err = oss.New(config.AppConfig.AliEndpoint, config.AppConfig.AliAccessKeyId, config.AppConfig.AliAccessKeySecret)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

// upload上传图片到阿里云
func Upload(fd multipart.File, ext string) (string, string) {
	// 获取存储空间。
	bucket, err := client.Bucket(config.AppConfig.AliBucketName)
	if err != nil {
		log.Println("Error:", err)
	}

	fileBytes, err := ioutil.ReadAll(fd) //读取内容
	if err != nil {
		log.Println("readfile err:", err)
	}
	//fd.Seek(0, os.SEEK_SET) //重置文件指针
	name := fmt.Sprintf("%x", md5.Sum(fileBytes)) // 计算MD5

	directory := fmt.Sprintf("test/%d/%d-%d/", time.Now().Year(), time.Now().Month(), time.Now().Day()) //以时间分隔路径

	//拼路径
	objectname := directory + name + ext
	//log.Println("objectname:", objectname)

	// 上传文件。
	err = bucket.PutObject(objectname, bytes.NewReader(fileBytes))
	imgurl := "err"
	if err != nil {
		log.Println("Error:", err)
	} else {
		imgurl = "https://" + config.AppConfig.AliBucketName + "." + config.AppConfig.AliEndpoint + "/" + objectname
		log.Println("URL:", imgurl)
	}
	return imgurl, objectname
}

// Delete删除一个文件
func Delete(objectName string) {
	// 获取存储空间。
	bucket, err := client.Bucket(config.AppConfig.AliBucketName)
	if err != nil {
		log.Println("Error:", err)
	}
	// 删除单个文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		log.Println("Error:", err)
	}
}
