// 处理高德地图天气查询的API

package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
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

// CastsResult 预报数据list结构，元素cast,按顺序为当天、第二天、第三天的预报数据,可预测未来三天的天气数据
type CastsResult struct {
	Data         string `json:"data"`
	Week         string `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
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
func GetWeatherForCity(city string) (weather WeatherResponse, err error) {
	u := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s",
		GetKey(),
		city,
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
func ConcurrentGetWeatherForCity(city string, wCh chan WeatherResponse, errCh chan error) {
	w, err := GetWeatherForCity(city)
	wCh <- w
	errCh <- err

}

// PrintWeatherResult 打印高德天气信息
func PrintWeatherResult(w interface{}, city string) {
	// 打印天气详情
	fmt.Printf("%s的天气:\n", city)
	for _, h := range w.([]LivesResult) {
		fmt.Print(h.Output())
	}
}
