package main

import (
	"flag"
	"fmt"
	"github.com/xblzbjs/goweather/cli"
	"github.com/xblzbjs/goweather/gaode"
	"github.com/xblzbjs/goweather/openweather"
	"log"
	"os"
	"strings"

	"net/http"
	"time"
)

var httpClient http.Client

const (
	WeatherPeriodCurrent  = "current"
	WeatherPeriodMinutely = "minutely"
	WeatherPeriodHourly   = "hourly"
	WeatherPeriodDaily    = "daily"
	UnitsImperial         = "imperial" //"standard", "metric" and "imperial" units are available
	UnitsMetric           = "metric"
)

func ExitInvalidArguments() {
	// 若不规范的命令行参数，退出
	println("\nUsage: goweather [ -period=current|hourly|daily ] [ -units=C(摄氏度)|F(华氏度) ] <地点>...\n")
	flag.Usage()
	println()
	os.Exit(2)
}

func main() {
	httpClient = http.Client{
		Timeout: time.Second * 10,
	}

	//命令行参数
	units := flag.String("units", "C", "C(摄氏度) | F(华氏度)")
	//period := flag.String("period", "current", "current | hourly | daily")
	flag.Parse()
	// 获取地址
	places := flag.Args() //地址

	if len(places) < 1 {
		cli.ExitInvalidArguments()
	}

	// un(单位) -> string
	// 判断摄氏度和华氏度
	//var un string
	//if strings.ToUpper(*units) == "C" {
	//	un = UnitsMetric
	//} else if strings.ToUpper(*units) == "F" {
	//	un = UnitsImperial
	//} else {
	//	cli.ExitInvalidArguments()
	//}
	//
	//// period错误处理
	//if *period != WeatherPeriodCurrent &&
	//	*period != WeatherPeriodHourly &&
	//	*period != WeatherPeriodDaily {
	//	cli.ExitInvalidArguments()
	//}

	// 异步
	chs := make([]chan openweather.OpenWeatherResponseOneCall, len(places))
	errChs := make([]chan error, len(places))

	start := time.Now()

	for i, place := range places {
		chs[i] = make(chan openweather.OpenWeatherResponseOneCall, 1)
		errChs[i] = make(chan error, 1)
		go concurrentGetWeatherForPlace(place, un, *period, chs[i], errChs[i])
	}

	for i, ch := range chs {
		w := <-ch
		err := <-errChs[i]
		if err != nil {
			log.Fatal(err)
		} else {
			switch *period {
			case WeatherPeriodCurrent:
				printWeatherResult(*w.Current, places[i], un)
			case WeatherPeriodHourly:
				printWeatherResult(*w.Hourly, places[i], un)
			case WeatherPeriodDaily:
				printWeatherResult(*w.Daily, places[i], un)
			}
		}
	}

	elapsed := time.Now().Sub(start)
	fmt.Printf("运行时间:%d\n", elapsed.Milliseconds())

}



func concurrentGetWeatherForPlace(place string, units string, period string, wCh chan openweather.OpenWeatherResponseOneCall, errCh chan error) {
	w, err := getWeatherForPlace(place, units, period)
	wCh <- w
	errCh <- err
}


