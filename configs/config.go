package configs

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// Global 全局配置
var Global = GetConfig()
// Key   密钥配置
var Key = GetConfig().Key

// Config 全局配置结构题
type Config struct {
	AppName string `ini:"app_name"`
	LogLevel string `ini:"log_level"`

	Key KeyConfig `ini:'key'`
}

// KeyConfig 密钥配置结构体
type KeyConfig struct {
	GaoDeKey string `ini:"gao_de_api"`
	OpenWeatherKey string `ini:"open_weather_api"`
}

// GetConfig 获取 my.ini 配置参数
func GetConfig()(c Config){
	cfg, err := ini.Load("/my.ini")
	if err != nil{
		fmt.Printf("配置config.ini文件读取错误：%v", err)
		os.Exit(1)
	}
	c = Config{}
	err = cfg.MapTo(&c)
	if err != nil{
		fmt.Printf("配置映射失败：%v", err)
	}
	return c
}