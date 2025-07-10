package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generates HTML documentation from Orus source and comments",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating documentation...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}
