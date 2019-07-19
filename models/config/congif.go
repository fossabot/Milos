package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AliBucketName      string
	AliAccessKeyId     string
	AliAccessKeySecret string
	AliEndpoint        string
	UserName           string
	Password           string
	HashKey            string
	BlockKey           string
}

var AppConfig Config

func init() {
	//读配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		log.Fatal("Error read config:", err)
	}
	viper.Unmarshal(&AppConfig)
}
