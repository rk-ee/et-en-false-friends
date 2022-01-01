package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type commWord struct {
	Value       string
	Definitions []string
	Popularity  int // where 1 is the least popular within the slice
	MorphCode   string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vvvv",
	Short: "finding false friends and working with lexitory APIs",
	Long:  `virkvirivirvavirn TODO:`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
