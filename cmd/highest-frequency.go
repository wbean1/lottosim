package cmd

import (
	"lottosim/pkg/draws"

	"github.com/spf13/cobra"
)

var numHighFreqDraws int
var highestFrequencyCmd = &cobra.Command{
	Use:   "highest-frequency",
	Short: "draw balls and make a ticket w/ high-frequency numbers",
	Long:  `draws X balls and creates a ticket from the highest-frequency numbers drawn; returns this as winner`,
	Run: func(cmd *cobra.Command, args []string) {
		draws.HighestFrequencyDraw(powerball, numHighFreqDraws)
	},
}

func init() {
	rootCmd.AddCommand(highestFrequencyCmd)
	highestFrequencyCmd.Flags().IntVar(&numHighFreqDraws, "numDraws", 10000, "number of draws (default: 5)")
}
