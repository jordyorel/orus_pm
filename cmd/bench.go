package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Integrates with Orus VM’s benchmarking system",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Benchmarking...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(benchCmd)
}
