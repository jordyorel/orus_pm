package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ka",
	Short: "A blazing-fast package manager for Orus.",
	Long:  `ka is a next-generation package manager for the Orus programming language, designed to be fast, reliable, and easy to use.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}
