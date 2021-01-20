package cmd

import (
	"lottosim/pkg/draws"

	"github.com/spf13/cobra"
)

var doublePickCmd = &cobra.Command{
	Use:   "double-pick",
	Short: "first numbers double-picked",
	Long:  `draws numbers from pool until the needed amount of numbers have been selected twice; returns this as winner`,
	Run: func(cmd *cobra.Command, args []string) {
		draws.DoubleDraw(powerball)
	},
}

func init() {
	rootCmd.AddCommand(doublePickCmd)
}
