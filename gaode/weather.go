// 处理高德地图天气查询的API

package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	WeatherPeriodBase = "base"
	WeatherPeriodAll  = "all"
)

// LivesResult 实况天气数据信息
type LivesResult struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}

func (live LivesResult) Output() string {
	var unitAbbr string
	unitAbbr = "C"

	return fmt.Sprintf("%s%s | 温度:%s%s | 风向与风力:%s %s | 湿度:%s | 更新时间:%s \n",
		live.Province, live.City,
		live.Temperature, unitAbbr,
		live.Winddirection, live.Windpower,
		live.Humidity,
		live.Reporttime,
	)
}

// ForecastResult 预报天气信息数据
type ForecastResult struct {
	City       string        `json:"city"`
	Adcode     string        `json:"adcode"`
	Province   string        `json:"province"`
	Reporttime string        `json:"reporttime"`
	Casts      []CastsResult `json:"casts"`
}

func (forecast ForecastResult) Output() string {
	forecast.Reporttime = forecast.Reporttime[0:10]
	return fmt.Sprintf("%s%s", forecast.City, forecast.Reporttime)
}

// CastsResult 预报数据list结构，元素cast,按顺序为当天、第二天、第三天的预报数据,可预测未来三天的天气数据
type CastsResult struct {
	Data         string `json:"data"`
	Week         string `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Nightwind    string `json:"nightwind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
}

func (cast CastsResult) Output() string {

	switch cast.Week {
	case "1":
		cast.Week = "一"
	case "2":
		cast.Week = "二"
	case "3":
		cast.Week = "三"
	case "4":
		cast.Week = "四"
	case "5":
		cast.Week = "五"
	case "6":
		cast.Week = "六"
	case "7":
		cast.Week = "日"

	}

	return fmt.Sprintf("(星期%s):\n"+
		"白天:\t%s°C %s | 风向与风力:%s %s \n"+
		"夜晚:\t%s°C %s | 风向与风力:%s %s \n",
		cast.Week,
		cast.Daytemp, cast.Dayweather, cast.Daywind, cast.Daypower,
		cast.Nighttemp, cast.Nightweather, cast.Nightwind, cast.Nightpower,
	)
}

// WeatherResponse	天气API响应结构
type WeatherResponse struct {
	Status    string            `json:"status"`
	Count     string            `json:"count"`
	Info      string            `json:"info"`
	Infocode  string            `json:"infocode"`
	Lives     *[]LivesResult    `json:"lives"`
	Forecasts *[]ForecastResult `json:"forecasts"`
}

// GetWeatherForCity 获取城市天气
func GetWeatherForCity(city string, extensions string) (weather WeatherResponse, err error) {
	u := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s&extensions=%s",
		GetKey(),
		city,
		extensions,
	)
	res, err := httpClient.Get(u)
	if err != nil {
		return weather, err
	}
	defer res.Body.Close()
	if res.Status != "200 OK" {
		return weather, errors.New(fmt.Sprintf("GaoDeWeatherRequest Failed: %s", res.Status))
	}
	return weather, json.NewDecoder(res.Body).Decode(&weather)
}

// ConcurrentGetWeatherForCity 异步获取城市天气
func ConcurrentGetWeatherForCity(city string, extensions string, wCh chan WeatherResponse, errCh chan error) {
	w, err := GetWeatherForCity(city, extensions)
	wCh <- w
	errCh <- err

}

// PrintWeatherResult 打印高德天气信息
func PrintWeatherResult(w interface{}, city string) {
	// 打印天气详情
	switch w.(type) {
	case []LivesResult:
		for _, e := range w.([]LivesResult) {
			fmt.Print(e.Output())
		}
	case []ForecastResult:
		for _, result := range w.([]ForecastResult) {
			for _, cast := range result.Casts {
				fmt.Print(result.Output())
				fmt.Print(cast.Output())
			}
		}
	}

}
