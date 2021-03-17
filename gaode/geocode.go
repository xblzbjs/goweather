// 处理高德地图地理编码的API
package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// GeocodeResult 地理编码结构体
type GeocodeResult struct {
	FormattedAddress string `json:"formatted_address"`
	Country          string `json:"country"`
	Province         string `json:"province"`
	Citycode         string `json:"citycode"`
	City             string `json:"city"`
	District         string `json:"district"`
	// TODO：添加township,neighborhood结构体,building结构体
	Adcode   string `json:"adcode"`
	Street   string `json:"street"`
	Number   string `json:"number"`
	Location string `json:"location"`
	Level    string `json:"level"`
}

// GeocodeResponse 高德地图地理编码响应对象
type GeocodeResponse struct {
	Status   string          `json:"status"`
	Count    string          `json:"count"`
	Info     string          `json:"info"`
	Infocode string          `json:"infocode"`
	Geocodes []GeocodeResult `json:"geocodes"`
}

// getLocationForAddress 将详细的结构化地址转换为高德经纬度坐标
func getLocationForAddress(address string, city string) (location string, err error) {
	escAddress := url.QueryEscape(address)
	escCity := url.QueryEscape(city)

	u := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s&city=%s",
		GetKey(),
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
