package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"goweather/gaode"
	"net/http"
	"net/url"
	"testing"
)

func TestGetLocationForAddress(t *testing.T) {
	var httpClient http.Client
	var geocode = gaode.GeocodeResponse{}

	escAddress := url.QueryEscape("北京市朝阳区阜通东大街6号")
	escCity := url.QueryEscape("北京")
	u := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s&city=%s",
		"5b17489a1ab8e8034e8546a7389e5ff6",
		escAddress,
		escCity,
	)
	res, err := httpClient.Get(u)
	if err != nil {
		t.Log(err)
	}

	defer res.Body.Close()

	// 定义高德地图地理编码响应对象
	err = json.NewDecoder(res.Body).Decode(&geocode)
	if err != nil {
		t.Log(err)
	}
	if geocode.Status != "1" || len(geocode.Geocodes) < 1 {
		errors.New(fmt.Sprintf("GetLocationRequest Failed: %s", geocode.Status))
	}
	t.Log(geocode.Geocodes[0].Location)
	t.Logf("Status:%s", geocode.Status)
}
