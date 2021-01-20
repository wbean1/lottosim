package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var debug bool
var powerball bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lottosim",
	Short: "making fortunes via ordering chaos",
	Long:  `lottosim finds winning lottery tickets from within the noise of rand()`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
	rootCmd.PersistentFlags().BoolVar(&powerball, "powerball", false, "draw according to powerball rules instead (default: megamillions)")
}
