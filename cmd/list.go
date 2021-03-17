package cmd

import (
	"fmt"
	"goweather/city"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print cities list",
	Long:  `From json file print city list`,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := city.ReadItems("/Users/xblzbjs/.city.json")
		if err != nil {
			log.Printf("%v", err)
		}
		fmt.Println(items)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
