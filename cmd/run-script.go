package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runScriptCmd = &cobra.Command{
	Use:   "run-script [script]",
	Short: "Runs a script from ka.toml",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Running script %s...\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runScriptCmd)
}

