package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show goweather version info.",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	output, err := ExecuteCommand("goweather", "version", args...)
	if err != nil {
		Error(cmd, args, err)
	}

	fmt.Fprint(os.Stdout, output)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
