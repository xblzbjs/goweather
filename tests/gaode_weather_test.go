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
		"5b17489a1ab8e8034e8546a7389e5ff6",
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
	if res.Status != "1" {
		t.Logf("GaoDeWeatherRequest Failed: %s", res.Status)
	}
	t.Log(weather)
	t.Log(json.NewDecoder(res.Body).Decode(&weather))
}
