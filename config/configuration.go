package config

import "github.com/spf13/viper"

// Configuration 全局配置
type Configuration struct {
	Api ApiConfiguration
}

var viperConfig *viper.Viper
