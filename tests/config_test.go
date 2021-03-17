package tests

import (
	"goweather/config"
	"log"
	"testing"

	"github.com/spf13/viper"
)

// TestGetConfig 测试获取配置
func TestGetConfig(t *testing.T) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("gaode api key is %s", configuration.Api.GaoDe)
	log.Printf("openweather apikey is %s", configuration.Api.OpenWeather)
}
