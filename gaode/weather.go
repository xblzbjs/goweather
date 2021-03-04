// @Title  			weather.go
// @Description  	处理高德地图天气查询的API
// @Author  		xblzbjs
// @Update  		2021-3-4
package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xblzbjs/goweather/config"
)


// LivesResult 实况天气数据信息
type LivesResult struct {
	Province      string
	City          string
	Adcode        string
	Weather       string
	Temperature   string
	Winddirection string
	Windpower     string
	Humidity      string
	Reporttime    string
}

// ForecastResult 预报天气信息数据
type ForecastResult struct {
	City       string
	Adcode     string
	Province   string
	Reporttime string
	Casts      []CastsResult
}

// CastsResult 预报数据list结构，元素cast,按顺序为当天、第二天、第三天的预报数据
type CastsResult struct {
	Data         string
	Week         string
	Dayweather   string
	Nightweather string
	Daytemp      string
	Nighttemp    string
	Daywind      string
	Daypower     string
	Nightpower   string
}

// WeatherResponse	天气API响应结构
type WeatherResponse struct {
	Status   string
	Count    string
	Info     string
	Infocode string
	Lives    []LivesResult
	Forecast []ForecastResult
}

// getWeatherForCity 获取城市天气
func getWeatherForCity(city string) (weather WeatherResponse, err error) {
	u := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s",
		config.GaoDeApiKey,
		city,
	)
	res, err := httpClient.Get(u)
	if err != nil {
		return weather, err
	}
	defer res.Body.Close()
	if res.Status != "1"{
		return weather, errors.New(fmt.Sprintf("GaoDeWeatherRequest Failed: %s", res.Status))
	}
	return weather, json.NewDecoder(res.Body).Decode(&weather)
}
