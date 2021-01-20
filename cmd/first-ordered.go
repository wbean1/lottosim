package cmd

import (
	"lottosim/pkg/draws"

	"github.com/spf13/cobra"
)

var firstOrderedCmd = &cobra.Command{
	Use:   "first-ordered",
	Short: "first ordered ticket drawn",
	Long:  `draws tickets (via simple-random) until a ticket was drawn in numerical order; returns this as winner`,
	Run: func(cmd *cobra.Command, args []string) {
		draws.FirstOrderedDraw(powerball)
	},
}

func init() {
	rootCmd.AddCommand(firstOrderedCmd)
}
