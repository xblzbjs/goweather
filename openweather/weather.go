package openweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xblzbjs/goweather/gaode"
	"strings"
	"time"
)

type OpenWeatherCondition struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type OpenWeatherResponseCurrent struct {
	Dt         int64
	Sunrise    int64
	Sunset     int64
	Temp       float32
	Feels_like float32
	Pressure   int
	Humidity   int
	Dew_point  float32
	Uvi        float32
	Clouds     int
	Visibility int
	Wind_speed float32
	//Wind_gust  float32
	Wind_deg int
	Weather  []OpenWeatherCondition
	Rain     struct {
		_1hr float32 `json:"1hr"`
	}
	Snow struct {
		_1hr float32 `json:"1hr"`
	}
}

func (w OpenWeatherResponseCurrent) Output(units string) string {
	var unitAbbr string

	switch units {
	case main.UnitsMetric:
		unitAbbr = "C"
	case main.UnitsImperial:
		unitAbbr = "F"
	}

	return fmt.Sprintf("Current: %g%s | Humidity: %d%% | %s\n",
		w.Temp,
		unitAbbr,
		w.Humidity,
		w.Weather[0].Description,
	)
}

type OpenWeatherResponseHourly struct {
	Dt         int64
	Temp       float32
	Feels_like float32
	Pressure   int
	Humidity   int
	Dew_point  float32
	Uvi        float32
	Clouds     int
	Visibility int
	Wind_speed float32
	Wind_gust  float32
	Wind_deg   int
	Weather    []OpenWeatherCondition
	Rain       struct {
		_1hr float32 `json:"1hr"`
	}
	Snow struct {
		_1hr float32 `json:"1hr"`
	}
}

func (w OpenWeatherResponseHourly) Output(units string) string {
	var unitAbbr string

	switch units {
	case main.UnitsMetric:
		unitAbbr = "C"
	case main.UnitsImperial:
		unitAbbr = "F"
	}

	t := time.Unix(w.Dt, 0)
	return fmt.Sprintf("%-9s %2d/%2d %02d:00   %5.2f%s | Humidity: %d%% | %s\n",
		t.Weekday().String(),
		t.Month(),
		t.Day(),
		t.Hour(),
		w.Temp,
		unitAbbr,
		w.Humidity,
		w.Weather[0].Description,
	)
}

type OpenWeatherResponseDaily struct {
	Dt      int64
	Sunrise int64
	Sunset  int64
	Temp    struct {
		Day   float32
		Min   float32
		Max   float32
		Night float32
		Eve   float32
		Morn  float32
	}
	Feels_like struct {
		Day   float32
		Night float32
		Eve   float32
		Morn  float32
	}
	Pressure   int
	Humidity   int
	Dew_point  float32
	Wind_speed float32
	Wind_deg   int
	Weather    []OpenWeatherCondition
	//Uvi        float32
	Clouds int
	Pop    float32
	//Visibility int
	//Wind_gust  float32
	Rain float32
	Snow float32
}

func (w OpenWeatherResponseDaily) Output(units string) string {
	var unitAbbr string

	switch units {
	case main.UnitsMetric:
		unitAbbr = "C"
	case main.UnitsImperial:
		unitAbbr = "F"
	}

	t := time.Unix(w.Dt, 0)
	return fmt.Sprintf("%-9s %2d/%2d   High: %5.2f%s Low: %5.2f%s | Humidity: %d%% | %s\n",
		t.Weekday().String(),
		t.Month(),
		t.Day(),
		w.Temp.Max,
		unitAbbr,
		w.Temp.Min,
		unitAbbr,
		w.Humidity,
		w.Weather[0].Description,
	)
}

type OpenWeatherResponseOneCall struct {
	Current *OpenWeatherResponseCurrent
	Hourly  *[]OpenWeatherResponseHourly
	Daily   *[]OpenWeatherResponseDaily
}

func getWeatherForLatLon(lat float64, lon float64, units string, period string) (weather OpenWeatherResponseOneCall, err error) {
	exclude := []string{main.WeatherPeriodMinutely}

	if period != main.WeatherPeriodCurrent {
		exclude = append(exclude, main.WeatherPeriodCurrent)
	}
	if period != main.WeatherPeriodHourly {
		exclude = append(exclude, main.WeatherPeriodHourly)
	}
	if period != main.WeatherPeriodDaily {
		exclude = append(exclude, main.WeatherPeriodDaily)
	}

	excString := strings.Join(exclude, ",")

	u := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?appid=%s&lat=%g&lon=%g&exclude=%s&units=%s&lang=zh_cn",
		OpenWeatherApiKey,
		lat,
		lon,
		excString,
		units,
	)
	r, err := main.httpClient.Get(u)
	if err != nil {
		return weather, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return weather, errors.New(fmt.Sprintf("OpenWeatherRequest Failed: %s", r.Status))
	}

	return weather, json.NewDecoder(r.Body).Decode(&weather)
}

func getWeatherForPlace(place string, units string, period string) (w openweather.OpenWeatherResponseOneCall, err error) {
	location, err := gaode.getLocationForPlace(place)
	if err != nil {
		return w, err
	}
	lat, lon := gaode.LocationToLatLon(location)
	return openweather.getWeatherForLatLon(lat, lon, units, period)
}

func concurrentGetWeatherForPlace(place string, units string, period string, wCh chan openweather.OpenWeatherResponseOneCall, errCh chan error) {
	w, err := getWeatherForPlace(place, units, period)
	wCh <- w
	errCh <- err
}

func printWeatherResult(w interface{}, place string, units string) {
	// 打印天气详情
	fmt.Printf("%s的天气:\n", place)

	switch w.(type) {
	case OpenWeatherResponseCurrent:
		fmt.Print(w.(OpenWeatherResponseCurrent).Output(units))
	case []OpenWeatherResponseHourly:
		for _, h := range w.([]OpenWeatherResponseHourly) {
			fmt.Print(h.Output(units))
		}
	case []OpenWeatherResponseDaily:
		for _, h := range w.([]OpenWeatherResponseDaily) {
			fmt.Print(h.Output(units))
		}
	}
}