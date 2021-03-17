package cmd

import (
	"fmt"
	"goweather/city"
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new city",
	Long:  `Add will create a new city to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := city.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		items = append(items, city.Item{Text: x})
	}
	err = city.SaveItems(dataFile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
