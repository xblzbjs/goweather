package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new city",
	Long:  `Add will create a new city to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	for _, x := range args {
		fmt.Println(x)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
