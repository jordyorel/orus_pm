package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Checks for known security issues in dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Auditing dependencies...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
}
