package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var CliVersion = "development"
var GoVersion = "1.xx"
var BuildDate = time.Now().Format("Mon Jan 2 15:04:05")

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"transitland-route-geometry-generator Version: %s - Go Version: %s - BuildDate: %s",
			CliVersion,
			GoVersion,
			BuildDate,
		)
	},
}
