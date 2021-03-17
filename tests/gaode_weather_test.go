package tests

import (
	"encoding/json"
	"fmt"
	"goweather/gaode"
	"net/http"
	"net/url"
	"testing"
)

// TestGetWeatherForCity 获取城市天气的测试
func TestGetWeatherForCity(t *testing.T) {
	var httpClient http.Client
	var weather = gaode.WeatherResponse{}
	escCity := url.QueryEscape("深圳")
	t.Log(gaode.GetKey())
	u := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?extensions=all&key=%s&city=%s",
		gaode.GetKey(),
		escCity,
	)
	res, err := httpClient.Get(u)
	if err != nil {
		t.Log(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		t.Log(err)
	}
	if res.Status != "200 OK" {
		t.Logf("GaoDeWeatherRequest Failed: %s", res.Status)
	}
	t.Logf("今天的天气:%v", (weather.Forecasts[0].Casts[0]))
	t.Logf("明天的天气:%v", weather.Forecasts[0].Casts[1])
	t.Logf("后天的天气:%v", weather.Forecasts[0].Casts[2])
	t.Logf("大后天的天气:%v", weather.Forecasts[0].Casts[3])
}
