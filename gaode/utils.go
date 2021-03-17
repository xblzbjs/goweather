// @Title  			utils.go
// @Description  	定义高德地图地理API的辅助函数
// @Author  		xblzbjs
// @Update  		2021-3-4
package gaode

import (
	"goweather/config"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var httpClient http.Client
var geocode GeocodeResponse

// GetKey 获取key
func GetKey() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	gaoDeKey := configuration.Api.GaoDe
	return gaoDeKey
}

// LocationToLatLon 将字符串的经度维度转化成浮点数经度和维度
func LocationToLatLon(location string) (lat float64, lon float64) {
	locationList := strings.Split(location, ",")
	lon, _ = strconv.ParseFloat(locationList[0], 6)
	lat, _ = strconv.ParseFloat(locationList[1], 6)
	return lat, lon
}
