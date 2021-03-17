package main

import (
	"flag"
	"goweather/config"
	"log"
	"os"

	"net/http"

	"github.com/spf13/viper"
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

// exitInvalidArguments 若不规范的命令行参数，退出
func exitInvalidArguments() {
	println("\nUsage: goweather [ -period=current|hourly|daily ] [ -units=C(摄氏度)|F(华氏度) ] <地点>...\n")
	flag.Usage()
	os.Exit(2)
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	gaoDeKey := configuration.Api.GaoDe
	openWeatherKey := configuration.Api.OpenWeather
	log.Printf("gaode api key is %s", gaoDeKey)
	log.Printf("openweather apikey is %s", openWeatherKey)
}

func main() {

	// 	httpClient = http.Client{
	// 		Timeout: time.Second * 10,
	// 	}

	// 	//命令行
	// 	units := flag.String("units", "C", "C(摄氏度) | F(华氏度)")
	// 	period := flag.String("period", "current", "current | hourly | daily")
	// 	flag.Parse()

	// 	places := flag.Args() //地址

	// 	if len(places) < 1 {
	// 		exitInvalidArguments()
	// 	}
	// 	// un(单位) -> string
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

	// 	// 异步
	// 	chs := make([]chan OpenWeatherResponseOneCall, len(places))
	// 	errChs := make([]chan error, len(places))

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

	// func getWeatherForPlace(place string, units string, period string) (w OpenWeatherResponseOneCall, err error) {
	// 	location, err := getLocationForPlace(place)
	// 	if err != nil {
	// 		return w, err
	// 	}
	// 	lat, lon := LocationToLatLon(location)
	// 	return getWeatherForLatLon(lat, lon, units, period)
}

// func concurrentGetWeatherForPlace(place string, units string, period string, wCh chan OpenWeatherResponseOneCall, errCh chan error) {
// 	w, err := getWeatherForPlace(place, units, period)
// 	wCh <- w
// 	errCh <- err
// }

// func printWeatherResult(w interface{}, place string, units string) {
// 	// 打印天气详情
// 	fmt.Printf("%s的天气:\n", place)

// 	switch w.(type) {
// 	case OpenWeatherResponseCurrent:
// 		fmt.Print(w.(OpenWeatherResponseCurrent).Output(units))
// 	case []OpenWeatherResponseHourly:
// 		for _, h := range w.([]OpenWeatherResponseHourly) {
// 			fmt.Print(h.Output(units))
// 		}
// 	case []OpenWeatherResponseDaily:
// 		for _, h := range w.([]OpenWeatherResponseDaily) {
// 			fmt.Print(h.Output(units))
// 		}
// 	}
// }
