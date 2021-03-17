package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var dataFile string
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
		"data file to store city")
}
