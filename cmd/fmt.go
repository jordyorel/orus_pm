package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Formats Orus code",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Formatting code...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
