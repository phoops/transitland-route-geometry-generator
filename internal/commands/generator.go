package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate routes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate")
	},
}
