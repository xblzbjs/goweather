package cmd

import (
	"fmt"
	"goweather/gaode"
	"log"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the weather in a city(citys)",
	Run: func(cmd *cobra.Command, cities []string) {
		cityChs := make([]chan gaode.WeatherResponse, len(cities))
		errChs := make([]chan error, len(cities))
		fmt.Println("城市信息如下", cities)
		for i, city := range cities {
			cityChs[i] = make(chan gaode.WeatherResponse, 1)
			errChs[i] = make(chan error, 1)
			go gaode.ConcurrentGetWeatherForCity(city, cityChs[i], errChs[i])
		}
		for i, ch := range cityChs {
			w := <-ch
			err := <-errChs[i]
			if err != nil {
				log.Fatal(err)
			} else {
				gaode.PrintWeatherResult(*w.Lives, cities[i])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
