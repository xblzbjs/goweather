package main

import (
	"goweather/cmd"

	"net/http"
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

func main() {
	cmd.Execute()
	// 	httpClient = http.Client{
	// 		Timeout: time.Second * 10,
	// 	}

	// 	// 判断摄氏度和华氏度
	// 	var un string
	// 	if strings.ToUpper(*units) == "C" {
	// 		un = UnitsMetric
	// 	} else if strings.ToUpper(*units) == "F" {
	// 		un = UnitsImperial
	// 	} else {
	// 		exitInvalidArguments()
	// 	}

	// 	// period错误处理
	// 	if *period != WeatherPeriodCurrent &&
	// 		*period != WeatherPeriodHourly &&
	// 		*period != WeatherPeriodDaily {
	// 		exitInvalidArguments()
	// 	}

	// 	start := time.Now()

	// 	for i, place := range places {
	// 		chs[i] = make(chan OpenWeatherResponseOneCall, 1)
	// 		errChs[i] = make(chan error, 1)
	// 		go concurrentGetWeatherForPlace(place, un, *period, chs[i], errChs[i])
	// 	}

	// 	for i, ch := range chs {
	// 		w := <-ch
	// 		err := <-errChs[i]
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		} else {
	// 			switch *period {
	// 			case WeatherPeriodCurrent:
	// 				printWeatherResult(*w.Current, places[i], un)
	// 			case WeatherPeriodHourly:
	// 				printWeatherResult(*w.Hourly, places[i], un)
	// 			case WeatherPeriodDaily:
	// 				printWeatherResult(*w.Daily, places[i], un)
	// 			}
	// 		}
	// 	}

	// 	elapsed := time.Now().Sub(start)
	// 	fmt.Printf("运行时间:%d\n", elapsed.Milliseconds())

	// }

}
