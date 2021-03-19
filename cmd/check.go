package cmd

import (
	"goweather/gaode"
	"log"

	"github.com/spf13/cobra"
)

var (
	weatherType string
	Unit        string
	Period      string
	Choice      string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the weather in a city(citys)",
	Run: func(cmd *cobra.Command, cities []string) {
		weatherType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}
		// 判断气象类型

		cityChs := make([]chan gaode.WeatherResponse, len(cities))
		errChs := make([]chan error, len(cities))
		for i, city := range cities {
			cityChs[i] = make(chan gaode.WeatherResponse, 1)
			errChs[i] = make(chan error, 1)
			go gaode.ConcurrentGetWeatherForCity(city, weatherType, cityChs[i], errChs[i])
		}
		for i, ch := range cityChs {
			w := <-ch
			err := <-errChs[i]
			if err != nil {
				log.Fatal(err)
			}
			if weatherType == "base" {
				gaode.PrintWeatherResult(*w.Lives, cities[i])
			} else if weatherType == "all" {
				gaode.PrintWeatherResult(*w.Forecasts, cities[i])
			}
		}

	},
}

func init() {
	checkCmd.Flags().StringVarP(&weatherType, "type", "t",
		"base", "Weather type(base:return live result all:return forecast result)")
	checkCmd.Flags().StringVarP(&Unit, "unit", "u", "C", "C｜F")
	checkCmd.Flags().StringVarP(&Choice, "choice", "c", "gaode",
		"Choice a weather service provider")
	checkCmd.Flags().StringVarP(&Period, "period", "p", "current",
		"current｜hourly｜daily")
	rootCmd.AddCommand(checkCmd)
}
