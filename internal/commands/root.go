package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "transitland-route-geometry-generator",
	Short: "transitland-route-geometry-generator",
	Long: `
Generate geoemtries for your gtfs routes in transistland.
Uses your trips geometries in order to compute their union as route shape
More information at https://github.com/phoops/transitland-route-geometry-generator`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
