package cmd

import (
	"lottosim/pkg/draws"

	"github.com/spf13/cobra"
)

var fullTicketDoubleCmd = &cobra.Command{
	Use:   "full-ticket-double",
	Short: "first ticket double drawn",
	Long:  `draws tickets (via simple-random) until the same ticket has been drawn twice; returns this as winner`,
	Run: func(cmd *cobra.Command, args []string) {
		draws.FullTicketDoubleDraw(powerball)
	},
}

func init() {
	rootCmd.AddCommand(fullTicketDoubleCmd)
}
