package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logs out from the registry",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Logging out...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
