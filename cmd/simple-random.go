package cmd

import (
	"lottosim/pkg/draws"

	"github.com/spf13/cobra"
)

var simpleRandomCmd = &cobra.Command{
	Use:   "simple-random",
	Short: "returns the first ticket created via rand()",
	Long:  `returns the first ticket created via rand()`,
	Run: func(cmd *cobra.Command, args []string) {
		draws.SimpleRandomDraw(powerball)
	},
}

func init() {
	rootCmd.AddCommand(simpleRandomCmd)
}
