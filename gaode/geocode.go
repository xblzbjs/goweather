// @Title  			geocode.go
// @Description  	处理高德地图地理编码的API
// @Author  		xblzbjs
// @Update  		2021-1-27

package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
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
	Adcode           string
	Street           string
	Number           string
	Location         string
	Level            string
}

//

// GeocodeResponse 高德地图地理编码响应对象
type GeocodeResponse struct {
	Status   string
	Count    string
	Info     string
	Infocode string
	Geocodes []GeocodeResult
}

// getLocationForPlace
func getLocationForPlace(address string) (location string, err error) {

	// @title     getLocationForPlace
	// @description  获取地址的经度纬度
	// @auth      xblzbjs             时间（2021/1/27）
	// @param     address         string        结构化地址信息:省份＋城市＋区县＋城镇＋乡村＋街道＋门牌号码
	// @return    location        string        经度,纬度
	//            err			  error			错误信息

	escAddress := url.QueryEscape(address)
	u := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s",
		main.GaoDeApiKey,
		escAddress,
	)
	r, err := main.httpClient.Get(u)
	if err != nil {
		return location, err
	}

	defer r.Body.Close()

	// 定义高德地图地理编码响应对象
	var geocode GeocodeResponse
	err = json.NewDecoder(r.Body).Decode(&geocode)
	if err != nil {
		return location, err
	}
	if geocode.Status != "1" || len(geocode.Geocodes) < 1 {
		return location, errors.New(fmt.Sprintf("GetLocationRequest Failed: %s", geocode.Status))
	}

	return geocode.Geocodes[0].Location, err
}

//func getLocation2(address string, city string) (location string, err error) {
//	escAddress := url.QueryEscape(address)
//	escCity := url.QueryEscape(city)
//	u := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s&city=%s",
//		GaoDeApiKey,
//		escAddress,
//		escCity,
//	)
//	r, err := httpClient.Get(u)
//	if err != nil {
//		return location, err
//	}
//
//	defer r.Body.Close()
//	// json
//	var geocode GaoDeGeocodeResponse
//
//	data, err := ioutil.ReadAll(r.Body)
//	err = json.Unmarshal(data, &geocode)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(data))
//	if geocode.Status != "1" || len(geocode.Geocodes) < 1 {
//		return location, err
//	}
//
//	return geocode.Geocodes[0].Location, err
//}

func LocationToLatLon(location string) (lat float64, lon float64) {

	// @title     LocationToLatLon
	// @description  将字符串的经度维度转化成浮点数经度和维度
	// @auth      xblzbjs             时间（2021/1/27）
	// @param     location        string        经度，纬度
	// @return    lat	          float64       经度
	//            lon			  float64		纬度

	locationList := strings.Split(location, ",")
	lon, _ = strconv.ParseFloat(locationList[0], 6)
	lat, _ = strconv.ParseFloat(locationList[1], 6)
	return lat, lon
}
