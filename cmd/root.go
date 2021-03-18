package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var cityArgs string

var (
	dataFile   string
	wetherType string
	cityList   string
	Unit       string
	Period     string
)

var rootCmd = &cobra.Command{
	Use:   "goweather",
	Short: "This is a weather application",
	Long:  `This will help you get live weather and forecast weather.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datefile.")
	}
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile",
		home+string(os.PathSeparator)+".city.json",
		"data file to store city weather info")
	rootCmd.PersistentFlags().StringVarP(&wetherType, "type", "t",
		"base", "气象类型(base:返回实况天气 all:返回预报天气)")
	rootCmd.PersistentFlags().StringVarP(&Unit, "unit", "u", "C", "C｜F")

	rootCmd.PersistentFlags().StringVarP(&Period, "current", "p", "current",
		"current｜hourly｜daily")
}
