// @Title  			geocode.go
// @Description  	处理高德地图地理编码的API
// @Author  		xblzbjs
// @Update  		2021-3-4

package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
	"goweather/config"
	"log"
	"net/url"

	"github.com/spf13/viper"
)

// GeocodeResult 地理编码结构体
type GeocodeResult struct {
	FormattedAddress string
	Country          string
	Province         string
	Citycode         string
	City             string
	District         string
	// TODO：添加township,neighborhood结构体,building结构体
	Adcode   string
	Street   string
	Number   string
	Location string
	Level    string
}

// GeocodeResponse 高德地图地理编码响应对象
type GeocodeResponse struct {
	Status   string
	Count    string
	Info     string
	Infocode string
	Geocodes []GeocodeResult
}

func init() {
	var configuration config.Configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

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

// getLocationForAddress 将详细的结构化地址转换为高德经纬度坐标
func getLocationForAddress(address string, city string) (location string, err error) {
	escAddress := url.QueryEscape(address)
	escCity := url.QueryEscape(city)
	u := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s&city=%s",
		gaoDeKey,
		escAddress,
		escCity,
	)
	res, err := httpClient.Get(u)
	if err != nil {
		return location, err
	}

	defer res.Body.Close()

	// 定义高德地图地理编码响应对象
	err = json.NewDecoder(res.Body).Decode(&geocode)
	if err != nil {
		return location, err
	}
	if geocode.Status != "1" || len(geocode.Geocodes) < 1 {
		return location, errors.New(fmt.Sprintf("GetLocationRequest Failed: %s", geocode.Status))
	}

	return geocode.Geocodes[0].Location, err
}
